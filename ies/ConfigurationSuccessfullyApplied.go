package ies

import (
	"github.com/lvdund/asn1go/per"
)

var configurationSuccessfullyAppliedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m-NG-RANNode-to-S-NG-RANNode-Container", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ConfigurationSuccessfullyApplied struct {
	MNGRANNodeToSNGRANNodeContainer []byte
	IEExtensions                    []byte
}

func (ie *ConfigurationSuccessfullyApplied) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(configurationSuccessfullyAppliedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.MNGRANNodeToSNGRANNodeContainer) > 0, false}); err != nil {
		return err
	}
	if len(ie.MNGRANNodeToSNGRANNodeContainer) > 0 {
		if err := e.EncodeOctetString(ie.MNGRANNodeToSNGRANNodeContainer, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
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

func (ie *ConfigurationSuccessfullyApplied) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(configurationSuccessfullyAppliedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.MNGRANNodeToSNGRANNodeContainer = val
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
