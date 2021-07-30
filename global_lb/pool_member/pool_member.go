package pool_member

import go_f5_soap "github.com/wule61/go-f5-soap"

const tns = "urn:iControl:GlobalLB/PoolMember"

type PoolMember struct {
	c *go_f5_soap.Client
}

func NewPoolMember(c *go_f5_soap.Client) *PoolMember {
	return &PoolMember{
		c: c,
	}
}

func (p *PoolMember) GetRatio()  {

}
