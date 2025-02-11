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
	"net/http"
	"net/url"
	"regexp"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// BuildsServer is a fake server for instances of the armappcontainers.BuildsClient type.
type BuildsServer struct {
	// BeginCreateOrUpdate is the fake for method BuildsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, builderName string, buildName string, buildEnvelope armappcontainers.BuildResource, options *armappcontainers.BuildsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcontainers.BuildsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BuildsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, builderName string, buildName string, options *armappcontainers.BuildsClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcontainers.BuildsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BuildsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, builderName string, buildName string, options *armappcontainers.BuildsClientGetOptions) (resp azfake.Responder[armappcontainers.BuildsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewBuildsServerTransport creates a new instance of BuildsServerTransport with the provided implementation.
// The returned BuildsServerTransport instance is connected to an instance of armappcontainers.BuildsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBuildsServerTransport(srv *BuildsServer) *BuildsServerTransport {
	return &BuildsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappcontainers.BuildsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappcontainers.BuildsClientDeleteResponse]](),
	}
}

// BuildsServerTransport connects instances of armappcontainers.BuildsClient to instances of BuildsServer.
// Don't use this type directly, use NewBuildsServerTransport instead.
type BuildsServerTransport struct {
	srv                 *BuildsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappcontainers.BuildsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappcontainers.BuildsClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for BuildsServerTransport.
func (b *BuildsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BuildsClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BuildsClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BuildsClient.Get":
		resp, err = b.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BuildsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builds/(?P<buildName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.BuildResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		buildNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, builderNameParam, buildNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BuildsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builds/(?P<buildName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		buildNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, builderNameParam, buildNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BuildsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builds/(?P<buildName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
	if err != nil {
		return nil, err
	}
	buildNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, builderNameParam, buildNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BuildResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
