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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// WorkspaceFeaturesServer is a fake server for instances of the armmachinelearning.WorkspaceFeaturesClient type.
type WorkspaceFeaturesServer struct {
	// NewListPager is the fake for method WorkspaceFeaturesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armmachinelearning.WorkspaceFeaturesClientListOptions) (resp azfake.PagerResponder[armmachinelearning.WorkspaceFeaturesClientListResponse])
}

// NewWorkspaceFeaturesServerTransport creates a new instance of WorkspaceFeaturesServerTransport with the provided implementation.
// The returned WorkspaceFeaturesServerTransport instance is connected to an instance of armmachinelearning.WorkspaceFeaturesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceFeaturesServerTransport(srv *WorkspaceFeaturesServer) *WorkspaceFeaturesServerTransport {
	return &WorkspaceFeaturesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.WorkspaceFeaturesClientListResponse]](),
	}
}

// WorkspaceFeaturesServerTransport connects instances of armmachinelearning.WorkspaceFeaturesClient to instances of WorkspaceFeaturesServer.
// Don't use this type directly, use NewWorkspaceFeaturesServerTransport instead.
type WorkspaceFeaturesServerTransport struct {
	srv          *WorkspaceFeaturesServer
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.WorkspaceFeaturesClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceFeaturesServerTransport.
func (w *WorkspaceFeaturesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceFeaturesClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceFeaturesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.WorkspaceFeaturesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}
