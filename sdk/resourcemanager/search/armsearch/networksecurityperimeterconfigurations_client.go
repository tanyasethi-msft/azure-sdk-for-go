// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// NetworkSecurityPerimeterConfigurationsClient contains the methods for the NetworkSecurityPerimeterConfigurations group.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsClient() instead.
type NetworkSecurityPerimeterConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a network security perimeter configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - nspConfigName - The network security perimeter configuration name.
//   - options - NetworkSecurityPerimeterConfigurationsClientGetOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.Get
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientGetOptions) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, searchServiceName, nspConfigName, options)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkSecurityPerimeterConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, _ *NetworkSecurityPerimeterConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/networkSecurityPerimeterConfigurations/{nspConfigName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if nspConfigName == "" {
		return nil, errors.New("parameter nspConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nspConfigName}", url.PathEscape(nspConfigName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkSecurityPerimeterConfigurationsClient) getHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfiguration); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Gets a list of network security perimeter configurations for a search service.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - options - NetworkSecurityPerimeterConfigurationsClientListByServiceOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.NewListByServicePager
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) NewListByServicePager(resourceGroupName string, searchServiceName string, options *NetworkSecurityPerimeterConfigurationsClientListByServiceOptions) *runtime.Pager[NetworkSecurityPerimeterConfigurationsClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkSecurityPerimeterConfigurationsClientListByServiceResponse]{
		More: func(page NetworkSecurityPerimeterConfigurationsClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkSecurityPerimeterConfigurationsClientListByServiceResponse) (NetworkSecurityPerimeterConfigurationsClientListByServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkSecurityPerimeterConfigurationsClient.NewListByServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServiceCreateRequest(ctx, resourceGroupName, searchServiceName, options)
			}, nil)
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListByServiceResponse{}, err
			}
			return client.listByServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NetworkSecurityPerimeterConfigurationsClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, _ *NetworkSecurityPerimeterConfigurationsClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/networkSecurityPerimeterConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *NetworkSecurityPerimeterConfigurationsClient) listByServiceHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientListByServiceResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationListResult); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientListByServiceResponse{}, err
	}
	return result, nil
}

// BeginReconcile - Reconcile network security perimeter configuration for the Azure AI Search resource provider. This triggers
// a manual resync with network security perimeter configurations by ensuring the search
// service carries the latest configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - nspConfigName - The network security perimeter configuration name.
//   - options - NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.BeginReconcile
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) BeginReconcile(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*runtime.Poller[NetworkSecurityPerimeterConfigurationsClientReconcileResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reconcile(ctx, resourceGroupName, searchServiceName, nspConfigName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reconcile - Reconcile network security perimeter configuration for the Azure AI Search resource provider. This triggers
// a manual resync with network security perimeter configurations by ensuring the search
// service carries the latest configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcile(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.BeginReconcile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reconcileCreateRequest(ctx, resourceGroupName, searchServiceName, nspConfigName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// reconcileCreateRequest creates the Reconcile request.
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcileCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, _ *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/networkSecurityPerimeterConfigurations/{nspConfigName}/reconcile"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if nspConfigName == "" {
		return nil, errors.New("parameter nspConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nspConfigName}", url.PathEscape(nspConfigName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
