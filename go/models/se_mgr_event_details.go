// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// SeMgrEventDetails se mgr event details
// swagger:model SeMgrEventDetails
type SeMgrEventDetails struct {

	// cloud_name of SeMgrEventDetails.
	CloudName *string `json:"cloud_name,omitempty"`

	// Unique object identifier of cloud.
	CloudUUID *string `json:"cloud_uuid,omitempty"`

	// enable_state of SeMgrEventDetails.
	EnableState *string `json:"enable_state,omitempty"`

	// Placeholder for description of property gcp_info of obj type SeMgrEventDetails field type str  type object
	GcpInfo *GcpInfo `json:"gcp_info,omitempty"`

	// host_name of SeMgrEventDetails.
	HostName *string `json:"host_name,omitempty"`

	// Unique object identifier of host.
	HostUUID *string `json:"host_uuid,omitempty"`

	// Number of memory.
	Memory *int32 `json:"memory,omitempty"`

	// migrate_state of SeMgrEventDetails.
	MigrateState *string `json:"migrate_state,omitempty"`

	// Name of the object.
	// Required: true
	Name *string `json:"name"`

	//  Field introduced in 20.1.3.
	NewMgmtIP *string `json:"new_mgmt_ip,omitempty"`

	//  Field introduced in 20.1.3.
	NewMgmtIp6 *string `json:"new_mgmt_ip6,omitempty"`

	//  Field introduced in 20.1.3.
	OldMgmtIP *string `json:"old_mgmt_ip,omitempty"`

	//  Field introduced in 20.1.3.
	OldMgmtIp6 *string `json:"old_mgmt_ip6,omitempty"`

	// reason of SeMgrEventDetails.
	Reason *string `json:"reason,omitempty"`

	// se_grp_name of SeMgrEventDetails.
	SeGrpName *string `json:"se_grp_name,omitempty"`

	// Unique object identifier of se_grp.
	SeGrpUUID *string `json:"se_grp_uuid,omitempty"`

	// Number of vcpus.
	Vcpus *int32 `json:"vcpus,omitempty"`

	// vs_name of SeMgrEventDetails.
	VsName []string `json:"vs_name,omitempty"`

	// Unique object identifier of vs.
	VsUUID []string `json:"vs_uuid,omitempty"`

	// vSphere HA on cluster enabled. Field introduced in 20.1.7, 21.1.3.
	VsphereHaEnabled *bool `json:"vsphere_ha_enabled,omitempty"`

	// This flag is set to true when Cloud Connector has detected an ESX host failure. This flag is set to false when the SE connects back to the controller, or when vSphere HA recovery timeout has occurred. Field introduced in 20.1.7, 21.1.3.
	VsphereHaInprogress *bool `json:"vsphere_ha_inprogress,omitempty"`
}
