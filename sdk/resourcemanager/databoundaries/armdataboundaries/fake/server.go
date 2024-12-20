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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoundaries/armdataboundaries"
	"net/http"
	"net/url"
	"regexp"
)

// Server is a fake server for instances of the armdataboundaries.Client type.
type Server struct {
	// GetScope is the fake for method Client.GetScope
	// HTTP status codes to indicate success: http.StatusOK
	GetScope func(ctx context.Context, scope string, defaultParam armdataboundaries.DefaultName, options *armdataboundaries.ClientGetScopeOptions) (resp azfake.Responder[armdataboundaries.ClientGetScopeResponse], errResp azfake.ErrorResponder)

	// GetTenant is the fake for method Client.GetTenant
	// HTTP status codes to indicate success: http.StatusOK
	GetTenant func(ctx context.Context, defaultParam armdataboundaries.DefaultName, options *armdataboundaries.ClientGetTenantOptions) (resp azfake.Responder[armdataboundaries.ClientGetTenantResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method Client.Put
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Put func(ctx context.Context, defaultParam armdataboundaries.DefaultName, dataBoundaryDefinition armdataboundaries.DataBoundaryDefinition, options *armdataboundaries.ClientPutOptions) (resp azfake.Responder[armdataboundaries.ClientPutResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armdataboundaries.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{srv: srv}
}

// ServerTransport connects instances of armdataboundaries.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv *Server
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.GetScope":
		resp, err = s.dispatchGetScope(req)
	case "Client.GetTenant":
		resp, err = s.dispatchGetTenant(req)
	case "Client.Put":
		resp, err = s.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchGetScope(req *http.Request) (*http.Response, error) {
	if s.srv.GetScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetScope not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Resources/dataBoundaries/(?P<default>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	defaultParamParam, err := parseWithCast(matches[regex.SubexpIndex("default")], func(v string) (armdataboundaries.DefaultName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armdataboundaries.DefaultName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetScope(req.Context(), scopeParam, defaultParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataBoundaryDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchGetTenant(req *http.Request) (*http.Response, error) {
	if s.srv.GetTenant == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTenant not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Resources/dataBoundaries/(?P<default>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	defaultParamParam, err := parseWithCast(matches[regex.SubexpIndex("default")], func(v string) (armdataboundaries.DefaultName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armdataboundaries.DefaultName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetTenant(req.Context(), defaultParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataBoundaryDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if s.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Resources/dataBoundaries/(?P<default>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataboundaries.DataBoundaryDefinition](req)
	if err != nil {
		return nil, err
	}
	defaultParamParam, err := parseWithCast(matches[regex.SubexpIndex("default")], func(v string) (armdataboundaries.DefaultName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armdataboundaries.DefaultName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Put(req.Context(), defaultParamParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataBoundaryDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
