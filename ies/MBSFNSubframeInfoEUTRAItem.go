package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodN1  int64 = 0
	MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodN2  int64 = 1
	MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodN4  int64 = 2
	MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodN8  int64 = 3
	MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodN16 int64 = 4
	MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodN32 int64 = 5
)

var mBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2, 3, 4, 5},
	ExtValues:  nil,
}

type MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriod struct {
	Value int64
}

func (ie *MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriod) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodConstraints)
}

func (ie *MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriod) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriodConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var mBSFNSubframeInfoEUTRAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "radioframeAllocationPeriod"},
		{Name: "radioframeAllocationOffset"},
		{Name: "subframeAllocation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSFNSubframeInfoEUTRAItem struct {
	RadioframeAllocationPeriod MBSFNSubframeInfoEUTRAItemRadioframeAllocationPeriod
	RadioframeAllocationOffset int64
	SubframeAllocation         MBSFNSubframeAllocationEUTRA
	IEExtensions               []byte
}

func (ie *MBSFNSubframeInfoEUTRAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSFNSubframeInfoEUTRAItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.RadioframeAllocationPeriod.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.RadioframeAllocationOffset, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(7)),
	}); err != nil {
		return err
	}
	if err := ie.SubframeAllocation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MBSFNSubframeInfoEUTRAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSFNSubframeInfoEUTRAItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RadioframeAllocationPeriod.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(7)),
		})
		if err != nil {
			return err
		}
		ie.RadioframeAllocationOffset = val
	}
	if err := ie.SubframeAllocation.Decode(d); err != nil {
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
