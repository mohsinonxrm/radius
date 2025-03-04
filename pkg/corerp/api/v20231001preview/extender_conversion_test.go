/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v20231001preview

import (
	"encoding/json"
	"testing"

	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/corerp/datamodel"
	"github.com/radius-project/radius/pkg/portableresources"
	rpv1 "github.com/radius-project/radius/pkg/rp/v1"
	"github.com/radius-project/radius/pkg/to"
	"github.com/radius-project/radius/test/testutil"
	"github.com/radius-project/radius/test/testutil/resourcetypeutil"
	"github.com/stretchr/testify/require"
)

func TestExtender_ConvertVersionedToDataModel(t *testing.T) {
	testset := []struct {
		desc     string
		file     string
		expected *datamodel.Extender
	}{
		{
			desc: "extender resource provisioning manual",
			file: "extender_manual.json",
			expected: &datamodel.Extender{
				BaseResource: v1.BaseResource{
					TrackedResource: v1.TrackedResource{
						ID:   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/extenders/extender0",
						Name: "extender0",
						Type: datamodel.ExtenderResourceType,
						Tags: map[string]string{},
					},
					InternalMetadata: v1.InternalMetadata{
						CreatedAPIVersion:      "",
						UpdatedAPIVersion:      "2023-10-01-preview",
						AsyncProvisioningState: v1.ProvisioningStateAccepted,
					},
					SystemData: v1.SystemData{},
				},
				Properties: datamodel.ExtenderProperties{
					BasicResourceProperties: rpv1.BasicResourceProperties{
						Application: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication",
						Environment: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0",
					},
					AdditionalProperties: map[string]any{"fromNumber": "222-222-2222"},
					ResourceProvisioning: portableresources.ResourceProvisioningManual,
					Secrets:              map[string]any{"accountSid": "sid", "authToken": "token"},
					ResourceRecipe:       portableresources.ResourceRecipe{Name: "default"},
				},
			},
		},
		{
			desc: "extender resource provisioning manual (no secrets)",
			file: "extender_manual_nosecrets.json",
			expected: &datamodel.Extender{
				BaseResource: v1.BaseResource{
					TrackedResource: v1.TrackedResource{
						ID:   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/extenders/extender0",
						Name: "extender0",
						Type: datamodel.ExtenderResourceType,
						Tags: map[string]string{},
					},
					InternalMetadata: v1.InternalMetadata{
						CreatedAPIVersion:      "",
						UpdatedAPIVersion:      "2023-10-01-preview",
						AsyncProvisioningState: v1.ProvisioningStateAccepted,
					},
					SystemData: v1.SystemData{},
				},
				Properties: datamodel.ExtenderProperties{
					BasicResourceProperties: rpv1.BasicResourceProperties{
						Application: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication",
						Environment: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0",
					},
					AdditionalProperties: map[string]any{"fromNumber": "222-222-2222"},
					ResourceProvisioning: portableresources.ResourceProvisioningManual,
					ResourceRecipe:       portableresources.ResourceRecipe{Name: "default"},
				},
			},
		},
		{
			desc: "extender resource recipe",
			file: "extender_recipe.json",
			expected: &datamodel.Extender{
				BaseResource: v1.BaseResource{
					TrackedResource: v1.TrackedResource{
						ID:   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/extenders/extender0",
						Name: "extender0",
						Type: datamodel.ExtenderResourceType,
						Tags: map[string]string{},
					},
					InternalMetadata: v1.InternalMetadata{
						CreatedAPIVersion:      "",
						UpdatedAPIVersion:      "2023-10-01-preview",
						AsyncProvisioningState: v1.ProvisioningStateAccepted,
					},
					SystemData: v1.SystemData{},
				},
				Properties: datamodel.ExtenderProperties{
					BasicResourceProperties: rpv1.BasicResourceProperties{
						Application: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication",
						Environment: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0",
					},
					ResourceProvisioning: portableresources.ResourceProvisioningRecipe,
					ResourceRecipe:       portableresources.ResourceRecipe{Name: "test-recipe"},
				},
			},
		},
	}

	for _, payload := range testset {
		// arrange
		rawPayload := testutil.ReadFixture(payload.file)
		versionedResource := &ExtenderResource{}
		err := json.Unmarshal(rawPayload, versionedResource)
		require.NoError(t, err)

		// act
		dm, err := versionedResource.ConvertTo()

		// assert
		require.NoError(t, err)
		convertedResource := dm.(*datamodel.Extender)

		require.Equal(t, payload.expected, convertedResource)
	}
}

func TestExtender_ConvertDataModelToVersioned(t *testing.T) {
	testset := []struct {
		desc     string
		file     string
		expected *ExtenderResource
	}{
		{
			desc: "extender resource provisioning manual datamodel",
			file: "extenderdatamodel_manual.json",
			expected: &ExtenderResource{
				Location: to.Ptr(""),
				Properties: &ExtenderProperties{
					Environment:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0"),
					Application:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication"),
					ResourceProvisioning: to.Ptr(ResourceProvisioningManual),
					ProvisioningState:    to.Ptr(ProvisioningStateAccepted),
					Recipe:               &Recipe{Name: to.Ptr(""), Parameters: nil},
					AdditionalProperties: map[string]any{"fromNumber": "222-222-2222"},
					Status:               resourcetypeutil.MustPopulateResourceStatus(&ResourceStatus{}),
				},
				Tags: map[string]*string{
					"env": to.Ptr("dev"),
				},
				ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/extenders/extender0"),
				Name: to.Ptr("extender0"),
				Type: to.Ptr(datamodel.ExtenderResourceType),
			},
		},
		{
			desc: "extender resource provisioning manual datamodel (no secrets)",
			file: "extenderdatamodel_manual_nosecrets.json",
			expected: &ExtenderResource{
				Location: to.Ptr(""),
				Properties: &ExtenderProperties{
					Environment:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0"),
					Application:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication"),
					ResourceProvisioning: to.Ptr(ResourceProvisioningManual),
					ProvisioningState:    to.Ptr(ProvisioningStateAccepted),
					Recipe:               &Recipe{Name: to.Ptr(""), Parameters: nil},
					AdditionalProperties: map[string]any{"fromNumber": "222-222-2222"},
					Status:               &ResourceStatus{},
				},
				Tags: map[string]*string{
					"env": to.Ptr("dev"),
				},
				ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/extenders/extender0"),
				Name: to.Ptr("extender0"),
				Type: to.Ptr(datamodel.ExtenderResourceType),
			},
		},
		{
			desc: "extender resource recipe datamodel",
			file: "extenderdatamodel_recipe.json",
			expected: &ExtenderResource{
				Location: to.Ptr(""),
				Properties: &ExtenderProperties{
					Environment:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0"),
					Application:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication"),
					ResourceProvisioning: to.Ptr(ResourceProvisioningRecipe),
					ProvisioningState:    to.Ptr(ProvisioningStateAccepted),
					Recipe:               &Recipe{Name: to.Ptr("test-recipe"), Parameters: nil},
					Status:               resourcetypeutil.MustPopulateResourceStatus(&ResourceStatus{}),
				},
				Tags: map[string]*string{
					"env": to.Ptr("dev"),
				},
				ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/extenders/extender0"),
				Name: to.Ptr("extender0"),
				Type: to.Ptr(datamodel.ExtenderResourceType),
			},
		},
	}

	for _, tc := range testset {
		t.Run(tc.desc, func(t *testing.T) {
			rawPayload := testutil.ReadFixture(tc.file)
			resource := &datamodel.Extender{}
			err := json.Unmarshal(rawPayload, resource)
			require.NoError(t, err)

			versionedResource := &ExtenderResource{}
			err = versionedResource.ConvertFrom(resource)
			require.NoError(t, err)

			// Skip system data comparison
			versionedResource.SystemData = nil

			require.Equal(t, tc.expected, versionedResource)
		})
	}
}

func TestExtender_ConvertFromValidation(t *testing.T) {
	validationTests := []struct {
		src v1.DataModelInterface
		err error
	}{
		{&resourcetypeutil.FakeResource{}, v1.ErrInvalidModelConversion},
		{nil, v1.ErrInvalidModelConversion},
	}

	for _, tc := range validationTests {
		versioned := &ExtenderResource{}
		err := versioned.ConvertFrom(tc.src)
		require.ErrorAs(t, tc.err, &err)
	}
}
