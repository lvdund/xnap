package ies

import (
	"github.com/lvdund/asn1go/per"
)

var servedCellInformationEUTRAFDDInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-earfcn"},
		{Name: "dl-earfcn"},
		{Name: "ul-e-utraTxBW"},
		{Name: "dl-e-utraTxBW"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellInformationEUTRAFDDInfo struct {
	UlEarfcn     EUTRAARFCN
	DlEarfcn     EUTRAARFCN
	UlEUtraTxBW  EUTRATransmissionBandwidth
	DlEUtraTxBW  EUTRATransmissionBandwidth
	IEExtensions []byte
}

func (ie *ServedCellInformationEUTRAFDDInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellInformationEUTRAFDDInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UlEarfcn.Encode(e); err != nil {
		return err
	}
	if err := ie.DlEarfcn.Encode(e); err != nil {
		return err
	}
	if err := ie.UlEUtraTxBW.Encode(e); err != nil {
		return err
	}
	if err := ie.DlEUtraTxBW.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ServedCellInformationEUTRAFDDInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellInformationEUTRAFDDInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UlEarfcn.Decode(d); err != nil {
		return err
	}
	if err := ie.DlEarfcn.Decode(d); err != nil {
		return err
	}
	if err := ie.UlEUtraTxBW.Decode(d); err != nil {
		return err
	}
	if err := ie.DlEUtraTxBW.Decode(d); err != nil {
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
