package dummy

import (
	"fmt"

	//	"coredns_wrapper/coredns/plugin"

	dnspb "dns_resolver_dummy/dnspb"
	"log"

	//log "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":50551", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := dnspb.NewDnsServiceClient(conn)
	domainName := parseQuery(r)

	if !(DNSauthenticate(domainName[:len(domainName)-1], c)) {
		return dns.RcodeRefused, nil
	}
	/*		DNSauthenticate("yahoo.com", c)
			DNSauthenticate("flipkart.com", c)
			DNSauthenticate("facebook.com", c)
			DNSauthenticate("infoblox.com", c)
			DNSauthenticate("twitter.com", c)
			DNSauthenticate("music.com", c)
	*/
	//	log.Println("END OF PROGRAM")

	return p.next.ServeDNS(ctx, w, r)
}

func DNSauthenticate(domainname string, c dnspb.DnsServiceClient) bool {
	response, err := c.DNSauthenticate(context.Background(), &dnspb.DomainReq{DomainName: domainname})
	if err != nil {
		log.Fatalf("Error when calling Server: %s", err)
	}
	log.Printf("Response from server %s for domain: %s", response.Action.String(), domainname)
	if response.Action.String() == "ALLOW" {
		return true
	}
	return false

}
