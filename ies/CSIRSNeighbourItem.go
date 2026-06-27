package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cSIRSNeighbourItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-cgi"},
		{Name: "csi-RS-MTC-Neighbour-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CSIRSNeighbourItem struct {
	NrCgi                 NRCGI
	CsiRSMTCNeighbourList *CSIRSMTCNeighbourList
	IEExtensions          []byte
}

func (ie *CSIRSNeighbourItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSNeighbourItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CsiRSMTCNeighbourList != nil, false}); err != nil {
		return err
	}
	if err := ie.NrCgi.Encode(e); err != nil {
		return err
	}
	if ie.CsiRSMTCNeighbourList != nil {
		if err := ie.CsiRSMTCNeighbourList.Encode(e); err != nil {
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

func (ie *CSIRSNeighbourItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSNeighbourItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrCgi.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.CsiRSMTCNeighbourList = new(CSIRSMTCNeighbourList)
		if err := ie.CsiRSMTCNeighbourList.Decode(d); err != nil {
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
