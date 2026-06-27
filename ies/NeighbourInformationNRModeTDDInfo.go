package ies

import (
	"github.com/lvdund/asn1go/per"
)

var neighbourInformationNRModeTDDInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-FreqInfo"},
		{Name: "ie-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NeighbourInformationNRModeTDDInfo struct {
	NrFreqInfo   NRFrequencyInfo
	IEExtensions []byte
}

func (ie *NeighbourInformationNRModeTDDInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(neighbourInformationNRModeTDDInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NrFreqInfo.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NeighbourInformationNRModeTDDInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(neighbourInformationNRModeTDDInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrFreqInfo.Decode(d); err != nil {
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
