// Copyright (C) 2019-2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package validator

import (
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/utils/constants"
)

// SubnetValidator validates a subnet on the Avalanche network.
type SubnetValidator struct {
	Validator `serialize:"true"`

	// ID of the subnet this validator is validating
	Subnet ids.ID `serialize:"true" json:"subnetID"`
}

// SubnetID is the ID of the subnet this validator is validating
func (v *SubnetValidator) SubnetID() ids.ID {
	return v.Subnet
}

// Verify this validator is valid
func (v *SubnetValidator) Verify() error {
	switch v.Subnet {
	case constants.PrimaryNetworkID:
		return errBadSubnetID
	default:
		return v.Validator.Verify()
	}
}
