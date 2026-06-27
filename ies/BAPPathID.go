package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bAPPathIDConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(10)),
	Max:        common.Ptr(int64(10)),
}

type BAPPathID struct {
	Value per.BitString
}

func (ie *BAPPathID) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, bAPPathIDConstraints)
}

func (ie *BAPPathID) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(bAPPathIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
