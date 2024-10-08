//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

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

// ClusterPoolUpgradeHistoriesClient contains the methods for the ClusterPoolUpgradeHistories group.
// Don't use this type directly, use NewClusterPoolUpgradeHistoriesClient() instead.
type ClusterPoolUpgradeHistoriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterPoolUpgradeHistoriesClient creates a new instance of ClusterPoolUpgradeHistoriesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterPoolUpgradeHistoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterPoolUpgradeHistoriesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterPoolUpgradeHistoriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Returns a list of upgrade history.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - options - ClusterPoolUpgradeHistoriesClientListOptions contains the optional parameters for the ClusterPoolUpgradeHistoriesClient.NewListPager
//     method.
func (client *ClusterPoolUpgradeHistoriesClient) NewListPager(resourceGroupName string, clusterPoolName string, options *ClusterPoolUpgradeHistoriesClientListOptions) *runtime.Pager[ClusterPoolUpgradeHistoriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterPoolUpgradeHistoriesClientListResponse]{
		More: func(page ClusterPoolUpgradeHistoriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterPoolUpgradeHistoriesClientListResponse) (ClusterPoolUpgradeHistoriesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ClusterPoolUpgradeHistoriesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, clusterPoolName, options)
			}, nil)
			if err != nil {
				return ClusterPoolUpgradeHistoriesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ClusterPoolUpgradeHistoriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClusterPoolUpgradeHistoriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/upgradeHistories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
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

// listHandleResponse handles the List response.
func (client *ClusterPoolUpgradeHistoriesClient) listHandleResponse(resp *http.Response) (ClusterPoolUpgradeHistoriesClientListResponse, error) {
	result := ClusterPoolUpgradeHistoriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPoolUpgradeHistoryListResult); err != nil {
		return ClusterPoolUpgradeHistoriesClientListResponse{}, err
	}
	return result, nil
}
