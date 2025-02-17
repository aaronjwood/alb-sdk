// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// DNSPolicy Dns policy
// swagger:model DnsPolicy
type DNSPolicy struct {

	// UNIX time since epoch in microseconds. Units(MICROSECONDS).
	// Read Only: true
	LastModified *string `json:"_last_modified,omitempty"`

	// Protobuf versioning for config pbs. Field introduced in 21.1.1.
	ConfigpbAttributes *ConfigPbAttributes `json:"configpb_attributes,omitempty"`

	// Creator name. Field introduced in 17.1.1.
	CreatedBy *string `json:"created_by,omitempty"`

	//  Field introduced in 17.1.1.
	Description *string `json:"description,omitempty"`

	// The DNS policy is created and modified by internal modules only. This should not be modified by users. Field introduced in 21.1.1.
	Internal *bool `json:"internal,omitempty"`

	// Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field deprecated in 20.1.5. Field introduced in 20.1.2. Maximum of 4 items allowed.
	Labels []*KeyValue `json:"labels,omitempty"`

	// List of labels to be used for granular RBAC. Field introduced in 20.1.5. Allowed in Basic edition, Essentials edition, Enterprise edition.
	Markers []*RoleFilterMatchLabel `json:"markers,omitempty"`

	// Name of the DNS Policy. Field introduced in 17.1.1.
	// Required: true
	Name *string `json:"name"`

	// DNS rules. Field introduced in 17.1.1.
	Rule []*DNSRule `json:"rule,omitempty"`

	//  It is a reference to an object of type Tenant. Field introduced in 17.1.1.
	TenantRef *string `json:"tenant_ref,omitempty"`

	// url
	// Read Only: true
	URL *string `json:"url,omitempty"`

	// UUID of the DNS Policy. Field introduced in 17.1.1.
	UUID *string `json:"uuid,omitempty"`
}
