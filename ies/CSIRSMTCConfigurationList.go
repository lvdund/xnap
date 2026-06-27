package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cSIRSMTCConfigurationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCSIRSconfigurations)),
}

type CSIRSMTCConfigurationList struct {
	Value []*CSIRSMTCConfigurationItem
}

func (ie *CSIRSMTCConfigurationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSIRSMTCConfigurationListConstraints)
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

func (ie *CSIRSMTCConfigurationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSIRSMTCConfigurationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CSIRSMTCConfigurationItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CSIRSMTCConfigurationItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
