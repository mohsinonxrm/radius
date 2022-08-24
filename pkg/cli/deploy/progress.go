// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package deploy

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gosuri/uilive"
	"github.com/mattn/go-isatty"
	"github.com/project-radius/radius/pkg/cli/clients"
	"github.com/project-radius/radius/pkg/cli/output"
)

func NewProgressListener(progressChan <-chan clients.ResourceProgress) ProgressListener {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return &InteractiveListener{
			progressChan: progressChan,
			writerDone:   &sync.WaitGroup{},
			Spinner:      output.ProgressDefaultSpinner,
		}
	} else {
		return &NoOpListener{
			progressChan: progressChan,
		}
	}
}

type ProgressListener interface {
	// Run is called to print progress to the command line. This should be called from
	// a goroutine because it blocks until the progress channel is closed.
	Run()
}

type NoOpListener struct {
	progressChan <-chan clients.ResourceProgress
}

func (listener *NoOpListener) Run() {
	for range listener.progressChan {
		// Do nothing except drain the updates.
	}
}

type InteractiveListener struct {
	progressChan <-chan clients.ResourceProgress
	Spinner      []string
	mutex        sync.Mutex
	entries      []Entry
	errors       []string
	spinnerIndex int
	writerDone   *sync.WaitGroup
}

type Entry struct {
	// FinalState is an optional token that will replace the spinner with static text.
	FinalState string

	// Format is the format string used to build the output line. It is expected to contain a placeholder
	// for the spinner/final-state token.
	Format string
}

func (listener *InteractiveListener) addEntry(format string) int {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	listener.entries = append(listener.entries, Entry{Format: format})
	return len(listener.entries) - 1
}

func (listener *InteractiveListener) addError(format string) int {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	listener.errors = append(listener.errors, format)
	return len(listener.errors) - 1
}

func (listener *InteractiveListener) updateEntry(index int, state string, format string) {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	listener.entries[index] = Entry{FinalState: state, Format: format}
}

func (listener *InteractiveListener) Run() {
	ticker := time.NewTicker(500 * time.Millisecond)

	progressDone := make(chan struct{})
	writerDone := make(chan struct{})

	// Main loop that updates spinner position and writes output. This runs concurrently with accepting updates.
	go func() {
		writer := uilive.New()
		writer.Start()

		paint := func() {
			listener.mutex.Lock()

			// Advance to next spinner position
			listener.spinnerIndex = (listener.spinnerIndex + 1) % len(listener.Spinner)

			// Replay all output lines
			for _, entry := range listener.entries {
				if entry.FinalState == "" {
					fmt.Fprintf(writer.Newline(), entry.Format+"\n", listener.Spinner[listener.spinnerIndex])
				} else {
					fmt.Fprintf(writer.Newline(), entry.Format+"\n", entry.FinalState)
				}
			}

			for _, er := range listener.errors {
				fmt.Fprintf(writer.Newline(), er+"\n")
			}

			listener.mutex.Unlock()
		}

	writer:
		for {
			select {
			case <-progressDone:
				paint() // Update UI once then terminate
				break writer
			case <-ticker.C:
				paint()
			}
		}

		writer.Stop()
		close(writerDone)
	}()

	// Storage for resources we've already 'seen'. This doesn't need to be accessed concurrently.
	resourceToLineIndexMap := map[string]int{}

	// Main loop that processes updates to resources. This runs concurrently with writing output.
	for update := range listener.progressChan {
		if !output.ShowResource(update.Resource) {
			continue
		}

		// NOTE: resources can go immediately to the Completed state without first
		// going to the started state.
		line, found := resourceToLineIndexMap[update.Resource.String()]
		if !found {
			line = listener.addEntry(output.FormatResourceForProgressDisplay(update.Resource))
			resourceToLineIndexMap[update.Resource.String()] = line
		}

		switch update.Status {
		case clients.StatusFailed:
			listener.updateEntry(line, output.ProgressFailed, output.FormatResourceForProgressDisplay(update.Resource))
		case clients.StatusCompleted:
			listener.updateEntry(line, output.ProgressCompleted, output.FormatResourceForProgressDisplay(update.Resource))
		case clients.StatusRetrying:
			listener.updateEntry(line, output.ProgressRetrying, output.FormatResourceForProgressDisplay(update.Resource))
			listener.addError(update.Message)
		}
	}

	// Force a final UI update and drain any updates in progress.
	ticker.Stop()
	close(progressDone)
	<-writerDone
}
