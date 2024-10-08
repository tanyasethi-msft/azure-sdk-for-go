//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

// RegistryDataReferencesClient contains the methods for the RegistryDataReferences group.
// Don't use this type directly, use NewRegistryDataReferencesClient() instead.
type RegistryDataReferencesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistryDataReferencesClient creates a new instance of RegistryDataReferencesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryDataReferencesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryDataReferencesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryDataReferencesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetBlobReferenceSAS - Get blob reference SAS Uri.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - name - Data reference name.
//   - version - Version identifier.
//   - body - Asset id and blob uri.
//   - options - RegistryDataReferencesClientGetBlobReferenceSASOptions contains the optional parameters for the RegistryDataReferencesClient.GetBlobReferenceSAS
//     method.
func (client *RegistryDataReferencesClient) GetBlobReferenceSAS(ctx context.Context, resourceGroupName string, registryName string, name string, version string, body GetBlobReferenceSASRequestDto, options *RegistryDataReferencesClientGetBlobReferenceSASOptions) (RegistryDataReferencesClientGetBlobReferenceSASResponse, error) {
	var err error
	const operationName = "RegistryDataReferencesClient.GetBlobReferenceSAS"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBlobReferenceSASCreateRequest(ctx, resourceGroupName, registryName, name, version, body, options)
	if err != nil {
		return RegistryDataReferencesClientGetBlobReferenceSASResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryDataReferencesClientGetBlobReferenceSASResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegistryDataReferencesClientGetBlobReferenceSASResponse{}, err
	}
	resp, err := client.getBlobReferenceSASHandleResponse(httpResp)
	return resp, err
}

// getBlobReferenceSASCreateRequest creates the GetBlobReferenceSAS request.
func (client *RegistryDataReferencesClient) getBlobReferenceSASCreateRequest(ctx context.Context, resourceGroupName string, registryName string, name string, version string, body GetBlobReferenceSASRequestDto, options *RegistryDataReferencesClientGetBlobReferenceSASOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/datareferences/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getBlobReferenceSASHandleResponse handles the GetBlobReferenceSAS response.
func (client *RegistryDataReferencesClient) getBlobReferenceSASHandleResponse(resp *http.Response) (RegistryDataReferencesClientGetBlobReferenceSASResponse, error) {
	result := RegistryDataReferencesClientGetBlobReferenceSASResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetBlobReferenceSASResponseDto); err != nil {
		return RegistryDataReferencesClientGetBlobReferenceSASResponse{}, err
	}
	return result, nil
}
