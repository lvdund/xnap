package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var iABQoSMappingInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dscp", Optional: true},
		{Name: "flow-label", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABQoSMappingInformation struct {
	Dscp         *per.BitString
	FlowLabel    *per.BitString
	IEExtensions []byte
}

func (ie *IABQoSMappingInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABQoSMappingInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Dscp != nil, ie.FlowLabel != nil, false}); err != nil {
		return err
	}
	if ie.Dscp != nil {
		if err := e.EncodeBitString(*ie.Dscp, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		}); err != nil {
			return err
		}
	}
	if ie.FlowLabel != nil {
		if err := e.EncodeBitString(*ie.FlowLabel, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(20)),
			Max:        common.Ptr(int64(20)),
		}); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABQoSMappingInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABQoSMappingInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		})
		if err != nil {
			return err
		}
		ie.Dscp = &val
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(20)),
			Max:        common.Ptr(int64(20)),
		})
		if err != nil {
			return err
		}
		ie.FlowLabel = &val
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
