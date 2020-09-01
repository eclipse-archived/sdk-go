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
	// LIVE is Live Migration kind
	LIVE string = "LIVE"

	// COLD is cold Migration kind
	COLD string = "COLD"

	//BARE is native FDU kind
	BARE string = "BARE"

	//KVM is KVM VM FDU kind
	KVM string = "KVM"

	//KVMUK is KVM unikernel FDU kind
	KVMUK string = "KVM_UK"

	//XEN is XEN VM FDU kind
	XEN string = "XEN"

	//XENUK is XEN Unikernel FDU kind
	XENUK string = "XEN_UK"

	//LXD is LXD Container FDU kind
	LXD string = "LXD"

	//DOCKER is OCI (containerd) FDU kind
	DOCKER string = "DOCKER"

	//MCU is microcontroller FDU kind
	MCU string = "MCU"

	//SCRIPT is script configuration kind
	SCRIPT string = "SCRIPT"

	//CLOUDINIT is cloud init configuration kind
	CLOUDINIT string = "CLOUD_INIT"

	//INTERNAL is internal interface kind
	INTERNAL string = "INTERNAL"

	//EXTERNAL is external interface kind
	EXTERNAL string = "EXTERNAL"

	//WLAN is WLAN interface kind
	WLAN string = "WLAN"

	//BLUETOOTH is Bluetooth interface kind
	BLUETOOTH string = "BLUETOOTH"

	//PARAVIRT is paravirtualised interface kind
	PARAVIRT string = "PARAVIRT"

	//FOSMGMT is fog05 management interface kind
	FOSMGMT string = "FOS_MGMT"

	//PCIPASSTHROUGH is PCI passthrough interface kind
	PCIPASSTHROUGH string = "PCI_PASSTHROUGH"

	//SRIOV is SR-IOV interface kind
	SRIOV string = "SR_IOV"

	//E1000 is e1000 interface kind
	E1000 string = "E1000"

	//RTL8139 is rtl8139 interface kind
	RTL8139 string = "RTL8139"

	//PHYSICAL  is physical interface kind
	PHYSICAL string = "PHYSICAL"

	//BRIDGED is bridged interface kind
	BRIDGED string = "BRIDGED"

	//GPIO is GPIO port kind
	GPIO string = "GPIO"

	//I2C is I2C port kind
	I2C string = "I2C"

	//BUS is BUS port kind
	BUS string = "BUS"

	//COM is COM port kind
	COM string = "COM"

	//CAN is CAN port kind
	CAN string = "CAN"

	BLOCK  string = "BLOCK"
	FILE   string = "FILE"
	OBJECT string = "OBJECT"

	DEFINE    string = "DEFINE"
	CONFIGURE string = "CONFIGURE"
	CLEAN     string = "CLEAN"
	RUN       string = "RUN"
	STARTING  string = "STARTING"
	STOP      string = "STOP"
	RESUME    string = "RESUME"
	PAUSE     string = "PAUSE"
	SCALE     string = "SCALE"
	TAKEOFF   string = "TAKE_OFF"
	LAND      string = "LAND"
	MIGRATE   string = "MIGRATE"
	UNDEFINE  string = "UNDEFINE"
	ERROR     string = "ERROR"
)

// FDUImage represents an FDU image
type FDUImage struct {
	UUID     *string `json:"uuid,omitempty"`
	Name     *string `json:"name,omitempty"`
	URI      string  `json:"uri"`
	Checksum string  `json:"checksum"` //SHA256SUM
	Format   string  `json:"format"`
}

// FDUCommand represents an FDU command in case of Native FDU
type FDUCommand struct {
	Binary string   `json:"binary"`
	Args   []string `json:"args"`
}

// FDUGeographicalRequirements represents the FDU Geographical Requirements
type FDUGeographicalRequirements struct {
	Position  *FDUPosition  `json:"position,omitempty"`
	Proximity *FDUProximity `json:"proximity,omitempty"`
}

// FDUPosition represents the FDU Position
type FDUPosition struct {
	Latitude  string  `json:"lat"`
	Longitude string  `json:"lon"`
	Radius    float64 `json:"radius"`
}

// FDUProximity represents the FDU Proximity
type FDUProximity struct {
	Neighbor string  `json:"neighbour"`
	Radius   float64 `json:"radius"`
}

// FDUEnergyRequirements represents the FDU Energy Requirements
type FDUEnergyRequirements struct {
	Key string `json:"key"`
}

// FDUComputationalRequirements represents the FDU Computational Requirements aka Flavor
type FDUComputationalRequirements struct {
	Name            *string  `json:"name,omitempty"`
	UUID            *string  `json:"uuid,omitempty"`
	CPUArch         string   `json:"cpu_arch"`
	CPUMinFrequency int      `json:"cpu_min_freq"`
	CPUMinCount     int      `json:"cpu_min_count"`
	GPUMinCount     *int     `json:"gpu_min_count,omitempty"`
	FPGAMinCount    *int     `json:"fpga_min_count,omitempty"`
	RAMSizeMB       float64  `json:"ram_size_mb"`
	StorageSizeGB   float64  `json:"storage_size_gb"`
	DutyCycle       *float64 `json:"duty_cycle,omitempty"`
}

// FDUConfiguration represents the FDU Configuration
type FDUConfiguration struct {
	ConfType string   `json:"conf_type"`
	Script   string   `json:"script"`
	SSHKeys  []string `json:"ssh_keys,omitempty"`
}

// FDUVirtualInterface represents the FDU Virtual Interface
type FDUVirtualInterface struct {
	InterfaceType string `json:"intf_type"`
	VPCI          string `json:"vpci"`
	Bandwidth     int    `json:"bandwidth"`
}

// FDUIOPort represenrs an FDU IO Port requirements
type FDUIOPort struct {
	Address    string `json:"address"`
	IOKind     string `json:"io_kind"`
	MinIOPorts int    `json:"min_io_ports"`
}

// FDUInterfaceDescriptor represent and FDU Network Interface descriptor
type FDUInterfaceDescriptor struct {
	Name             string              `json:"name"`
	IsMGMT           bool                `json:"is_mgmt"`
	InterfaceType    string              `json:"if_type"`
	MACAddress       *string             `json:"mac_address,omitempty"`
	VirtualInterface FDUVirtualInterface `json:"virtual_interface"`
	CPID             *string             `json:"cp_id,omitempty"`
	ExtCPID          *string             `json:"ext_cp_id,omitempty"`
}

// FDUStorageDescriptor represents an FDU Storage Descriptor
type FDUStorageDescriptor struct {
	ID                 string  `json:"id"`
	StorageType        string  `json:"storage_type"`
	Size               int     `json:"size"`
	FileSystemProtocol *string `json:"file_system_protocol,omitempty"`
	CPID               *string `json:"cp_id,omitempty"`
}

// FDU represent and FDU descriptor
type FDU struct {
	ID                       string                       `json:"id"`
	Name                     string                       `json:"name"`
	UUID                     *string                      `json:"uuid,omitempty"`
	Description              *string                      `json:"description,omitempty"`
	ComputationRequirements  FDUComputationalRequirements `json:"computation_requirements"`
	Image                    *FDUImage                    `json:"image,omitempty"`
	Command                  *FDUCommand                  `json:"command,omitempty"`
	Storage                  []FDUStorageDescriptor       `json:"storage"`
	GeographicalRequirements *FDUGeographicalRequirements `json:"geographical_requirements,omitempty"`
	EnergyRequirements       *FDUEnergyRequirements       `json:"energy_requirements,omitempty"`
	Hypervisor               string                       `json:"hypervisor"`
	MigrationKind            string                       `json:"migration_kind"`
	Configuration            *FDUConfiguration            `json:"configuration,omitempty"`
	Interfaces               []FDUInterfaceDescriptor     `json:"interfaces"`
	IOPorts                  []FDUIOPort                  `json:"io_ports"`
	ConnectionPoints         []ConnectionPointDescriptor  `json:"connection_points"`
	DependsOn                []string                     `json:"depends_on"`
}

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
