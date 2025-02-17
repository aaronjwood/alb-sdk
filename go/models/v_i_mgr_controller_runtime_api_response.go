// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// VIMgrControllerRuntimeAPIResponse v i mgr controller runtime Api response
// swagger:model VIMgrControllerRuntimeApiResponse
type VIMgrControllerRuntimeAPIResponse struct {

	// count
	// Required: true
	Count *int32 `json:"count"`

	// next
	Next *string `json:"next,omitempty"`

	// results
	// Required: true
	Results []*VIMgrControllerRuntime `json:"results,omitempty"`
}
