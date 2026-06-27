package ies

import (
	"github.com/lvdund/asn1go/per"
)

var allowedPNINPNIDItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-id"},
		{Name: "pni-npn-restricted-information"},
		{Name: "allowed-CAG-id-list-per-plmn"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AllowedPNINPNIDItem struct {
	PlmnId                      PLMNIdentity
	PniNpnRestrictedInformation PNINPNRestrictedInformation
	AllowedCAGIdListPerPlmn     AllowedCAGIDListPerPLMN
	IEExtensions                []byte
}

func (ie *AllowedPNINPNIDItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(allowedPNINPNIDItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnId.Encode(e); err != nil {
		return err
	}
	if err := ie.PniNpnRestrictedInformation.Encode(e); err != nil {
		return err
	}
	if err := ie.AllowedCAGIdListPerPlmn.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AllowedPNINPNIDItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(allowedPNINPNIDItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnId.Decode(d); err != nil {
		return err
	}
	if err := ie.PniNpnRestrictedInformation.Decode(d); err != nil {
		return err
	}
	if err := ie.AllowedCAGIdListPerPlmn.Decode(d); err != nil {
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
