// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package converter

import (
	"encoding/json"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp/api/v20230415preview"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
)

// RabbitMQMessageQueueDataModelFromVersioned converts version agnostic RabbitMQMessageQueue datamodel to versioned model.
func RabbitMQMessageQueueDataModelToVersioned(model *datamodel.RabbitMQMessageQueue, version string) (v1.VersionedModelInterface, error) {
	switch version {
	case v20230415preview.Version:
		versioned := &v20230415preview.RabbitMQMessageQueueResource{}
		err := versioned.ConvertFrom(model)
		return versioned, err
	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}

// RabbitMQMessageQueueDataModelToVersioned converts versioned RabbitMQMessageQueue model to datamodel.
func RabbitMQMessageQueueDataModelFromVersioned(content []byte, version string) (*datamodel.RabbitMQMessageQueue, error) {
	switch version {
	case v20230415preview.Version:
		versioned := &v20230415preview.RabbitMQMessageQueueResource{}
		if err := json.Unmarshal(content, versioned); err != nil {
			return nil, err
		}
		dm, err := versioned.ConvertTo()
		return dm.(*datamodel.RabbitMQMessageQueue), err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}

// RabbitMQSecretsDataModelFromVersioned converts version agnostic RabbitMQSecrets datamodel to versioned model.
func RabbitMQSecretsDataModelToVersioned(model *datamodel.RabbitMQSecrets, version string) (v1.VersionedModelInterface, error) {
	switch version {
	case v20230415preview.Version:
		versioned := &v20230415preview.RabbitMQSecrets{}
		err := versioned.ConvertFrom(model)
		return versioned, err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}
