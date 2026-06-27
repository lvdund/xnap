package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var servedCellsToUpdateNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "served-Cells-ToAdd-NR", Optional: true},
		{Name: "served-Cells-ToModify-NR", Optional: true},
		{Name: "served-Cells-ToDelete-NR", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellsToUpdateNR struct {
	ServedCellsToAddNR    *ServedCellsNR
	ServedCellsToModifyNR *ServedCellsToModifyNR
	ServedCellsToDeleteNR []*NRCGI
	IEExtensions          []byte
}

func (ie *ServedCellsToUpdateNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellsToUpdateNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ServedCellsToAddNR != nil, ie.ServedCellsToModifyNR != nil, len(ie.ServedCellsToDeleteNR) > 0, false}); err != nil {
		return err
	}
	if ie.ServedCellsToAddNR != nil {
		if err := ie.ServedCellsToAddNR.Encode(e); err != nil {
			return err
		}
	}
	if ie.ServedCellsToModifyNR != nil {
		if err := ie.ServedCellsToModifyNR.Encode(e); err != nil {
			return err
		}
	}
	if len(ie.ServedCellsToDeleteNR) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.ServedCellsToDeleteNR))); err != nil {
			return err
		}
		for _, item := range ie.ServedCellsToDeleteNR {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ServedCellsToUpdateNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellsToUpdateNRConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ServedCellsToAddNR = new(ServedCellsNR)
		if err := ie.ServedCellsToAddNR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ServedCellsToModifyNR = new(ServedCellsToModifyNR)
		if err := ie.ServedCellsToModifyNR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.ServedCellsToDeleteNR = make([]*NRCGI, n)
		for i := range ie.ServedCellsToDeleteNR {
			ie.ServedCellsToDeleteNR[i] = new(NRCGI)
			if err := ie.ServedCellsToDeleteNR[i].Decode(d); err != nil {
				return err
			}
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
