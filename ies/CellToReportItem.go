package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellToReportItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cell-ID"},
		{Name: "sSBToReport-List", Optional: true},
		{Name: "sliceToReport-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellToReportItem struct {
	CellID            GlobalNGRANCellID
	SSBToReportList   *SSBToReportList
	SliceToReportList *SliceToReportList
	IEExtensions      []byte
}

func (ie *CellToReportItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellToReportItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SSBToReportList != nil, ie.SliceToReportList != nil, false}); err != nil {
		return err
	}
	if err := ie.CellID.Encode(e); err != nil {
		return err
	}
	if ie.SSBToReportList != nil {
		if err := ie.SSBToReportList.Encode(e); err != nil {
			return err
		}
	}
	if ie.SliceToReportList != nil {
		if err := ie.SliceToReportList.Encode(e); err != nil {
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

func (ie *CellToReportItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellToReportItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CellID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SSBToReportList = new(SSBToReportList)
		if err := ie.SSBToReportList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.SliceToReportList = new(SliceToReportList)
		if err := ie.SliceToReportList.Decode(d); err != nil {
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
