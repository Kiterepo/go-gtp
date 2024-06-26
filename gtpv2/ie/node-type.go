// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewNodeType creates a new NodeType IE.
func NewNodeType(nodeType uint8) *IE {
	return NewUint8IE(NodeType, nodeType)
}

// NodeType returns NodeType in uint8 if the type of IE matches.
func (i *IE) NodeType() (uint8, error) {
	if i.Type != NodeType {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustNodeType returns NodeType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustNodeType() uint8 {
	v, _ := i.NodeType()
	return v
}
