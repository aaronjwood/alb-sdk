// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// Tenant tenant
// swagger:model Tenant
type Tenant struct {

	// UNIX time since epoch in microseconds. Units(MICROSECONDS).
	// Read Only: true
	LastModified *string `json:"_last_modified,omitempty"`

	// Placeholder for description of property config_settings of obj type Tenant field type str  type object
	ConfigSettings *TenantConfiguration `json:"config_settings,omitempty"`

	// Protobuf versioning for config pbs. Field introduced in 21.1.1.
	ConfigpbAttributes *ConfigPbAttributes `json:"configpb_attributes,omitempty"`

	// Creator of this tenant.
	CreatedBy *string `json:"created_by,omitempty"`

	// User defined description for the object.
	Description *string `json:"description,omitempty"`

	// The referred label groups are enforced on the tenant if this is set to true.If this is set to false, the label groups are suggested for the tenant. Field introduced in 20.1.5.
	EnforceLabelGroup *bool `json:"enforce_label_group,omitempty"`

	// The label_groups to be enforced on the tenant. This is strictly enforced only if enforce_label_group is set to True. It is a reference to an object of type LabelGroup. Field introduced in 20.1.5.
	LabelGroupRefs []string `json:"label_group_refs,omitempty"`

	// Placeholder for description of property local of obj type Tenant field type str  type boolean
	Local *bool `json:"local,omitempty"`

	// Name of the object.
	// Required: true
	Name *string `json:"name"`

	// Suggestive pool of key value pairs for recommending assignment of labels to objects in the User Interface. Every entry is unique in both key and value. Field deprecated in 20.1.5. Field introduced in 20.1.2. Maximum of 256 items allowed.
	SuggestedObjectLabels []*TenantLabel `json:"suggested_object_labels,omitempty"`

	// url
	// Read Only: true
	URL *string `json:"url,omitempty"`

	// Unique object identifier of the object.
	UUID *string `json:"uuid,omitempty"`
}
