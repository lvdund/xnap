package ies

import (
	"github.com/lvdund/asn1go/per"
)

var successfulPSCellChangeReportListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "successfulPSCellChangeReport"},
		{Name: "sNMobilityInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SuccessfulPSCellChangeReportListItem struct {
	SuccessfulPSCellChangeReport SuccessfulPSCellChangeReportContainer
	SNMobilityInformation        *SNMobilityInformation
	IEExtensions                 []byte
}

func (ie *SuccessfulPSCellChangeReportListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(successfulPSCellChangeReportListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SNMobilityInformation != nil, false}); err != nil {
		return err
	}
	if err := ie.SuccessfulPSCellChangeReport.Encode(e); err != nil {
		return err
	}
	if ie.SNMobilityInformation != nil {
		if err := ie.SNMobilityInformation.Encode(e); err != nil {
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

func (ie *SuccessfulPSCellChangeReportListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(successfulPSCellChangeReportListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SuccessfulPSCellChangeReport.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SNMobilityInformation = new(SNMobilityInformation)
		if err := ie.SNMobilityInformation.Decode(d); err != nil {
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
