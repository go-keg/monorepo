package ent

import (
	"github.com/go-keg/keg/contrib/ent/driver"
)

func (c *Client) DebugWithOptions(opts ...driver.DebugOption) *Client {
	cfg := c.config
	cfg.driver = driver.Debug(c.driver, opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}
