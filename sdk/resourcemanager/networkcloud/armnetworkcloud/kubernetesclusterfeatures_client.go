// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// KubernetesClusterFeaturesClient contains the methods for the KubernetesClusterFeatures group.
// Don't use this type directly, use NewKubernetesClusterFeaturesClient() instead.
type KubernetesClusterFeaturesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKubernetesClusterFeaturesClient creates a new instance of KubernetesClusterFeaturesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKubernetesClusterFeaturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KubernetesClusterFeaturesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KubernetesClusterFeaturesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new Kubernetes cluster feature or update properties of the Kubernetes cluster feature if
// it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - featureName - The name of the feature.
//   - kubernetesClusterFeatureParameters - The request body.
//   - options - KubernetesClusterFeaturesClientBeginCreateOrUpdateOptions contains the optional parameters for the KubernetesClusterFeaturesClient.BeginCreateOrUpdate
//     method.
func (client *KubernetesClusterFeaturesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureParameters KubernetesClusterFeature, options *KubernetesClusterFeaturesClientBeginCreateOrUpdateOptions) (*runtime.Poller[KubernetesClusterFeaturesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, kubernetesClusterName, featureName, kubernetesClusterFeatureParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClusterFeaturesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClusterFeaturesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new Kubernetes cluster feature or update properties of the Kubernetes cluster feature if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
func (client *KubernetesClusterFeaturesClient) createOrUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureParameters KubernetesClusterFeature, options *KubernetesClusterFeaturesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClusterFeaturesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, kubernetesClusterName, featureName, kubernetesClusterFeatureParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *KubernetesClusterFeaturesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureParameters KubernetesClusterFeature, options *KubernetesClusterFeaturesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/features/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if err := runtime.MarshalAsJSON(req, kubernetesClusterFeatureParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided Kubernetes cluster feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - featureName - The name of the feature.
//   - options - KubernetesClusterFeaturesClientBeginDeleteOptions contains the optional parameters for the KubernetesClusterFeaturesClient.BeginDelete
//     method.
func (client *KubernetesClusterFeaturesClient) BeginDelete(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, options *KubernetesClusterFeaturesClientBeginDeleteOptions) (*runtime.Poller[KubernetesClusterFeaturesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, kubernetesClusterName, featureName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClusterFeaturesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClusterFeaturesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the provided Kubernetes cluster feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
func (client *KubernetesClusterFeaturesClient) deleteOperation(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, options *KubernetesClusterFeaturesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClusterFeaturesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, kubernetesClusterName, featureName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KubernetesClusterFeaturesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, options *KubernetesClusterFeaturesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/features/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// Get - Get properties of the provided the Kubernetes cluster feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - featureName - The name of the feature.
//   - options - KubernetesClusterFeaturesClientGetOptions contains the optional parameters for the KubernetesClusterFeaturesClient.Get
//     method.
func (client *KubernetesClusterFeaturesClient) Get(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, options *KubernetesClusterFeaturesClientGetOptions) (KubernetesClusterFeaturesClientGetResponse, error) {
	var err error
	const operationName = "KubernetesClusterFeaturesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, kubernetesClusterName, featureName, options)
	if err != nil {
		return KubernetesClusterFeaturesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KubernetesClusterFeaturesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return KubernetesClusterFeaturesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *KubernetesClusterFeaturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, _ *KubernetesClusterFeaturesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/features/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KubernetesClusterFeaturesClient) getHandleResponse(resp *http.Response) (KubernetesClusterFeaturesClientGetResponse, error) {
	result := KubernetesClusterFeaturesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubernetesClusterFeature); err != nil {
		return KubernetesClusterFeaturesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByKubernetesClusterPager - Get a list of features for the provided Kubernetes cluster.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - options - KubernetesClusterFeaturesClientListByKubernetesClusterOptions contains the optional parameters for the KubernetesClusterFeaturesClient.NewListByKubernetesClusterPager
//     method.
func (client *KubernetesClusterFeaturesClient) NewListByKubernetesClusterPager(resourceGroupName string, kubernetesClusterName string, options *KubernetesClusterFeaturesClientListByKubernetesClusterOptions) *runtime.Pager[KubernetesClusterFeaturesClientListByKubernetesClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[KubernetesClusterFeaturesClientListByKubernetesClusterResponse]{
		More: func(page KubernetesClusterFeaturesClientListByKubernetesClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KubernetesClusterFeaturesClientListByKubernetesClusterResponse) (KubernetesClusterFeaturesClientListByKubernetesClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KubernetesClusterFeaturesClient.NewListByKubernetesClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByKubernetesClusterCreateRequest(ctx, resourceGroupName, kubernetesClusterName, options)
			}, nil)
			if err != nil {
				return KubernetesClusterFeaturesClientListByKubernetesClusterResponse{}, err
			}
			return client.listByKubernetesClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByKubernetesClusterCreateRequest creates the ListByKubernetesCluster request.
func (client *KubernetesClusterFeaturesClient) listByKubernetesClusterCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, _ *KubernetesClusterFeaturesClientListByKubernetesClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/features"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByKubernetesClusterHandleResponse handles the ListByKubernetesCluster response.
func (client *KubernetesClusterFeaturesClient) listByKubernetesClusterHandleResponse(resp *http.Response) (KubernetesClusterFeaturesClientListByKubernetesClusterResponse, error) {
	result := KubernetesClusterFeaturesClientListByKubernetesClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubernetesClusterFeatureList); err != nil {
		return KubernetesClusterFeaturesClientListByKubernetesClusterResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch properties of the provided Kubernetes cluster feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - featureName - The name of the feature.
//   - kubernetesClusterFeatureUpdateParameters - The request body.
//   - options - KubernetesClusterFeaturesClientBeginUpdateOptions contains the optional parameters for the KubernetesClusterFeaturesClient.BeginUpdate
//     method.
func (client *KubernetesClusterFeaturesClient) BeginUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureUpdateParameters KubernetesClusterFeaturePatchParameters, options *KubernetesClusterFeaturesClientBeginUpdateOptions) (*runtime.Poller[KubernetesClusterFeaturesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, kubernetesClusterName, featureName, kubernetesClusterFeatureUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClusterFeaturesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClusterFeaturesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch properties of the provided Kubernetes cluster feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
func (client *KubernetesClusterFeaturesClient) update(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureUpdateParameters KubernetesClusterFeaturePatchParameters, options *KubernetesClusterFeaturesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClusterFeaturesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, kubernetesClusterName, featureName, kubernetesClusterFeatureUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *KubernetesClusterFeaturesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, featureName string, kubernetesClusterFeatureUpdateParameters KubernetesClusterFeaturePatchParameters, options *KubernetesClusterFeaturesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/features/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if err := runtime.MarshalAsJSON(req, kubernetesClusterFeatureUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
