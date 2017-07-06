package libvirt

import "github.com/godbus/dbus"

// Conn libvirt connection
type Conn struct {
	conn *dbus.Conn
}

// NewSessionConn creates session conn used for unprivileged access
func NewSessionConn() (*Conn, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn}, nil
}

// NewSystemConn creates system conn for privileged access
func NewSystemConn() (*Conn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn}, nil
}

// Close closes conn
func (c *Conn) Close() error {
	return c.conn.Close()
}
