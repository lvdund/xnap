package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nodeMeasurementInitiationResultItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nodemeasurementFailedReportCharacteristics"},
		{Name: "cause"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NodeMeasurementInitiationResultItem struct {
	NodemeasurementFailedReportCharacteristics per.BitString
	Cause                                      Cause
	IEExtensions                               []byte
}

func (ie *NodeMeasurementInitiationResultItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nodeMeasurementInitiationResultItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.NodemeasurementFailedReportCharacteristics, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(32)),
		Max:        common.Ptr(int64(32)),
	}); err != nil {
		return err
	}
	if err := ie.Cause.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NodeMeasurementInitiationResultItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nodeMeasurementInitiationResultItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(32)),
			Max:        common.Ptr(int64(32)),
		})
		if err != nil {
			return err
		}
		ie.NodemeasurementFailedReportCharacteristics = val
	}
	if err := ie.Cause.Decode(d); err != nil {
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
