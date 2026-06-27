package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOExecutionConditionItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measObjectContainer"},
		{Name: "reportConfigContainer"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOExecutionConditionItem struct {
	MeasObjectContainer   MeasObjectContainer
	ReportConfigContainer ReportConfigContainer
	IEExtensions          []byte
}

func (ie *CHOExecutionConditionItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOExecutionConditionItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.MeasObjectContainer.Encode(e); err != nil {
		return err
	}
	if err := ie.ReportConfigContainer.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CHOExecutionConditionItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOExecutionConditionItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MeasObjectContainer.Decode(d); err != nil {
		return err
	}
	if err := ie.ReportConfigContainer.Decode(d); err != nil {
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
