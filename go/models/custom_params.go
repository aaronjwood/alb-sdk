// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// CustomParams custom params
// swagger:model CustomParams
type CustomParams struct {

	// Placeholder for description of property is_dynamic of obj type CustomParams field type str  type boolean
	IsDynamic *bool `json:"is_dynamic,omitempty"`

	// Placeholder for description of property is_sensitive of obj type CustomParams field type str  type boolean
	IsSensitive *bool `json:"is_sensitive,omitempty"`

	// Name of the object.
	// Required: true
	Name *string `json:"name"`

	// value of CustomParams.
	Value *string `json:"value,omitempty"`
}
