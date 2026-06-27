package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mIMOPRBusageInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-GBR-PRB-usage-for-MIMO"},
		{Name: "ul-GBR-PRB-usage-for-MIMO"},
		{Name: "dl-non-GBR-PRB-usage-for-MIMO"},
		{Name: "ul-non-GBR-PRB-usage-for-MIMO"},
		{Name: "dl-Total-PRB-usage-for-MIMO"},
		{Name: "ul-Total-PRB-usage-for-MIMO"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MIMOPRBusageInformation struct {
	DlGBRPRBUsageForMIMO    DLGBRPRBUsageForMIMO
	UlGBRPRBUsageForMIMO    ULGBRPRBUsageForMIMO
	DlNonGBRPRBUsageForMIMO DLNonGBRPRBUsageForMIMO
	UlNonGBRPRBUsageForMIMO ULNonGBRPRBUsageForMIMO
	DlTotalPRBUsageForMIMO  DLTotalPRBUsageForMIMO
	UlTotalPRBUsageForMIMO  ULTotalPRBUsageForMIMO
	IEExtensions            []byte
}

func (ie *MIMOPRBusageInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mIMOPRBusageInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DlGBRPRBUsageForMIMO.Encode(e); err != nil {
		return err
	}
	if err := ie.UlGBRPRBUsageForMIMO.Encode(e); err != nil {
		return err
	}
	if err := ie.DlNonGBRPRBUsageForMIMO.Encode(e); err != nil {
		return err
	}
	if err := ie.UlNonGBRPRBUsageForMIMO.Encode(e); err != nil {
		return err
	}
	if err := ie.DlTotalPRBUsageForMIMO.Encode(e); err != nil {
		return err
	}
	if err := ie.UlTotalPRBUsageForMIMO.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MIMOPRBusageInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mIMOPRBusageInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DlGBRPRBUsageForMIMO.Decode(d); err != nil {
		return err
	}
	if err := ie.UlGBRPRBUsageForMIMO.Decode(d); err != nil {
		return err
	}
	if err := ie.DlNonGBRPRBUsageForMIMO.Decode(d); err != nil {
		return err
	}
	if err := ie.UlNonGBRPRBUsageForMIMO.Decode(d); err != nil {
		return err
	}
	if err := ie.DlTotalPRBUsageForMIMO.Decode(d); err != nil {
		return err
	}
	if err := ie.UlTotalPRBUsageForMIMO.Decode(d); err != nil {
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
