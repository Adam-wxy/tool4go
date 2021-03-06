// Copyright 2019 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !jsoniter

package json

import "encoding/json"

// PKG package name imported
const PKG = "encoding/json"

var (
	// Marshal is exported by tool4go/json package.
	Marshal = json.Marshal

	// Unmarshal is exported by tool4go/json package.
	Unmarshal = json.Unmarshal

	// MarshalIndent is exported by tool4go/json package.
	MarshalIndent = json.MarshalIndent

	// NewEncoder is exported by tool4go/json package.
	NewEncoder = json.NewEncoder

	// NewDecoder is exported by tool4go/json package.
	NewDecoder = json.NewDecoder

	// Valid is exported by tool4go/json package.
	Valid = json.Valid
)
