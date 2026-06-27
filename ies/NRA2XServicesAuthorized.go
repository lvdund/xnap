package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRA2XServicesAuthorizedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aerialUE", Optional: true},
		{Name: "aerialControllerUE", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRA2XServicesAuthorized struct {
	AerialUE           *AerialUE
	AerialControllerUE *AerialControllerUE
	IEExtensions       []byte
}

func (ie *NRA2XServicesAuthorized) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRA2XServicesAuthorizedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AerialUE != nil, ie.AerialControllerUE != nil, false}); err != nil {
		return err
	}
	if ie.AerialUE != nil {
		if err := ie.AerialUE.Encode(e); err != nil {
			return err
		}
	}
	if ie.AerialControllerUE != nil {
		if err := ie.AerialControllerUE.Encode(e); err != nil {
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

func (ie *NRA2XServicesAuthorized) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRA2XServicesAuthorizedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.AerialUE = new(AerialUE)
		if err := ie.AerialUE.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.AerialControllerUE = new(AerialControllerUE)
		if err := ie.AerialControllerUE.Decode(d); err != nil {
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
