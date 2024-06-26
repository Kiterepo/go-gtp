// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRATType creates a new RATType IE.
func NewRATType(rat uint8) *IE {
	return NewUint8IE(RATType, rat)
}

// RATType returns RATType in uint8 if the type of IE matches.
func (i *IE) RATType() (uint8, error) {
	switch i.Type {
	case RATType:
		return i.ValueAsUint8()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustRATType returns RATType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustRATType() uint8 {
	v, _ := i.RATType()
	return v
}
