// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewAuthorizedApplicationsClient creates a new instance of AuthorizedApplicationsClient.
func (c *ClientFactory) NewAuthorizedApplicationsClient() *AuthorizedApplicationsClient {
	return &AuthorizedApplicationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCustomRolloutsClient creates a new instance of CustomRolloutsClient.
func (c *ClientFactory) NewCustomRolloutsClient() *CustomRolloutsClient {
	return &CustomRolloutsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDefaultRolloutsClient creates a new instance of DefaultRolloutsClient.
func (c *ClientFactory) NewDefaultRolloutsClient() *DefaultRolloutsClient {
	return &DefaultRolloutsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNewRegionFrontloadReleaseClient creates a new instance of NewRegionFrontloadReleaseClient.
func (c *ClientFactory) NewNewRegionFrontloadReleaseClient() *NewRegionFrontloadReleaseClient {
	return &NewRegionFrontloadReleaseClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNotificationRegistrationsClient creates a new instance of NotificationRegistrationsClient.
func (c *ClientFactory) NewNotificationRegistrationsClient() *NotificationRegistrationsClient {
	return &NotificationRegistrationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProviderMonitorSettingsClient creates a new instance of ProviderMonitorSettingsClient.
func (c *ClientFactory) NewProviderMonitorSettingsClient() *ProviderMonitorSettingsClient {
	return &ProviderMonitorSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProviderRegistrationsClient creates a new instance of ProviderRegistrationsClient.
func (c *ClientFactory) NewProviderRegistrationsClient() *ProviderRegistrationsClient {
	return &ProviderRegistrationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewResourceActionsClient creates a new instance of ResourceActionsClient.
func (c *ClientFactory) NewResourceActionsClient() *ResourceActionsClient {
	return &ResourceActionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewResourceTypeRegistrationsClient creates a new instance of ResourceTypeRegistrationsClient.
func (c *ClientFactory) NewResourceTypeRegistrationsClient() *ResourceTypeRegistrationsClient {
	return &ResourceTypeRegistrationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSKUsClient creates a new instance of SKUsClient.
func (c *ClientFactory) NewSKUsClient() *SKUsClient {
	return &SKUsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
