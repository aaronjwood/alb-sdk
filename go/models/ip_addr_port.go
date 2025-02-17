// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// IPAddrPort Ip addr port
// swagger:model IpAddrPort
type IPAddrPort struct {

	// Hostname of server. One of IP address or hostname should be set.
	Hostname *string `json:"hostname,omitempty"`

	// IP Address of host. One of IP address or hostname should be set.
	IP *IPAddr `json:"ip,omitempty"`

	// Name of the object.
	Name *string `json:"name,omitempty"`

	// Port number of server. Allowed values are 1-65535.
	// Required: true
	Port *int32 `json:"port"`
}
