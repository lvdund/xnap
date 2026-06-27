package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEAggregateMaximumBitRateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-UE-AMBR"},
		{Name: "ul-UE-AMBR"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEAggregateMaximumBitRate struct {
	DlUEAMBR     BitRate
	UlUEAMBR     BitRate
	IEExtensions []byte
}

func (ie *UEAggregateMaximumBitRate) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAggregateMaximumBitRateConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DlUEAMBR.Encode(e); err != nil {
		return err
	}
	if err := ie.UlUEAMBR.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UEAggregateMaximumBitRate) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAggregateMaximumBitRateConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DlUEAMBR.Decode(d); err != nil {
		return err
	}
	if err := ie.UlUEAMBR.Decode(d); err != nil {
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
