package client

import "github.com/weibocom/ipc/apis/database"

func (c *Client) GetDynamicGlobalProperties(creator string, name string, fee int, jsonMeta string) (*database.DynamicGlobalProperties, error) {
	return c.Database.GetDynamicGlobalProperties()
}
