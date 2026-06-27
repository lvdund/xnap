package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sDTSRBsToBeSetupListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srb-ID"},
		{Name: "sRB-RLC-Bearer-Configuration"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SDTSRBsToBeSetupListItem struct {
	SrbID                     SRBID
	SRBRLCBearerConfiguration []byte
	IEExtensions              []byte
}

func (ie *SDTSRBsToBeSetupListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTSRBsToBeSetupListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SrbID.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.SRBRLCBearerConfiguration, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SDTSRBsToBeSetupListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTSRBsToBeSetupListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SrbID.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.SRBRLCBearerConfiguration = val
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
