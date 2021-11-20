// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package cmd

import (
	"github.com/spf13/cobra"
)

var envInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a Radius environment",
	Long:  `Create a Radius environment`,
}

func init() {
	envCmd.AddCommand(envInitCmd)
}
