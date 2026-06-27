package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERLFReportContainerLTEExtensionConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ueRLFReportContainerLTE"},
		{Name: "ueRLFReportContainerLTEExtendBand"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UERLFReportContainerLTEExtension struct {
	UeRLFReportContainerLTE           UERLFReportContainerLTE
	UeRLFReportContainerLTEExtendBand UERLFReportContainerLTEExtendBand
	IEExtensions                      []byte
}

func (ie *UERLFReportContainerLTEExtension) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uERLFReportContainerLTEExtensionConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UeRLFReportContainerLTE.Encode(e); err != nil {
		return err
	}
	if err := ie.UeRLFReportContainerLTEExtendBand.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UERLFReportContainerLTEExtension) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uERLFReportContainerLTEExtensionConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UeRLFReportContainerLTE.Decode(d); err != nil {
		return err
	}
	if err := ie.UeRLFReportContainerLTEExtendBand.Decode(d); err != nil {
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
