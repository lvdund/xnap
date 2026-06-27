package ies

import (
	"github.com/lvdund/asn1go/per"
)

var averagePacketDelayConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uL-AveragePacketDelay"},
		{Name: "dL-AveragePacketDelay"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AveragePacketDelay struct {
	ULAveragePacketDelay AveragePacketDelayValue
	DLAveragePacketDelay AveragePacketDelayValue
	IEExtensions         []byte
}

func (ie *AveragePacketDelay) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(averagePacketDelayConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.ULAveragePacketDelay.Encode(e); err != nil {
		return err
	}
	if err := ie.DLAveragePacketDelay.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AveragePacketDelay) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(averagePacketDelayConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ULAveragePacketDelay.Decode(d); err != nil {
		return err
	}
	if err := ie.DLAveragePacketDelay.Decode(d); err != nil {
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
