//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcomputeschedule

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// ScheduledActionsClientVirtualMachinesCancelOperationsResponse contains the response from method ScheduledActionsClient.VirtualMachinesCancelOperations.
type ScheduledActionsClientVirtualMachinesCancelOperationsResponse struct {
	// This is the response from a cancel operations request
	CancelOperationsResponse
}

// ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse contains the response from method ScheduledActionsClient.VirtualMachinesExecuteDeallocate.
type ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse struct {
	// The response from a deallocate request
	DeallocateResourceOperationResponse
}

// ScheduledActionsClientVirtualMachinesExecuteHibernateResponse contains the response from method ScheduledActionsClient.VirtualMachinesExecuteHibernate.
type ScheduledActionsClientVirtualMachinesExecuteHibernateResponse struct {
	// The response from a Hibernate request
	HibernateResourceOperationResponse
}

// ScheduledActionsClientVirtualMachinesExecuteStartResponse contains the response from method ScheduledActionsClient.VirtualMachinesExecuteStart.
type ScheduledActionsClientVirtualMachinesExecuteStartResponse struct {
	// The response from a start request
	StartResourceOperationResponse
}

// ScheduledActionsClientVirtualMachinesGetOperationErrorsResponse contains the response from method ScheduledActionsClient.VirtualMachinesGetOperationErrors.
type ScheduledActionsClientVirtualMachinesGetOperationErrorsResponse struct {
	// This is the response from a get operations errors request
	GetOperationErrorsResponse
}

// ScheduledActionsClientVirtualMachinesGetOperationStatusResponse contains the response from method ScheduledActionsClient.VirtualMachinesGetOperationStatus.
type ScheduledActionsClientVirtualMachinesGetOperationStatusResponse struct {
	// This is the response from a get operations status request
	GetOperationStatusResponse
}

// ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse contains the response from method ScheduledActionsClient.VirtualMachinesSubmitDeallocate.
type ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse struct {
	// The response from a deallocate request
	DeallocateResourceOperationResponse
}

// ScheduledActionsClientVirtualMachinesSubmitHibernateResponse contains the response from method ScheduledActionsClient.VirtualMachinesSubmitHibernate.
type ScheduledActionsClientVirtualMachinesSubmitHibernateResponse struct {
	// The response from a Hibernate request
	HibernateResourceOperationResponse
}

// ScheduledActionsClientVirtualMachinesSubmitStartResponse contains the response from method ScheduledActionsClient.VirtualMachinesSubmitStart.
type ScheduledActionsClientVirtualMachinesSubmitStartResponse struct {
	// The response from a start request
	StartResourceOperationResponse
}
