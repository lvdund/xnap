package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var protectedEUTRAResourceIndicationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "activationSFN"},
		{Name: "protectedResourceList"},
		{Name: "mbsfnControlRegionLength", Optional: true},
		{Name: "pDCCHRegionLength"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ProtectedEUTRAResourceIndication struct {
	ActivationSFN            ActivationSFN
	ProtectedResourceList    ProtectedEUTRAResourceList
	MbsfnControlRegionLength *MBSFNControlRegionLength
	PDCCHRegionLength        int64
	IEExtensions             []byte
}

func (ie *ProtectedEUTRAResourceIndication) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(protectedEUTRAResourceIndicationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MbsfnControlRegionLength != nil, false}); err != nil {
		return err
	}
	if err := ie.ActivationSFN.Encode(e); err != nil {
		return err
	}
	if err := ie.ProtectedResourceList.Encode(e); err != nil {
		return err
	}
	if ie.MbsfnControlRegionLength != nil {
		if err := ie.MbsfnControlRegionLength.Encode(e); err != nil {
			return err
		}
	}
	if err := e.EncodeInteger(ie.PDCCHRegionLength, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(3)),
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ProtectedEUTRAResourceIndication) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(protectedEUTRAResourceIndicationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ActivationSFN.Decode(d); err != nil {
		return err
	}
	if err := ie.ProtectedResourceList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.MbsfnControlRegionLength = new(MBSFNControlRegionLength)
		if err := ie.MbsfnControlRegionLength.Decode(d); err != nil {
			return err
		}
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(3)),
		})
		if err != nil {
			return err
		}
		ie.PDCCHRegionLength = val
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
