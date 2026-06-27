package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dataTrafficResourceIndicationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "activationSFN"},
		{Name: "sharedResourceType"},
		{Name: "reservedSubframePattern", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DataTrafficResourceIndication struct {
	ActivationSFN           ActivationSFN
	SharedResourceType      SharedResourceType
	ReservedSubframePattern *ReservedSubframePattern
	IEExtensions            []byte
}

func (ie *DataTrafficResourceIndication) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataTrafficResourceIndicationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ReservedSubframePattern != nil, false}); err != nil {
		return err
	}
	if err := ie.ActivationSFN.Encode(e); err != nil {
		return err
	}
	if err := ie.SharedResourceType.Encode(e); err != nil {
		return err
	}
	if ie.ReservedSubframePattern != nil {
		if err := ie.ReservedSubframePattern.Encode(e); err != nil {
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

func (ie *DataTrafficResourceIndication) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataTrafficResourceIndicationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ActivationSFN.Decode(d); err != nil {
		return err
	}
	if err := ie.SharedResourceType.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.ReservedSubframePattern = new(ReservedSubframePattern)
		if err := ie.ReservedSubframePattern.Decode(d); err != nil {
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
