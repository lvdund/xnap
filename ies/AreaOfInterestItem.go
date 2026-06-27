package ies

import (
	"github.com/lvdund/asn1go/per"
)

var areaOfInterestItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "listOfTAIsinAoI", Optional: true},
		{Name: "listOfCellsinAoI", Optional: true},
		{Name: "listOfRANNodesinAoI", Optional: true},
		{Name: "requestReferenceID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AreaOfInterestItem struct {
	ListOfTAIsinAoI     *ListOfTAIsinAoI
	ListOfCellsinAoI    *ListOfCells
	ListOfRANNodesinAoI *ListOfRANNodesinAoI
	RequestReferenceID  RequestReferenceID
	IEExtensions        []byte
}

func (ie *AreaOfInterestItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(areaOfInterestItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ListOfTAIsinAoI != nil, ie.ListOfCellsinAoI != nil, ie.ListOfRANNodesinAoI != nil, false}); err != nil {
		return err
	}
	if ie.ListOfTAIsinAoI != nil {
		if err := ie.ListOfTAIsinAoI.Encode(e); err != nil {
			return err
		}
	}
	if ie.ListOfCellsinAoI != nil {
		if err := ie.ListOfCellsinAoI.Encode(e); err != nil {
			return err
		}
	}
	if ie.ListOfRANNodesinAoI != nil {
		if err := ie.ListOfRANNodesinAoI.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.RequestReferenceID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AreaOfInterestItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(areaOfInterestItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ListOfTAIsinAoI = new(ListOfTAIsinAoI)
		if err := ie.ListOfTAIsinAoI.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ListOfCellsinAoI = new(ListOfCells)
		if err := ie.ListOfCellsinAoI.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.ListOfRANNodesinAoI = new(ListOfRANNodesinAoI)
		if err := ie.ListOfRANNodesinAoI.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.RequestReferenceID.Decode(d); err != nil {
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
