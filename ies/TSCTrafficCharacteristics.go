package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tSCTrafficCharacteristicsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tSCAssistanceInformationDownlink", Optional: true},
		{Name: "tSCAssistanceInformationUplink", Optional: true},
		{Name: "ie-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TSCTrafficCharacteristics struct {
	TSCAssistanceInformationDownlink *TSCAssistanceInformation
	TSCAssistanceInformationUplink   *TSCAssistanceInformation
	IEExtensions                     []byte
}

func (ie *TSCTrafficCharacteristics) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tSCTrafficCharacteristicsConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.TSCAssistanceInformationDownlink != nil, ie.TSCAssistanceInformationUplink != nil, false}); err != nil {
		return err
	}
	if ie.TSCAssistanceInformationDownlink != nil {
		if err := ie.TSCAssistanceInformationDownlink.Encode(e); err != nil {
			return err
		}
	}
	if ie.TSCAssistanceInformationUplink != nil {
		if err := ie.TSCAssistanceInformationUplink.Encode(e); err != nil {
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

func (ie *TSCTrafficCharacteristics) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tSCTrafficCharacteristicsConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.TSCAssistanceInformationDownlink = new(TSCAssistanceInformation)
		if err := ie.TSCAssistanceInformationDownlink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.TSCAssistanceInformationUplink = new(TSCAssistanceInformation)
		if err := ie.TSCAssistanceInformationUplink.Decode(d); err != nil {
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
