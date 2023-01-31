// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/ava-labs/avalanchego/utils/set"
=======
>>>>>>> d6c7e2094 (Track subnet uptimes (#1427))
=======
	"github.com/ava-labs/avalanchego/utils/set"
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var _ validatorUptimes = (*uptimes)(nil)

type uptimeAndReward struct {
	UpDuration      time.Duration `serialize:"true"`
	LastUpdated     uint64        `serialize:"true"` // Unix time in seconds
	PotentialReward uint64        `serialize:"true"`

	txID        ids.ID
	lastUpdated time.Time
}

type validatorUptimes interface {
	// LoadUptime sets the uptime measurements of [vdrID] on [subnetID] to
	// [uptime]. GetUptime and SetUptime will return an error if the [vdrID] and
	// [subnetID] hasn't been loaded. This call will not result in a write to disk.
	LoadUptime(
		vdrID ids.NodeID,
		subnetID ids.ID,
		uptime *uptimeAndReward,
	)

	// GetUptime returns the current uptime measurements of [vdrID] on
	// [subnetID].
	GetUptime(
		vdrID ids.NodeID,
		subnetID ids.ID,
	) (upDuration time.Duration, lastUpdated time.Time, err error)

	// SetUptime updates the uptime measurements of [vdrID] on [subnetID].
	// Unless these measurements are deleted first, the next call to
	// WriteUptimes will write this update to disk.
	SetUptime(
		vdrID ids.NodeID,
		subnetID ids.ID,
		upDuration time.Duration,
		lastUpdated time.Time,
	) error

	// DeleteUptime removes in-memory references to the uptimes measurements of
	// [vdrID] on [subnetID]. If there were staged updates from a prior call to
	// SetUptime, the updates will be dropped. This call will not result in a
	// write to disk.
	DeleteUptime(vdrID ids.NodeID, subnetID ids.ID)

	// WriteUptimes writes all staged updates from a prior call to SetUptime.
	WriteUptimes(
		dbPrimary database.KeyValueWriter,
		dbSubnet database.KeyValueWriter,
	) error
}

type uptimes struct {
	uptimes map[ids.NodeID]map[ids.ID]*uptimeAndReward // vdrID -> subnetID -> uptimes
	// updatedUptimes tracks the updates since the last call to WriteUptimes
<<<<<<< HEAD
<<<<<<< HEAD
	updatedUptimes map[ids.NodeID]set.Set[ids.ID] // vdrID -> subnetIDs
=======
	updatedUptimes map[ids.NodeID]map[ids.ID]struct{} // vdrID -> subnetID -> nil
>>>>>>> d6c7e2094 (Track subnet uptimes (#1427))
=======
	updatedUptimes map[ids.NodeID]set.Set[ids.ID] // vdrID -> subnetIDs
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
}

func newValidatorUptimes() validatorUptimes {
	return &uptimes{
		uptimes:        make(map[ids.NodeID]map[ids.ID]*uptimeAndReward),
<<<<<<< HEAD
<<<<<<< HEAD
		updatedUptimes: make(map[ids.NodeID]set.Set[ids.ID]),
=======
		updatedUptimes: make(map[ids.NodeID]map[ids.ID]struct{}),
>>>>>>> d6c7e2094 (Track subnet uptimes (#1427))
=======
		updatedUptimes: make(map[ids.NodeID]set.Set[ids.ID]),
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
	}
}

func (u *uptimes) LoadUptime(
	vdrID ids.NodeID,
	subnetID ids.ID,
	uptime *uptimeAndReward,
) {
	subnetUptimes, ok := u.uptimes[vdrID]
	if !ok {
		subnetUptimes = make(map[ids.ID]*uptimeAndReward)
		u.uptimes[vdrID] = subnetUptimes
	}
	subnetUptimes[subnetID] = uptime
}

func (u *uptimes) GetUptime(
	vdrID ids.NodeID,
	subnetID ids.ID,
) (upDuration time.Duration, lastUpdated time.Time, err error) {
	uptime, exists := u.uptimes[vdrID][subnetID]
	if !exists {
		return 0, time.Time{}, database.ErrNotFound
	}
	return uptime.UpDuration, uptime.lastUpdated, nil
}

func (u *uptimes) SetUptime(
	vdrID ids.NodeID,
	subnetID ids.ID,
	upDuration time.Duration,
	lastUpdated time.Time,
) error {
	uptime, exists := u.uptimes[vdrID][subnetID]
	if !exists {
		return database.ErrNotFound
	}
	uptime.UpDuration = upDuration
	uptime.lastUpdated = lastUpdated

	updatedSubnetUptimes, ok := u.updatedUptimes[vdrID]
	if !ok {
<<<<<<< HEAD
<<<<<<< HEAD
		updatedSubnetUptimes = set.Set[ids.ID]{}
		u.updatedUptimes[vdrID] = updatedSubnetUptimes
	}
	updatedSubnetUptimes.Add(subnetID)
=======
		updatedSubnetUptimes = make(map[ids.ID]struct{})
		u.updatedUptimes[vdrID] = updatedSubnetUptimes
	}
	updatedSubnetUptimes[subnetID] = struct{}{}
>>>>>>> d6c7e2094 (Track subnet uptimes (#1427))
=======
		updatedSubnetUptimes = set.Set[ids.ID]{}
		u.updatedUptimes[vdrID] = updatedSubnetUptimes
	}
	updatedSubnetUptimes.Add(subnetID)
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
	return nil
}

func (u *uptimes) DeleteUptime(vdrID ids.NodeID, subnetID ids.ID) {
	subnetUptimes := u.uptimes[vdrID]
	delete(subnetUptimes, subnetID)
	if len(subnetUptimes) == 0 {
		delete(u.uptimes, vdrID)
	}

	subnetUpdatedUptimes := u.updatedUptimes[vdrID]
	delete(subnetUpdatedUptimes, subnetID)
	if len(subnetUpdatedUptimes) == 0 {
		delete(u.updatedUptimes, vdrID)
	}
}

func (u *uptimes) WriteUptimes(
	dbPrimary database.KeyValueWriter,
	dbSubnet database.KeyValueWriter,
) error {
	for vdrID, updatedSubnets := range u.updatedUptimes {
		for subnetID := range updatedSubnets {
			uptime := u.uptimes[vdrID][subnetID]
			uptime.LastUpdated = uint64(uptime.lastUpdated.Unix())

			uptimeBytes, err := genesis.Codec.Marshal(txs.Version, uptime)
			if err != nil {
				return err
			}
			db := dbSubnet
			if subnetID == constants.PrimaryNetworkID {
				db = dbPrimary
			}
			if err := db.Put(uptime.txID[:], uptimeBytes); err != nil {
				return err
			}
		}
		delete(u.updatedUptimes, vdrID)
	}
	return nil
}
