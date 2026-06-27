package ies

import (
	"github.com/lvdund/asn1go/per"
)

var specialSubframeInfoEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "specialSubframePattern"},
		{Name: "cyclicPrefixDL"},
		{Name: "cyclicPrefixUL"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SpecialSubframeInfoEUTRA struct {
	SpecialSubframePattern SpecialSubframePatternsEUTRA
	CyclicPrefixDL         CyclicPrefixEUTRADL
	CyclicPrefixUL         CyclicPrefixEUTRAUL
	IEExtensions           []byte
}

func (ie *SpecialSubframeInfoEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(specialSubframeInfoEUTRAConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SpecialSubframePattern.Encode(e); err != nil {
		return err
	}
	if err := ie.CyclicPrefixDL.Encode(e); err != nil {
		return err
	}
	if err := ie.CyclicPrefixUL.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SpecialSubframeInfoEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(specialSubframeInfoEUTRAConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SpecialSubframePattern.Decode(d); err != nil {
		return err
	}
	if err := ie.CyclicPrefixDL.Decode(d); err != nil {
		return err
	}
	if err := ie.CyclicPrefixUL.Decode(d); err != nil {
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
