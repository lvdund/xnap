package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRModeInfoTDDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrFrequencyInfo"},
		{Name: "nrTransmissonBandwidth"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRModeInfoTDD struct {
	NrFrequencyInfo        NRFrequencyInfo
	NrTransmissonBandwidth NRTransmissionBandwidth
	IEExtensions           []byte
}

func (ie *NRModeInfoTDD) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRModeInfoTDDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NrFrequencyInfo.Encode(e); err != nil {
		return err
	}
	if err := ie.NrTransmissonBandwidth.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRModeInfoTDD) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRModeInfoTDDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrFrequencyInfo.Decode(d); err != nil {
		return err
	}
	if err := ie.NrTransmissonBandwidth.Decode(d); err != nil {
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
