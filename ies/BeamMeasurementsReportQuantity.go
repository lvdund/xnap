package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	BeamMeasurementsReportQuantityRSRPTrue int64 = 0
)

var beamMeasurementsReportQuantityRSRPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BeamMeasurementsReportQuantityRSRP struct {
	Value int64
}

func (ie *BeamMeasurementsReportQuantityRSRP) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, beamMeasurementsReportQuantityRSRPConstraints)
}

func (ie *BeamMeasurementsReportQuantityRSRP) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(beamMeasurementsReportQuantityRSRPConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	BeamMeasurementsReportQuantityRSRQTrue int64 = 0
)

var beamMeasurementsReportQuantityRSRQConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BeamMeasurementsReportQuantityRSRQ struct {
	Value int64
}

func (ie *BeamMeasurementsReportQuantityRSRQ) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, beamMeasurementsReportQuantityRSRQConstraints)
}

func (ie *BeamMeasurementsReportQuantityRSRQ) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(beamMeasurementsReportQuantityRSRQConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	BeamMeasurementsReportQuantitySINRTrue int64 = 0
)

var beamMeasurementsReportQuantitySINRConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BeamMeasurementsReportQuantitySINR struct {
	Value int64
}

func (ie *BeamMeasurementsReportQuantitySINR) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, beamMeasurementsReportQuantitySINRConstraints)
}

func (ie *BeamMeasurementsReportQuantitySINR) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(beamMeasurementsReportQuantitySINRConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var beamMeasurementsReportQuantityConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rSRP"},
		{Name: "rSRQ"},
		{Name: "sINR"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BeamMeasurementsReportQuantity struct {
	RSRP         BeamMeasurementsReportQuantityRSRP
	RSRQ         BeamMeasurementsReportQuantityRSRQ
	SINR         BeamMeasurementsReportQuantitySINR
	IEExtensions []byte
}

func (ie *BeamMeasurementsReportQuantity) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamMeasurementsReportQuantityConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.RSRP.Encode(e); err != nil {
		return err
	}
	if err := ie.RSRQ.Encode(e); err != nil {
		return err
	}
	if err := ie.SINR.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BeamMeasurementsReportQuantity) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamMeasurementsReportQuantityConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RSRP.Decode(d); err != nil {
		return err
	}
	if err := ie.RSRQ.Decode(d); err != nil {
		return err
	}
	if err := ie.SINR.Decode(d); err != nil {
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
