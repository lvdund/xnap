package ies

import (
	"github.com/lvdund/asn1go/per"
)

var m4ConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m4period"},
		{Name: "m4-links-to-log"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type M4Configuration struct {
	M4period     M4period
	M4LinksToLog LinksToLog
	IEExtensions []byte
}

func (ie *M4Configuration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(m4ConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.M4period.Encode(e); err != nil {
		return err
	}
	if err := ie.M4LinksToLog.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *M4Configuration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(m4ConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.M4period.Decode(d); err != nil {
		return err
	}
	if err := ie.M4LinksToLog.Decode(d); err != nil {
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
