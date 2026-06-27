package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sULInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sulFrequencyInfo"},
		{Name: "sulTransmissionBandwidth"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SULInformation struct {
	SulFrequencyInfo         NRARFCN
	SulTransmissionBandwidth NRTransmissionBandwidth
	IEExtensions             []byte
}

func (ie *SULInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sULInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SulFrequencyInfo.Encode(e); err != nil {
		return err
	}
	if err := ie.SulTransmissionBandwidth.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SULInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sULInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SulFrequencyInfo.Decode(d); err != nil {
		return err
	}
	if err := ie.SulTransmissionBandwidth.Decode(d); err != nil {
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
