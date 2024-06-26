//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

import "encoding/json"

func unmarshalClusterAvailableUpgradePropertiesClassification(rawMsg json.RawMessage) (ClusterAvailableUpgradePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClusterAvailableUpgradePropertiesClassification
	switch m["upgradeType"] {
	case string(ClusterAvailableUpgradeTypeAKSPatchUpgrade):
		b = &ClusterAvailableUpgradeAksPatchUpgradeProperties{}
	case string(ClusterAvailableUpgradeTypeHotfixUpgrade):
		b = &ClusterAvailableUpgradeHotfixUpgradeProperties{}
	default:
		b = &ClusterAvailableUpgradeProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalClusterJobPropertiesClassification(rawMsg json.RawMessage) (ClusterJobPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClusterJobPropertiesClassification
	switch m["jobType"] {
	case string(JobTypeFlinkJob):
		b = &FlinkJobProperties{}
	default:
		b = &ClusterJobProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalClusterPoolAvailableUpgradePropertiesClassification(rawMsg json.RawMessage) (ClusterPoolAvailableUpgradePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClusterPoolAvailableUpgradePropertiesClassification
	switch m["upgradeType"] {
	case string(ClusterPoolAvailableUpgradeTypeAKSPatchUpgrade):
		b = &ClusterPoolAvailableUpgradeAksPatchUpgradeProperties{}
	case string(ClusterPoolAvailableUpgradeTypeNodeOsUpgrade):
		b = &ClusterPoolAvailableUpgradeNodeOsUpgradeProperties{}
	default:
		b = &ClusterPoolAvailableUpgradeProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalClusterPoolUpgradePropertiesClassification(rawMsg json.RawMessage) (ClusterPoolUpgradePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClusterPoolUpgradePropertiesClassification
	switch m["upgradeType"] {
	case string(ClusterPoolUpgradeTypeAKSPatchUpgrade):
		b = &ClusterPoolAKSPatchVersionUpgradeProperties{}
	case string(ClusterPoolUpgradeTypeNodeOsUpgrade):
		b = &ClusterPoolNodeOsImageUpdateProperties{}
	default:
		b = &ClusterPoolUpgradeProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalClusterUpgradePropertiesClassification(rawMsg json.RawMessage) (ClusterUpgradePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClusterUpgradePropertiesClassification
	switch m["upgradeType"] {
	case string(ClusterUpgradeTypeAKSPatchUpgrade):
		b = &ClusterAKSPatchVersionUpgradeProperties{}
	case string(ClusterUpgradeTypeHotfixUpgrade):
		b = &ClusterHotfixUpgradeProperties{}
	default:
		b = &ClusterUpgradeProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
