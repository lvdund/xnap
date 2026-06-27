package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rVQoEConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "availableRANVisibleQoEMetrics", Optional: true},
		{Name: "reportingPeriodicity", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RVQoEConfig struct {
	AvailableRANVisibleQoEMetrics *AvailableRVQoEMetrics
	ReportingPeriodicity          *RVQoEReportingPeriodicity
	IEExtensions                  []byte
}

func (ie *RVQoEConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rVQoEConfigConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AvailableRANVisibleQoEMetrics != nil, ie.ReportingPeriodicity != nil, false}); err != nil {
		return err
	}
	if ie.AvailableRANVisibleQoEMetrics != nil {
		if err := ie.AvailableRANVisibleQoEMetrics.Encode(e); err != nil {
			return err
		}
	}
	if ie.ReportingPeriodicity != nil {
		if err := ie.ReportingPeriodicity.Encode(e); err != nil {
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

func (ie *RVQoEConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rVQoEConfigConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.AvailableRANVisibleQoEMetrics = new(AvailableRVQoEMetrics)
		if err := ie.AvailableRANVisibleQoEMetrics.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ReportingPeriodicity = new(RVQoEReportingPeriodicity)
		if err := ie.ReportingPeriodicity.Decode(d); err != nil {
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
