// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhardwaresecuritymodules_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
	"log"
)

// Generated from example definition: 2025-03-31/CloudHsmCluster_CreateOrValidate_Backup_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_BeginBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginBackup(ctx, "rgcloudhsm", "chsm1", &armhardwaresecuritymodules.CloudHsmClustersClientBeginBackupOptions{
		BackupRequestProperties: &armhardwaresecuritymodules.BackupRequestProperties{
			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
			Token:                        to.Ptr("se=2018-02-01T00%3A00Z&spr=https&sv=2017-04-17&sr=b&sig=REDACTED"),
		}})
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientBackupResponse{
	// 	BackupResult: &armhardwaresecuritymodules.BackupResult{
	// 		Properties: &armhardwaresecuritymodules.BackupResultProperties{
	// 			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			JobID: to.Ptr("572a45927fc240e1ac075de27371680b"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			Status: to.Ptr(armhardwaresecuritymodules.BackupRestoreOperationStatusInProgress),
	// 			StatusDetails: to.Ptr("Backup operation is in progress"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_CreateOrUpdate_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginCreateOrUpdate(ctx, "rgcloudhsm", "chsm1", armhardwaresecuritymodules.CloudHsmCluster{
		Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
			Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": {},
			},
		},
		Location: to.Ptr("eastus2"),
		Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
			PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
		},
		SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
			Name:   to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
			Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
		},
		Tags: map[string]*string{
			"Dept":        to.Ptr("hsm"),
			"Environment": to.Ptr("dogfood"),
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientCreateOrUpdateResponse{
	// 	CloudHsmCluster: &armhardwaresecuritymodules.CloudHsmCluster{
	// 		Name: to.Ptr("chsm1"),
	// 		Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1"),
	// 		Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": &armhardwaresecuritymodules.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
	// 			ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
	// 			AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
	// 			ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
	// 			StatusMessage: to.Ptr("This is a status message"),
	// 		},
	// 		SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
	// 			Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
	// 			Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
	// 		},
	// 		SystemData: &armhardwaresecuritymodules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			CreatedBy: to.Ptr("CHsmUser1"),
	// 			CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("CHsmUser2"),
	// 			LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("hsm"),
	// 			"Environment": to.Ptr("dogfood"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_Delete_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginDelete(ctx, "rgcloudhsm", "chsm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_Get_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudHsmClustersClient().Get(ctx, "rgcloudhsm", "chsm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhardwaresecuritymodules.CloudHsmClustersClientGetResponse{
	// 	CloudHsmCluster: &armhardwaresecuritymodules.CloudHsmCluster{
	// 		Name: to.Ptr("chsm1"),
	// 		Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1"),
	// 		Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": &armhardwaresecuritymodules.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
	// 			ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
	// 			AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
	// 			ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
	// 			StatusMessage: to.Ptr("This is a status message"),
	// 		},
	// 		SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
	// 			Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
	// 			Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
	// 		},
	// 		SystemData: &armhardwaresecuritymodules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			CreatedBy: to.Ptr("CHsmUser1"),
	// 			CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("CHsmUser2"),
	// 			LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("hsm"),
	// 			"Environment": to.Ptr("dogfood"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_ListByResourceGroup_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudHsmClustersClient().NewListByResourceGroupPager("rgcloudhsm", nil)
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
		// page = armhardwaresecuritymodules.CloudHsmClustersClientListByResourceGroupResponse{
		// 	CloudHsmClusterListResult: armhardwaresecuritymodules.CloudHsmClusterListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters?api-version=2022-03-31&$skiptoken=dmF1bHQtcGVza3ktanVyeS03MzA3Ng=="),
		// 		Value: []*armhardwaresecuritymodules.CloudHsmCluster{
		// 			{
		// 				Name: to.Ptr("chsm1"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1"),
		// 				Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
		// 					Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": &armhardwaresecuritymodules.UserAssignedIdentity{
		// 							ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus2"),
		// 				Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
		// 					ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
		// 					AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
		// 					ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
		// 					StatusMessage: to.Ptr("This is a status message"),
		// 				},
		// 				SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
		// 					Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
		// 					Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
		// 				},
		// 				SystemData: &armhardwaresecuritymodules.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					CreatedBy: to.Ptr("CHsmUser1"),
		// 					CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("CHsmUser2"),
		// 					LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("dogfood"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("chsm2"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm2"),
		// 				Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
		// 					Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-2": &armhardwaresecuritymodules.UserAssignedIdentity{
		// 							ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus2"),
		// 				Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
		// 					ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
		// 					AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
		// 					ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
		// 				},
		// 				SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
		// 					Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
		// 					Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
		// 				},
		// 				SystemData: &armhardwaresecuritymodules.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					CreatedBy: to.Ptr("CHsmUser1"),
		// 					CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("CHsmUser2"),
		// 					LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("dogfood"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_ListBySubscription_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudHsmClustersClient().NewListBySubscriptionPager(nil)
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
		// page = armhardwaresecuritymodules.CloudHsmClustersClientListBySubscriptionResponse{
		// 	CloudHsmClusterListResult: armhardwaresecuritymodules.CloudHsmClusterListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters?api-version=2022-07-01&$skiptoken=dmF1bHQtcGVza3ktanVyeS03MzA3Ng=="),
		// 		Value: []*armhardwaresecuritymodules.CloudHsmCluster{
		// 			{
		// 				Name: to.Ptr("chsm1"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1"),
		// 				Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
		// 					Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": &armhardwaresecuritymodules.UserAssignedIdentity{
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus2"),
		// 				Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
		// 					ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
		// 					AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
		// 					ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
		// 					StatusMessage: to.Ptr("This is a status message"),
		// 				},
		// 				SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
		// 					Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
		// 					Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
		// 				},
		// 				SystemData: &armhardwaresecuritymodules.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					CreatedBy: to.Ptr("CHsmUser1"),
		// 					CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("CHsmUser2"),
		// 					LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("dogfood"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("chsm2"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm2"),
		// 				Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
		// 					Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-2": &armhardwaresecuritymodules.UserAssignedIdentity{
		// 							ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 							PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus2"),
		// 				Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
		// 					ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
		// 					AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
		// 					ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
		// 					StatusMessage: to.Ptr("This is a status message"),
		// 				},
		// 				SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
		// 					Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
		// 					Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
		// 				},
		// 				SystemData: &armhardwaresecuritymodules.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					CreatedBy: to.Ptr("CHsmUser1"),
		// 					CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("CHsmUser2"),
		// 					LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("dogfood"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_RequestOrValidate_Restore_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_BeginRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginRestore(ctx, "rgcloudhsm", "chsm1", armhardwaresecuritymodules.RestoreRequestProperties{
		AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
		BackupID:                     to.Ptr("backupId"),
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientRestoreResponse{
	// 	RestoreResult: &armhardwaresecuritymodules.RestoreResult{
	// 		Properties: &armhardwaresecuritymodules.BackupRestoreBaseResultProperties{
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			JobID: to.Ptr("572a45927fc240e1ac075de27371680b"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			Status: to.Ptr(armhardwaresecuritymodules.BackupRestoreOperationStatusInProgress),
	// 			StatusDetails: to.Ptr("Restore operation is in progress"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_Update_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginUpdate(ctx, "rgcloudhsm", "chsm1", armhardwaresecuritymodules.CloudHsmClusterPatchParameters{
		Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
			Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": {},
			},
		},
		Tags: map[string]*string{
			"Dept":        to.Ptr("hsm"),
			"Environment": to.Ptr("dogfood"),
			"Slice":       to.Ptr("A"),
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientUpdateResponse{
	// 	CloudHsmCluster: &armhardwaresecuritymodules.CloudHsmCluster{
	// 		Name: to.Ptr("chsm1"),
	// 		Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1"),
	// 		Identity: &armhardwaresecuritymodules.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhardwaresecuritymodules.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armhardwaresecuritymodules.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1": &armhardwaresecuritymodules.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armhardwaresecuritymodules.CloudHsmClusterProperties{
	// 			ActivationState: to.Ptr(armhardwaresecuritymodules.ActivationState("null")),
	// 			AutoGeneratedDomainNameLabelScope: to.Ptr(armhardwaresecuritymodules.AutoGeneratedDomainNameLabelScopeTenantReuse),
	// 			ProvisioningState: to.Ptr(armhardwaresecuritymodules.ProvisioningStateUpdating),
	// 			PublicNetworkAccess: to.Ptr(armhardwaresecuritymodules.PublicNetworkAccessDisabled),
	// 			StatusMessage: to.Ptr("This is a status message"),
	// 		},
	// 		SKU: &armhardwaresecuritymodules.CloudHsmClusterSKU{
	// 			Name: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUNameStandardB1),
	// 			Family: to.Ptr(armhardwaresecuritymodules.CloudHsmClusterSKUFamilyB),
	// 		},
	// 		SystemData: &armhardwaresecuritymodules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			CreatedBy: to.Ptr("CHsmUser1"),
	// 			CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("CHsmUser2"),
	// 			LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("hsm"),
	// 			"Environment": to.Ptr("dogfood"),
	// 			"Slice": to.Ptr("A"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_Create_Backup_MaximumSet_Gen_ValidateBackupProperties.json
func ExampleCloudHsmClustersClient_BeginValidateBackupProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginValidateBackupProperties(ctx, "rgcloudhsm", "chsm1", &armhardwaresecuritymodules.CloudHsmClustersClientBeginValidateBackupPropertiesOptions{
		BackupRequestProperties: &armhardwaresecuritymodules.BackupRequestProperties{
			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
			Token:                        to.Ptr("se=2018-02-01T00%3A00Z&spr=https&sv=2017-04-17&sr=b&sig=REDACTED"),
		}})
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientValidateBackupPropertiesResponse{
	// 	BackupResult: &armhardwaresecuritymodules.BackupResult{
	// 		Properties: &armhardwaresecuritymodules.BackupResultProperties{
	// 			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			JobID: to.Ptr("572a45927fc240e1ac075de27371680b"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			Status: to.Ptr(armhardwaresecuritymodules.BackupRestoreOperationStatusInProgress),
	// 			StatusDetails: to.Ptr("Backup operation is in progress"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-31/CloudHsmCluster_RequestOrValidate_Restore_MaximumSet_Gen_ValidateRestoreProperties.json
func ExampleCloudHsmClustersClient_BeginValidateRestoreProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginValidateRestoreProperties(ctx, "rgcloudhsm", "chsm1", &armhardwaresecuritymodules.CloudHsmClustersClientBeginValidateRestorePropertiesOptions{
		RestoreRequestProperties: &armhardwaresecuritymodules.RestoreRequestProperties{
			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
			BackupID:                     to.Ptr("backupId"),
		}})
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientValidateRestorePropertiesResponse{
	// 	RestoreResult: &armhardwaresecuritymodules.RestoreResult{
	// 		Properties: &armhardwaresecuritymodules.BackupRestoreBaseResultProperties{
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			JobID: to.Ptr("572a45927fc240e1ac075de27371680b"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			Status: to.Ptr(armhardwaresecuritymodules.BackupRestoreOperationStatusInProgress),
	// 			StatusDetails: to.Ptr("Restore operation is in progress"),
	// 		},
	// 	},
	// }
}
