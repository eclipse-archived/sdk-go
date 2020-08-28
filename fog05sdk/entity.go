/*
* Copyright (c) 2014,2019 Contributors to the Eclipse Foundation
* See the NOTICE file(s) distributed with this work for additional
* information regarding copyright ownership.
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors: Gabriele Baldoni, ADLINK Technology Inc.
* golang APIs
 */

package fog05sdk

const (
	// FIM is FIM EntityKind
	FIM string = "FIM"

	// CLOUD is CLOUD EntityKind
	CLOUD string = "CLOUD"

	// CPU is ScalingMetric
	CPU string = "CPU"

	//MEMORY is ScalingMetric
	MEMORY string = "MEMORY"

	//DISK is ScalingMetric
	DISK string = "DISK"

	//CUSTOM is ScalingMetric
	CUSTOM string = "CUSTOM"
)

// ScalingPolicy represents a scaling policy
type ScalingPolicy struct {
	Metric               string  `json:"metric"`
	ScaleUpThreshold     float32 `json:"scale_up_threshold"`
	ScaleDownThreshold   float32 `json:"scale_down_threshold"`
	ThresholdSensibility float32 `json:"threshold_sensibility"`
	MinReplicas          uint8   `json:"min_replicas"`
	MaxReplicas          uint8   `json:"max_replicas`
	FDUID                string  `json:"fdu_id"`
}

// ConstituentFDU is an FDU composing the entity
type ConstituentFDU struct {
	ID    string `json:"id"`
	Index uint8  `json:"index"`
}

// ConstituentFDU is an FDU composing the entity instance
type ConstituentRecordFDU struct {
	ID    string `json:"id"`
	UUID  string `json:"uuid"`
	Index uint8  `json:"index"`
}

// ConstituentVirtualLinkRecord is an virtual link composing the entity instance
type ConstituentVirtualLinkRecord struct {
	ID   string `json:"id"`
	UUID string `json:"uuid"`
}

// EntityDescriptor represent an Entity descriptor
type EntityDescriptor struct {
	UUID            *string           `json:"uuid,omitempty"`
	ID              string            `json:"id"`
	Version         string            `json:"version"`
	EntityVersion   string            `json:"entity_version"`
	Kind            string            `json:"kind"`
	Description     *string           `json:"description,omitempty`
	FDUs            *[]ConstituentFDU `json:"fdus,omitempty"`
	ScalingPolicies *[]ScalingPolicy  `json:"scaling_policies,omitempty`
	VirtualLinks    *[]VirtualNetwork `json:"virtual_links,omitempty"`
	CloudDescriptor *string           `json:"cloud_descriptor"`
}

// EntityRecord represent an Enitity instance record
type EntityRecord struct {
	UUID         string                          `json:"uuid"`
	ID           string                          `json:"id"`
	FDUs         *[]ConstituentRecordFDU         `json:"fdus,omitempty"`
	VirtualLinks *[]ConstituentVirtualLinkRecord `json:"virtual_links,omitempy`
	CloudRecord  *string                         `json:"cloud_record,omitempy"`
}
