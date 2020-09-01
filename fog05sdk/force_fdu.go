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
	//ENV is environment configuration kind
	ENV string = "ENV"

	//VIRTUAL is interface kind (default)
	VIRTUAL string = "VIRTUAL"

	//VIRTIO is VIRTIO interface kind (default)
	VIRTIO string = "VIRTIO"

	// Scaling Metrics

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
	ProbeInterval        float32 `json:"probe_interval"`
	MinReplicas          uint8   `json:"min_replicas"`
	MaxReplicas          uint8   `json:"max_replicas"`
}

// FOrCEFDUComputationalRequirements represents the FDU Computational Requirements aka Flavor for FOrCE
type FOrCEFDUComputationalRequirements struct {
	Name            *string `json:"name,omitempty"`
	UUID            *string `json:"uuid,omitempty"`
	CPUArch         string  `json:"cpu_arch"`
	CPUMinFrequency int     `json:"cpu_min_freq"`             //default 0
	CPUMinCount     int     `json:"cpu_min_count"`            //default 1
	GPUMinCount     *int    `json:"gpu_min_count,omitempty"`  //default 0
	FPGAMinCount    *int    `json:"fpga_min_count,omitempty"` //default 0
	RAMSizeMB       uint32  `json:"ram_size_mb"`
	StorageSizeMB   uint32  `json:"storage_size_mb"`
	// DutyCycle       *float64 `json:"duty_cycle,omitempty"`
}

// FOrCEFDUConfiguration represents the FDU Configuration
type FOrCEFDUConfiguration struct {
	ConfType string    `json:"conf_type"`
	Script   *string   `json:"script,omitempty"`
	Env      *[]string `json:"env,omitempty"` //[VAR=VALUE,VAR2=VALUE2,...]
	SSHKeys  []string  `json:"ssh_keys,omitempty"`
}

// FOrcEFDUVirtualInterface represents the FDU Virtual Interface
type FOrcEFDUVirtualInterface struct {
	InterfaceKind string  `json:"vif_kind"`
	Parent        *string `json:"parent,omitempty"` //PCI address, bridge name, interface name
	Bandwidth     *uint8  `json:"bandwidth,omitempty"`
}

// FOrcEFDUInterfaceDescriptor represent and FDU Network Interface descriptor
type FOrcEFDUInterfaceDescriptor struct {
	Name             string                   `json:"name"`
	Kind             string                   `json:"kind"`
	MACAddress       *string                  `json:"mac_address,omitempty"`
	VirtualInterface FOrcEFDUVirtualInterface `json:"virtual_interface"`
	CPID             *string                  `json:"cp_id,omitempty"`
}

// FOrcEConnectionPointDescriptor represents a Connection Point
type FOrcEConnectionPointDescriptor struct {
	UUID   *string `json:"uuid,omitempty"`
	Name   string  `json:"name"`
	ID     string  `json:"id"`
	VLDRef string  `json:"vld_ref"`
}

// FOrcEFDUStorageDescriptor represents an FDU Storage Descriptor
type FOrcEFDUStorageDescriptor struct {
	ID          string `json:"id"`
	StorageType string `json:"storage_type"`
	Size        int    `json:"size"`
	// FileSystemProtocol *string `json:"file_system_protocol,omitempty"`
	// CPID               *string `json:"cp_id,omitempty"`
}

// FOrCEFDUDescriptor represent and FDU descriptor
type FOrCEFDUDescriptor struct {
	UUID                     *string                           `json:"uuid,omitempty"`
	ID                       string                            `json:"id"`
	Name                     string                            `json:"name"`
	Version                  string                            `json:"version"`
	FDUVersion               string                            `json:"fdu_version"`
	Description              *string                           `json:"description,omitempty"`
	Hypervisor               string                            `json:"hypervisor"`
	Image                    *FDUImage                         `json:"image,omitempty"`
	HypervisorSpecific       *string                           `json:"hypervisor_specific,omitempty"`
	ComputationRequirements  FOrCEFDUComputationalRequirements `json:"computation_requirements"`
	GeographicalRequirements *FDUGeographicalRequirements      `json:"geographical_requirements,omitempty"`
	Interfaces               []FOrcEFDUInterfaceDescriptor     `json:"interfaces"`
	Storage                  []FOrcEFDUStorageDescriptor       `json:"storage"`
	ConnectionPoints         []FOrcEConnectionPointDescriptor  `json:"connection_points"`
	Configuration            *FOrCEFDUConfiguration            `json:"configuration,omitempty"`
	MigrationKind            string                            `json:"migration_kind"`
	ScalingPolicies          *[]ScalingPolicy                  `json:"scaling_policies,omitempty"`
	DependsOn                []string                          `json:"depends_on"`
}

//FOrcEFDURecord represent a record
type FOrcEFDURecord struct {
	UUID   string `json:"uuid"`
	ID     string `json:"id"` //ref to FOrCEFDUDescriptor.UUID
	status string `json:"status"`
}

/*
// FDUStorageRecord represent an FDU Storage Record
type FDUStorageRecord struct {
	UUID               string  `json:"uuid"`
	StorageID          string  `json:"storage_id"`
	StorageType        string  `json:"storage_type"`
	Size               int     `json:"size"`
	FileSystemProtocol *string `json:"file_system_protocol,omitempty"`
	CPID               *string `json:"cp_id,omitempty"`
}

// FDUInterfaceRecord represent an FDU Interface Record
type FDUInterfaceRecord struct {
	Name                 string               `json:"name"`
	IsMGMT               bool                 `json:"is_mgmt"`
	InterfaceType        string               `json:"if_type"`
	MACAddress           *string              `json:"mac_address,omitempty"`
	VirtualInterface     *FDUVirtualInterface `json:"virtual_interface,omitempty"`
	CPID                 *string              `json:"cp_id,omitempty"`
	ExtCPID              *string              `json:"ext_cp_id,omitempty"`
	VirtualInterfaceName string               `json:"vintf_name"`
	Status               string               `json:"status"`
	PhysicalFace         *string              `json:"phy_face,omitempty"`
	VEthFaceName         *string              `json:"veth_face_name,omitempty"`
	Properties           *jsont               `json:"properties,omitempty"`
}

// FDUMigrationProperties represent FDU Migration Properties used during migration
type FDUMigrationProperties struct {
	Destination string `json:"destination"`
	Source      string `json:"source"`
}

// FDURecord represent an FDU instance record
type FDURecord struct {
	UUID                     string                       `json:"uuid"`
	FDUID                    string                       `json:"fdu_id"`
	Status                   string                       `json:"status"`
	Image                    *FDUImage                    `json:"image,omitempty"`
	Command                  *FDUCommand                  `json:"command,omitempty"`
	Storage                  []FDUStorageRecord           `json:"storage"`
	ComputationRequirements  FDUComputationalRequirements `json:"computation_requirements"`
	GeographicalRequirements *FDUGeographicalRequirements `json:"geographical_requirements,omitempty"`
	EnergyRequirements       *FDUEnergyRequirements       `json:"energy_requirements,omitempty"`
	Hypervisor               string                       `json:"hypervisor"`
	MigrationKind            string                       `json:"migration_kind"`
	Configuration            *FDUConfiguration            `json:"configuration,omitempty"`
	Interfaces               *[]FDUInterfaceRecord        `json:"interfaces,omitempty"`
	IOPorts                  []FDUIOPort                  `json:"io_ports"`
	ConnectionPoints         []ConnectionPointRecord      `json:"connection_points"`
	DependsOn                []string                     `json:"depends_on"`
	ErrorCode                *int                         `json:"error_code,omitempty"`
	ErrorMsg                 *string                      `json:"error_msg,omitempty"`
	MigrationProperties      *FDUMigrationProperties      `json:"migration_properties,omitempty"`
	HypervisorInfo           *jsont                       `json:"hypervisor_info,omitempty"`
}
*/
