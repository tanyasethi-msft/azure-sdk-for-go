//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
	"net/http"
)

// DomainRegistrationProviderServer is a fake server for instances of the armappservice.DomainRegistrationProviderClient type.
type DomainRegistrationProviderServer struct {
	// NewListOperationsPager is the fake for method DomainRegistrationProviderClient.NewListOperationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListOperationsPager func(options *armappservice.DomainRegistrationProviderClientListOperationsOptions) (resp azfake.PagerResponder[armappservice.DomainRegistrationProviderClientListOperationsResponse])
}

// NewDomainRegistrationProviderServerTransport creates a new instance of DomainRegistrationProviderServerTransport with the provided implementation.
// The returned DomainRegistrationProviderServerTransport instance is connected to an instance of armappservice.DomainRegistrationProviderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDomainRegistrationProviderServerTransport(srv *DomainRegistrationProviderServer) *DomainRegistrationProviderServerTransport {
	return &DomainRegistrationProviderServerTransport{
		srv:                    srv,
		newListOperationsPager: newTracker[azfake.PagerResponder[armappservice.DomainRegistrationProviderClientListOperationsResponse]](),
	}
}

// DomainRegistrationProviderServerTransport connects instances of armappservice.DomainRegistrationProviderClient to instances of DomainRegistrationProviderServer.
// Don't use this type directly, use NewDomainRegistrationProviderServerTransport instead.
type DomainRegistrationProviderServerTransport struct {
	srv                    *DomainRegistrationProviderServer
	newListOperationsPager *tracker[azfake.PagerResponder[armappservice.DomainRegistrationProviderClientListOperationsResponse]]
}

// Do implements the policy.Transporter interface for DomainRegistrationProviderServerTransport.
func (d *DomainRegistrationProviderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DomainRegistrationProviderClient.NewListOperationsPager":
		resp, err = d.dispatchNewListOperationsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DomainRegistrationProviderServerTransport) dispatchNewListOperationsPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListOperationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListOperationsPager not implemented")}
	}
	newListOperationsPager := d.newListOperationsPager.get(req)
	if newListOperationsPager == nil {
		resp := d.srv.NewListOperationsPager(nil)
		newListOperationsPager = &resp
		d.newListOperationsPager.add(req, newListOperationsPager)
		server.PagerResponderInjectNextLinks(newListOperationsPager, req, func(page *armappservice.DomainRegistrationProviderClientListOperationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListOperationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListOperationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListOperationsPager) {
		d.newListOperationsPager.remove(req)
	}
	return resp, nil
}
