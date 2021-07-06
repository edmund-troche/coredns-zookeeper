package zookeeper

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register(Name, setupZookeeper) }

func setupZookeeper(c *caddy.Controller) error {
	c.Next() // 'zookeeper'
	if c.NextArg() {
		return plugin.Error(Name, c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Zookeeper{}
	})

	return nil
}
