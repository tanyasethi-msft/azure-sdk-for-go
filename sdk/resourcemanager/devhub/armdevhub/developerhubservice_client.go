//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevhub

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

// DeveloperHubServiceClient contains the methods for the DeveloperHubServiceClient group.
// Don't use this type directly, use NewDeveloperHubServiceClient() instead.
type DeveloperHubServiceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeveloperHubServiceClient creates a new instance of DeveloperHubServiceClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeveloperHubServiceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeveloperHubServiceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeveloperHubServiceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GeneratePreviewArtifacts - Generate preview dockerfile and manifests.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - location - The name of the Azure region.
//   - options - DeveloperHubServiceClientGeneratePreviewArtifactsOptions contains the optional parameters for the DeveloperHubServiceClient.GeneratePreviewArtifacts
//     method.
func (client *DeveloperHubServiceClient) GeneratePreviewArtifacts(ctx context.Context, location string, parameters ArtifactGenerationProperties, options *DeveloperHubServiceClientGeneratePreviewArtifactsOptions) (DeveloperHubServiceClientGeneratePreviewArtifactsResponse, error) {
	var err error
	const operationName = "DeveloperHubServiceClient.GeneratePreviewArtifacts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generatePreviewArtifactsCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return DeveloperHubServiceClientGeneratePreviewArtifactsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeveloperHubServiceClientGeneratePreviewArtifactsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeveloperHubServiceClientGeneratePreviewArtifactsResponse{}, err
	}
	resp, err := client.generatePreviewArtifactsHandleResponse(httpResp)
	return resp, err
}

// generatePreviewArtifactsCreateRequest creates the GeneratePreviewArtifacts request.
func (client *DeveloperHubServiceClient) generatePreviewArtifactsCreateRequest(ctx context.Context, location string, parameters ArtifactGenerationProperties, options *DeveloperHubServiceClientGeneratePreviewArtifactsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/generatePreviewArtifacts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// generatePreviewArtifactsHandleResponse handles the GeneratePreviewArtifacts response.
func (client *DeveloperHubServiceClient) generatePreviewArtifactsHandleResponse(resp *http.Response) (DeveloperHubServiceClientGeneratePreviewArtifactsResponse, error) {
	result := DeveloperHubServiceClientGeneratePreviewArtifactsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DeveloperHubServiceClientGeneratePreviewArtifactsResponse{}, err
	}
	return result, nil
}

// GitHubOAuth - Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - location - The name of the Azure region.
//   - options - DeveloperHubServiceClientGitHubOAuthOptions contains the optional parameters for the DeveloperHubServiceClient.GitHubOAuth
//     method.
func (client *DeveloperHubServiceClient) GitHubOAuth(ctx context.Context, location string, options *DeveloperHubServiceClientGitHubOAuthOptions) (DeveloperHubServiceClientGitHubOAuthResponse, error) {
	var err error
	const operationName = "DeveloperHubServiceClient.GitHubOAuth"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.gitHubOAuthCreateRequest(ctx, location, options)
	if err != nil {
		return DeveloperHubServiceClientGitHubOAuthResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeveloperHubServiceClientGitHubOAuthResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeveloperHubServiceClientGitHubOAuthResponse{}, err
	}
	resp, err := client.gitHubOAuthHandleResponse(httpResp)
	return resp, err
}

// gitHubOAuthCreateRequest creates the GitHubOAuth request.
func (client *DeveloperHubServiceClient) gitHubOAuthCreateRequest(ctx context.Context, location string, options *DeveloperHubServiceClientGitHubOAuthOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/githuboauth/default/getGitHubOAuthInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// gitHubOAuthHandleResponse handles the GitHubOAuth response.
func (client *DeveloperHubServiceClient) gitHubOAuthHandleResponse(resp *http.Response) (DeveloperHubServiceClientGitHubOAuthResponse, error) {
	result := DeveloperHubServiceClientGitHubOAuthResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubOAuthInfoResponse); err != nil {
		return DeveloperHubServiceClientGitHubOAuthResponse{}, err
	}
	return result, nil
}

// GitHubOAuthCallback - Callback URL to hit once authenticated with GitHub App to have the service store the OAuth token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - location - The name of the Azure region.
//   - code - The code response from authenticating the GitHub App.
//   - state - The state response from authenticating the GitHub App.
//   - options - DeveloperHubServiceClientGitHubOAuthCallbackOptions contains the optional parameters for the DeveloperHubServiceClient.GitHubOAuthCallback
//     method.
func (client *DeveloperHubServiceClient) GitHubOAuthCallback(ctx context.Context, location string, code string, state string, options *DeveloperHubServiceClientGitHubOAuthCallbackOptions) (DeveloperHubServiceClientGitHubOAuthCallbackResponse, error) {
	var err error
	const operationName = "DeveloperHubServiceClient.GitHubOAuthCallback"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.gitHubOAuthCallbackCreateRequest(ctx, location, code, state, options)
	if err != nil {
		return DeveloperHubServiceClientGitHubOAuthCallbackResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeveloperHubServiceClientGitHubOAuthCallbackResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeveloperHubServiceClientGitHubOAuthCallbackResponse{}, err
	}
	resp, err := client.gitHubOAuthCallbackHandleResponse(httpResp)
	return resp, err
}

// gitHubOAuthCallbackCreateRequest creates the GitHubOAuthCallback request.
func (client *DeveloperHubServiceClient) gitHubOAuthCallbackCreateRequest(ctx context.Context, location string, code string, state string, options *DeveloperHubServiceClientGitHubOAuthCallbackOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/githuboauth/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	reqQP.Set("code", code)
	reqQP.Set("state", state)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// gitHubOAuthCallbackHandleResponse handles the GitHubOAuthCallback response.
func (client *DeveloperHubServiceClient) gitHubOAuthCallbackHandleResponse(resp *http.Response) (DeveloperHubServiceClientGitHubOAuthCallbackResponse, error) {
	result := DeveloperHubServiceClientGitHubOAuthCallbackResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubOAuthResponse); err != nil {
		return DeveloperHubServiceClientGitHubOAuthCallbackResponse{}, err
	}
	return result, nil
}

// ListGitHubOAuth - Callback URL to hit once authenticated with GitHub App to have the service store the OAuth token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - location - The name of the Azure region.
//   - options - DeveloperHubServiceClientListGitHubOAuthOptions contains the optional parameters for the DeveloperHubServiceClient.ListGitHubOAuth
//     method.
func (client *DeveloperHubServiceClient) ListGitHubOAuth(ctx context.Context, location string, options *DeveloperHubServiceClientListGitHubOAuthOptions) (DeveloperHubServiceClientListGitHubOAuthResponse, error) {
	var err error
	const operationName = "DeveloperHubServiceClient.ListGitHubOAuth"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listGitHubOAuthCreateRequest(ctx, location, options)
	if err != nil {
		return DeveloperHubServiceClientListGitHubOAuthResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeveloperHubServiceClientListGitHubOAuthResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeveloperHubServiceClientListGitHubOAuthResponse{}, err
	}
	resp, err := client.listGitHubOAuthHandleResponse(httpResp)
	return resp, err
}

// listGitHubOAuthCreateRequest creates the ListGitHubOAuth request.
func (client *DeveloperHubServiceClient) listGitHubOAuthCreateRequest(ctx context.Context, location string, options *DeveloperHubServiceClientListGitHubOAuthOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/githuboauth"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listGitHubOAuthHandleResponse handles the ListGitHubOAuth response.
func (client *DeveloperHubServiceClient) listGitHubOAuthHandleResponse(resp *http.Response) (DeveloperHubServiceClientListGitHubOAuthResponse, error) {
	result := DeveloperHubServiceClientListGitHubOAuthResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubOAuthListResponse); err != nil {
		return DeveloperHubServiceClientListGitHubOAuthResponse{}, err
	}
	return result, nil
}
