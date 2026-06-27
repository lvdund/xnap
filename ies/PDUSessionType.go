package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDUSessionTypeIpv4         int64 = 0
	PDUSessionTypeIpv6         int64 = 1
	PDUSessionTypeIpv4v6       int64 = 2
	PDUSessionTypeEthernet     int64 = 3
	PDUSessionTypeUnstructured int64 = 4
)

var pDUSessionTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type PDUSessionType struct {
	Value int64
}

func (ie *PDUSessionType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDUSessionTypeConstraints)
}

func (ie *PDUSessionType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDUSessionTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
