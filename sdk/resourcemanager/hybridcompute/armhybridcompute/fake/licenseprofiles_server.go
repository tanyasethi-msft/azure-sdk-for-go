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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
	"net/http"
	"net/url"
	"regexp"
)

// LicenseProfilesServer is a fake server for instances of the armhybridcompute.LicenseProfilesClient type.
type LicenseProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method LicenseProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, machineName string, parameters armhybridcompute.LicenseProfile, options *armhybridcompute.LicenseProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridcompute.LicenseProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LicenseProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, machineName string, options *armhybridcompute.LicenseProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridcompute.LicenseProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LicenseProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, machineName string, options *armhybridcompute.LicenseProfilesClientGetOptions) (resp azfake.Responder[armhybridcompute.LicenseProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LicenseProfilesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, machineName string, options *armhybridcompute.LicenseProfilesClientListOptions) (resp azfake.PagerResponder[armhybridcompute.LicenseProfilesClientListResponse])

	// BeginUpdate is the fake for method LicenseProfilesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, machineName string, parameters armhybridcompute.LicenseProfileUpdate, options *armhybridcompute.LicenseProfilesClientBeginUpdateOptions) (resp azfake.PollerResponder[armhybridcompute.LicenseProfilesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewLicenseProfilesServerTransport creates a new instance of LicenseProfilesServerTransport with the provided implementation.
// The returned LicenseProfilesServerTransport instance is connected to an instance of armhybridcompute.LicenseProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLicenseProfilesServerTransport(srv *LicenseProfilesServer) *LicenseProfilesServerTransport {
	return &LicenseProfilesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armhybridcompute.LicenseProfilesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armhybridcompute.LicenseProfilesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armhybridcompute.LicenseProfilesClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armhybridcompute.LicenseProfilesClientUpdateResponse]](),
	}
}

// LicenseProfilesServerTransport connects instances of armhybridcompute.LicenseProfilesClient to instances of LicenseProfilesServer.
// Don't use this type directly, use NewLicenseProfilesServerTransport instead.
type LicenseProfilesServerTransport struct {
	srv                 *LicenseProfilesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armhybridcompute.LicenseProfilesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armhybridcompute.LicenseProfilesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armhybridcompute.LicenseProfilesClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armhybridcompute.LicenseProfilesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for LicenseProfilesServerTransport.
func (l *LicenseProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LicenseProfilesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if licenseProfilesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = licenseProfilesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LicenseProfilesClient.BeginCreateOrUpdate":
				res.resp, res.err = l.dispatchBeginCreateOrUpdate(req)
			case "LicenseProfilesClient.BeginDelete":
				res.resp, res.err = l.dispatchBeginDelete(req)
			case "LicenseProfilesClient.Get":
				res.resp, res.err = l.dispatchGet(req)
			case "LicenseProfilesClient.NewListPager":
				res.resp, res.err = l.dispatchNewListPager(req)
			case "LicenseProfilesClient.BeginUpdate":
				res.resp, res.err = l.dispatchBeginUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (l *LicenseProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/licenseProfiles/(?P<licenseProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcompute.LicenseProfile](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, machineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LicenseProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/licenseProfiles/(?P<licenseProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameParam, machineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LicenseProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/licenseProfiles/(?P<licenseProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, machineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LicenseProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LicenseProfilesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/licenseProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(resourceGroupNameParam, machineNameParam, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armhybridcompute.LicenseProfilesClientListResponse, createLink func() string) {
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

func (l *LicenseProfilesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/licenseProfiles/(?P<licenseProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcompute.LicenseProfileUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceGroupNameParam, machineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		l.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		l.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to LicenseProfilesServerTransport
var licenseProfilesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
