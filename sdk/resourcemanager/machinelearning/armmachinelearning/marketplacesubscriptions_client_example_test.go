//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/MarketplaceSubscription/list.json
func ExampleMarketplaceSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMarketplaceSubscriptionsClient().NewListPager("test-rg", "my-aml-workspace", &armmachinelearning.MarketplaceSubscriptionsClientListOptions{Skip: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MarketplaceSubscriptionResourceArmPaginatedResult = armmachinelearning.MarketplaceSubscriptionResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.MarketplaceSubscription{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:28:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:28:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
		// 			},
		// 			Properties: &armmachinelearning.MarketplaceSubscriptionProperties{
		// 				MarketplacePlan: &armmachinelearning.MarketplacePlan{
		// 					OfferID: to.Ptr("string"),
		// 					PlanID: to.Ptr("string"),
		// 					PublisherID: to.Ptr("string"),
		// 				},
		// 				MarketplaceSubscriptionStatus: to.Ptr(armmachinelearning.MarketplaceSubscriptionStatusSuspended),
		// 				ModelID: to.Ptr("string"),
		// 				ProvisioningState: to.Ptr(armmachinelearning.MarketplaceSubscriptionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/MarketplaceSubscription/delete.json
func ExampleMarketplaceSubscriptionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMarketplaceSubscriptionsClient().BeginDelete(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/MarketplaceSubscription/get.json
func ExampleMarketplaceSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMarketplaceSubscriptionsClient().Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MarketplaceSubscription = armmachinelearning.MarketplaceSubscription{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:26:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:26:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 	},
	// 	Properties: &armmachinelearning.MarketplaceSubscriptionProperties{
	// 		MarketplacePlan: &armmachinelearning.MarketplacePlan{
	// 			OfferID: to.Ptr("string"),
	// 			PlanID: to.Ptr("string"),
	// 			PublisherID: to.Ptr("string"),
	// 		},
	// 		MarketplaceSubscriptionStatus: to.Ptr(armmachinelearning.MarketplaceSubscriptionStatusSubscribed),
	// 		ModelID: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armmachinelearning.MarketplaceSubscriptionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/MarketplaceSubscription/createOrUpdate.json
func ExampleMarketplaceSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMarketplaceSubscriptionsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.MarketplaceSubscription{
		Properties: &armmachinelearning.MarketplaceSubscriptionProperties{
			ModelID: to.Ptr("string"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MarketplaceSubscription = armmachinelearning.MarketplaceSubscription{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:26:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:26:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 	},
	// 	Properties: &armmachinelearning.MarketplaceSubscriptionProperties{
	// 		MarketplacePlan: &armmachinelearning.MarketplacePlan{
	// 			OfferID: to.Ptr("string"),
	// 			PlanID: to.Ptr("string"),
	// 			PublisherID: to.Ptr("string"),
	// 		},
	// 		MarketplaceSubscriptionStatus: to.Ptr(armmachinelearning.MarketplaceSubscriptionStatusSuspended),
	// 		ModelID: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armmachinelearning.MarketplaceSubscriptionProvisioningStateSucceeded),
	// 	},
	// }
}
