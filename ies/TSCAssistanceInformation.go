package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tSCAssistanceInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicity"},
		{Name: "burstArrivalTime", Optional: true},
		{Name: "ie-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TSCAssistanceInformation struct {
	Periodicity      int64
	BurstArrivalTime []byte
	IEExtensions     []byte
}

func (ie *TSCAssistanceInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tSCAssistanceInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.BurstArrivalTime) > 0, false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.Periodicity, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(640000)),
	}); err != nil {
		return err
	}
	if len(ie.BurstArrivalTime) > 0 {
		if err := e.EncodeOctetString(ie.BurstArrivalTime, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
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

func (ie *TSCAssistanceInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tSCAssistanceInformationConstraints)
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
			UpperBound: common.Ptr(int64(640000)),
		})
		if err != nil {
			return err
		}
		ie.Periodicity = val
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.BurstArrivalTime = val
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
