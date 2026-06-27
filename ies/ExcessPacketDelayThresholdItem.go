package ies

import (
	"github.com/lvdund/asn1go/per"
)

var excessPacketDelayThresholdItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fiveQI"},
		{Name: "excessPacketDelayThresholdValue"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExcessPacketDelayThresholdItem struct {
	FiveQI                          FiveQI
	ExcessPacketDelayThresholdValue ExcessPacketDelayThresholdValue
	IEExtensions                    []byte
}

func (ie *ExcessPacketDelayThresholdItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(excessPacketDelayThresholdItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.FiveQI.Encode(e); err != nil {
		return err
	}
	if err := ie.ExcessPacketDelayThresholdValue.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ExcessPacketDelayThresholdItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(excessPacketDelayThresholdItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FiveQI.Decode(d); err != nil {
		return err
	}
	if err := ie.ExcessPacketDelayThresholdValue.Decode(d); err != nil {
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
