package dummy

import (
	"github.com/miekg/dns"
)

func parseQuery(r *dns.Msg) string {
	qName, _ := getNameAndType(r)
	return qName
}

func getNameAndType(r *dns.Msg) (string, uint16) {
	if r == nil || len(r.Question) <= 0 {
		return ".", dns.TypeNone
	}

	q := r.Question[0]
	return q.Name, q.Qtype
}
