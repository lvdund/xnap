package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pCIListForMDTConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofNeighPCIforMDT)),
}

type PCIListForMDT struct {
	Value []*NRPCI
}

func (ie *PCIListForMDT) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pCIListForMDTConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PCIListForMDT) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pCIListForMDTConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NRPCI, n)
	for i := range ie.Value {
		ie.Value[i] = new(NRPCI)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
