//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboundaries

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the DataBoundaries group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal *arm.Client
}

// NewClient creates a new instance of Client with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		internal: cl,
	}
	return client, nil
}

// GetScope - Get data boundary at specified scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - scope - The scope at which the operation is performed.
//   - defaultParam - Default string modeled as parameter for auto generation to work correctly.
//   - options - ClientGetScopeOptions contains the optional parameters for the Client.GetScope method.
func (client *Client) GetScope(ctx context.Context, scope string, defaultParam DefaultName, options *ClientGetScopeOptions) (ClientGetScopeResponse, error) {
	var err error
	const operationName = "Client.GetScope"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getScopeCreateRequest(ctx, scope, defaultParam, options)
	if err != nil {
		return ClientGetScopeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetScopeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetScopeResponse{}, err
	}
	resp, err := client.getScopeHandleResponse(httpResp)
	return resp, err
}

// getScopeCreateRequest creates the GetScope request.
func (client *Client) getScopeCreateRequest(ctx context.Context, scope string, defaultParam DefaultName, options *ClientGetScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/dataBoundaries/{default}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getScopeHandleResponse handles the GetScope response.
func (client *Client) getScopeHandleResponse(resp *http.Response) (ClientGetScopeResponse, error) {
	result := ClientGetScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataBoundaryDefinition); err != nil {
		return ClientGetScopeResponse{}, err
	}
	return result, nil
}

// GetTenant - Get data boundary of tenant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - defaultParam - Default string modeled as parameter for auto generation to work correctly.
//   - options - ClientGetTenantOptions contains the optional parameters for the Client.GetTenant method.
func (client *Client) GetTenant(ctx context.Context, defaultParam DefaultName, options *ClientGetTenantOptions) (ClientGetTenantResponse, error) {
	var err error
	const operationName = "Client.GetTenant"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getTenantCreateRequest(ctx, defaultParam, options)
	if err != nil {
		return ClientGetTenantResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetTenantResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetTenantResponse{}, err
	}
	resp, err := client.getTenantHandleResponse(httpResp)
	return resp, err
}

// getTenantCreateRequest creates the GetTenant request.
func (client *Client) getTenantCreateRequest(ctx context.Context, defaultParam DefaultName, options *ClientGetTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/dataBoundaries/{default}"
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTenantHandleResponse handles the GetTenant response.
func (client *Client) getTenantHandleResponse(resp *http.Response) (ClientGetTenantResponse, error) {
	result := ClientGetTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataBoundaryDefinition); err != nil {
		return ClientGetTenantResponse{}, err
	}
	return result, nil
}

// Put - Opt-in tenant to data boundary.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - defaultParam - Default string modeled as parameter for auto generation to work correctly.
//   - dataBoundaryDefinition - The data boundary to opt the tenant to.
//   - options - ClientPutOptions contains the optional parameters for the Client.Put method.
func (client *Client) Put(ctx context.Context, defaultParam DefaultName, dataBoundaryDefinition DataBoundaryDefinition, options *ClientPutOptions) (ClientPutResponse, error) {
	var err error
	const operationName = "Client.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, defaultParam, dataBoundaryDefinition, options)
	if err != nil {
		return ClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *Client) putCreateRequest(ctx context.Context, defaultParam DefaultName, dataBoundaryDefinition DataBoundaryDefinition, options *ClientPutOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/dataBoundaries/{default}"
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dataBoundaryDefinition); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *Client) putHandleResponse(resp *http.Response) (ClientPutResponse, error) {
	result := ClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataBoundaryDefinition); err != nil {
		return ClientPutResponse{}, err
	}
	return result, nil
}
