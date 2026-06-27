package ies

import (
	"github.com/lvdund/asn1go/per"
)

var m6ConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m6report-Interval"},
		{Name: "m6-links-to-log"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type M6Configuration struct {
	M6reportInterval M6reportInterval
	M6LinksToLog     LinksToLog
	IEExtensions     []byte
}

func (ie *M6Configuration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(m6ConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.M6reportInterval.Encode(e); err != nil {
		return err
	}
	if err := ie.M6LinksToLog.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *M6Configuration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(m6ConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.M6reportInterval.Decode(d); err != nil {
		return err
	}
	if err := ie.M6LinksToLog.Decode(d); err != nil {
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
