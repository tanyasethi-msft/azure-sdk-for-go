//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/VerifiedPartners_Get.json
func ExampleVerifiedPartnersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVerifiedPartnersClient().Get(ctx, "Contoso.Finance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VerifiedPartner = armeventgrid.VerifiedPartner{
	// 	Name: to.Ptr("Contoso.Finance"),
	// 	Type: to.Ptr("Microsoft.EventGrid/verifiedPartners"),
	// 	ID: to.Ptr("/providers/Microsoft.EventGrid/verifiedPartners/Contoso.Finance"),
	// 	Properties: &armeventgrid.VerifiedPartnerProperties{
	// 		OrganizationName: to.Ptr("Contoso"),
	// 		PartnerDestinationDetails: &armeventgrid.PartnerDetails{
	// 			Description: to.Ptr("This is custom description"),
	// 			LongDescription: to.Ptr("This is long custom description"),
	// 			SetupURI: to.Ptr("https://www.example.com/"),
	// 		},
	// 		PartnerDisplayName: to.Ptr("Contoso - Finance Department"),
	// 		PartnerRegistrationImmutableID: to.Ptr("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
	// 		PartnerTopicDetails: &armeventgrid.PartnerDetails{
	// 			Description: to.Ptr("This is short description"),
	// 			LongDescription: to.Ptr("This is really long long... long description"),
	// 			SetupURI: to.Ptr("https://www.example.com/"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/VerifiedPartners_List.json
func ExampleVerifiedPartnersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVerifiedPartnersClient().NewListPager(&armeventgrid.VerifiedPartnersClientListOptions{Filter: nil,
		Top: nil,
	})
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
		// page.VerifiedPartnersListResult = armeventgrid.VerifiedPartnersListResult{
		// 	Value: []*armeventgrid.VerifiedPartner{
		// 		{
		// 			Name: to.Ptr("Contoso.Finance"),
		// 			Type: to.Ptr("Microsoft.EventGrid/verifiedPartners"),
		// 			ID: to.Ptr("/providers/Microsoft.EventGrid/verifiedPartners/Contoso.Finance"),
		// 			Properties: &armeventgrid.VerifiedPartnerProperties{
		// 				OrganizationName: to.Ptr("Contoso"),
		// 				PartnerDestinationDetails: &armeventgrid.PartnerDetails{
		// 					Description: to.Ptr("This is custom description"),
		// 					LongDescription: to.Ptr("This is long custom description"),
		// 					SetupURI: to.Ptr("https://www.example.com/"),
		// 				},
		// 				PartnerDisplayName: to.Ptr("Contoso - Finance Department"),
		// 				PartnerRegistrationImmutableID: to.Ptr("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
		// 				PartnerTopicDetails: &armeventgrid.PartnerDetails{
		// 					Description: to.Ptr("This is short description"),
		// 					LongDescription: to.Ptr("This is really long long... long description"),
		// 					SetupURI: to.Ptr("https://www.example.com/"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
