package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ServedCellsToModifyNRItemDeactivationIndicationDeactivated int64 = 0
)

var servedCellsToModifyNRItemDeactivationIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ServedCellsToModifyNRItemDeactivationIndication struct {
	Value int64
}

func (ie *ServedCellsToModifyNRItemDeactivationIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, servedCellsToModifyNRItemDeactivationIndicationConstraints)
}

func (ie *ServedCellsToModifyNRItemDeactivationIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(servedCellsToModifyNRItemDeactivationIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var servedCellsToModifyNRItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "old-NR-CGI"},
		{Name: "served-cell-info-NR"},
		{Name: "neighbour-info-NR", Optional: true},
		{Name: "neighbour-info-E-UTRA", Optional: true},
		{Name: "deactivation-indication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellsToModifyNRItem struct {
	OldNRCGI               NRCGI
	ServedCellInfoNR       ServedCellInformationNR
	NeighbourInfoNR        *NeighbourInformationNR
	NeighbourInfoEUTRA     *NeighbourInformationEUTRA
	DeactivationIndication *ServedCellsToModifyNRItemDeactivationIndication
	IEExtensions           []byte
}

func (ie *ServedCellsToModifyNRItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellsToModifyNRItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NeighbourInfoNR != nil, ie.NeighbourInfoEUTRA != nil, ie.DeactivationIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.OldNRCGI.Encode(e); err != nil {
		return err
	}
	if err := ie.ServedCellInfoNR.Encode(e); err != nil {
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

func (ie *ServedCellsToModifyNRItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellsToModifyNRItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.OldNRCGI.Decode(d); err != nil {
		return err
	}
	if err := ie.ServedCellInfoNR.Decode(d); err != nil {
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
		ie.DeactivationIndication = new(ServedCellsToModifyNRItemDeactivationIndication)
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
