// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// ControllerDiscontinuousTimeChangeEventDetails controller discontinuous time change event details
// swagger:model ControllerDiscontinuousTimeChangeEventDetails
type ControllerDiscontinuousTimeChangeEventDetails struct {

	// Time stamp before the discontinuous jump in time.
	FromTime *string `json:"from_time,omitempty"`

	// Name of the Controller responsible for this event.
	NodeName *string `json:"node_name,omitempty"`

	// System Peer and Candidate NTP Servers active at the point of time jump.
	NtpServers *string `json:"ntp_servers,omitempty"`

	// Time stamp to which the time has discontinuously jumped.
	ToTime *string `json:"to_time,omitempty"`
}
