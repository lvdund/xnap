package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rRCSetupInitiatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rRRCSetup-Initiated-Reporting"},
		{Name: "uERLFReportContainer", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RRCSetupInitiated struct {
	RRRCSetupInitiatedReporting RRCSetupInitiatedReporting
	UERLFReportContainer        *UERLFReportContainer
	IEExtensions                []byte
}

func (ie *RRCSetupInitiated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupInitiatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.UERLFReportContainer != nil, false}); err != nil {
		return err
	}
	if err := ie.RRRCSetupInitiatedReporting.Encode(e); err != nil {
		return err
	}
	if ie.UERLFReportContainer != nil {
		if err := ie.UERLFReportContainer.Encode(e); err != nil {
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

func (ie *RRCSetupInitiated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupInitiatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RRRCSetupInitiatedReporting.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.UERLFReportContainer = new(UERLFReportContainer)
		if err := ie.UERLFReportContainer.Decode(d); err != nil {
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
