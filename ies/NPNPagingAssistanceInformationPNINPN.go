package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nPNPagingAssistanceInformationPNINPNConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "allowedPNI-NPN-ID-List"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NPNPagingAssistanceInformationPNINPN struct {
	AllowedPNINPNIDList AllowedPNINPNIDList
	IEExtensions        []byte
}

func (ie *NPNPagingAssistanceInformationPNINPN) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nPNPagingAssistanceInformationPNINPNConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.AllowedPNINPNIDList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NPNPagingAssistanceInformationPNINPN) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nPNPagingAssistanceInformationPNINPNConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.AllowedPNINPNIDList.Decode(d); err != nil {
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
