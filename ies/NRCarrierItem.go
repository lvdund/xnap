package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRCarrierItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierSCS"},
		{Name: "offsetToCarrier"},
		{Name: "carrierBandwidth"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRCarrierItem struct {
	CarrierSCS       NRSCS
	OffsetToCarrier  int64
	CarrierBandwidth int64
	IEExtensions     []byte
}

func (ie *NRCarrierItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRCarrierItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.CarrierSCS.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.OffsetToCarrier, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(2199)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.CarrierBandwidth, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(common.MaxnoofPhysicalResourceBlocks)),
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

func (ie *NRCarrierItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRCarrierItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CarrierSCS.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(2199)),
		})
		if err != nil {
			return err
		}
		ie.OffsetToCarrier = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(common.MaxnoofPhysicalResourceBlocks)),
		})
		if err != nil {
			return err
		}
		ie.CarrierBandwidth = val
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
