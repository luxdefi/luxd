// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package gsubnetlookup

import (
	"context"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/snow"

	subnetlookuppb "github.com/luxdefi/node/proto/pb/subnetlookup"
)

var _ snow.SubnetLookup = (*Client)(nil)

// Client is a subnet lookup that talks over RPC.
type Client struct {
	client subnetlookuppb.SubnetLookupClient
}

// NewClient returns an alias lookup connected to a remote alias lookup
func NewClient(client subnetlookuppb.SubnetLookupClient) *Client {
	return &Client{client: client}
}

func (c *Client) SubnetID(chainID ids.ID) (ids.ID, error) {
	resp, err := c.client.SubnetID(context.Background(), &subnetlookuppb.SubnetIDRequest{
		ChainId: chainID[:],
	})
	if err != nil {
		return ids.ID{}, err
	}
	return ids.ToID(resp.Id)
}
