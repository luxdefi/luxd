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
// See the file LICENSE for licensing terms.

package lux

import (
<<<<<<< HEAD
=======
<<<<<<< HEAD:snow/engine/avalanche/traced_engine.go
	"context"

	"go.opentelemetry.io/otel/attribute"

	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/snow/consensus/avalanche"
	"github.com/luxdefi/node/snow/engine/common"
	"github.com/luxdefi/node/trace"
=======
>>>>>>> 53a8245a8 (Update consensus)
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/snow/consensus/lux"
	"github.com/luxdefi/node/snow/engine/common"
	"github.com/luxdefi/node/trace"
<<<<<<< HEAD
=======
>>>>>>> 04d685aa2 (Update consensus):snow/engine/lux/traced_engine.go
>>>>>>> 53a8245a8 (Update consensus)
)

var _ Engine = (*tracedEngine)(nil)

type tracedEngine struct {
	common.Engine
	engine Engine
<<<<<<< HEAD
=======
	tracer trace.Tracer
>>>>>>> 53a8245a8 (Update consensus)
}

func TraceEngine(engine Engine, tracer trace.Tracer) Engine {
	return &tracedEngine{
		Engine: common.TraceEngine(engine, tracer),
		engine: engine,
<<<<<<< HEAD
	}
}

func (e *tracedEngine) GetVtx(vtxID ids.ID) (lux.Vertex, error) {
	return e.engine.GetVtx(vtxID)
=======
		tracer: tracer,
	}
}

<<<<<<< HEAD:snow/engine/avalanche/traced_engine.go
func (e *tracedEngine) GetVtx(ctx context.Context, vtxID ids.ID) (avalanche.Vertex, error) {
	ctx, span := e.tracer.Start(ctx, "tracedEngine.GetVtx", oteltrace.WithAttributes(
		attribute.Stringer("vtxID", vtxID),
	))
	defer span.End()

	return e.engine.GetVtx(ctx, vtxID)
=======
func (e *tracedEngine) GetVtx(vtxID ids.ID) (lux.Vertex, error) {
	return e.engine.GetVtx(vtxID)
>>>>>>> 04d685aa2 (Update consensus):snow/engine/lux/traced_engine.go
>>>>>>> 53a8245a8 (Update consensus)
}
