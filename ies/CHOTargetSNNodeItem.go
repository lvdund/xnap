package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOTargetSNNodeItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "target-S-NG-RANnodeID"},
		{Name: "pduSessionResourcesAdmittedList"},
		{Name: "cho-Candidate-PSCells-list"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOTargetSNNodeItem struct {
	TargetSNGRANnodeID              GlobalNGRANNodeID
	PduSessionResourcesAdmittedList PDUSessionResourcesAdmittedList
	ChoCandidatePSCellsList         CHOCandidatePSCellsList
	IEExtensions                    []byte
}

func (ie *CHOTargetSNNodeItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOTargetSNNodeItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Encode(e); err != nil {
		return err
	}
	if err := ie.PduSessionResourcesAdmittedList.Encode(e); err != nil {
		return err
	}
	if err := ie.ChoCandidatePSCellsList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CHOTargetSNNodeItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOTargetSNNodeItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Decode(d); err != nil {
		return err
	}
	if err := ie.PduSessionResourcesAdmittedList.Decode(d); err != nil {
		return err
	}
	if err := ie.ChoCandidatePSCellsList.Decode(d); err != nil {
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
