//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoleAssignmentsClient contains the methods for the BillingRoleAssignments group.
// Don't use this type directly, use NewRoleAssignmentsClient() instead.
type RoleAssignmentsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleAssignmentsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RoleAssignmentsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// DeleteByBillingAccount - Deletes a role assignment for the caller on a billing account. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingRoleAssignmentName - The ID that uniquely identifies a role assignment.
// options - RoleAssignmentsClientDeleteByBillingAccountOptions contains the optional parameters for the RoleAssignmentsClient.DeleteByBillingAccount
// method.
func (client *RoleAssignmentsClient) DeleteByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string, options *RoleAssignmentsClientDeleteByBillingAccountOptions) (RoleAssignmentsClientDeleteByBillingAccountResponse, error) {
	req, err := client.deleteByBillingAccountCreateRequest(ctx, billingAccountName, billingRoleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientDeleteByBillingAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientDeleteByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientDeleteByBillingAccountResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteByBillingAccountHandleResponse(resp)
}

// deleteByBillingAccountCreateRequest creates the DeleteByBillingAccount request.
func (client *RoleAssignmentsClient) deleteByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, billingRoleAssignmentName string, options *RoleAssignmentsClientDeleteByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments/{billingRoleAssignmentName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingRoleAssignmentName == "" {
		return nil, errors.New("parameter billingRoleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleAssignmentName}", url.PathEscape(billingRoleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteByBillingAccountHandleResponse handles the DeleteByBillingAccount response.
func (client *RoleAssignmentsClient) deleteByBillingAccountHandleResponse(resp *http.Response) (RoleAssignmentsClientDeleteByBillingAccountResponse, error) {
	result := RoleAssignmentsClientDeleteByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientDeleteByBillingAccountResponse{}, err
	}
	return result, nil
}

// DeleteByBillingProfile - Deletes a role assignment for the caller on a billing profile. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// billingRoleAssignmentName - The ID that uniquely identifies a role assignment.
// options - RoleAssignmentsClientDeleteByBillingProfileOptions contains the optional parameters for the RoleAssignmentsClient.DeleteByBillingProfile
// method.
func (client *RoleAssignmentsClient) DeleteByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string, options *RoleAssignmentsClientDeleteByBillingProfileOptions) (RoleAssignmentsClientDeleteByBillingProfileResponse, error) {
	req, err := client.deleteByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, billingRoleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientDeleteByBillingProfileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientDeleteByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientDeleteByBillingProfileResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteByBillingProfileHandleResponse(resp)
}

// deleteByBillingProfileCreateRequest creates the DeleteByBillingProfile request.
func (client *RoleAssignmentsClient) deleteByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string, options *RoleAssignmentsClientDeleteByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments/{billingRoleAssignmentName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if billingRoleAssignmentName == "" {
		return nil, errors.New("parameter billingRoleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleAssignmentName}", url.PathEscape(billingRoleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteByBillingProfileHandleResponse handles the DeleteByBillingProfile response.
func (client *RoleAssignmentsClient) deleteByBillingProfileHandleResponse(resp *http.Response) (RoleAssignmentsClientDeleteByBillingProfileResponse, error) {
	result := RoleAssignmentsClientDeleteByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientDeleteByBillingProfileResponse{}, err
	}
	return result, nil
}

// DeleteByInvoiceSection - Deletes a role assignment for the caller on an invoice section. The operation is supported for
// billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// invoiceSectionName - The ID that uniquely identifies an invoice section.
// billingRoleAssignmentName - The ID that uniquely identifies a role assignment.
// options - RoleAssignmentsClientDeleteByInvoiceSectionOptions contains the optional parameters for the RoleAssignmentsClient.DeleteByInvoiceSection
// method.
func (client *RoleAssignmentsClient) DeleteByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string, options *RoleAssignmentsClientDeleteByInvoiceSectionOptions) (RoleAssignmentsClientDeleteByInvoiceSectionResponse, error) {
	req, err := client.deleteByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientDeleteByInvoiceSectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientDeleteByInvoiceSectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientDeleteByInvoiceSectionResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteByInvoiceSectionHandleResponse(resp)
}

// deleteByInvoiceSectionCreateRequest creates the DeleteByInvoiceSection request.
func (client *RoleAssignmentsClient) deleteByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string, options *RoleAssignmentsClientDeleteByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments/{billingRoleAssignmentName}"
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
	if billingRoleAssignmentName == "" {
		return nil, errors.New("parameter billingRoleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleAssignmentName}", url.PathEscape(billingRoleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteByInvoiceSectionHandleResponse handles the DeleteByInvoiceSection response.
func (client *RoleAssignmentsClient) deleteByInvoiceSectionHandleResponse(resp *http.Response) (RoleAssignmentsClientDeleteByInvoiceSectionResponse, error) {
	result := RoleAssignmentsClientDeleteByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientDeleteByInvoiceSectionResponse{}, err
	}
	return result, nil
}

// GetByBillingAccount - Gets a role assignment for the caller on a billing account. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingRoleAssignmentName - The ID that uniquely identifies a role assignment.
// options - RoleAssignmentsClientGetByBillingAccountOptions contains the optional parameters for the RoleAssignmentsClient.GetByBillingAccount
// method.
func (client *RoleAssignmentsClient) GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string, options *RoleAssignmentsClientGetByBillingAccountOptions) (RoleAssignmentsClientGetByBillingAccountResponse, error) {
	req, err := client.getByBillingAccountCreateRequest(ctx, billingAccountName, billingRoleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientGetByBillingAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientGetByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientGetByBillingAccountResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByBillingAccountHandleResponse(resp)
}

// getByBillingAccountCreateRequest creates the GetByBillingAccount request.
func (client *RoleAssignmentsClient) getByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, billingRoleAssignmentName string, options *RoleAssignmentsClientGetByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments/{billingRoleAssignmentName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingRoleAssignmentName == "" {
		return nil, errors.New("parameter billingRoleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleAssignmentName}", url.PathEscape(billingRoleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingAccountHandleResponse handles the GetByBillingAccount response.
func (client *RoleAssignmentsClient) getByBillingAccountHandleResponse(resp *http.Response) (RoleAssignmentsClientGetByBillingAccountResponse, error) {
	result := RoleAssignmentsClientGetByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientGetByBillingAccountResponse{}, err
	}
	return result, nil
}

// GetByBillingProfile - Gets a role assignment for the caller on a billing profile. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// billingRoleAssignmentName - The ID that uniquely identifies a role assignment.
// options - RoleAssignmentsClientGetByBillingProfileOptions contains the optional parameters for the RoleAssignmentsClient.GetByBillingProfile
// method.
func (client *RoleAssignmentsClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string, options *RoleAssignmentsClientGetByBillingProfileOptions) (RoleAssignmentsClientGetByBillingProfileResponse, error) {
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, billingRoleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientGetByBillingProfileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientGetByBillingProfileResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByBillingProfileHandleResponse(resp)
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *RoleAssignmentsClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string, options *RoleAssignmentsClientGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments/{billingRoleAssignmentName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if billingRoleAssignmentName == "" {
		return nil, errors.New("parameter billingRoleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleAssignmentName}", url.PathEscape(billingRoleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *RoleAssignmentsClient) getByBillingProfileHandleResponse(resp *http.Response) (RoleAssignmentsClientGetByBillingProfileResponse, error) {
	result := RoleAssignmentsClientGetByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientGetByBillingProfileResponse{}, err
	}
	return result, nil
}

// GetByInvoiceSection - Gets a role assignment for the caller on an invoice section. The operation is supported for billing
// accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// invoiceSectionName - The ID that uniquely identifies an invoice section.
// billingRoleAssignmentName - The ID that uniquely identifies a role assignment.
// options - RoleAssignmentsClientGetByInvoiceSectionOptions contains the optional parameters for the RoleAssignmentsClient.GetByInvoiceSection
// method.
func (client *RoleAssignmentsClient) GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string, options *RoleAssignmentsClientGetByInvoiceSectionOptions) (RoleAssignmentsClientGetByInvoiceSectionResponse, error) {
	req, err := client.getByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientGetByInvoiceSectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientGetByInvoiceSectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientGetByInvoiceSectionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByInvoiceSectionHandleResponse(resp)
}

// getByInvoiceSectionCreateRequest creates the GetByInvoiceSection request.
func (client *RoleAssignmentsClient) getByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string, options *RoleAssignmentsClientGetByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments/{billingRoleAssignmentName}"
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
	if billingRoleAssignmentName == "" {
		return nil, errors.New("parameter billingRoleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleAssignmentName}", url.PathEscape(billingRoleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByInvoiceSectionHandleResponse handles the GetByInvoiceSection response.
func (client *RoleAssignmentsClient) getByInvoiceSectionHandleResponse(resp *http.Response) (RoleAssignmentsClientGetByInvoiceSectionResponse, error) {
	result := RoleAssignmentsClientGetByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientGetByInvoiceSectionResponse{}, err
	}
	return result, nil
}

// NewListByBillingAccountPager - Lists the role assignments for the caller on a billing account. The operation is supported
// for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// options - RoleAssignmentsClientListByBillingAccountOptions contains the optional parameters for the RoleAssignmentsClient.ListByBillingAccount
// method.
func (client *RoleAssignmentsClient) NewListByBillingAccountPager(billingAccountName string, options *RoleAssignmentsClientListByBillingAccountOptions) *runtime.Pager[RoleAssignmentsClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListByBillingAccountResponse]{
		More: func(page RoleAssignmentsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListByBillingAccountResponse) (RoleAssignmentsClientListByBillingAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListByBillingAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleAssignmentsClientListByBillingAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListByBillingAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *RoleAssignmentsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *RoleAssignmentsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *RoleAssignmentsClient) listByBillingAccountHandleResponse(resp *http.Response) (RoleAssignmentsClientListByBillingAccountResponse, error) {
	result := RoleAssignmentsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the role assignments for the caller on a billing profile. The operation is supported
// for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// options - RoleAssignmentsClientListByBillingProfileOptions contains the optional parameters for the RoleAssignmentsClient.ListByBillingProfile
// method.
func (client *RoleAssignmentsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *RoleAssignmentsClientListByBillingProfileOptions) *runtime.Pager[RoleAssignmentsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListByBillingProfileResponse]{
		More: func(page RoleAssignmentsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListByBillingProfileResponse) (RoleAssignmentsClientListByBillingProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListByBillingProfileResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleAssignmentsClientListByBillingProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListByBillingProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *RoleAssignmentsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *RoleAssignmentsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *RoleAssignmentsClient) listByBillingProfileHandleResponse(resp *http.Response) (RoleAssignmentsClientListByBillingProfileResponse, error) {
	result := RoleAssignmentsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// NewListByInvoiceSectionPager - Lists the role assignments for the caller on an invoice section. The operation is supported
// for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// invoiceSectionName - The ID that uniquely identifies an invoice section.
// options - RoleAssignmentsClientListByInvoiceSectionOptions contains the optional parameters for the RoleAssignmentsClient.ListByInvoiceSection
// method.
func (client *RoleAssignmentsClient) NewListByInvoiceSectionPager(billingAccountName string, billingProfileName string, invoiceSectionName string, options *RoleAssignmentsClientListByInvoiceSectionOptions) *runtime.Pager[RoleAssignmentsClientListByInvoiceSectionResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListByInvoiceSectionResponse]{
		More: func(page RoleAssignmentsClientListByInvoiceSectionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListByInvoiceSectionResponse) (RoleAssignmentsClientListByInvoiceSectionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListByInvoiceSectionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleAssignmentsClientListByInvoiceSectionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListByInvoiceSectionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInvoiceSectionHandleResponse(resp)
		},
	})
}

// listByInvoiceSectionCreateRequest creates the ListByInvoiceSection request.
func (client *RoleAssignmentsClient) listByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *RoleAssignmentsClientListByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInvoiceSectionHandleResponse handles the ListByInvoiceSection response.
func (client *RoleAssignmentsClient) listByInvoiceSectionHandleResponse(resp *http.Response) (RoleAssignmentsClientListByInvoiceSectionResponse, error) {
	result := RoleAssignmentsClientListByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListByInvoiceSectionResponse{}, err
	}
	return result, nil
}
