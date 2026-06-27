package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ServedCellsToModifyEUTRAItemDeactivationIndicationDeactivated int64 = 0
)

var servedCellsToModifyEUTRAItemDeactivationIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ServedCellsToModifyEUTRAItemDeactivationIndication struct {
	Value int64
}

func (ie *ServedCellsToModifyEUTRAItemDeactivationIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, servedCellsToModifyEUTRAItemDeactivationIndicationConstraints)
}

func (ie *ServedCellsToModifyEUTRAItemDeactivationIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(servedCellsToModifyEUTRAItemDeactivationIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var servedCellsToModifyEUTRAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "old-ECGI"},
		{Name: "served-cell-info-E-UTRA"},
		{Name: "neighbour-info-NR", Optional: true},
		{Name: "neighbour-info-E-UTRA", Optional: true},
		{Name: "deactivation-indication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellsToModifyEUTRAItem struct {
	OldECGI                EUTRACGI
	ServedCellInfoEUTRA    ServedCellInformationEUTRA
	NeighbourInfoNR        *NeighbourInformationNR
	NeighbourInfoEUTRA     *NeighbourInformationEUTRA
	DeactivationIndication *ServedCellsToModifyEUTRAItemDeactivationIndication
	IEExtensions           []byte
}

func (ie *ServedCellsToModifyEUTRAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellsToModifyEUTRAItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NeighbourInfoNR != nil, ie.NeighbourInfoEUTRA != nil, ie.DeactivationIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.OldECGI.Encode(e); err != nil {
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
	if ie.DeactivationIndication != nil {
		if err := ie.DeactivationIndication.Encode(e); err != nil {
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

func (ie *ServedCellsToModifyEUTRAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellsToModifyEUTRAItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.OldECGI.Decode(d); err != nil {
		return err
	}
	if err := ie.ServedCellInfoEUTRA.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.NeighbourInfoNR = new(NeighbourInformationNR)
		if err := ie.NeighbourInfoNR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.NeighbourInfoEUTRA = new(NeighbourInformationEUTRA)
		if err := ie.NeighbourInfoEUTRA.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.DeactivationIndication = new(ServedCellsToModifyEUTRAItemDeactivationIndication)
		if err := ie.DeactivationIndication.Decode(d); err != nil {
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
