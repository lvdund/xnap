package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ReservedSubframePatternSubframeTypeMbsfn    int64 = 0
	ReservedSubframePatternSubframeTypeNonMbsfn int64 = 1
)

var reservedSubframePatternSubframeTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type ReservedSubframePatternSubframeType struct {
	Value int64
}

func (ie *ReservedSubframePatternSubframeType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reservedSubframePatternSubframeTypeConstraints)
}

func (ie *ReservedSubframePatternSubframeType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reservedSubframePatternSubframeTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var reservedSubframePatternConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subframeType"},
		{Name: "reservedSubframePattern"},
		{Name: "mbsfnControlRegionLength", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ReservedSubframePattern struct {
	SubframeType             ReservedSubframePatternSubframeType
	ReservedSubframePattern  per.BitString
	MbsfnControlRegionLength *MBSFNControlRegionLength
	IEExtensions             []byte
}

func (ie *ReservedSubframePattern) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reservedSubframePatternConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MbsfnControlRegionLength != nil, false}); err != nil {
		return err
	}
	if err := ie.SubframeType.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.ReservedSubframePattern, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(10)),
		Max:        common.Ptr(int64(160)),
	}); err != nil {
		return err
	}
	if ie.MbsfnControlRegionLength != nil {
		if err := ie.MbsfnControlRegionLength.Encode(e); err != nil {
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

func (ie *ReservedSubframePattern) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reservedSubframePatternConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SubframeType.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(10)),
			Max:        common.Ptr(int64(160)),
		})
		if err != nil {
			return err
		}
		ie.ReservedSubframePattern = val
	}
	if seq.IsComponentPresent(2) {
		ie.MbsfnControlRegionLength = new(MBSFNControlRegionLength)
		if err := ie.MbsfnControlRegionLength.Decode(d); err != nil {
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
