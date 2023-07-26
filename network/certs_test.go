// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"crypto/tls"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/network/peer"
	"github.com/luxdefi/node/staking"
)

var (
	certLock   sync.Mutex
	tlsCerts   []*tls.Certificate
	tlsConfigs []*tls.Config
)

func getTLS(t *testing.T, index int) (ids.NodeID, *tls.Certificate, *tls.Config) {
	certLock.Lock()
	defer certLock.Unlock()

	for len(tlsCerts) <= index {
		cert, err := staking.NewTLSCert()
		require.NoError(t, err)
		tlsConfig := peer.TLSConfig(*cert, nil)

		tlsCerts = append(tlsCerts, cert)
		tlsConfigs = append(tlsConfigs, tlsConfig)
	}

	cert := tlsCerts[index]
	return ids.NodeIDFromCert(cert.Leaf), cert, tlsConfigs[index]
}
