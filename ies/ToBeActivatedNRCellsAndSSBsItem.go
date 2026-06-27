package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var toBeActivatedNRCellsAndSSBsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrCGI"},
		{Name: "sSBstobeActivatedList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ToBeActivatedNRCellsAndSSBsItem struct {
	NrCGI                 NRCGI
	SSBstobeActivatedList []*SSBsToBeActivatedItem
	IEExtensions          []byte
}

func (ie *ToBeActivatedNRCellsAndSSBsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(toBeActivatedNRCellsAndSSBsItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.SSBstobeActivatedList) > 0, false}); err != nil {
		return err
	}
	if err := ie.NrCGI.Encode(e); err != nil {
		return err
	}
	if len(ie.SSBstobeActivatedList) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.SSBstobeActivatedList))); err != nil {
			return err
		}
		for _, item := range ie.SSBstobeActivatedList {
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

func (ie *ToBeActivatedNRCellsAndSSBsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(toBeActivatedNRCellsAndSSBsItemConstraints)
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
		ie.SSBstobeActivatedList = make([]*SSBsToBeActivatedItem, n)
		for i := range ie.SSBstobeActivatedList {
			ie.SSBstobeActivatedList[i] = new(SSBsToBeActivatedItem)
			if err := ie.SSBstobeActivatedList[i].Decode(d); err != nil {
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
