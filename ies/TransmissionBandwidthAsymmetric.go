package ies

import (
	"github.com/lvdund/asn1go/per"
)

var transmissionBandwidthAsymmetricConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-Transmission-Bandwidth"},
		{Name: "dl-Transmission-Bandwidth"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TransmissionBandwidthAsymmetric struct {
	UlTransmissionBandwidth NRTransmissionBandwidth
	DlTransmissionBandwidth NRTransmissionBandwidth
	IEExtensions            []byte
}

func (ie *TransmissionBandwidthAsymmetric) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(transmissionBandwidthAsymmetricConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UlTransmissionBandwidth.Encode(e); err != nil {
		return err
	}
	if err := ie.DlTransmissionBandwidth.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TransmissionBandwidthAsymmetric) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(transmissionBandwidthAsymmetricConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UlTransmissionBandwidth.Decode(d); err != nil {
		return err
	}
	if err := ie.DlTransmissionBandwidth.Decode(d); err != nil {
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
