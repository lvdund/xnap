package ies

import (
	"github.com/lvdund/asn1go/per"
)

var m5ConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m5period"},
		{Name: "m5-links-to-log"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type M5Configuration struct {
	M5period     M5period
	M5LinksToLog LinksToLog
	IEExtensions []byte
}

func (ie *M5Configuration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(m5ConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.M5period.Encode(e); err != nil {
		return err
	}
	if err := ie.M5LinksToLog.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *M5Configuration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(m5ConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.M5period.Decode(d); err != nil {
		return err
	}
	if err := ie.M5LinksToLog.Decode(d); err != nil {
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
