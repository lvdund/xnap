package ies

import (
	"github.com/lvdund/asn1go/per"
)

var m1PeriodicReportingConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportInterval"},
		{Name: "reportAmount"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type M1PeriodicReporting struct {
	ReportInterval ReportIntervalMDT
	ReportAmount   ReportAmountMDT
	IEExtensions   []byte
}

func (ie *M1PeriodicReporting) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(m1PeriodicReportingConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.ReportInterval.Encode(e); err != nil {
		return err
	}
	if err := ie.ReportAmount.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *M1PeriodicReporting) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(m1PeriodicReportingConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ReportInterval.Decode(d); err != nil {
		return err
	}
	if err := ie.ReportAmount.Decode(d); err != nil {
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
