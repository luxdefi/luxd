// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/stretchr/testify/require"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/utils/constants"
)

func Test_newOutboundBuilder(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	mb, err := newMsgBuilder("test", prometheus.NewRegistry(), int64(constants.DefaultMaxMessageSize), 5*time.Second)
	require.NoError(err)

	builder := newOutboundBuilder(true /*compress*/, mb)

	outMsg, err := builder.GetAcceptedStateSummary(ids.GenerateTestID(), uint32(12345), time.Hour, []uint64{1000, 2000})
	require.NoError(err)

	t.Logf("outbound message built %q with size %d", outMsg.Op().String(), len(outMsg.Bytes()))
}
