package dummy

import (
	"fmt"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

var vlog = clog.NewWithPlugin("dummy")

func init() {
	fmt.Println("init called in dummy")
	caddy.RegisterPlugin("dummy", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	clog.Info("in Dummy setup")
	dummyPlugin, err := dummyParse(c)

	if err != nil {
		return plugin.Error("dummy", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		dummyPlugin.next = next
		return dummyPlugin
	})

	return nil
}

func dummyParse(c *caddy.Controller) (*dummyPlugin, error) {
	ap := newdummyPlugin()

	for c.Next() {
		remainingArgs := c.RemainingArgs()
		if len(remainingArgs) > 0 {
			return nil, c.Err("too many arguments")
		}
	}
	return ap, nil
}

const (
	defPass  = "infoblox"
	defNonce = "infoblox"
)
