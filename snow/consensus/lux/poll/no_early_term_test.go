<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
=======
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
>>>>>>> 53a8245a8 (Update consensus)
=======
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
=======
// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
>>>>>>> 34554f662 (Update LICENSE)
>>>>>>> c5eafdb72 (Update LICENSE)
=======
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
>>>>>>> 8fb2bec88 (Must keep bloodline pure)
// See the file LICENSE for licensing terms.

package poll

import (
	"testing"

<<<<<<< HEAD
	"github.com/luxdefi/node/ids"
=======
	"github.com/luxdefi/node/ids"
>>>>>>> 53a8245a8 (Update consensus)
)

func TestNoEarlyTermResults(t *testing.T) {
	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	vdr1 := ids.NodeID{1} // k = 1

	vdrs := ids.NodeIDBag{}
	vdrs.Add(vdr1)

	factory := NewNoEarlyTermFactory()
	poll := factory.New(vdrs)

	poll.Vote(vdr1, votes)
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate after receiving k votes")
	}

	result := poll.Result()
	if list := result.List(); len(list) != 1 {
		t.Fatalf("Wrong number of vertices returned")
	} else if retVtxID := list[0]; retVtxID != vtxID {
		t.Fatalf("Wrong vertex returned")
	} else if set := result.GetSet(vtxID); set.Len() != 1 {
		t.Fatalf("Wrong number of votes returned")
	}
}

func TestNoEarlyTermString(t *testing.T) {
	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	vdr1 := ids.NodeID{1}
	vdr2 := ids.NodeID{2} // k = 2

	vdrs := ids.NodeIDBag{}
	vdrs.Add(
		vdr1,
		vdr2,
	)

	factory := NewNoEarlyTermFactory()
	poll := factory.New(vdrs)

	poll.Vote(vdr1, votes)

	expected := `waiting on Bag: (Size = 1)
    ID[NodeID-BaMPFdqMUQ46BV8iRcwbVfsam55kMqcp]: Count = 1
received UniqueBag: (Size = 1)
    ID[SYXsAycDPUu4z2ZksJD5fh5nTDcH3vCFHnpcVye5XuJ2jArg]: Members = 0000000000000002`
	if result := poll.String(); expected != result {
		t.Fatalf("Poll should have returned %s but returned %s", expected, result)
	}
}

func TestNoEarlyTermDropsDuplicatedVotes(t *testing.T) {
	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	vdr1 := ids.NodeID{1}
	vdr2 := ids.NodeID{2} // k = 2

	vdrs := ids.NodeIDBag{}
	vdrs.Add(
		vdr1,
		vdr2,
	)

	factory := NewNoEarlyTermFactory()
	poll := factory.New(vdrs)

	poll.Vote(vdr1, votes)
	if poll.Finished() {
		t.Fatalf("Poll finished after less than alpha votes")
	}
	poll.Vote(vdr1, votes)
	if poll.Finished() {
		t.Fatalf("Poll finished after getting a duplicated vote")
	}
	poll.Vote(vdr2, votes)
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate after receiving k votes")
	}
}
