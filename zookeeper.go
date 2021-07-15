// Package zookeeper implements a plugin that returns details about znodes
package zookeeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/coredns/coredns/request"
	"github.com/go-zookeeper/zk"

	"github.com/miekg/dns"
)

const Name = "zookeeper"

// Zookeeper is a plugin that returns data associated with a znode
type Zookeeper struct{}

// ServeDNS implements the plugin.Handler interface.
func (z Zookeeper) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Authoritative = true

	/*
		ip := state.IP()
		var rr dns.RR

		switch state.Family() {
		case 1:
			rr = new(dns.A)
			rr.(*dns.A).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeA, Class: state.QClass()}
			rr.(*dns.A).A = net.ParseIP(ip).To4()
		case 2:
			rr = new(dns.AAAA)
			rr.(*dns.AAAA).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeAAAA, Class: state.QClass()}
			rr.(*dns.AAAA).AAAA = net.ParseIP(ip)
		}
	*/

	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		fmt.Println("failed to get CPU utilization")
		panic(err)
	}
	cpuUtil, _, err := c.Get("/node/1/performance/cpu")
	if err != nil {
		fmt.Println("failed to get CPU utilization")
	} else {
		fmt.Println("CPU utilization: " + string(cpuUtil))
	}

	srv := new(dns.SRV)
	srv.Hdr = dns.RR_Header{Name: "_" + state.Proto() + "." + state.QName(), Rrtype: dns.TypeSRV, Class: state.QClass()}
	if state.QName() == "." {
		srv.Hdr.Name = "_" + state.Proto() + state.QName()
	}
	port, _ := strconv.Atoi(state.Port())
	srv.Port = uint16(port)
	srv.Target = "."

	a.Extra = []dns.RR{srv}

	w.WriteMsg(a)

	return 0, nil
}

// Name implements the Handler interface.
func (z Zookeeper) Name() string { return Name }
