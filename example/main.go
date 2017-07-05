package main

import (
	"fmt"

	libvirt "github.com/vtolstov/go-libvirt"
)

var (
	domxml = string(`
    <domain type="kvm">
      <name>foo</name>
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
	fmt.Printf("%#+v\n", domains)
	/*
		obj := conn.Object("org.libvirt", "/org/libvirt/Manager")
		call := obj.Call("org.libvirt.Manager.ListDomains", 0, uint(0))
		if call.Err != nil {
			panic(call.Err)
		}
		var domains interface{}
		if err := call.Store(&domains); err != nil {
			panic(err)
		}
		fmt.Printf("%#+v\n", domains)
		call = obj.Call("org.libvirt.Manager.CreateXML", 0, domxml, uint(0))
		if call.Err != nil {
			panic(call.Err)
		}
		var res interface{}
		if err := call.Store(&res); err != nil {
			panic(err)
		}
		fmt.Printf("%#+v\n", res)
		call = obj.Call("org.libvirt.Manager.ListDomains", 0, uint(0))
		if call.Err != nil {
			panic(call.Err)
		}
		if err := call.Store(&domains); err != nil {
			panic(err)
		}
		fmt.Printf("%#+v\n", domains)
	*/
}
