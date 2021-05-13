package dummy

import (
	"fmt"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

var log = clog.NewWithPlugin("dummy")

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
		if len(remainingArgs) > 2 {
			return nil, c.Err("too many arguments")
		}
		if len(remainingArgs) > 0 {
			ap.pass = remainingArgs[0]
		}
		if len(remainingArgs) > 1 {
			ap.nonce = remainingArgs[1]
		}
	}
	return ap, nil
}

const (
	defPass  = "infoblox"
	defNonce = "infoblox"
)
