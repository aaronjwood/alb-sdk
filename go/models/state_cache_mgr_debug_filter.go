// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// StateCacheMgrDebugFilter state cache mgr debug filter
// swagger:model StateCacheMgrDebugFilter
type StateCacheMgrDebugFilter struct {

	// Pool UUID. It is a reference to an object of type Pool.
	PoolRef *string `json:"pool_ref,omitempty"`

	// VirtualService UUID. It is a reference to an object of type VirtualService.
	VsRef *string `json:"vs_ref,omitempty"`
}
