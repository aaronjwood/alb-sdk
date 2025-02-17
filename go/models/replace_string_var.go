// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// ReplaceStringVar replace *string var
// swagger:model ReplaceStringVar
type ReplaceStringVar struct {

	// Type of replacement *string - can be a variable exposed from datascript, value of an HTTP header or a custom user-input literal string. Enum options - DATASCRIPT_VAR, HTTP_HEADER_VAR, LITERAL_STRING.
	Type *string `json:"type,omitempty"`

	// Value of the replacement *string - name of variable exposed from datascript, name of the HTTP header or a custom user-input literal string.
	Val *string `json:"val,omitempty"`
}
