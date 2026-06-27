package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pSCellListContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type PSCellListContainer struct {
	Value []byte
}

func (ie *PSCellListContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, pSCellListContainerConstraints)
}

func (ie *PSCellListContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(pSCellListContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
