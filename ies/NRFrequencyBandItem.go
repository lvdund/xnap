package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRFrequencyBandItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-frequency-band"},
		{Name: "supported-SUL-Band-List", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRFrequencyBandItem struct {
	NrFrequencyBand      NRFrequencyBand
	SupportedSULBandList *SupportedSULBandList
	IEExtensions         []byte
}

func (ie *NRFrequencyBandItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRFrequencyBandItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SupportedSULBandList != nil, false}); err != nil {
		return err
	}
	if err := ie.NrFrequencyBand.Encode(e); err != nil {
		return err
	}
	if ie.SupportedSULBandList != nil {
		if err := ie.SupportedSULBandList.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRFrequencyBandItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRFrequencyBandItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrFrequencyBand.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SupportedSULBandList = new(SupportedSULBandList)
		if err := ie.SupportedSULBandList.Decode(d); err != nil {
			return err
		}
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
