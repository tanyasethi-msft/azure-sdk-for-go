// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armchaos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ExperimentsClient contains the methods for the Experiments group.
// Don't use this type directly, use NewExperimentsClient() instead.
type ExperimentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExperimentsClient creates a new instance of ExperimentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExperimentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExperimentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExperimentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - Cancel a running Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientBeginCancelOptions contains the optional parameters for the ExperimentsClient.BeginCancel method.
func (client *ExperimentsClient) BeginCancel(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientBeginCancelOptions) (*runtime.Poller[ExperimentsClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, experimentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExperimentsClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExperimentsClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Cancel - Cancel a running Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ExperimentsClient) cancel(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientBeginCancelOptions) (*http.Response, error) {
	var err error
	const operationName = "ExperimentsClient.BeginCancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, experimentName, options)
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

// cancelCreateRequest creates the Cancel request.
func (client *ExperimentsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, _ *ExperimentsClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateOrUpdate - Create or update a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - resource - Experiment resource to be created or updated.
//   - options - ExperimentsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExperimentsClient.BeginCreateOrUpdate
//     method.
func (client *ExperimentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, experimentName string, resource Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExperimentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, experimentName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExperimentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExperimentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ExperimentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, experimentName string, resource Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExperimentsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, experimentName, resource, options)
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
func (client *ExperimentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, resource Experiment, _ *ExperimentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientBeginDeleteOptions contains the optional parameters for the ExperimentsClient.BeginDelete method.
func (client *ExperimentsClient) BeginDelete(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*runtime.Poller[ExperimentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, experimentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExperimentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExperimentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ExperimentsClient) deleteOperation(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ExperimentsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, experimentName, options)
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
func (client *ExperimentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, _ *ExperimentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ExecutionDetails - Execution details of an experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - executionID - GUID that represents a Experiment execution detail.
//   - options - ExperimentsClientExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.ExecutionDetails
//     method.
func (client *ExperimentsClient) ExecutionDetails(ctx context.Context, resourceGroupName string, experimentName string, executionID string, options *ExperimentsClientExecutionDetailsOptions) (ExperimentsClientExecutionDetailsResponse, error) {
	var err error
	const operationName = "ExperimentsClient.ExecutionDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executionDetailsCreateRequest(ctx, resourceGroupName, experimentName, executionID, options)
	if err != nil {
		return ExperimentsClientExecutionDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientExecutionDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExperimentsClientExecutionDetailsResponse{}, err
	}
	resp, err := client.executionDetailsHandleResponse(httpResp)
	return resp, err
}

// executionDetailsCreateRequest creates the ExecutionDetails request.
func (client *ExperimentsClient) executionDetailsCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, executionID string, _ *ExperimentsClientExecutionDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executions/{executionId}/executionDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	if executionID == "" {
		return nil, errors.New("parameter executionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{executionId}", url.PathEscape(executionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// executionDetailsHandleResponse handles the ExecutionDetails response.
func (client *ExperimentsClient) executionDetailsHandleResponse(resp *http.Response) (ExperimentsClientExecutionDetailsResponse, error) {
	result := ExperimentsClientExecutionDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecutionDetails); err != nil {
		return ExperimentsClientExecutionDetailsResponse{}, err
	}
	return result, nil
}

// Get - Get a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientGetOptions contains the optional parameters for the ExperimentsClient.Get method.
func (client *ExperimentsClient) Get(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientGetOptions) (ExperimentsClientGetResponse, error) {
	var err error
	const operationName = "ExperimentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExperimentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExperimentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, _ *ExperimentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExperimentsClient) getHandleResponse(resp *http.Response) (ExperimentsClientGetResponse, error) {
	result := ExperimentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Experiment); err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	return result, nil
}

// GetExecution - Get an execution of an Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - executionID - GUID that represents a Experiment execution detail.
//   - options - ExperimentsClientGetExecutionOptions contains the optional parameters for the ExperimentsClient.GetExecution
//     method.
func (client *ExperimentsClient) GetExecution(ctx context.Context, resourceGroupName string, experimentName string, executionID string, options *ExperimentsClientGetExecutionOptions) (ExperimentsClientGetExecutionResponse, error) {
	var err error
	const operationName = "ExperimentsClient.GetExecution"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExecutionCreateRequest(ctx, resourceGroupName, experimentName, executionID, options)
	if err != nil {
		return ExperimentsClientGetExecutionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientGetExecutionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExperimentsClientGetExecutionResponse{}, err
	}
	resp, err := client.getExecutionHandleResponse(httpResp)
	return resp, err
}

// getExecutionCreateRequest creates the GetExecution request.
func (client *ExperimentsClient) getExecutionCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, executionID string, _ *ExperimentsClientGetExecutionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executions/{executionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	if executionID == "" {
		return nil, errors.New("parameter executionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{executionId}", url.PathEscape(executionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getExecutionHandleResponse handles the GetExecution response.
func (client *ExperimentsClient) getExecutionHandleResponse(resp *http.Response) (ExperimentsClientGetExecutionResponse, error) {
	result := ExperimentsClientGetExecutionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecution); err != nil {
		return ExperimentsClientGetExecutionResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Experiment resources in a resource group.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ExperimentsClientListOptions contains the optional parameters for the ExperimentsClient.NewListPager method.
func (client *ExperimentsClient) NewListPager(resourceGroupName string, options *ExperimentsClientListOptions) *runtime.Pager[ExperimentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListResponse]{
		More: func(page ExperimentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListResponse) (ExperimentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExperimentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ExperimentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExperimentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ExperimentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.Running != nil {
		reqQP.Set("running", strconv.FormatBool(*options.Running))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExperimentsClient) listHandleResponse(resp *http.Response) (ExperimentsClientListResponse, error) {
	result := ExperimentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentListResult); err != nil {
		return ExperimentsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Get a list of Experiment resources in a subscription.
//
// Generated from API version 2025-01-01
//   - options - ExperimentsClientListAllOptions contains the optional parameters for the ExperimentsClient.NewListAllPager method.
func (client *ExperimentsClient) NewListAllPager(options *ExperimentsClientListAllOptions) *runtime.Pager[ExperimentsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListAllResponse]{
		More: func(page ExperimentsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListAllResponse) (ExperimentsClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExperimentsClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ExperimentsClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *ExperimentsClient) listAllCreateRequest(ctx context.Context, options *ExperimentsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.Running != nil {
		reqQP.Set("running", strconv.FormatBool(*options.Running))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ExperimentsClient) listAllHandleResponse(resp *http.Response) (ExperimentsClientListAllResponse, error) {
	result := ExperimentsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentListResult); err != nil {
		return ExperimentsClientListAllResponse{}, err
	}
	return result, nil
}

// NewListAllExecutionsPager - Get a list of executions of an Experiment resource.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientListAllExecutionsOptions contains the optional parameters for the ExperimentsClient.NewListAllExecutionsPager
//     method.
func (client *ExperimentsClient) NewListAllExecutionsPager(resourceGroupName string, experimentName string, options *ExperimentsClientListAllExecutionsOptions) *runtime.Pager[ExperimentsClientListAllExecutionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListAllExecutionsResponse]{
		More: func(page ExperimentsClientListAllExecutionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListAllExecutionsResponse) (ExperimentsClientListAllExecutionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExperimentsClient.NewListAllExecutionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllExecutionsCreateRequest(ctx, resourceGroupName, experimentName, options)
			}, nil)
			if err != nil {
				return ExperimentsClientListAllExecutionsResponse{}, err
			}
			return client.listAllExecutionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllExecutionsCreateRequest creates the ListAllExecutions request.
func (client *ExperimentsClient) listAllExecutionsCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, _ *ExperimentsClientListAllExecutionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllExecutionsHandleResponse handles the ListAllExecutions response.
func (client *ExperimentsClient) listAllExecutionsHandleResponse(resp *http.Response) (ExperimentsClientListAllExecutionsResponse, error) {
	result := ExperimentsClientListAllExecutionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecutionListResult); err != nil {
		return ExperimentsClientListAllExecutionsResponse{}, err
	}
	return result, nil
}

// BeginStart - Start a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientBeginStartOptions contains the optional parameters for the ExperimentsClient.BeginStart method.
func (client *ExperimentsClient) BeginStart(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientBeginStartOptions) (*runtime.Poller[ExperimentsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, experimentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExperimentsClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExperimentsClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Start a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ExperimentsClient) start(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "ExperimentsClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, experimentName, options)
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

// startCreateRequest creates the Start request.
func (client *ExperimentsClient) startCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, _ *ExperimentsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - The operation to update an experiment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - experimentName - String that represents a Experiment resource name.
//   - properties - Parameters supplied to the Update experiment operation.
//   - options - ExperimentsClientBeginUpdateOptions contains the optional parameters for the ExperimentsClient.BeginUpdate method.
func (client *ExperimentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, experimentName string, properties ExperimentUpdate, options *ExperimentsClientBeginUpdateOptions) (*runtime.Poller[ExperimentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, experimentName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExperimentsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExperimentsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation to update an experiment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ExperimentsClient) update(ctx context.Context, resourceGroupName string, experimentName string, properties ExperimentUpdate, options *ExperimentsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExperimentsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, experimentName, properties, options)
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
func (client *ExperimentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, properties ExperimentUpdate, _ *ExperimentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
