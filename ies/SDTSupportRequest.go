package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sDTSupportRequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sdtindicator"},
		{Name: "sdtAssistantInfo", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SDTSupportRequest struct {
	Sdtindicator     SDTIndicator
	SdtAssistantInfo *SDTAssistantInfo
	IEExtensions     []byte
}

func (ie *SDTSupportRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTSupportRequestConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SdtAssistantInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.Sdtindicator.Encode(e); err != nil {
		return err
	}
	if ie.SdtAssistantInfo != nil {
		if err := ie.SdtAssistantInfo.Encode(e); err != nil {
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

func (ie *SDTSupportRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTSupportRequestConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Sdtindicator.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SdtAssistantInfo = new(SDTAssistantInfo)
		if err := ie.SdtAssistantInfo.Decode(d); err != nil {
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
