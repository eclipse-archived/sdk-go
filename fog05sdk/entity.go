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

	//ONBOARDING is Entity status
	ONBOARDING string = "ONBOARDING"

	//ONBOARDED is Entity status
	ONBOARDED string = "ONBOARDED"

	//STARTING is Entity status
	//STARTING string = "STARTING"

	//RUNNING is Entity status
	RUNNING string = "RUNNING"

	//STOPPING is Entity status
	STOPPING string = "STOPPING"

	//STOPPED is Entity status
	STOPPED string = "STOPPED"

	//OFFLOADING is Entity status
	OFFLOADING string = "OFFLOADING"

	//OFFLOADED is Entity status
	OFFLOADED string = "OFFLOADED"

	//INVALID is Entity status
	INVALID string = "INVALID"

	//ERROR is Entity status
	//ERROR string = "ERROR"

	//RECOVERING is Entity status
	RECOVERING string = "RECOVERING"

	//L2 is L2 link kind (default) Multicast VXLAN
	L2 string = "L2"

	//L3 is L3 link kind (tree-based GRE)
	L3 string = "L3"

	// ELINE is Point-to-Point VXLAN
	ELINE string = "ELINE"

	// ELAN is ELAN multicast VXLAN
	ELAN string = "ELAN"
)

// IPConfiguration represent the IP Configuration
type IPConfiguration struct {
	Subnet    *string `json:"subnet,omitempty"`
	Gateway   *string `json:"gateway,omitempty"`
	DHCPRange *string `json:"dhcp_range,omitempty"`
	DNS       *string `json:"dns,omitempty"`
}

// VirtualLinkDescriptor represent a Virtual Link
type VirtualLinkDescriptor struct {
	UUID            *string          `json:"uuid,omitempty`
	ID              string           `json:"id"`
	Name            *string          `json:"name,omitempty"`
	IsMgmt          bool             `json:"is_mgmt"`
	LinkKind        string           `json:"link_kind"`
	IPVersion       string           `json:"ip_version"`
	IPConfiguration *IPConfiguration `json:"ip_configuration,omitempty"`
}

// EntityDescriptor represent an Entity descriptor
type EntityDescriptor struct {
	UUID          *string          `json:"uuid,omitempty"`
	ID            string           `json:"id"`
	Name          *string          `json:"name,omitempty"`
	Version       string           `json:"version"`
	EntityVersion string           `json:"entity_version"`
	Description   *string          `json:"description,omitempty"`
	FDUs          []FDU            `json:"fdus"`
	VirtualLinks  []VirtualNetwork `json:"virtual_links"`
}

// EntityRecord represent an Enitity instance record
type EntityRecord struct {
	UUID         string   `json:"uuid"`
	ID           string   `json:"id"`
	Status       string   `json:"status"`
	FDUs         []string `json:"fdus"`
	VirtualLinks []string `json:"virtual_links"`
}
