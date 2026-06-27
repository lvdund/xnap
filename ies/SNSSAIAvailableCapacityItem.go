package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNSSAIAvailableCapacityItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sNSSAI"},
		{Name: "sliceAvailableCapacityValueDownlink"},
		{Name: "sliceAvailableCapacityValueUplink"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type SNSSAIAvailableCapacityItem struct {
	SNSSAI                              SNSSAI
	SliceAvailableCapacityValueDownlink int64
	SliceAvailableCapacityValueUplink   int64
	IEExtensions                        []byte
}

func (ie *SNSSAIAvailableCapacityItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNSSAIAvailableCapacityItemConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SNSSAI.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SliceAvailableCapacityValueDownlink, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(100)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SliceAvailableCapacityValueUplink, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(100)),
	}); err != nil {
		return err
	}
	return nil
}

func (ie *SNSSAIAvailableCapacityItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNSSAIAvailableCapacityItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SNSSAI.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(100)),
		})
		if err != nil {
			return err
		}
		ie.SliceAvailableCapacityValueDownlink = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(100)),
		})
		if err != nil {
			return err
		}
		ie.SliceAvailableCapacityValueUplink = val
	}
	return nil
}
