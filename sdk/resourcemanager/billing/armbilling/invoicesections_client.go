//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// InvoiceSectionsClient contains the methods for the InvoiceSections group.
// Don't use this type directly, use NewInvoiceSectionsClient() instead.
type InvoiceSectionsClient struct {
	internal *arm.Client
}

// NewInvoiceSectionsClient creates a new instance of InvoiceSectionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInvoiceSectionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*InvoiceSectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InvoiceSectionsClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an invoice section. The operation is supported only for billing accounts with
// agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - parameters - An invoice section.
//   - options - InvoiceSectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the InvoiceSectionsClient.BeginCreateOrUpdate
//     method.
func (client *InvoiceSectionsClient) BeginCreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection, options *InvoiceSectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[InvoiceSectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InvoiceSectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InvoiceSectionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an invoice section. The operation is supported only for billing accounts with agreement
// type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *InvoiceSectionsClient) createOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection, options *InvoiceSectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InvoiceSectionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InvoiceSectionsClient) createOrUpdateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection, options *InvoiceSectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an invoice section. The operation is supported for billing accounts with agreement type Microsoft
// Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - options - InvoiceSectionsClientBeginDeleteOptions contains the optional parameters for the InvoiceSectionsClient.BeginDelete
//     method.
func (client *InvoiceSectionsClient) BeginDelete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientBeginDeleteOptions) (*runtime.Poller[InvoiceSectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InvoiceSectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InvoiceSectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an invoice section. The operation is supported for billing accounts with agreement type Microsoft Customer
// Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *InvoiceSectionsClient) deleteOperation(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "InvoiceSectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
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
func (client *InvoiceSectionsClient) deleteCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an invoice section by its ID. The operation is supported only for billing accounts with agreement type Microsoft
// Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - options - InvoiceSectionsClientGetOptions contains the optional parameters for the InvoiceSectionsClient.Get method.
func (client *InvoiceSectionsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientGetOptions) (InvoiceSectionsClientGetResponse, error) {
	var err error
	const operationName = "InvoiceSectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
	if err != nil {
		return InvoiceSectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InvoiceSectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InvoiceSectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InvoiceSectionsClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InvoiceSectionsClient) getHandleResponse(resp *http.Response) (InvoiceSectionsClientGetResponse, error) {
	result := InvoiceSectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceSection); err != nil {
		return InvoiceSectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the invoice sections that a user has access to. The operation is supported only for
// billing accounts with agreement type Microsoft Customer Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - InvoiceSectionsClientListByBillingProfileOptions contains the optional parameters for the InvoiceSectionsClient.NewListByBillingProfilePager
//     method.
func (client *InvoiceSectionsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *InvoiceSectionsClientListByBillingProfileOptions) *runtime.Pager[InvoiceSectionsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[InvoiceSectionsClientListByBillingProfileResponse]{
		More: func(page InvoiceSectionsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InvoiceSectionsClientListByBillingProfileResponse) (InvoiceSectionsClientListByBillingProfileResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InvoiceSectionsClient.NewListByBillingProfilePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			}, nil)
			if err != nil {
				return InvoiceSectionsClientListByBillingProfileResponse{}, err
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *InvoiceSectionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *InvoiceSectionsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.IncludeDeleted != nil {
		reqQP.Set("includeDeleted", strconv.FormatBool(*options.IncludeDeleted))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", *options.OrderBy)
	}
	if options != nil && options.Search != nil {
		reqQP.Set("search", *options.Search)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(*options.Top, 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *InvoiceSectionsClient) listByBillingProfileHandleResponse(resp *http.Response) (InvoiceSectionsClientListByBillingProfileResponse, error) {
	result := InvoiceSectionsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceSectionListResult); err != nil {
		return InvoiceSectionsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// ValidateDeleteEligibility - Validates if the invoice section can be deleted. The operation is supported for billing accounts
// with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - options - InvoiceSectionsClientValidateDeleteEligibilityOptions contains the optional parameters for the InvoiceSectionsClient.ValidateDeleteEligibility
//     method.
func (client *InvoiceSectionsClient) ValidateDeleteEligibility(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientValidateDeleteEligibilityOptions) (InvoiceSectionsClientValidateDeleteEligibilityResponse, error) {
	var err error
	const operationName = "InvoiceSectionsClient.ValidateDeleteEligibility"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateDeleteEligibilityCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
	if err != nil {
		return InvoiceSectionsClientValidateDeleteEligibilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InvoiceSectionsClientValidateDeleteEligibilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InvoiceSectionsClientValidateDeleteEligibilityResponse{}, err
	}
	resp, err := client.validateDeleteEligibilityHandleResponse(httpResp)
	return resp, err
}

// validateDeleteEligibilityCreateRequest creates the ValidateDeleteEligibility request.
func (client *InvoiceSectionsClient) validateDeleteEligibilityCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *InvoiceSectionsClientValidateDeleteEligibilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/validateDeleteEligibility"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// validateDeleteEligibilityHandleResponse handles the ValidateDeleteEligibility response.
func (client *InvoiceSectionsClient) validateDeleteEligibilityHandleResponse(resp *http.Response) (InvoiceSectionsClientValidateDeleteEligibilityResponse, error) {
	result := InvoiceSectionsClientValidateDeleteEligibilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeleteInvoiceSectionEligibilityResult); err != nil {
		return InvoiceSectionsClientValidateDeleteEligibilityResponse{}, err
	}
	return result, nil
}
