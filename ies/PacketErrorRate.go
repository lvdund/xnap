package ies

import (
	"github.com/lvdund/asn1go/per"
)

var packetErrorRateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pER-Scalar"},
		{Name: "pER-Exponent"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PacketErrorRate struct {
	PERScalar    PERScalar
	PERExponent  PERExponent
	IEExtensions []byte
}

func (ie *PacketErrorRate) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(packetErrorRateConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PERScalar.Encode(e); err != nil {
		return err
	}
	if err := ie.PERExponent.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PacketErrorRate) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(packetErrorRateConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PERScalar.Decode(d); err != nil {
		return err
	}
	if err := ie.PERExponent.Decode(d); err != nil {
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
