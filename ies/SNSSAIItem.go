package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sNSSAIItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sNSSAI"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type SNSSAIItem struct {
	SNSSAI       SNSSAI
	IEExtensions []byte
}

func (ie *SNSSAIItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNSSAIItemConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SNSSAI.Encode(e); err != nil {
		return err
	}
	return nil
}

func (ie *SNSSAIItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNSSAIItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SNSSAI.Decode(d); err != nil {
		return err
	}
	return nil
}
