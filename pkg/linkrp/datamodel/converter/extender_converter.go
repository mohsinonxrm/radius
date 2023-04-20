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

// ExtenderDataModelFromVersioned converts version agnostic Extender datamodel to versioned model.
func ExtenderDataModelToVersioned(model *datamodel.Extender, version string) (v1.VersionedModelInterface, error) {
	switch version {
	case v20230415preview.Version:
		versioned := &v20230415preview.ExtenderResource{}
		err := versioned.ConvertFrom(model)
		if err != nil {
			return nil, err
		}

		return versioned, nil

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}

// ExtenderDataModelToVersioned converts versioned Extender model to datamodel.
func ExtenderDataModelFromVersioned(content []byte, version string) (*datamodel.Extender, error) {
	switch version {
	case v20230415preview.Version:
		am := &v20230415preview.ExtenderResource{}
		if err := json.Unmarshal(content, am); err != nil {
			return nil, err
		}
		dm, err := am.ConvertTo()
		return dm.(*datamodel.Extender), err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}
