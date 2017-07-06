package libvirt

import "github.com/godbus/dbus"

// CreateXML creates new domain using xmldesc and flags
func (c *Conn) CreateXML(xmldesc string, flags uint) (*Domain, error) {
	var iface dbus.ObjectPath
	obj := c.conn.Object(destObject, managerPath)
	call := obj.Call(managerDest+".CreateXML", 0, xmldesc, flags)
	if call.Err != nil {
		return nil, call.Err
	}
	if err := call.Store(&iface); err != nil {
		return nil, call.Err
	}
	return &Domain{conn: c.conn, path: iface}, nil
}

// DefineXML creates new domain using xmldesc and flags
func (c *Conn) DefineXML(xmldesc string, flags uint) (*Domain, error) {
	var iface dbus.ObjectPath
	obj := c.conn.Object(destObject, managerPath)
	call := obj.Call(managerDest+".DefineXML", 0, xmldesc, flags)
	if call.Err != nil {
		return nil, call.Err
	}
	if err := call.Store(&iface); err != nil {
		return nil, call.Err
	}
	return &Domain{conn: c.conn, path: iface}, nil
}

// ListAllDomains lists all domains using flags
func (c *Conn) ListAllDomains(flags uint) ([]*Domain, error) {
	var iface []dbus.ObjectPath
	var domains []*Domain
	obj := c.conn.Object(destObject, managerPath)
	call := obj.Call(managerDest+".ListDomains", 0, flags)
	if call.Err != nil {
		return nil, call.Err
	}
	if err := call.Store(&iface); err != nil {
		return nil, err
	}

	for _, d := range iface {
		domains = append(domains, &Domain{conn: c.conn, path: d})
	}

	return domains, nil
}

// Domain struct explains domain
type Domain struct {
	conn *dbus.Conn
	path dbus.ObjectPath
}

// Name returns domain name
func (d *Domain) Name() (string, error) {
	obj := d.conn.Object(destObject, d.path)
	v, err := obj.GetProperty(domainDest + ".Name")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// ID returns domain id
func (d *Domain) ID() (uint32, error) {
	obj := d.conn.Object(destObject, d.path)
	v, err := obj.GetProperty(domainDest + ".Id")
	if err != nil {
		return 0, err
	}
	return v.Value().(uint32), nil
}

// UUID return domain uuid
func (d *Domain) UUID() (string, error) {
	obj := d.conn.Object(destObject, d.path)
	v, err := obj.GetProperty(domainDest + ".UUID")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// State returns domain state
func (d *Domain) State() (string, error) {
	obj := d.conn.Object(destObject, d.path)
	v, err := obj.GetProperty(domainDest + ".State")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// Reboot reboots domain
func (d *Domain) Reboot(flags uint) error {
	obj := d.conn.Object(destObject, d.path)
	call := obj.Call(managerDest+".Reboot", 0, flags)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

// XMLDesc return domain xmldesc
func (d *Domain) XMLDesc(flags uint) (string, error) {
	var xmldesc string
	obj := d.conn.Object(destObject, d.path)
	call := obj.Call(domainDest+".GetXMLDesc", 0, flags)
	if call.Err != nil {
		return "", call.Err
	}

	if err := call.Store(&xmldesc); err != nil {
		return "", err
	}
	return xmldesc, nil
}

// Shutdown shutdown domain
func (d *Domain) Shutdown(flags uint) error {
	obj := d.conn.Object(destObject, d.path)
	call := obj.Call(domainDest+".Shutdown", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

// Destroy destroys domain
func (d *Domain) Destroy() error {
	obj := d.conn.Object(destObject, d.path)
	call := obj.Call(domainDest+".Destroy", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

// Create create domain
func (d *Domain) Create() error {
	obj := d.conn.Object(destObject, d.path)
	call := obj.Call(domainDest+".Create", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

// Undefine undefines domain
func (d *Domain) Undefine() error {
	obj := d.conn.Object(destObject, d.path)
	call := obj.Call(domainDest+".Undefine", 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}
