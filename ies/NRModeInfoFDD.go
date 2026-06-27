package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRModeInfoFDDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ulNRFrequencyInfo"},
		{Name: "dlNRFrequencyInfo"},
		{Name: "ulNRTransmissonBandwidth"},
		{Name: "dlNRTransmissonBandwidth"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRModeInfoFDD struct {
	UlNRFrequencyInfo        NRFrequencyInfo
	DlNRFrequencyInfo        NRFrequencyInfo
	UlNRTransmissonBandwidth NRTransmissionBandwidth
	DlNRTransmissonBandwidth NRTransmissionBandwidth
	IEExtensions             []byte
}

func (ie *NRModeInfoFDD) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRModeInfoFDDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UlNRFrequencyInfo.Encode(e); err != nil {
		return err
	}
	if err := ie.DlNRFrequencyInfo.Encode(e); err != nil {
		return err
	}
	if err := ie.UlNRTransmissonBandwidth.Encode(e); err != nil {
		return err
	}
	if err := ie.DlNRTransmissonBandwidth.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRModeInfoFDD) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRModeInfoFDDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UlNRFrequencyInfo.Decode(d); err != nil {
		return err
	}
	if err := ie.DlNRFrequencyInfo.Decode(d); err != nil {
		return err
	}
	if err := ie.UlNRTransmissonBandwidth.Decode(d); err != nil {
		return err
	}
	if err := ie.DlNRTransmissonBandwidth.Decode(d); err != nil {
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
