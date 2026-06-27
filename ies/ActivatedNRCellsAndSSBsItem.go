package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var activatedNRCellsAndSSBsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrCGI"},
		{Name: "sSBsActivatedList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ActivatedNRCellsAndSSBsItem struct {
	NrCGI             NRCGI
	SSBsActivatedList []*SSBsActivatedItem
	IEExtensions      []byte
}

func (ie *ActivatedNRCellsAndSSBsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(activatedNRCellsAndSSBsItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.SSBsActivatedList) > 0, false}); err != nil {
		return err
	}
	if err := ie.NrCGI.Encode(e); err != nil {
		return err
	}
	if len(ie.SSBsActivatedList) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.SSBsActivatedList))); err != nil {
			return err
		}
		for _, item := range ie.SSBsActivatedList {
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

func (ie *ActivatedNRCellsAndSSBsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(activatedNRCellsAndSSBsItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrCGI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.SSBsActivatedList = make([]*SSBsActivatedItem, n)
		for i := range ie.SSBsActivatedList {
			ie.SSBsActivatedList[i] = new(SSBsActivatedItem)
			if err := ie.SSBsActivatedList[i].Decode(d); err != nil {
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
