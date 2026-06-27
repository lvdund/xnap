package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mBSServiceAreaTAIItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-ID"},
		{Name: "tAC"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSServiceAreaTAIItem struct {
	PlmnID       PLMNIdentity
	TAC          TAC
	IEExtensions []byte
}

func (ie *MBSServiceAreaTAIItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSServiceAreaTAIItemConstraints)
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
	if err := ie.TAC.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MBSServiceAreaTAIItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSServiceAreaTAIItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnID.Decode(d); err != nil {
		return err
	}
	if err := ie.TAC.Decode(d); err != nil {
		return err
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
