package libvirt

func (c *Conn) ListAllDomains(flags uint) (interface{}, error) {
	var domains interface{}
	obj := c.conn.Object(destObject, managerPath)
	call := obj.Call(managerDest+".ListDomains", 0, flags)
	if call.Err != nil {
		return nil, call.Err
	}
	if err := call.Store(&domains); err != nil {
		return nil, err
	}
	return domains, nil
}
