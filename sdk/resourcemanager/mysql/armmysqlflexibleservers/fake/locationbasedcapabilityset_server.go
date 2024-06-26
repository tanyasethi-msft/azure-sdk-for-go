//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
	"net/http"
	"net/url"
	"regexp"
)

// LocationBasedCapabilitySetServer is a fake server for instances of the armmysqlflexibleservers.LocationBasedCapabilitySetClient type.
type LocationBasedCapabilitySetServer struct {
	// Get is the fake for method LocationBasedCapabilitySetClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, locationName string, capabilitySetName string, options *armmysqlflexibleservers.LocationBasedCapabilitySetClientGetOptions) (resp azfake.Responder[armmysqlflexibleservers.LocationBasedCapabilitySetClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LocationBasedCapabilitySetClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(locationName string, options *armmysqlflexibleservers.LocationBasedCapabilitySetClientListOptions) (resp azfake.PagerResponder[armmysqlflexibleservers.LocationBasedCapabilitySetClientListResponse])
}

// NewLocationBasedCapabilitySetServerTransport creates a new instance of LocationBasedCapabilitySetServerTransport with the provided implementation.
// The returned LocationBasedCapabilitySetServerTransport instance is connected to an instance of armmysqlflexibleservers.LocationBasedCapabilitySetClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationBasedCapabilitySetServerTransport(srv *LocationBasedCapabilitySetServer) *LocationBasedCapabilitySetServerTransport {
	return &LocationBasedCapabilitySetServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmysqlflexibleservers.LocationBasedCapabilitySetClientListResponse]](),
	}
}

// LocationBasedCapabilitySetServerTransport connects instances of armmysqlflexibleservers.LocationBasedCapabilitySetClient to instances of LocationBasedCapabilitySetServer.
// Don't use this type directly, use NewLocationBasedCapabilitySetServerTransport instead.
type LocationBasedCapabilitySetServerTransport struct {
	srv          *LocationBasedCapabilitySetServer
	newListPager *tracker[azfake.PagerResponder[armmysqlflexibleservers.LocationBasedCapabilitySetClientListResponse]]
}

// Do implements the policy.Transporter interface for LocationBasedCapabilitySetServerTransport.
func (l *LocationBasedCapabilitySetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocationBasedCapabilitySetClient.Get":
		resp, err = l.dispatchGet(req)
	case "LocationBasedCapabilitySetClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocationBasedCapabilitySetServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capabilitySets/(?P<capabilitySetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	capabilitySetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capabilitySetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), locationNameParam, capabilitySetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Capability, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocationBasedCapabilitySetServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capabilitySets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(locationNameParam, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmysqlflexibleservers.LocationBasedCapabilitySetClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		l.newListPager.remove(req)
	}
	return resp, nil
}
