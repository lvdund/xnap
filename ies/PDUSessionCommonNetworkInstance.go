package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionCommonNetworkInstanceConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type PDUSessionCommonNetworkInstance struct {
	Value []byte
}

func (ie *PDUSessionCommonNetworkInstance) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, pDUSessionCommonNetworkInstanceConstraints)
}

func (ie *PDUSessionCommonNetworkInstance) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(pDUSessionCommonNetworkInstanceConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
