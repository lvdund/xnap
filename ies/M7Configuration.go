package ies

import (
	"github.com/lvdund/asn1go/per"
)

var m7ConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m7period"},
		{Name: "m7-links-to-log"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type M7Configuration struct {
	M7period     M7period
	M7LinksToLog LinksToLog
	IEExtensions []byte
}

func (ie *M7Configuration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(m7ConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.M7period.Encode(e); err != nil {
		return err
	}
	if err := ie.M7LinksToLog.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *M7Configuration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(m7ConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.M7period.Decode(d); err != nil {
		return err
	}
	if err := ie.M7LinksToLog.Decode(d); err != nil {
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
