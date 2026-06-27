package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ClockQualityAcceptanceCriteriaTraceabletoUTCTrue int64 = 0
)

var clockQualityAcceptanceCriteriaTraceabletoUTCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ClockQualityAcceptanceCriteriaTraceabletoUTC struct {
	Value int64
}

func (ie *ClockQualityAcceptanceCriteriaTraceabletoUTC) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, clockQualityAcceptanceCriteriaTraceabletoUTCConstraints)
}

func (ie *ClockQualityAcceptanceCriteriaTraceabletoUTC) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(clockQualityAcceptanceCriteriaTraceabletoUTCConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	ClockQualityAcceptanceCriteriaTraceabletoGNSSTrue int64 = 0
)

var clockQualityAcceptanceCriteriaTraceabletoGNSSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ClockQualityAcceptanceCriteriaTraceabletoGNSS struct {
	Value int64
}

func (ie *ClockQualityAcceptanceCriteriaTraceabletoGNSS) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, clockQualityAcceptanceCriteriaTraceabletoGNSSConstraints)
}

func (ie *ClockQualityAcceptanceCriteriaTraceabletoGNSS) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(clockQualityAcceptanceCriteriaTraceabletoGNSSConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var clockQualityAcceptanceCriteriaConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "synchronisationState"},
		{Name: "traceabletoUTC", Optional: true},
		{Name: "traceabletoGNSS", Optional: true},
		{Name: "clockFrequencyStability", Optional: true},
		{Name: "clockAccuracy"},
		{Name: "parentTimeSource"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ClockQualityAcceptanceCriteria struct {
	SynchronisationState    per.BitString
	TraceabletoUTC          *ClockQualityAcceptanceCriteriaTraceabletoUTC
	TraceabletoGNSS         *ClockQualityAcceptanceCriteriaTraceabletoGNSS
	ClockFrequencyStability *per.BitString
	ClockAccuracy           int64
	ParentTimeSource        per.BitString
	IEExtensions            []byte
}

func (ie *ClockQualityAcceptanceCriteria) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(clockQualityAcceptanceCriteriaConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.TraceabletoUTC != nil, ie.TraceabletoGNSS != nil, ie.ClockFrequencyStability != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.SynchronisationState, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if ie.TraceabletoUTC != nil {
		if err := ie.TraceabletoUTC.Encode(e); err != nil {
			return err
		}
	}
	if ie.TraceabletoGNSS != nil {
		if err := ie.TraceabletoGNSS.Encode(e); err != nil {
			return err
		}
	}
	if ie.ClockFrequencyStability != nil {
		if err := e.EncodeBitString(*ie.ClockFrequencyStability, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(16)),
			Max:        common.Ptr(int64(16)),
		}); err != nil {
			return err
		}
	}
	if err := e.EncodeInteger(ie.ClockAccuracy, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(40000000)),
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.ParentTimeSource, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
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

func (ie *ClockQualityAcceptanceCriteria) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(clockQualityAcceptanceCriteriaConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.SynchronisationState = val
	}
	if seq.IsComponentPresent(1) {
		ie.TraceabletoUTC = new(ClockQualityAcceptanceCriteriaTraceabletoUTC)
		if err := ie.TraceabletoUTC.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.TraceabletoGNSS = new(ClockQualityAcceptanceCriteriaTraceabletoGNSS)
		if err := ie.TraceabletoGNSS.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(16)),
			Max:        common.Ptr(int64(16)),
		})
		if err != nil {
			return err
		}
		ie.ClockFrequencyStability = &val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(40000000)),
		})
		if err != nil {
			return err
		}
		ie.ClockAccuracy = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.ParentTimeSource = val
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
