package ies

import (
	"github.com/lvdund/asn1go/per"
)

var bAPControlPDURLCCHItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bHRLCCHID"},
		{Name: "nexthopBAPAddress"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BAPControlPDURLCCHItem struct {
	BHRLCCHID         BHRLCChannelID
	NexthopBAPAddress BAPAddress
	IEExtensions      []byte
}

func (ie *BAPControlPDURLCCHItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bAPControlPDURLCCHItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.BHRLCCHID.Encode(e); err != nil {
		return err
	}
	if err := ie.NexthopBAPAddress.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BAPControlPDURLCCHItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bAPControlPDURLCCHItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BHRLCCHID.Decode(d); err != nil {
		return err
	}
	if err := ie.NexthopBAPAddress.Decode(d); err != nil {
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
