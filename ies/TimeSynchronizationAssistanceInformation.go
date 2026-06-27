package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	TimeSynchronizationAssistanceInformationTimeDistributionIndicationEnabled  int64 = 0
	TimeSynchronizationAssistanceInformationTimeDistributionIndicationDisabled int64 = 1
)

var timeSynchronizationAssistanceInformationTimeDistributionIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type TimeSynchronizationAssistanceInformationTimeDistributionIndication struct {
	Value int64
}

func (ie *TimeSynchronizationAssistanceInformationTimeDistributionIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, timeSynchronizationAssistanceInformationTimeDistributionIndicationConstraints)
}

func (ie *TimeSynchronizationAssistanceInformationTimeDistributionIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(timeSynchronizationAssistanceInformationTimeDistributionIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var timeSynchronizationAssistanceInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "timeDistributionIndication"},
		{Name: "uuTimeSynchronizationErrorBudget"},
		{Name: "ie-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TimeSynchronizationAssistanceInformation struct {
	TimeDistributionIndication       TimeSynchronizationAssistanceInformationTimeDistributionIndication
	UuTimeSynchronizationErrorBudget int64
	IEExtensions                     []byte
}

func (ie *TimeSynchronizationAssistanceInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(timeSynchronizationAssistanceInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TimeDistributionIndication.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.UuTimeSynchronizationErrorBudget, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(1000000)),
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

func (ie *TimeSynchronizationAssistanceInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(timeSynchronizationAssistanceInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TimeDistributionIndication.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(1000000)),
		})
		if err != nil {
			return err
		}
		ie.UuTimeSynchronizationErrorBudget = val
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
