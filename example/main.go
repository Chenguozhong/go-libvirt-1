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

	domains, err := conn.ListAllDomains(uint(0))
	if err != nil {
		panic(err)
	}
	if err := domains[0].Destroy(); err != nil {
		panic(err)
	}
	for _, domain := range domains {
		id, err := domain.ID()
		if err != nil {
			panic(err)
		}
		name, err := domain.Name()
		if err != nil {
			panic(err)
		}
		state, err := domain.State()
		if err != nil {
			panic(err)
		}
		uuid, err := domain.UUID()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\t%s\t%s\t\t\t%s\n", id, uuid, name, state)
	}
}
