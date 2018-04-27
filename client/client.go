package client

import (
	// RPC

	"github.com/weibocom/ipc/apis/condenser"
	"github.com/weibocom/ipc/apis/database"
	"github.com/weibocom/ipc/apis/follow"
	"github.com/weibocom/ipc/apis/login"
	"github.com/weibocom/ipc/apis/networkbroadcast"
	"github.com/weibocom/ipc/interfaces"
)

// Client can be used to access Steem remote APIs.
//
// There is a public field for every Steem API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc interfaces.CallCloser

	// Login represents login_api.
	Login *login.API

	// Database represents database_api.
	Database *database.API

	// Follow represents follow_api.
	Follow *follow.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *networkbroadcast.API

	// Condenser represents condenser_api.
	Condenser *condenser.API
}

// NewClient creates a new RPC client that use the given CallCloser internally.
func NewClient(cc interfaces.CallCloser) (*Client, error) {
	client := &Client{cc: cc}
	client.Login = login.NewAPI(client.cc)
	client.Database = database.NewAPI(client.cc)

	followAPI, err := follow.NewAPI(client.cc)
	if err != nil {
		return nil, err
	}
	client.Follow = followAPI

	networkBroadcastAPI, err := networkbroadcast.NewAPI(client.cc)
	if err != nil {
		return nil, err
	}
	client.NetworkBroadcast = networkBroadcastAPI

	condenserAPI := condenser.NewAPI(client.cc)
	if err != nil {
		return nil, err
	}
	client.Condenser = condenserAPI

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (c *Client) Close() error {
	return c.cc.Close()
}
