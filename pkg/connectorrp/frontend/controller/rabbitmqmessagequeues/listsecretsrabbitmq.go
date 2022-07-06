// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package rabbitmqmessagequeues

import (
	"context"
	"errors"
	"net/http"

	manager "github.com/project-radius/radius/pkg/armrpc/asyncoperation/statusmanager"
	ctrl "github.com/project-radius/radius/pkg/armrpc/frontend/controller"
	"github.com/project-radius/radius/pkg/armrpc/servicecontext"
	"github.com/project-radius/radius/pkg/connectorrp/datamodel"
	"github.com/project-radius/radius/pkg/connectorrp/datamodel/converter"
	"github.com/project-radius/radius/pkg/connectorrp/frontend/deployment"
	"github.com/project-radius/radius/pkg/connectorrp/renderers"
	"github.com/project-radius/radius/pkg/radrp/rest"
	"github.com/project-radius/radius/pkg/ucp/store"
)

var _ ctrl.Controller = (*ListSecretsRabbitMQMessageQueue)(nil)

// ListSecretsRabbitMQMessageQueue is the controller implementation to list secrets for the to access the connected rabbitMQ resource resource id passed in the request body.
type ListSecretsRabbitMQMessageQueue struct {
	ctrl.BaseController
}

// NewListSecretsRabbitMQMessageQueue creates a new instance of ListSecretsRabbitMQMessageQueue.
func NewListSecretsRabbitMQMessageQueue(ds store.StorageClient, sm manager.StatusManager, dp deployment.DeploymentProcessor) (ctrl.Controller, error) {
	return &ListSecretsRabbitMQMessageQueue{ctrl.NewBaseController(ds, sm, dp)}, nil
}

// Run returns secrets values for the specified RabbitMQMessageQueue resource
func (ctrl *ListSecretsRabbitMQMessageQueue) Run(ctx context.Context, req *http.Request) (rest.Response, error) {
	sCtx := servicecontext.ARMRequestContextFromContext(ctx)

	resource := &datamodel.RabbitMQMessageQueue{}
	// Request route for listsecrets has name of the operation as suffix which should be removed to get the resource id.
	// route id format: subscriptions/<subscription_id>/resourceGroups/<resource_group>/providers/Applications.Connector/mongoDatabases/<resource_name>/listsecrets
	parsedResourceID := sCtx.ResourceID.Truncate()
	_, err := ctrl.GetResource(ctx, parsedResourceID.String(), resource)
	if err != nil {
		if errors.Is(&store.ErrNotFound{}, err) {
			return rest.NewNotFoundResponse(sCtx.ResourceID), nil
		}
		return nil, err
	}

	secrets, err := ctrl.DeploymentProcessor.FetchSecrets(ctx, deployment.ResourceData{ID: sCtx.ResourceID, Resource: resource, OutputResources: resource.Properties.Status.OutputResources, ComputedValues: resource.ComputedValues, SecretValues: resource.SecretValues})
	if err != nil {
		return nil, err
	}

	redisSecrets := datamodel.RabbitMQSecrets{
		ConnectionString: secrets[renderers.ConnectionStringValue].(string),
	}

	versioned, _ := converter.RabbitMQSecretsDataModelToVersioned(&redisSecrets, sCtx.APIVersion)
	return rest.NewOKResponse(versioned), nil
}
