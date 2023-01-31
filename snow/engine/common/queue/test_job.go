// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package queue

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	errExecute                = errors.New("unexpectedly called Execute")
	errHasMissingDependencies = errors.New("unexpectedly called HasMissingDependencies")
)

// TestJob is a test Job
type TestJob struct {
	T *testing.T

	CantID,
	CantMissingDependencies,
	CantExecute,
	CantBytes,
	CantHasMissingDependencies bool

	IDF                     func() ids.ID
<<<<<<< HEAD
<<<<<<< HEAD
	MissingDependenciesF    func(context.Context) (set.Set[ids.ID], error)
=======
	MissingDependenciesF    func(context.Context) (ids.Set, error)
>>>>>>> 5be92660b (Pass message context through the VM interface (#2219))
=======
	MissingDependenciesF    func(context.Context) (set.Set[ids.ID], error)
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
	ExecuteF                func(context.Context) error
	BytesF                  func() []byte
	HasMissingDependenciesF func(context.Context) (bool, error)
}

func (j *TestJob) Default(cant bool) {
	j.CantID = cant
	j.CantMissingDependencies = cant
	j.CantExecute = cant
	j.CantBytes = cant
	j.CantHasMissingDependencies = cant
}

func (j *TestJob) ID() ids.ID {
	if j.IDF != nil {
		return j.IDF()
	}
	if j.CantID && j.T != nil {
		j.T.Fatalf("Unexpectedly called ID")
	}
	return ids.ID{}
}

<<<<<<< HEAD
<<<<<<< HEAD
func (j *TestJob) MissingDependencies(ctx context.Context) (set.Set[ids.ID], error) {
=======
func (j *TestJob) MissingDependencies(ctx context.Context) (ids.Set, error) {
>>>>>>> 5be92660b (Pass message context through the VM interface (#2219))
=======
func (j *TestJob) MissingDependencies(ctx context.Context) (set.Set[ids.ID], error) {
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
	if j.MissingDependenciesF != nil {
		return j.MissingDependenciesF(ctx)
	}
	if j.CantMissingDependencies && j.T != nil {
		j.T.Fatalf("Unexpectedly called MissingDependencies")
	}
	return set.Set[ids.ID]{}, nil
}

func (j *TestJob) Execute(ctx context.Context) error {
	if j.ExecuteF != nil {
		return j.ExecuteF(ctx)
	}
	if j.CantExecute && j.T != nil {
		j.T.Fatal(errExecute)
	}
	return errExecute
}

func (j *TestJob) Bytes() []byte {
	if j.BytesF != nil {
		return j.BytesF()
	}
	if j.CantBytes && j.T != nil {
		j.T.Fatalf("Unexpectedly called Bytes")
	}
	return nil
}

func (j *TestJob) HasMissingDependencies(ctx context.Context) (bool, error) {
	if j.HasMissingDependenciesF != nil {
		return j.HasMissingDependenciesF(ctx)
	}
	if j.CantHasMissingDependencies && j.T != nil {
		j.T.Fatal(errHasMissingDependencies)
	}
	return false, errHasMissingDependencies
}
