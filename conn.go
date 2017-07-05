package libvirt

import "github.com/godbus/dbus"

const (
	destObject  = "org.libvirt"
	managerDest = "org.libvirt.Manager"
	managerPath = "/org/libvirt/Manager"
	domainDest  = "org.libvirt.domain"
	domainPath  = "/org/libvirt/domain"
)

type Conn struct {
	conn *dbus.Conn
}

func NewSessionConn() (*Conn, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn}, nil
}

func NewSystemConn() (*Conn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn}, nil
}

func (c *Conn) Close() error {
	return c.conn.Close()
}
