package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"

	// Include necessary plugins.
	_ "plugin_check/dummy"

	_ "github.com/caddyserver/caddy/onevent"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/chaos"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/dnstap"
	_ "github.com/coredns/coredns/plugin/erratic"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/pprof"
	_ "github.com/coredns/coredns/plugin/tls"
	_ "github.com/coredns/coredns/plugin/whoami"
	_ "github.com/coredns/rrl/plugins/rrl"
)

// Directives are registered in the order they should be executed.
var directives = []string{
	"tls",
	"bind",
	"debug",
	"health",
	"pprof",
	"prometheus",
	"errors",
	"log",
	"dnstap",
	"ibdnstap",
	"chaos",
	"rrl",
	"file",
	"auth",
	"redirect",
	"policy",
	"cache",
	"dummy",
	"forward",
	"erratic",
	"whoami",
	"on",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}
