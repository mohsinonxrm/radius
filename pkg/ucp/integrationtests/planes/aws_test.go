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
	awsPlaneCollectionURL          = "/planes/aws?api-version=2023-04-15-preview"
	awsPlaneResourceURL            = "/planes/aws/aws?api-version=2023-04-15-preview"
	awsPlaneRequestFixture         = "testdata/awsplane_v20230415preview_requestbody.json"
	awsPlaneResponseFixture        = "testdata/awsplane_v20230415preview_responsebody.json"
	awsPlaneListResponseFixture    = "testdata/awsplane_v20230415preview_list_responsebody.json"
	awsPlaneUpdatedRequestFixture  = "testdata/awsplane_updated_v20230415preview_requestbody.json"
	awsPlaneUpdatedResponseFixture = "testdata/awsplane_updated_v20230415preview_responsebody.json"
)

func Test_AWSPlane_PUT_Create(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneRequestFixture)
	response.EqualsFixture(200, awsPlaneResponseFixture)
}

func Test_AWSPlane_PUT_Update(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneRequestFixture)
	response.EqualsFixture(200, awsPlaneResponseFixture)

	response = server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneUpdatedRequestFixture)
	response.EqualsFixture(200, awsPlaneUpdatedResponseFixture)
}

func Test_AWSPlane_GET_Empty(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeRequest("GET", awsPlaneResourceURL, nil)
	response.EqualsErrorCode(404, v1.CodeNotFound)
}

func Test_AWSPlane_GET_Found(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneRequestFixture)
	response.EqualsFixture(200, awsPlaneResponseFixture)

	response = server.MakeRequest("GET", awsPlaneResourceURL, nil)
	response.EqualsFixture(200, awsPlaneResponseFixture)
}

func Test_AWSPlane_LIST(t *testing.T) {
	t.Skip("This functionality is currently broken. See https://github.com/project-radius/radius/issues/4878")

	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneRequestFixture)
	response.EqualsFixture(200, awsPlaneResponseFixture)

	response = server.MakeRequest("GET", awsPlaneCollectionURL, nil)
	response.EqualsFixture(200, awsPlaneListResponseFixture)
}

func Test_AWSPlane_DELETE_DoesNotExist(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeRequest("DELETE", awsPlaneResourceURL, nil)
	response.EqualsResponse(204, nil)
}

func Test_AWSPlane_DELETE_Found(t *testing.T) {
	server := testserver.Start(t)
	defer server.Close()

	response := server.MakeFixtureRequest("PUT", awsPlaneResourceURL, awsPlaneRequestFixture)
	response.EqualsFixture(200, awsPlaneResponseFixture)

	response = server.MakeRequest("DELETE", awsPlaneResourceURL, nil)
	response.EqualsResponse(200, nil)
}
