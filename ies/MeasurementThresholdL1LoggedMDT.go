package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MeasurementThresholdL1LoggedMDTChThresholdRSRP = 0
	MeasurementThresholdL1LoggedMDTChThresholdRSRQ = 1
	MeasurementThresholdL1LoggedMDTExtension       = -1
)

var measurementThresholdL1LoggedMDTConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "threshold-RSRP"},
		{Name: "threshold-RSRQ"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "choice-extension"},
	},
}

type MeasurementThresholdL1LoggedMDT struct {
	Choice          int
	ThresholdRSRP   *ThresholdRSRP
	ThresholdRSRQ   *ThresholdRSRQ
	ChoiceExtension []byte
}

func (ie *MeasurementThresholdL1LoggedMDT) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(measurementThresholdL1LoggedMDTConstraints)
	switch ie.Choice {
	case 0: // threshold-RSRP
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.ThresholdRSRP.Encode(e); err != nil {
			return err
		}
	case 1: // threshold-RSRQ
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ThresholdRSRQ.Encode(e); err != nil {
			return err
		}
	default: // extension
		if err := choice.EncodeChoice(0, true, ie.ChoiceExtension); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MeasurementThresholdL1LoggedMDT) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(measurementThresholdL1LoggedMDTConstraints)
	idx, isExt, extBytes, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	if isExt {
		ie.Choice = MeasurementThresholdL1LoggedMDTExtension
		ie.ChoiceExtension = extBytes
		return nil
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // threshold-RSRP
		ie.ThresholdRSRP = new(ThresholdRSRP)
		if err := ie.ThresholdRSRP.Decode(d); err != nil {
			return err
		}
	case 1: // threshold-RSRQ
		ie.ThresholdRSRQ = new(ThresholdRSRQ)
		if err := ie.ThresholdRSRQ.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
