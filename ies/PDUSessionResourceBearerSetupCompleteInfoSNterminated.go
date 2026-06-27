package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionResourceBearerSetupCompleteInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dRBsToBeSetupList"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceBearerSetupCompleteInfoSNterminated struct {
	DRBsToBeSetupList []*DRBsToBeSetupListBearerSetupCompleteSNterminatedItem
	IEExtensions      []byte
}

func (ie *PDUSessionResourceBearerSetupCompleteInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceBearerSetupCompleteInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(1)),
		Max:        common.Ptr(int64(common.MaxnoofDRBs)),
	})
	if err := soEnc.EncodeLength(int64(len(ie.DRBsToBeSetupList))); err != nil {
		return err
	}
	for _, item := range ie.DRBsToBeSetupList {
		if err := item.Encode(e); err != nil {
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

func (ie *PDUSessionResourceBearerSetupCompleteInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceBearerSetupCompleteInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofDRBs)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.DRBsToBeSetupList = make([]*DRBsToBeSetupListBearerSetupCompleteSNterminatedItem, n)
		for i := range ie.DRBsToBeSetupList {
			ie.DRBsToBeSetupList[i] = new(DRBsToBeSetupListBearerSetupCompleteSNterminatedItem)
			if err := ie.DRBsToBeSetupList[i].Decode(d); err != nil {
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
