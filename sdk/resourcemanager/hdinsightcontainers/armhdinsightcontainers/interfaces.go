//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

// ClusterAvailableUpgradePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetClusterAvailableUpgradeProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClusterAvailableUpgradeAksPatchUpgradeProperties, *ClusterAvailableUpgradeHotfixUpgradeProperties, *ClusterAvailableUpgradeProperties
type ClusterAvailableUpgradePropertiesClassification interface {
	// GetClusterAvailableUpgradeProperties returns the ClusterAvailableUpgradeProperties content of the underlying type.
	GetClusterAvailableUpgradeProperties() *ClusterAvailableUpgradeProperties
}

// ClusterJobPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetClusterJobProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClusterJobProperties, *FlinkJobProperties
type ClusterJobPropertiesClassification interface {
	// GetClusterJobProperties returns the ClusterJobProperties content of the underlying type.
	GetClusterJobProperties() *ClusterJobProperties
}

// ClusterPoolAvailableUpgradePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetClusterPoolAvailableUpgradeProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClusterPoolAvailableUpgradeAksPatchUpgradeProperties, *ClusterPoolAvailableUpgradeNodeOsUpgradeProperties, *ClusterPoolAvailableUpgradeProperties
type ClusterPoolAvailableUpgradePropertiesClassification interface {
	// GetClusterPoolAvailableUpgradeProperties returns the ClusterPoolAvailableUpgradeProperties content of the underlying type.
	GetClusterPoolAvailableUpgradeProperties() *ClusterPoolAvailableUpgradeProperties
}

// ClusterPoolUpgradePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetClusterPoolUpgradeProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClusterPoolAKSPatchVersionUpgradeProperties, *ClusterPoolNodeOsImageUpdateProperties, *ClusterPoolUpgradeProperties
type ClusterPoolUpgradePropertiesClassification interface {
	// GetClusterPoolUpgradeProperties returns the ClusterPoolUpgradeProperties content of the underlying type.
	GetClusterPoolUpgradeProperties() *ClusterPoolUpgradeProperties
}

// ClusterUpgradePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetClusterUpgradeProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClusterAKSPatchVersionUpgradeProperties, *ClusterHotfixUpgradeProperties, *ClusterUpgradeProperties
type ClusterUpgradePropertiesClassification interface {
	// GetClusterUpgradeProperties returns the ClusterUpgradeProperties content of the underlying type.
	GetClusterUpgradeProperties() *ClusterUpgradeProperties
}
