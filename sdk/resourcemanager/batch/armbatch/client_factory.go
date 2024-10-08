//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAccountClient creates a new instance of AccountClient.
func (c *ClientFactory) NewAccountClient() *AccountClient {
	return &AccountClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationClient creates a new instance of ApplicationClient.
func (c *ClientFactory) NewApplicationClient() *ApplicationClient {
	return &ApplicationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationPackageClient creates a new instance of ApplicationPackageClient.
func (c *ClientFactory) NewApplicationPackageClient() *ApplicationPackageClient {
	return &ApplicationPackageClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCertificateClient creates a new instance of CertificateClient.
func (c *ClientFactory) NewCertificateClient() *CertificateClient {
	return &CertificateClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLocationClient creates a new instance of LocationClient.
func (c *ClientFactory) NewLocationClient() *LocationClient {
	return &LocationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkSecurityPerimeterClient creates a new instance of NetworkSecurityPerimeterClient.
func (c *ClientFactory) NewNetworkSecurityPerimeterClient() *NetworkSecurityPerimeterClient {
	return &NetworkSecurityPerimeterClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPoolClient creates a new instance of PoolClient.
func (c *ClientFactory) NewPoolClient() *PoolClient {
	return &PoolClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionClient creates a new instance of PrivateEndpointConnectionClient.
func (c *ClientFactory) NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient {
	return &PrivateEndpointConnectionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourceClient creates a new instance of PrivateLinkResourceClient.
func (c *ClientFactory) NewPrivateLinkResourceClient() *PrivateLinkResourceClient {
	return &PrivateLinkResourceClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
