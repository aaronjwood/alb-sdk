// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// ConfigPbAttributes config pb attributes
// swagger:model ConfigPbAttributes
type ConfigPbAttributes struct {

	// Protobuf version number. Gets incremented if there is se Diff of federated diff in config pbs.This field will be a monotonically increasing number indicating the number of Config Update operations. Field introduced in 21.1.1.
	Version *int32 `json:"version,omitempty"`
}
