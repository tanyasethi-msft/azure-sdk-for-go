//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevopsinfrastructure

import "time"

// AgentProfile - The agent profile of the machines in the pool.
type AgentProfile struct {
	// REQUIRED; Discriminator property for AgentProfile.
	Kind *string

	// Defines pool buffer/stand-by agents.
	ResourcePredictions any

	// Defines how the pool buffer/stand-by agents is provided.
	ResourcePredictionsProfile ResourcePredictionsProfileClassification
}

// GetAgentProfile implements the AgentProfileClassification interface for type AgentProfile.
func (a *AgentProfile) GetAgentProfile() *AgentProfile { return a }

// AgentProfileUpdate - The agent profile of the machines in the pool.
type AgentProfileUpdate struct {
	// REQUIRED; Discriminator property for AgentProfile.
	Kind *string

	// Defines pool buffer/stand-by agents.
	ResourcePredictions any

	// Defines how the pool buffer/stand-by agents is provided.
	ResourcePredictionsProfile ResourcePredictionsProfileUpdateClassification
}

// GetAgentProfileUpdate implements the AgentProfileUpdateClassification interface for type AgentProfileUpdate.
func (a *AgentProfileUpdate) GetAgentProfileUpdate() *AgentProfileUpdate { return a }

// AutomaticResourcePredictionsProfile - The stand-by agent scheme is determined based on historical demand.
type AutomaticResourcePredictionsProfile struct {
	// REQUIRED; Determines how the stand-by scheme should be provided.
	Kind *ResourcePredictionsProfileType

	// Determines the balance between cost and performance.
	PredictionPreference *PredictionPreference
}

// GetResourcePredictionsProfile implements the ResourcePredictionsProfileClassification interface for type AutomaticResourcePredictionsProfile.
func (a *AutomaticResourcePredictionsProfile) GetResourcePredictionsProfile() *ResourcePredictionsProfile {
	return &ResourcePredictionsProfile{
		Kind: a.Kind,
	}
}

// AutomaticResourcePredictionsProfileUpdate - The stand-by agent scheme is determined based on historical demand.
type AutomaticResourcePredictionsProfileUpdate struct {
	// REQUIRED; Determines how the stand-by scheme should be provided.
	Kind *ResourcePredictionsProfileType

	// Determines the balance between cost and performance.
	PredictionPreference *PredictionPreference
}

// GetResourcePredictionsProfileUpdate implements the ResourcePredictionsProfileUpdateClassification interface for type AutomaticResourcePredictionsProfileUpdate.
func (a *AutomaticResourcePredictionsProfileUpdate) GetResourcePredictionsProfileUpdate() *ResourcePredictionsProfileUpdate {
	return &ResourcePredictionsProfileUpdate{
		Kind: a.Kind,
	}
}

// AzureDevOpsOrganizationProfile - Azure DevOps organization profile
type AzureDevOpsOrganizationProfile struct {
	// REQUIRED; Discriminator property for OrganizationProfile.
	Kind *string

	// REQUIRED; The list of Azure DevOps organizations the pool should be present in.
	Organizations []*Organization

	// The type of permission which determines which accounts are admins on the Azure DevOps pool.
	PermissionProfile *AzureDevOpsPermissionProfile
}

// GetOrganizationProfile implements the OrganizationProfileClassification interface for type AzureDevOpsOrganizationProfile.
func (a *AzureDevOpsOrganizationProfile) GetOrganizationProfile() *OrganizationProfile {
	return &OrganizationProfile{
		Kind: a.Kind,
	}
}

// AzureDevOpsPermissionProfile - Defines the type of Azure DevOps pool permission.
type AzureDevOpsPermissionProfile struct {
	// REQUIRED; Determines who has admin permissions to the Azure DevOps pool.
	Kind *AzureDevOpsPermissionType

	// Group email addresses
	Groups []*string

	// User email addresses
	Users []*string
}

// DataDisk - The data disk of the VMSS.
type DataDisk struct {
	// The type of caching to be enabled for the data disks. The default value for caching is readwrite. For information about
	// the caching options see:
	// https://blogs.msdn.microsoft.com/windowsazurestorage/2012/06/27/exploring-windows-azure-drives-disks-and-images/.
	Caching *CachingType

	// The initial disk size in gigabytes.
	DiskSizeGiB *int32

	// The drive letter for the empty data disk. If not specified, it will be the first available letter.
	DriveLetter *string

	// The storage Account type to be used for the data disk. If omitted, the default is "standard_lrs".
	StorageAccountType *StorageAccountType
}

// DevOpsAzureSKU - The Azure SKU of the machines in the pool.
type DevOpsAzureSKU struct {
	// REQUIRED; The Azure SKU name of the machines in the pool.
	Name *string
}

// FabricProfile - Defines the type of fabric the agent will run on.
type FabricProfile struct {
	// REQUIRED; Discriminator property for FabricProfile.
	Kind *string
}

// GetFabricProfile implements the FabricProfileClassification interface for type FabricProfile.
func (f *FabricProfile) GetFabricProfile() *FabricProfile { return f }

// GitHubOrganization - Defines a GitHub organization
type GitHubOrganization struct {
	// REQUIRED; The GitHub organization URL in which the pool should be created.
	URL *string

	// Optional list of repositories in which the pool should be created.
	Repositories []*string
}

// GitHubOrganizationProfile - GitHub organization profile
type GitHubOrganizationProfile struct {
	// REQUIRED; Discriminator property for OrganizationProfile.
	Kind *string

	// REQUIRED; The list of GitHub organizations/repositories the pool should be present in.
	Organizations []*GitHubOrganization
}

// GetOrganizationProfile implements the OrganizationProfileClassification interface for type GitHubOrganizationProfile.
func (g *GitHubOrganizationProfile) GetOrganizationProfile() *OrganizationProfile {
	return &OrganizationProfile{
		Kind: g.Kind,
	}
}

// ImageVersion - An image version object
type ImageVersion struct {
	// The resource-specific properties for this resource.
	Properties *ImageVersionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ImageVersionListResult - The response of a ImageVersion list operation.
type ImageVersionListResult struct {
	// REQUIRED; The ImageVersion items on this page
	Value []*ImageVersion

	// The link to the next page of items
	NextLink *string
}

// ImageVersionProperties - Details of the ImageVersionProperties.
type ImageVersionProperties struct {
	// REQUIRED; Version of the image.
	Version *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// ManualResourcePredictionsProfile - Customer provides the stand-by agent scheme.
type ManualResourcePredictionsProfile struct {
	// REQUIRED; Determines how the stand-by scheme should be provided.
	Kind *ResourcePredictionsProfileType
}

// GetResourcePredictionsProfile implements the ResourcePredictionsProfileClassification interface for type ManualResourcePredictionsProfile.
func (m *ManualResourcePredictionsProfile) GetResourcePredictionsProfile() *ResourcePredictionsProfile {
	return &ResourcePredictionsProfile{
		Kind: m.Kind,
	}
}

// ManualResourcePredictionsProfileUpdate - Customer provides the stand-by agent scheme.
type ManualResourcePredictionsProfileUpdate struct {
	// REQUIRED; Determines how the stand-by scheme should be provided.
	Kind *ResourcePredictionsProfileType
}

// GetResourcePredictionsProfileUpdate implements the ResourcePredictionsProfileUpdateClassification interface for type ManualResourcePredictionsProfileUpdate.
func (m *ManualResourcePredictionsProfileUpdate) GetResourcePredictionsProfileUpdate() *ResourcePredictionsProfileUpdate {
	return &ResourcePredictionsProfileUpdate{
		Kind: m.Kind,
	}
}

// NetworkProfile - The network profile of the machines in the pool.
type NetworkProfile struct {
	// REQUIRED; The subnet id on which to put all machines created in the pool.
	SubnetID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// Organization - Defines an Azure DevOps organization.
type Organization struct {
	// REQUIRED; The Azure DevOps organization URL in which the pool should be created.
	URL *string

	// How many machines can be created at maximum in this organization out of the maximumConcurrency of the pool.
	Parallelism *int32

	// Optional list of projects in which the pool should be created.
	Projects []*string
}

// OrganizationProfile - Defines the organization in which the pool will be used.
type OrganizationProfile struct {
	// REQUIRED; Discriminator property for OrganizationProfile.
	Kind *string
}

// GetOrganizationProfile implements the OrganizationProfileClassification interface for type OrganizationProfile.
func (o *OrganizationProfile) GetOrganizationProfile() *OrganizationProfile { return o }

// OsProfile - The OS profile of the machines in the pool.
type OsProfile struct {
	// Determines how the service should be run. By default, this will be set to Service.
	LogonType *LogonType

	// The secret management settings of the machines in the pool.
	SecretsManagementSettings *SecretsManagementSettings
}

// Pool - Concrete tracked resource types can be created by aliasing this type using a specific property type.
type Pool struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *PoolProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PoolImage - The VM image of the machines in the pool.
type PoolImage struct {
	// List of aliases to reference the image by.
	Aliases []*string

	// The percentage of the buffer to be allocated to this image.
	Buffer *string

	// The resource id of the image.
	ResourceID *string

	// The image to use from a well-known set of images made available to customers.
	WellKnownImageName *string
}

// PoolListResult - The response of a Pool list operation.
type PoolListResult struct {
	// REQUIRED; The Pool items on this page
	Value []*Pool

	// The link to the next page of items
	NextLink *string
}

// PoolProperties - Pool properties
type PoolProperties struct {
	// REQUIRED; Defines how the machine will be handled once it executed a job.
	AgentProfile AgentProfileClassification

	// REQUIRED; The resource id of the DevCenter Project the pool belongs to.
	DevCenterProjectResourceID *string

	// REQUIRED; Defines the type of fabric the agent will run on.
	FabricProfile FabricProfileClassification

	// REQUIRED; Defines how many resources can there be created at any given time.
	MaximumConcurrency *int32

	// REQUIRED; Defines the organization in which the pool will be used.
	OrganizationProfile OrganizationProfileClassification

	// The status of the current operation.
	ProvisioningState *ProvisioningState
}

// PoolUpdate - The type used for update operations of the Pool.
type PoolUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The updatable properties of the Pool.
	Properties *PoolUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// PoolUpdateProperties - The updatable properties of the Pool.
type PoolUpdateProperties struct {
	// Defines how the machine will be handled once it executed a job.
	AgentProfile AgentProfileUpdateClassification

	// The resource id of the DevCenter Project the pool belongs to.
	DevCenterProjectResourceID *string

	// Defines the type of fabric the agent will run on.
	FabricProfile FabricProfileClassification

	// Defines how many resources can there be created at any given time.
	MaximumConcurrency *int32

	// Defines the organization in which the pool will be used.
	OrganizationProfile OrganizationProfileClassification

	// The status of the current operation.
	ProvisioningState *ProvisioningState
}

// Quota - Describes Resource Quota
type Quota struct {
	// The resource-specific properties for this resource.
	Properties *QuotaProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// QuotaListResult - The response of a Quota list operation.
type QuotaListResult struct {
	// REQUIRED; The Quota items on this page
	Value []*Quota

	// The link to the next page of items
	NextLink *string
}

// QuotaName - The Quota Names
type QuotaName struct {
	// The localized name of the resource.
	LocalizedValue *string

	// The name of the resource.
	Value *string
}

// QuotaProperties - Describes Resource Quota properties
type QuotaProperties struct {
	// REQUIRED; The current usage of the resource.
	CurrentValue *int64

	// REQUIRED; The maximum permitted usage of the resource.
	Limit *int64

	// REQUIRED; The details of the quota.
	Name *QuotaName

	// REQUIRED; The unit of usage measurement.
	Unit *string
}

// ResourceDetailsObject - A ResourceDetailsObject
type ResourceDetailsObject struct {
	// The resource-specific properties for this resource.
	Properties *ResourceDetailsObjectProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceDetailsObjectListResult - The response of a ResourceDetailsObject list operation.
type ResourceDetailsObjectListResult struct {
	// REQUIRED; The ResourceDetailsObject items on this page
	Value []*ResourceDetailsObject

	// The link to the next page of items
	NextLink *string
}

// ResourceDetailsObjectProperties - Details of the ResourceDetailsObject.
type ResourceDetailsObjectProperties struct {
	// REQUIRED; The image name of the resource.
	Image *string

	// REQUIRED; The version of the image running on the resource.
	ImageVersion *string

	// REQUIRED; The status of the resource.
	Status *ResourceStatus
}

// ResourcePredictionsProfile - Determines how the stand-by scheme should be provided.
type ResourcePredictionsProfile struct {
	// REQUIRED; Determines how the stand-by scheme should be provided.
	Kind *ResourcePredictionsProfileType
}

// GetResourcePredictionsProfile implements the ResourcePredictionsProfileClassification interface for type ResourcePredictionsProfile.
func (r *ResourcePredictionsProfile) GetResourcePredictionsProfile() *ResourcePredictionsProfile {
	return r
}

// ResourcePredictionsProfileUpdate - Determines how the stand-by scheme should be provided.
type ResourcePredictionsProfileUpdate struct {
	// REQUIRED; Determines how the stand-by scheme should be provided.
	Kind *ResourcePredictionsProfileType
}

// GetResourcePredictionsProfileUpdate implements the ResourcePredictionsProfileUpdateClassification interface for type ResourcePredictionsProfileUpdate.
func (r *ResourcePredictionsProfileUpdate) GetResourcePredictionsProfileUpdate() *ResourcePredictionsProfileUpdate {
	return r
}

// ResourceSKU - A ResourceSku
type ResourceSKU struct {
	// The resource-specific properties for this resource.
	Properties *ResourceSKUProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceSKUCapabilities - Describes The SKU capabilities object.
type ResourceSKUCapabilities struct {
	// REQUIRED; The name of the SKU capability.
	Name *string

	// REQUIRED; The value of the SKU capability.
	Value *string
}

// ResourceSKUListResult - The response of a ResourceSku list operation.
type ResourceSKUListResult struct {
	// REQUIRED; The ResourceSku items on this page
	Value []*ResourceSKU

	// The link to the next page of items
	NextLink *string
}

// ResourceSKULocationInfo - Describes an available Compute SKU Location Information.
type ResourceSKULocationInfo struct {
	// REQUIRED; Location of the SKU
	Location *string

	// REQUIRED; Gets details of capabilities available to a SKU in specific zones.
	ZoneDetails []*ResourceSKUZoneDetails

	// REQUIRED; List of availability zones where the SKU is supported.
	Zones []*string
}

// ResourceSKUProperties - Properties of a ResourceSku
type ResourceSKUProperties struct {
	// REQUIRED; Name value pairs to describe the capability.
	Capabilities []*ResourceSKUCapabilities

	// REQUIRED; The family of the SKU.
	Family *string

	// REQUIRED; A list of locations and availability zones in those locations where the SKU is available
	LocationInfo []*ResourceSKULocationInfo

	// REQUIRED; The set of locations that the SKU is available.
	Locations []*string

	// REQUIRED; The type of resource the SKU applies to.
	ResourceType *string

	// REQUIRED; The restrictions of the SKU.
	Restrictions []*ResourceSKURestrictions

	// REQUIRED; The size of the SKU.
	Size *string

	// REQUIRED; The tier of virtual machines in a scale set
	Tier *string
}

// ResourceSKURestrictionInfo - Describes an available Compute SKU Restriction Information.
type ResourceSKURestrictionInfo struct {
	// Locations where the SKU is restricted
	Locations []*string

	// List of availability zones where the SKU is restricted.
	Zones []*string
}

// ResourceSKURestrictions - The restrictions of the SKU.
type ResourceSKURestrictions struct {
	// REQUIRED; The information about the restriction where the SKU cannot be used.
	RestrictionInfo *ResourceSKURestrictionInfo

	// REQUIRED; The value of restrictions. If the restriction type is set to location. This would be different locations where
	// the SKU is restricted.
	Values []*string

	// the reason for restriction.
	ReasonCode *ResourceSKURestrictionsReasonCode

	// the type of restrictions.
	Type *ResourceSKURestrictionsType
}

// ResourceSKUZoneDetails - Describes The zonal capabilities of a SKU.
type ResourceSKUZoneDetails struct {
	// REQUIRED; A list of capabilities that are available for the SKU in the specified list of zones.
	Capabilities []*ResourceSKUCapabilities

	// REQUIRED; Gets the set of zones that the SKU is available in with the specified capabilities.
	Name []*string
}

// SecretsManagementSettings - The secret management settings of the machines in the pool.
type SecretsManagementSettings struct {
	// REQUIRED; Defines if the key of the certificates should be exportable.
	KeyExportable *bool

	// REQUIRED; The list of certificates to install on all machines in the pool.
	ObservedCertificates []*string

	// Where to store certificates on the machine.
	CertificateStoreLocation *string
}

// Stateful profile meaning that the machines will be returned to the pool after running a job.
type Stateful struct {
	// REQUIRED; Discriminator property for AgentProfile.
	Kind *string

	// How long should the machine be kept around after it ran a workload when there are no stand-by agents. The maximum is one
	// week.
	GracePeriodTimeSpan *string

	// How long should stateful machines be kept around. The maximum is one week.
	MaxAgentLifetime *string

	// Defines pool buffer/stand-by agents.
	ResourcePredictions any

	// Defines how the pool buffer/stand-by agents is provided.
	ResourcePredictionsProfile ResourcePredictionsProfileClassification
}

// GetAgentProfile implements the AgentProfileClassification interface for type Stateful.
func (s *Stateful) GetAgentProfile() *AgentProfile {
	return &AgentProfile{
		Kind:                       s.Kind,
		ResourcePredictions:        s.ResourcePredictions,
		ResourcePredictionsProfile: s.ResourcePredictionsProfile,
	}
}

// StatefulUpdate - Stateful profile meaning that the machines will be returned to the pool after running a job.
type StatefulUpdate struct {
	// REQUIRED; Discriminator property for AgentProfile.
	Kind *string

	// How long should the machine be kept around after it ran a workload when there are no stand-by agents. The maximum is one
	// week.
	GracePeriodTimeSpan *string

	// How long should stateful machines be kept around. The maximum is one week.
	MaxAgentLifetime *string

	// Defines pool buffer/stand-by agents.
	ResourcePredictions any

	// Defines how the pool buffer/stand-by agents is provided.
	ResourcePredictionsProfile ResourcePredictionsProfileUpdateClassification
}

// GetAgentProfileUpdate implements the AgentProfileUpdateClassification interface for type StatefulUpdate.
func (s *StatefulUpdate) GetAgentProfileUpdate() *AgentProfileUpdate {
	return &AgentProfileUpdate{
		Kind:                       s.Kind,
		ResourcePredictions:        s.ResourcePredictions,
		ResourcePredictionsProfile: s.ResourcePredictionsProfile,
	}
}

// StatelessAgentProfile - Stateless profile meaning that the machines will be cleaned up after running a job.
type StatelessAgentProfile struct {
	// REQUIRED; Discriminator property for AgentProfile.
	Kind *string

	// Defines pool buffer/stand-by agents.
	ResourcePredictions any

	// Defines how the pool buffer/stand-by agents is provided.
	ResourcePredictionsProfile ResourcePredictionsProfileClassification
}

// GetAgentProfile implements the AgentProfileClassification interface for type StatelessAgentProfile.
func (s *StatelessAgentProfile) GetAgentProfile() *AgentProfile {
	return &AgentProfile{
		Kind:                       s.Kind,
		ResourcePredictions:        s.ResourcePredictions,
		ResourcePredictionsProfile: s.ResourcePredictionsProfile,
	}
}

// StatelessAgentProfileUpdate - Stateless profile meaning that the machines will be cleaned up after running a job.
type StatelessAgentProfileUpdate struct {
	// REQUIRED; Discriminator property for AgentProfile.
	Kind *string

	// Defines pool buffer/stand-by agents.
	ResourcePredictions any

	// Defines how the pool buffer/stand-by agents is provided.
	ResourcePredictionsProfile ResourcePredictionsProfileUpdateClassification
}

// GetAgentProfileUpdate implements the AgentProfileUpdateClassification interface for type StatelessAgentProfileUpdate.
func (s *StatelessAgentProfileUpdate) GetAgentProfileUpdate() *AgentProfileUpdate {
	return &AgentProfileUpdate{
		Kind:                       s.Kind,
		ResourcePredictions:        s.ResourcePredictions,
		ResourcePredictionsProfile: s.ResourcePredictionsProfile,
	}
}

// StorageProfile - The storage profile of the VMSS.
type StorageProfile struct {
	// A list of empty data disks to attach.
	DataDisks []*DataDisk

	// The Azure SKU name of the machines in the pool.
	OSDiskStorageAccountType *OsDiskStorageAccountType
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// VmssFabricProfile - The agents will run on Virtual Machine Scale Sets.
type VmssFabricProfile struct {
	// REQUIRED; The VM images of the machines in the pool.
	Images []*PoolImage

	// REQUIRED; Discriminator property for FabricProfile.
	Kind *string

	// REQUIRED; The Azure SKU of the machines in the pool.
	SKU *DevOpsAzureSKU

	// The network profile of the machines in the pool.
	NetworkProfile *NetworkProfile

	// The OS profile of the machines in the pool.
	OSProfile *OsProfile

	// The storage profile of the machines in the pool.
	StorageProfile *StorageProfile
}

// GetFabricProfile implements the FabricProfileClassification interface for type VmssFabricProfile.
func (v *VmssFabricProfile) GetFabricProfile() *FabricProfile {
	return &FabricProfile{
		Kind: v.Kind,
	}
}
