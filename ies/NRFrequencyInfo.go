package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRFrequencyInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrARFCN"},
		{Name: "sul-information", Optional: true},
		{Name: "frequencyBand-List"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRFrequencyInfo struct {
	NrARFCN           NRARFCN
	SulInformation    *SULInformation
	FrequencyBandList NRFrequencyBandList
	IEExtensions      []byte
}

func (ie *NRFrequencyInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRFrequencyInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SulInformation != nil, false}); err != nil {
		return err
	}
	if err := ie.NrARFCN.Encode(e); err != nil {
		return err
	}
	if ie.SulInformation != nil {
		if err := ie.SulInformation.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.FrequencyBandList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRFrequencyInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRFrequencyInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrARFCN.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SulInformation = new(SULInformation)
		if err := ie.SulInformation.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.FrequencyBandList.Decode(d); err != nil {
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
