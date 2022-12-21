// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package datamodel

import (
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/rp"
)

// RabbitMQMessageQueue represents RabbitMQMessageQueue link resource.
type RabbitMQMessageQueue struct {
	v1.BaseResource

	// Properties is the properties of the resource.
	Properties RabbitMQMessageQueueProperties `json:"properties"`

	// LinkMetadata represents internal DataModel properties common to all link types.
	LinkMetadata
}

func (rabbitmq RabbitMQMessageQueue) ResourceTypeName() string {
	return "Applications.Link/rabbitMQMessageQueues"
}

// RabbitMQMessageQueueProperties represents the properties of RabbitMQMessageQueue response resource.
type RabbitMQMessageQueueProperties struct {
	rp.BasicResourceProperties
	ProvisioningState v1.ProvisioningState `json:"provisioningState,omitempty"`
	Queue             string               `json:"queue"`
	Recipe            LinkRecipe           `json:"recipe,omitempty"`
	Secrets           RabbitMQSecrets      `json:"secrets,omitempty"`
	Mode              LinkMode             `json:"mode,omitempty"`
}

// Secrets values consisting of secrets provided for the resource
type RabbitMQSecrets struct {
	ConnectionString string `json:"connectionString"`
}

func (rabbitmq RabbitMQSecrets) ResourceTypeName() string {
	return "Applications.Link/rabbitMQMessageQueues"
}
