// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package worker

import (
	"context"

	"github.com/go-logr/logr"
	manager "github.com/project-radius/radius/pkg/armrpc/asyncoperation/statusmanager"
	"github.com/project-radius/radius/pkg/armrpc/hostoptions"
	corerp_deployment "github.com/project-radius/radius/pkg/corerp/backend/deployment"
	corerp_model "github.com/project-radius/radius/pkg/corerp/model"
	"github.com/project-radius/radius/pkg/deployment"
	"github.com/project-radius/radius/pkg/ucp/dataprovider"
	queue "github.com/project-radius/radius/pkg/ucp/queue/client"
	qprovider "github.com/project-radius/radius/pkg/ucp/queue/provider"
)

// Service is the base worker service implementation to initialize and start worker.
type Service struct {
	// ProviderName is the name of provider namespace.
	ProviderName string
	// Options is the server hosting options.
	Options hostoptions.HostOptions
	// StorageProvider is the provider of storage client.
	StorageProvider dataprovider.DataStorageProvider
	// OperationStatusManager is the manager of the operation status.
	OperationStatusManager manager.StatusManager
	// Controllers is the registry of the async operation controllers.
	Controllers *ControllerRegistry
	// RequestQueue is the queue client for async operation request message.
	RequestQueue queue.Client
	// DeploymentProcessors is the map of deployment processors available in the program.
	DeploymentProcessors map[string]deployment.DeploymentProcessor
}

// Init initializes worker service.
func (s *Service) Init(ctx context.Context) error {
	s.StorageProvider = dataprovider.NewStorageProvider(s.Options.Config.StorageProvider)
	qp := qprovider.New(s.ProviderName, s.Options.Config.QueueProvider)
	opSC, err := s.StorageProvider.GetStorageClient(ctx, s.ProviderName+"/operationstatuses")
	if err != nil {
		return err
	}
	s.RequestQueue, err = qp.GetClient(ctx)
	if err != nil {
		return err
	}
	s.OperationStatusManager = manager.New(opSC, s.RequestQueue, s.ProviderName, s.Options.Config.Env.RoleLocation)
	s.Controllers = NewControllerRegistry(s.StorageProvider)

	// Should we create DPs here?
	// DeploymentProcessors
	coreDP, err := corerp_deployment.NewCoreRPDeploymentProcessor(corerp_model.ApplicationModel{}, s.StorageProvider, nil, nil)
	if err != nil {
		return err
	}
	s.DeploymentProcessors["core-rp"] = coreDP
	// ConnectorRP will also be added here

	return nil
}

// Start starts the worker.
func (s *Service) Start(ctx context.Context, opt Options) error {
	logger := logr.FromContextOrDiscard(ctx)
	ctx = hostoptions.WithContext(ctx, s.Options.Config)

	// Create and start worker.
	worker := New(opt, s.OperationStatusManager, s.RequestQueue, s.Controllers)

	logger.Info("Start Worker...")
	if err := worker.Start(ctx); err != nil {
		logger.Error(err, "failed to start worker...")
	}

	logger.Info("Worker stopped...")
	return nil
}
