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

	"github.com/stretchr/testify/require"

	"github.com/prometheus/client_golang/prometheus"

<<<<<<< HEAD
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/utils/logging"
	"github.com/luxdefi/node/utils/wrappers"
=======
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/utils/logging"
	"github.com/luxdefi/node/utils/wrappers"
>>>>>>> 53a8245a8 (Update consensus)
)

func TestNewSetErrorOnMetrics(t *testing.T) {
	factory := NewNoEarlyTermFactory()
	log := logging.NoLog{}
	namespace := ""
	registerer := prometheus.NewRegistry()

	errs := wrappers.Errs{}
	errs.Add(
		registerer.Register(prometheus.NewCounter(prometheus.CounterOpts{
			Name: "polls",
		})),
		registerer.Register(prometheus.NewCounter(prometheus.CounterOpts{
			Name: "poll_duration",
		})),
	)
	if errs.Errored() {
		t.Fatal(errs.Err)
	}

	if s := NewSet(factory, log, namespace, registerer); s == nil {
		t.Fatalf("shouldn't have errored due to metrics failures")
	}
}

func TestCreateAndFinishPoll(t *testing.T) {
	factory := NewNoEarlyTermFactory()
	log := logging.NoLog{}
	namespace := ""
	registerer := prometheus.NewRegistry()
	s := NewSet(factory, log, namespace, registerer)

	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	vdr1 := ids.NodeID{1}
	vdr2 := ids.NodeID{2} // k = 2

	vdrs := ids.NodeIDBag{}
	vdrs.Add(
		vdr1,
		vdr2,
	)

	if s.Len() != 0 {
		t.Fatalf("Shouldn't have any active polls yet")
	} else if !s.Add(0, vdrs) {
		t.Fatalf("Should have been able to add a new poll")
	} else if s.Len() != 1 {
		t.Fatalf("Should only have one active poll")
	} else if s.Add(0, vdrs) {
		t.Fatalf("Shouldn't have been able to add a duplicated poll")
	} else if s.Len() != 1 {
		t.Fatalf("Should only have one active poll")
	} else if results := s.Vote(1, vdr1, votes); len(results) > 0 {
		t.Fatalf("Shouldn't have been able to finish a non-existent poll")
<<<<<<< HEAD
	} else if results = s.Vote(0, vdr1, votes); len(results) > 0 {
		t.Fatalf("Shouldn't have been able to finish an ongoing poll")
	} else if results = s.Vote(0, vdr1, votes); len(results) > 0 {
		t.Fatalf("Should have dropped a duplicated poll")
	} else if results = s.Vote(0, vdr2, votes); len(results) == 0 {
=======
	} else if results := s.Vote(0, vdr1, votes); len(results) > 0 {
		t.Fatalf("Shouldn't have been able to finish an ongoing poll")
	} else if results := s.Vote(0, vdr1, votes); len(results) > 0 {
		t.Fatalf("Should have dropped a duplicated poll")
	} else if results := s.Vote(0, vdr2, votes); len(results) == 0 {
>>>>>>> 53a8245a8 (Update consensus)
		t.Fatalf("Should have finished the poll")
	} else if len(results) != 1 {
		t.Fatalf("Wrong number of results returned")
	} else if list := results[0].List(); len(list) != 1 {
		t.Fatalf("Wrong number of vertices returned")
	} else if retVtxID := list[0]; retVtxID != vtxID {
		t.Fatalf("Wrong vertex returned")
	} else if set := results[0].GetSet(vtxID); set.Len() != 2 {
		t.Fatalf("Wrong number of votes returned")
	}
}

func TestCreateAndFinishPollOutOfOrder_OlderFinishesFirst(t *testing.T) {
	factory := NewNoEarlyTermFactory()
	log := logging.NoLog{}
	namespace := ""
	registerer := prometheus.NewRegistry()
	s := NewSet(factory, log, namespace, registerer)

	// create validators
	vdr1 := ids.NodeID{1}
	vdr2 := ids.NodeID{2}
	vdr3 := ids.NodeID{3}

	vdrs := []ids.NodeID{vdr1, vdr2, vdr3}

	// create two polls for the two vtxs
	vdrBag := ids.NodeIDBag{}
	vdrBag.Add(vdrs...)
	added := s.Add(1, vdrBag)
	require.True(t, added)

	vdrBag = ids.NodeIDBag{}
	vdrBag.Add(vdrs...)
	added = s.Add(2, vdrBag)
	require.True(t, added)
	require.Equal(t, s.Len(), 2)

	// vote vtx1 for poll 1
	// vote vtx2 for poll 2
	vtx1 := ids.ID{1}
	vtx2 := ids.ID{2}

	var results []ids.UniqueBag

	// vote out of order
	results = s.Vote(1, vdr1, []ids.ID{vtx1})
	require.Len(t, results, 0)
	results = s.Vote(2, vdr2, []ids.ID{vtx2})
	require.Len(t, results, 0)
	results = s.Vote(2, vdr3, []ids.ID{vtx2})
	require.Len(t, results, 0)

	results = s.Vote(1, vdr2, []ids.ID{vtx1})
	require.Len(t, results, 0)

	results = s.Vote(1, vdr3, []ids.ID{vtx1}) // poll 1 finished, poll 2 still remaining
	require.Len(t, results, 1)                // because 1 is the oldest
	require.Equal(t, vtx1, results[0].List()[0])

	results = s.Vote(2, vdr1, []ids.ID{vtx2}) // poll 2 finished
	require.Len(t, results, 1)                // because 2 is the oldest now
	require.Equal(t, vtx2, results[0].List()[0])
}

func TestCreateAndFinishPollOutOfOrder_UnfinishedPollsGaps(t *testing.T) {
	factory := NewNoEarlyTermFactory()
	log := logging.NoLog{}
	namespace := ""
	registerer := prometheus.NewRegistry()
	s := NewSet(factory, log, namespace, registerer)

	// create validators
	vdr1 := ids.NodeID{1}
	vdr2 := ids.NodeID{2}
	vdr3 := ids.NodeID{3}

	vdrs := []ids.NodeID{vdr1, vdr2, vdr3}

	// create three polls for the two vtxs
	vdrBag := ids.NodeIDBag{}
	vdrBag.Add(vdrs...)
	added := s.Add(1, vdrBag)
	require.True(t, added)

	vdrBag = ids.NodeIDBag{}
	vdrBag.Add(vdrs...)
	added = s.Add(2, vdrBag)
	require.True(t, added)

	vdrBag = ids.NodeIDBag{}
	vdrBag.Add(vdrs...)
	added = s.Add(3, vdrBag)
	require.True(t, added)
	require.Equal(t, s.Len(), 3)

	// vote vtx1 for poll 1
	// vote vtx2 for poll 2
	// vote vtx3 for poll 3
	vtx1 := ids.ID{1}
	vtx2 := ids.ID{2}
	vtx3 := ids.ID{3}

	var results []ids.UniqueBag

	// vote out of order
	// 2 finishes first to create a gap of finished poll between two unfinished polls 1 and 3
	results = s.Vote(2, vdr3, []ids.ID{vtx2})
	require.Len(t, results, 0)
	results = s.Vote(2, vdr2, []ids.ID{vtx2})
	require.Len(t, results, 0)
	results = s.Vote(2, vdr1, []ids.ID{vtx2})
	require.Len(t, results, 0)

	// 3 finishes now, 2 has already finished but 1 is not finished so we expect to receive no results still
	results = s.Vote(3, vdr2, []ids.ID{vtx3})
	require.Len(t, results, 0)
	results = s.Vote(3, vdr3, []ids.ID{vtx3})
	require.Len(t, results, 0)
	results = s.Vote(3, vdr1, []ids.ID{vtx3})
	require.Len(t, results, 0)

	// 1 finishes now, 2 and 3 have already finished so we expect 3 items in results
	results = s.Vote(1, vdr1, []ids.ID{vtx1})
	require.Len(t, results, 0)
	results = s.Vote(1, vdr2, []ids.ID{vtx1})
	require.Len(t, results, 0)
	results = s.Vote(1, vdr3, []ids.ID{vtx1})
	require.Len(t, results, 3)
	require.Equal(t, vtx1.String(), results[0].List()[0].String())
	require.Equal(t, vtx2.String(), results[1].List()[0].String())
	require.Equal(t, vtx3.String(), results[2].List()[0].String())
}

func TestSetString(t *testing.T) {
	factory := NewNoEarlyTermFactory()
	log := logging.NoLog{}
	namespace := ""
	registerer := prometheus.NewRegistry()
	s := NewSet(factory, log, namespace, registerer)

	vdr1 := ids.NodeID{1} // k = 1

	vdrs := ids.NodeIDBag{}
	vdrs.Add(vdr1)

	expected := `current polls: (Size = 1)
    RequestID 0:
        waiting on Bag: (Size = 1)
            ID[NodeID-6HgC8KRBEhXYbF4riJyJFLSHt37UNuRt]: Count = 1
        received UniqueBag: (Size = 0)`
	if !s.Add(0, vdrs) {
		t.Fatalf("Should have been able to add a new poll")
	} else if str := s.String(); expected != str {
		t.Fatalf("Set return wrong string, Expected:\n%s\nReturned:\n%s",
			expected,
			str)
	}
}
