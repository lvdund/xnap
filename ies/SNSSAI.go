package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNSSAIConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sst"},
		{Name: "sd", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SNSSAI struct {
	Sst          []byte
	Sd           []byte
	IEExtensions []byte
}

func (ie *SNSSAI) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNSSAIConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.Sd) > 0, false}); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.Sst, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(1)),
		Max:        common.Ptr(int64(1)),
	}); err != nil {
		return err
	}
	if len(ie.Sd) > 0 {
		if err := e.EncodeOctetString(ie.Sd, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(3)),
			Max:        common.Ptr(int64(3)),
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

func (ie *SNSSAI) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNSSAIConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(1)),
		})
		if err != nil {
			return err
		}
		ie.Sst = val
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(3)),
			Max:        common.Ptr(int64(3)),
		})
		if err != nil {
			return err
		}
		ie.Sd = val
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
