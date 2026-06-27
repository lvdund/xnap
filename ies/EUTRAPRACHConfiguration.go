package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	EUTRAPRACHConfigurationHighSpeedFlagTrue  int64 = 0
	EUTRAPRACHConfigurationHighSpeedFlagFalse int64 = 1
)

var eUTRAPRACHConfigurationHighSpeedFlagConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type EUTRAPRACHConfigurationHighSpeedFlag struct {
	Value int64
}

func (ie *EUTRAPRACHConfigurationHighSpeedFlag) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eUTRAPRACHConfigurationHighSpeedFlagConstraints)
}

func (ie *EUTRAPRACHConfigurationHighSpeedFlag) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eUTRAPRACHConfigurationHighSpeedFlagConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var eUTRAPRACHConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rootSequenceIndex"},
		{Name: "zeroCorrelationIndex"},
		{Name: "highSpeedFlag"},
		{Name: "prach-FreqOffset"},
		{Name: "prach-ConfigIndex", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type EUTRAPRACHConfiguration struct {
	RootSequenceIndex    int64
	ZeroCorrelationIndex int64
	HighSpeedFlag        EUTRAPRACHConfigurationHighSpeedFlag
	PrachFreqOffset      int64
	PrachConfigIndex     *int64
	IEExtensions         []byte
}

func (ie *EUTRAPRACHConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAPRACHConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PrachConfigIndex != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.RootSequenceIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(837)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.ZeroCorrelationIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(15)),
	}); err != nil {
		return err
	}
	if err := ie.HighSpeedFlag.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.PrachFreqOffset, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(94)),
	}); err != nil {
		return err
	}
	if ie.PrachConfigIndex != nil {
		if err := e.EncodeInteger(*ie.PrachConfigIndex, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(63)),
		}); err != nil {
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

func (ie *EUTRAPRACHConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAPRACHConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(837)),
		})
		if err != nil {
			return err
		}
		ie.RootSequenceIndex = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(15)),
		})
		if err != nil {
			return err
		}
		ie.ZeroCorrelationIndex = val
	}
	if err := ie.HighSpeedFlag.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(94)),
		})
		if err != nil {
			return err
		}
		ie.PrachFreqOffset = val
	}
	if seq.IsComponentPresent(4) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(63)),
		})
		if err != nil {
			return err
		}
		ie.PrachConfigIndex = &val
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
