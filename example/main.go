package main

import (
	"fmt"

	libvirt "github.com/vtolstov/go-libvirt"
)

var (
	domxml = string(`
    <domain type="kvm">
      <name>fii</name>
      <memory>1024</memory>
      <os><type>hvm</type></os>
    </domain>
`)
)

func main() {

	conn, err := libvirt.NewSessionConn()
	if err != nil {
		panic(err)
	}

	domains, err := conn.ListAllDomains(0)
	if err != nil {
		panic(err)
	}
	/*
		domain, err := conn.CreateXML(domxml, 0)
		if err != nil {
			panic(err)
		}
	*/
	stats, err := domains[0].Stats(0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#+v\n", stats)
}
