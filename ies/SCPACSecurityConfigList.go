package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sCPACSecurityConfigListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSecurityConfigurations)),
}

type SCPACSecurityConfigList struct {
	Value []*SCPACSecurityConfigItem
}

func (ie *SCPACSecurityConfigList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sCPACSecurityConfigListConstraints)
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

func (ie *SCPACSecurityConfigList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sCPACSecurityConfigListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SCPACSecurityConfigItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SCPACSecurityConfigItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
