package ies

import (
	"github.com/lvdund/asn1go/per"
)

var neighbourInformationEUTRAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "e-utra-PCI"},
		{Name: "e-utra-cgi"},
		{Name: "earfcn"},
		{Name: "tac"},
		{Name: "ranac", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NeighbourInformationEUTRAItem struct {
	EUtraPCI     EUTRAPCI
	EUtraCgi     EUTRACGI
	Earfcn       EUTRAARFCN
	Tac          TAC
	Ranac        *RANAC
	IEExtensions []byte
}

func (ie *NeighbourInformationEUTRAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(neighbourInformationEUTRAItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Ranac != nil, false}); err != nil {
		return err
	}
	if err := ie.EUtraPCI.Encode(e); err != nil {
		return err
	}
	if err := ie.EUtraCgi.Encode(e); err != nil {
		return err
	}
	if err := ie.Earfcn.Encode(e); err != nil {
		return err
	}
	if err := ie.Tac.Encode(e); err != nil {
		return err
	}
	if ie.Ranac != nil {
		if err := ie.Ranac.Encode(e); err != nil {
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

func (ie *NeighbourInformationEUTRAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(neighbourInformationEUTRAItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EUtraPCI.Decode(d); err != nil {
		return err
	}
	if err := ie.EUtraCgi.Decode(d); err != nil {
		return err
	}
	if err := ie.Earfcn.Decode(d); err != nil {
		return err
	}
	if err := ie.Tac.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.Ranac = new(RANAC)
		if err := ie.Ranac.Decode(d); err != nil {
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
