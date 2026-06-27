package ies

import (
	"github.com/lvdund/asn1go/per"
)

var locationReportingInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventType"},
		{Name: "reportArea"},
		{Name: "areaOfInterest", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type LocationReportingInformation struct {
	EventType      EventType
	ReportArea     ReportArea
	AreaOfInterest *AreaOfInterestInformation
	IEExtensions   []byte
}

func (ie *LocationReportingInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(locationReportingInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AreaOfInterest != nil, false}); err != nil {
		return err
	}
	if err := ie.EventType.Encode(e); err != nil {
		return err
	}
	if err := ie.ReportArea.Encode(e); err != nil {
		return err
	}
	if ie.AreaOfInterest != nil {
		if err := ie.AreaOfInterest.Encode(e); err != nil {
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

func (ie *LocationReportingInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(locationReportingInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EventType.Decode(d); err != nil {
		return err
	}
	if err := ie.ReportArea.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.AreaOfInterest = new(AreaOfInterestInformation)
		if err := ie.AreaOfInterest.Decode(d); err != nil {
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
