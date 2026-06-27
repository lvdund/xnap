package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uETrajectoryCollectionConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "collectionTimeDurationForUETrajectory"},
		{Name: "numberOfVisitedCells"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UETrajectoryCollectionConfiguration struct {
	CollectionTimeDurationForUETrajectory int64
	NumberOfVisitedCells                  int64
	IEExtensions                          []byte
}

func (ie *UETrajectoryCollectionConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uETrajectoryCollectionConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.CollectionTimeDurationForUETrajectory, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(4096)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.NumberOfVisitedCells, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(16)),
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

func (ie *UETrajectoryCollectionConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uETrajectoryCollectionConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(4096)),
		})
		if err != nil {
			return err
		}
		ie.CollectionTimeDurationForUETrajectory = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(16)),
		})
		if err != nil {
			return err
		}
		ie.NumberOfVisitedCells = val
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
