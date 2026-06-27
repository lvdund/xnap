package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	NEDCTDMPatternSubframeAssignmentSa0 int64 = 0
	NEDCTDMPatternSubframeAssignmentSa1 int64 = 1
	NEDCTDMPatternSubframeAssignmentSa2 int64 = 2
	NEDCTDMPatternSubframeAssignmentSa3 int64 = 3
	NEDCTDMPatternSubframeAssignmentSa4 int64 = 4
	NEDCTDMPatternSubframeAssignmentSa5 int64 = 5
	NEDCTDMPatternSubframeAssignmentSa6 int64 = 6
)

var nEDCTDMPatternSubframeAssignmentConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  nil,
}

type NEDCTDMPatternSubframeAssignment struct {
	Value int64
}

func (ie *NEDCTDMPatternSubframeAssignment) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nEDCTDMPatternSubframeAssignmentConstraints)
}

func (ie *NEDCTDMPatternSubframeAssignment) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nEDCTDMPatternSubframeAssignmentConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var nEDCTDMPatternConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subframeAssignment"},
		{Name: "harqOffset"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NEDCTDMPattern struct {
	SubframeAssignment NEDCTDMPatternSubframeAssignment
	HarqOffset         int64
	IEExtensions       []byte
}

func (ie *NEDCTDMPattern) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nEDCTDMPatternConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SubframeAssignment.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.HarqOffset, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(9)),
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

func (ie *NEDCTDMPattern) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nEDCTDMPatternConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SubframeAssignment.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(9)),
		})
		if err != nil {
			return err
		}
		ie.HarqOffset = val
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
