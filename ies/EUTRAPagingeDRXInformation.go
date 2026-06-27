package ies

import (
	"github.com/lvdund/asn1go/per"
)

var eUTRAPagingeDRXInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eutrapaging-eDRX-Cycle"},
		{Name: "eutrapaging-Time-Window", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type EUTRAPagingeDRXInformation struct {
	EutrapagingEDRXCycle  EUTRAPagingEDRXCycle
	EutrapagingTimeWindow *EUTRAPagingTimeWindow
	IEExtensions          []byte
}

func (ie *EUTRAPagingeDRXInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAPagingeDRXInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.EutrapagingTimeWindow != nil, false}); err != nil {
		return err
	}
	if err := ie.EutrapagingEDRXCycle.Encode(e); err != nil {
		return err
	}
	if ie.EutrapagingTimeWindow != nil {
		if err := ie.EutrapagingTimeWindow.Encode(e); err != nil {
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

func (ie *EUTRAPagingeDRXInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAPagingeDRXInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EutrapagingEDRXCycle.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.EutrapagingTimeWindow = new(EUTRAPagingTimeWindow)
		if err := ie.EutrapagingTimeWindow.Decode(d); err != nil {
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
