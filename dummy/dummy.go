package dummy

import (
	"context"
	"fmt"

	//	"coredns_wrapper/coredns/plugin"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

type dummyPlugin struct {
	next  plugin.Handler
	pass  string
	nonce string
}

func newdummyPlugin() *dummyPlugin {
	return &dummyPlugin{
		pass:  defPass,
		nonce: defNonce,
	}
}

// Name implements the Handler interface
func (p *dummyPlugin) Name() string {
	return "dummy"
}

// ServeDNS implements the Handler interface
func (p *dummyPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	fmt.Println("in serveDNS call")
	/*
		if qName, ok := checkQuery(r); ok {
			// Parse the domain and access code
			domainName, accessCode, err := parseQuery(qName, p.nonce, p.pass)
			if err != nil {
				// Cannot parse domain and access code
				// forge access code to get redirection to error server
				log.Warningf("Failed to decode: qName=%q, error=%q", qName, err)
				accessCode = fakeAccessCode
			}
			addEdnsOption(r, ednsQueryName, []byte(domainName))
			addEdnsOption(r, ednsAccessCode, []byte(accessCode))

			ts := time.Now().Unix()
			tsBuf := make([]byte, 8)
			binary.BigEndian.PutUint64(tsBuf, uint64(ts))
			addEdnsOption(r, ednsTimeStamp, tsBuf)
		}
	*/
	return p.next.ServeDNS(ctx, w, r)
}
