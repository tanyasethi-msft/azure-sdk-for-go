// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armpurestorageblock

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

// AvsVMVolumesClient contains the methods for the AvsVMVolumes group.
// Don't use this type directly, use NewAvsVMVolumesClient() instead.
type AvsVMVolumesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvsVMVolumesClient creates a new instance of AvsVMVolumesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvsVMVolumesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvsVMVolumesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvsVMVolumesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Delete a volume in an AVS VM
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storagePoolName - Name of the storage pool
//   - avsVMID - ID of the AVS VM
//   - volumeID - ID of the volume in the AVS VM
//   - options - AvsVMVolumesClientBeginDeleteOptions contains the optional parameters for the AvsVMVolumesClient.BeginDelete
//     method.
func (client *AvsVMVolumesClient) BeginDelete(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, options *AvsVMVolumesClientBeginDeleteOptions) (*runtime.Poller[AvsVMVolumesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storagePoolName, avsVMID, volumeID, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AvsVMVolumesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AvsVMVolumesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a volume in an AVS VM
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AvsVMVolumesClient) deleteOperation(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, options *AvsVMVolumesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AvsVMVolumesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storagePoolName, avsVMID, volumeID, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AvsVMVolumesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, _ *AvsVMVolumesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PureStorage.Block/storagePools/{storagePoolName}/avsVms/{avsVmId}/avsVmVolumes/{volumeId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storagePoolName == "" {
		return nil, errors.New("parameter storagePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storagePoolName}", url.PathEscape(storagePoolName))
	if avsVMID == "" {
		return nil, errors.New("parameter avsVMID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{avsVmId}", url.PathEscape(avsVMID))
	if volumeID == "" {
		return nil, errors.New("parameter volumeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeId}", url.PathEscape(volumeID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a volume in an AVS VM
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storagePoolName - Name of the storage pool
//   - avsVMID - ID of the AVS VM
//   - volumeID - ID of the volume in the AVS VM
//   - options - AvsVMVolumesClientGetOptions contains the optional parameters for the AvsVMVolumesClient.Get method.
func (client *AvsVMVolumesClient) Get(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, options *AvsVMVolumesClientGetOptions) (AvsVMVolumesClientGetResponse, error) {
	var err error
	const operationName = "AvsVMVolumesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, storagePoolName, avsVMID, volumeID, options)
	if err != nil {
		return AvsVMVolumesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvsVMVolumesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AvsVMVolumesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AvsVMVolumesClient) getCreateRequest(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, _ *AvsVMVolumesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PureStorage.Block/storagePools/{storagePoolName}/avsVms/{avsVmId}/avsVmVolumes/{volumeId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storagePoolName == "" {
		return nil, errors.New("parameter storagePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storagePoolName}", url.PathEscape(storagePoolName))
	if avsVMID == "" {
		return nil, errors.New("parameter avsVMID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{avsVmId}", url.PathEscape(avsVMID))
	if volumeID == "" {
		return nil, errors.New("parameter volumeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeId}", url.PathEscape(volumeID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AvsVMVolumesClient) getHandleResponse(resp *http.Response) (AvsVMVolumesClientGetResponse, error) {
	result := AvsVMVolumesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvsVMVolume); err != nil {
		return AvsVMVolumesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAvsVMPager - List volumes in an AVS VM
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storagePoolName - Name of the storage pool
//   - avsVMID - ID of the AVS VM
//   - options - AvsVMVolumesClientListByAvsVMOptions contains the optional parameters for the AvsVMVolumesClient.NewListByAvsVMPager
//     method.
func (client *AvsVMVolumesClient) NewListByAvsVMPager(resourceGroupName string, storagePoolName string, avsVMID string, options *AvsVMVolumesClientListByAvsVMOptions) *runtime.Pager[AvsVMVolumesClientListByAvsVMResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvsVMVolumesClientListByAvsVMResponse]{
		More: func(page AvsVMVolumesClientListByAvsVMResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvsVMVolumesClientListByAvsVMResponse) (AvsVMVolumesClientListByAvsVMResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvsVMVolumesClient.NewListByAvsVMPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAvsVMCreateRequest(ctx, resourceGroupName, storagePoolName, avsVMID, options)
			}, nil)
			if err != nil {
				return AvsVMVolumesClientListByAvsVMResponse{}, err
			}
			return client.listByAvsVMHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAvsVMCreateRequest creates the ListByAvsVM request.
func (client *AvsVMVolumesClient) listByAvsVMCreateRequest(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, _ *AvsVMVolumesClientListByAvsVMOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PureStorage.Block/storagePools/{storagePoolName}/avsVms/{avsVmId}/avsVmVolumes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storagePoolName == "" {
		return nil, errors.New("parameter storagePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storagePoolName}", url.PathEscape(storagePoolName))
	if avsVMID == "" {
		return nil, errors.New("parameter avsVMID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{avsVmId}", url.PathEscape(avsVMID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAvsVMHandleResponse handles the ListByAvsVM response.
func (client *AvsVMVolumesClient) listByAvsVMHandleResponse(resp *http.Response) (AvsVMVolumesClientListByAvsVMResponse, error) {
	result := AvsVMVolumesClientListByAvsVMResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvsVMVolumeListResult); err != nil {
		return AvsVMVolumesClientListByAvsVMResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a volume in an AVS VM
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storagePoolName - Name of the storage pool
//   - avsVMID - ID of the AVS VM
//   - volumeID - ID of the volume in the AVS VM
//   - properties - The resource properties to be updated.
//   - options - AvsVMVolumesClientBeginUpdateOptions contains the optional parameters for the AvsVMVolumesClient.BeginUpdate
//     method.
func (client *AvsVMVolumesClient) BeginUpdate(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, properties AvsVMVolumeUpdate, options *AvsVMVolumesClientBeginUpdateOptions) (*runtime.Poller[AvsVMVolumesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, storagePoolName, avsVMID, volumeID, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AvsVMVolumesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AvsVMVolumesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a volume in an AVS VM
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AvsVMVolumesClient) update(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, properties AvsVMVolumeUpdate, options *AvsVMVolumesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AvsVMVolumesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storagePoolName, avsVMID, volumeID, properties, options)
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
func (client *AvsVMVolumesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storagePoolName string, avsVMID string, volumeID string, properties AvsVMVolumeUpdate, _ *AvsVMVolumesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PureStorage.Block/storagePools/{storagePoolName}/avsVms/{avsVmId}/avsVmVolumes/{volumeId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storagePoolName == "" {
		return nil, errors.New("parameter storagePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storagePoolName}", url.PathEscape(storagePoolName))
	if avsVMID == "" {
		return nil, errors.New("parameter avsVMID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{avsVmId}", url.PathEscape(avsVMID))
	if volumeID == "" {
		return nil, errors.New("parameter volumeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeId}", url.PathEscape(volumeID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
