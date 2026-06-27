package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBListWithCauseItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-id"},
		{Name: "cause"},
		{Name: "rLC-Mode", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBListWithCauseItem struct {
	DrbId        DRBID
	Cause        Cause
	RLCMode      *RLCMode
	IEExtensions []byte
}

func (ie *DRBListWithCauseItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBListWithCauseItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RLCMode != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbId.Encode(e); err != nil {
		return err
	}
	if err := ie.Cause.Encode(e); err != nil {
		return err
	}
	if ie.RLCMode != nil {
		if err := ie.RLCMode.Encode(e); err != nil {
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

func (ie *DRBListWithCauseItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBListWithCauseItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbId.Decode(d); err != nil {
		return err
	}
	if err := ie.Cause.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.RLCMode = new(RLCMode)
		if err := ie.RLCMode.Decode(d); err != nil {
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
