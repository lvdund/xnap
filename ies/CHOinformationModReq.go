package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CHOinformationModReqConditionalReconfigIntraMnCho int64 = 0
)

var cHOinformationModReqConditionalReconfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CHOinformationModReqConditionalReconfig struct {
	Value int64
}

func (ie *CHOinformationModReqConditionalReconfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cHOinformationModReqConditionalReconfigConstraints)
}

func (ie *CHOinformationModReqConditionalReconfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cHOinformationModReqConditionalReconfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var cHOinformationModReqConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "conditionalReconfig"},
		{Name: "cHO-EstimatedArrivalProbability", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOinformationModReq struct {
	ConditionalReconfig            CHOinformationModReqConditionalReconfig
	CHOEstimatedArrivalProbability *CHOProbability
	IEExtensions                   []byte
}

func (ie *CHOinformationModReq) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOinformationModReqConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CHOEstimatedArrivalProbability != nil, false}); err != nil {
		return err
	}
	if err := ie.ConditionalReconfig.Encode(e); err != nil {
		return err
	}
	if ie.CHOEstimatedArrivalProbability != nil {
		if err := ie.CHOEstimatedArrivalProbability.Encode(e); err != nil {
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

func (ie *CHOinformationModReq) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOinformationModReqConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ConditionalReconfig.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.CHOEstimatedArrivalProbability = new(CHOProbability)
		if err := ie.CHOEstimatedArrivalProbability.Decode(d); err != nil {
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
