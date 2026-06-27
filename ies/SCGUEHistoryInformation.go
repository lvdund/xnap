package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sCGUEHistoryInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lastVisitedPSCellList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SCGUEHistoryInformation struct {
	LastVisitedPSCellList *LastVisitedPSCellList
	IEExtensions          []byte
}

func (ie *SCGUEHistoryInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGUEHistoryInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.LastVisitedPSCellList != nil, false}); err != nil {
		return err
	}
	if ie.LastVisitedPSCellList != nil {
		if err := ie.LastVisitedPSCellList.Encode(e); err != nil {
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

func (ie *SCGUEHistoryInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGUEHistoryInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.LastVisitedPSCellList = new(LastVisitedPSCellList)
		if err := ie.LastVisitedPSCellList.Decode(d); err != nil {
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
