package ies

import (
	"github.com/lvdund/asn1go/per"
)

var configurationRejectedByMNGRANNodeConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cause"},
		{Name: "m-NG-RANNode-to-S-NG-RANNode-Container", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ConfigurationRejectedByMNGRANNode struct {
	Cause                           Cause
	MNGRANNodeToSNGRANNodeContainer []byte
	IEExtensions                    []byte
}

func (ie *ConfigurationRejectedByMNGRANNode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(configurationRejectedByMNGRANNodeConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.MNGRANNodeToSNGRANNodeContainer) > 0, false}); err != nil {
		return err
	}
	if err := ie.Cause.Encode(e); err != nil {
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

func (ie *ConfigurationRejectedByMNGRANNode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(configurationRejectedByMNGRANNodeConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Cause.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
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
