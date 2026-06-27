package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cAGListforMDTItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmnID"},
		{Name: "cAGID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CAGListforMDTItem struct {
	PlmnID       PLMNIdentity
	CAGID        CAGIdentifier
	IEExtensions []byte
}

func (ie *CAGListforMDTItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAGListforMDTItemConstraints)
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
	if err := ie.CAGID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CAGListforMDTItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAGListforMDTItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnID.Decode(d); err != nil {
		return err
	}
	if err := ie.CAGID.Decode(d); err != nil {
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
