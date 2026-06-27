package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	RBsetConfigurationRBsetSizeRb2  int64 = 0
	RBsetConfigurationRBsetSizeRb4  int64 = 1
	RBsetConfigurationRBsetSizeRb8  int64 = 2
	RBsetConfigurationRBsetSizeRb16 int64 = 3
	RBsetConfigurationRBsetSizeRb32 int64 = 4
	RBsetConfigurationRBsetSizeRb64 int64 = 5
)

var rBsetConfigurationRBsetSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2, 3, 4, 5},
	ExtValues:  nil,
}

type RBsetConfigurationRBsetSize struct {
	Value int64
}

func (ie *RBsetConfigurationRBsetSize) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rBsetConfigurationRBsetSizeConstraints)
}

func (ie *RBsetConfigurationRBsetSize) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rBsetConfigurationRBsetSizeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var rBsetConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subcarrierSpacing"},
		{Name: "rBsetSize"},
		{Name: "numberofRBSets"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RBsetConfiguration struct {
	SubcarrierSpacing SubcarrierSpacing
	RBsetSize         RBsetConfigurationRBsetSize
	NumberofRBSets    int64
	IEExtensions      []byte
}

func (ie *RBsetConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rBsetConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SubcarrierSpacing.Encode(e); err != nil {
		return err
	}
	if err := ie.RBsetSize.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.NumberofRBSets, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(common.MaxnoofRBsetsPerCell)),
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

func (ie *RBsetConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rBsetConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SubcarrierSpacing.Decode(d); err != nil {
		return err
	}
	if err := ie.RBsetSize.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(common.MaxnoofRBsetsPerCell)),
		})
		if err != nil {
			return err
		}
		ie.NumberofRBSets = val
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
