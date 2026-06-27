package ies

import (
	"github.com/lvdund/asn1go/per"
)

var beamMeasurementsReportConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "beamMeasurementsReportQuantity", Optional: true},
		{Name: "maxNrofRS-IndexesToReport", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BeamMeasurementsReportConfiguration struct {
	BeamMeasurementsReportQuantity *BeamMeasurementsReportQuantity
	MaxNrofRSIndexesToReport       *MaxNrofRSIndexesToReport
	IEExtensions                   []byte
}

func (ie *BeamMeasurementsReportConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamMeasurementsReportConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.BeamMeasurementsReportQuantity != nil, ie.MaxNrofRSIndexesToReport != nil, false}); err != nil {
		return err
	}
	if ie.BeamMeasurementsReportQuantity != nil {
		if err := ie.BeamMeasurementsReportQuantity.Encode(e); err != nil {
			return err
		}
	}
	if ie.MaxNrofRSIndexesToReport != nil {
		if err := ie.MaxNrofRSIndexesToReport.Encode(e); err != nil {
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

func (ie *BeamMeasurementsReportConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamMeasurementsReportConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.BeamMeasurementsReportQuantity = new(BeamMeasurementsReportQuantity)
		if err := ie.BeamMeasurementsReportQuantity.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.MaxNrofRSIndexesToReport = new(MaxNrofRSIndexesToReport)
		if err := ie.MaxNrofRSIndexesToReport.Decode(d); err != nil {
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
