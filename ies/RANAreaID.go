package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rANAreaIDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tAC"},
		{Name: "rANAC", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RANAreaID struct {
	TAC          TAC
	RANAC        *RANAC
	IEExtensions []byte
}

func (ie *RANAreaID) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rANAreaIDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RANAC != nil, false}); err != nil {
		return err
	}
	if err := ie.TAC.Encode(e); err != nil {
		return err
	}
	if ie.RANAC != nil {
		if err := ie.RANAC.Encode(e); err != nil {
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

func (ie *RANAreaID) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rANAreaIDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TAC.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.RANAC = new(RANAC)
		if err := ie.RANAC.Decode(d); err != nil {
			return err
		}
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
