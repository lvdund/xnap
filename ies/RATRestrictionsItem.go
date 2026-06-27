package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rATRestrictionsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity"},
		{Name: "rat-RestrictionInformation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RATRestrictionsItem struct {
	PlmnIdentity              PLMNIdentity
	RatRestrictionInformation RATRestrictionInformation
	IEExtensions              []byte
}

func (ie *RATRestrictionsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rATRestrictionsItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Encode(e); err != nil {
		return err
	}
	if err := ie.RatRestrictionInformation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RATRestrictionsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rATRestrictionsItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Decode(d); err != nil {
		return err
	}
	if err := ie.RatRestrictionInformation.Decode(d); err != nil {
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
