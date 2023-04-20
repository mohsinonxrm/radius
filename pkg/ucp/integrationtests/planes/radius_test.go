// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package planes

import (
	"testing"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/ucp/integrationtests/testserver"
)

const (
	radiusPlaneCollectionURL          = "/planes/radius?api-version=2023-04-15-preview"
	radiusPlaneResourceURL            = "/planes/radius/local?api-version=2023-04-15-preview"
	radiusPlaneRequestFixture         = "testdata/radiusplane_v20230415preview_requestbody.json"
	radiusPlaneResponseFixture        = "testdata/radiusplane_v20230415preview_responsebody.json"
	radiusPlaneListResponseFixture    = "testdata/radiusplane_v20230415preview_list_responsebody.json"
	radiusPlaneUpdatedRequestFixture  = "testdata/radiusplane_updated_v20230415preview_requestbody.json"
	radiusPlaneUpdatedResponseFixture = "testdata/radiusplane_updated_v20230415preview_responsebody.json"
)

func Test_RadiusPlane_PUT_Create(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", radiusPlaneResourceURL, radiusPlaneRequestFixture)
	response.EqualsFixture(200, radiusPlaneResponseFixture)
}

func Test_RadiusPlane_PUT_Update(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", radiusPlaneResourceURL, radiusPlaneRequestFixture)
	response.EqualsFixture(200, radiusPlaneResponseFixture)

	response = server.MakeFixtureRequest("PUT", radiusPlaneResourceURL, radiusPlaneUpdatedRequestFixture)
	response.EqualsFixture(200, radiusPlaneUpdatedResponseFixture)
}

func Test_RadiusPlane_GET_Empty(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeRequest("GET", radiusPlaneResourceURL, nil)
	response.EqualsErrorCode(404, v1.CodeNotFound)
}

func Test_RadiusPlane_GET_Found(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", radiusPlaneResourceURL, radiusPlaneRequestFixture)
	response.EqualsFixture(200, radiusPlaneResponseFixture)

	response = server.MakeRequest("GET", radiusPlaneResourceURL, nil)
	response.EqualsFixture(200, radiusPlaneResponseFixture)
}

func Test_RadiusPlane_LIST(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	// Add a radius plane
	response := server.MakeFixtureRequest("PUT", radiusPlaneResourceURL, radiusPlaneRequestFixture)
	response.EqualsFixture(200, radiusPlaneResponseFixture)

	// Add an AWS plane
	response = server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneRequestFixture)
	response.EqualsFixture(200, awsPlaneResponseFixture)

	// Verify that /planes/radius URL returns planes only with the radius plane type.
	response = server.MakeRequest("GET", radiusPlaneCollectionURL, nil)
	response.EqualsFixture(200, radiusPlaneListResponseFixture)
}

func Test_RadiusPlane_DELETE_DoesNotExist(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeRequest("DELETE", radiusPlaneResourceURL, nil)
	response.EqualsResponse(204, nil)
}

func Test_RadiusPlane_DELETE_Found(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", radiusPlaneResourceURL, radiusPlaneRequestFixture)
	response.EqualsFixture(200, radiusPlaneResponseFixture)

	response = server.MakeRequest("DELETE", radiusPlaneResourceURL, nil)
	response.EqualsResponse(200, nil)
}
