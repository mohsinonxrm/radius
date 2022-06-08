//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationsClient contains the methods for the Applications group.
// Don't use this type directly, use NewApplicationsClient() instead.
type ApplicationsClient struct {
	ep string
	pl runtime.Pipeline
	rootScope string
}

// NewApplicationsClient creates a new instance of ApplicationsClient with the specified values.
func NewApplicationsClient(con *arm.Connection, rootScope string) *ApplicationsClient {
	return &ApplicationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), rootScope: rootScope}
}

// CreateOrUpdate - Create or update an Application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationsClient) CreateOrUpdate(ctx context.Context, applicationName string, applicationResource ApplicationResource, options *ApplicationsCreateOrUpdateOptions) (ApplicationsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, applicationName, applicationResource, options)
	if err != nil {
		return ApplicationsCreateOrUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return ApplicationsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ApplicationsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationsClient) createOrUpdateCreateRequest(ctx context.Context, applicationName string, applicationResource ApplicationResource, options *ApplicationsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/applications/{applicationName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, applicationResource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationsClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationsCreateOrUpdateResponse, error) {
	result := ApplicationsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResource); err != nil {
		return ApplicationsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ApplicationsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete an Application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationsClient) Delete(ctx context.Context, applicationName string, options *ApplicationsDeleteOptions) (ApplicationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, applicationName, options)
	if err != nil {
		return ApplicationsDeleteResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return ApplicationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ApplicationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ApplicationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationsClient) deleteCreateRequest(ctx context.Context, applicationName string, options *ApplicationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/applications/{applicationName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the properties of an Application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationsClient) Get(ctx context.Context, applicationName string, options *ApplicationsGetOptions) (ApplicationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, applicationName, options)
	if err != nil {
		return ApplicationsGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return ApplicationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationsClient) getCreateRequest(ctx context.Context, applicationName string, options *ApplicationsGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/applications/{applicationName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationsClient) getHandleResponse(resp *http.Response) (ApplicationsGetResponse, error) {
	result := ApplicationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResource); err != nil {
		return ApplicationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByScope - List all applications in the given scope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationsClient) ListByScope(options *ApplicationsListByScopeOptions) (*ApplicationsListByScopePager) {
	return &ApplicationsListByScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByScopeCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ApplicationsListByScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationResourceList.NextLink)
		},
	}
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *ApplicationsClient) listByScopeCreateRequest(ctx context.Context, options *ApplicationsListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/applications"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByScopeHandleResponse handles the ListByScope response.
func (client *ApplicationsClient) listByScopeHandleResponse(resp *http.Response) (ApplicationsListByScopeResponse, error) {
	result := ApplicationsListByScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceList); err != nil {
		return ApplicationsListByScopeResponse{}, err
	}
	return result, nil
}

// listByScopeHandleError handles the ListByScope error response.
func (client *ApplicationsClient) listByScopeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update the properties of an existing Application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationsClient) Update(ctx context.Context, applicationName string, applicationResource ApplicationResource, options *ApplicationsUpdateOptions) (ApplicationsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, applicationName, applicationResource, options)
	if err != nil {
		return ApplicationsUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return ApplicationsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ApplicationsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationsClient) updateCreateRequest(ctx context.Context, applicationName string, applicationResource ApplicationResource, options *ApplicationsUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/applications/{applicationName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", url.PathEscape(client.rootScope))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, applicationResource)
}

// updateHandleResponse handles the Update response.
func (client *ApplicationsClient) updateHandleResponse(resp *http.Response) (ApplicationsUpdateResponse, error) {
	result := ApplicationsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResource); err != nil {
		return ApplicationsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ApplicationsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

