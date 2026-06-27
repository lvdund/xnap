package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var gUAMIConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-ID"},
		{Name: "amf-region-id"},
		{Name: "amf-set-id"},
		{Name: "amf-pointer"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type GUAMI struct {
	PlmnID       PLMNIdentity
	AmfRegionId  per.BitString
	AmfSetId     per.BitString
	AmfPointer   per.BitString
	IEExtensions []byte
}

func (ie *GUAMI) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gUAMIConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnID.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.AmfRegionId, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(8)),
		Max:        common.Ptr(int64(8)),
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.AmfSetId, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(10)),
		Max:        common.Ptr(int64(10)),
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.AmfPointer, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(6)),
		Max:        common.Ptr(int64(6)),
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

func (ie *GUAMI) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gUAMIConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnID.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(8)),
			Max:        common.Ptr(int64(8)),
		})
		if err != nil {
			return err
		}
		ie.AmfRegionId = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(10)),
			Max:        common.Ptr(int64(10)),
		})
		if err != nil {
			return err
		}
		ie.AmfSetId = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		})
		if err != nil {
			return err
		}
		ie.AmfPointer = val
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
