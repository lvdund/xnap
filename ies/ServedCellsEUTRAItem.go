package ies

import (
	"github.com/lvdund/asn1go/per"
)

var servedCellsEUTRAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "served-cell-info-E-UTRA"},
		{Name: "neighbour-info-NR", Optional: true},
		{Name: "neighbour-info-E-UTRA", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellsEUTRAItem struct {
	ServedCellInfoEUTRA ServedCellInformationEUTRA
	NeighbourInfoNR     *NeighbourInformationNR
	NeighbourInfoEUTRA  *NeighbourInformationEUTRA
	IEExtensions        []byte
}

func (ie *ServedCellsEUTRAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellsEUTRAItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NeighbourInfoNR != nil, ie.NeighbourInfoEUTRA != nil, false}); err != nil {
		return err
	}
	if err := ie.ServedCellInfoEUTRA.Encode(e); err != nil {
		return err
	}
	if ie.NeighbourInfoNR != nil {
		if err := ie.NeighbourInfoNR.Encode(e); err != nil {
			return err
		}
	}
	if ie.NeighbourInfoEUTRA != nil {
		if err := ie.NeighbourInfoEUTRA.Encode(e); err != nil {
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

func (ie *ServedCellsEUTRAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellsEUTRAItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ServedCellInfoEUTRA.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.NeighbourInfoNR = new(NeighbourInformationNR)
		if err := ie.NeighbourInfoNR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.NeighbourInfoEUTRA = new(NeighbourInformationEUTRA)
		if err := ie.NeighbourInfoEUTRA.Decode(d); err != nil {
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
