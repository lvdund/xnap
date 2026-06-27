package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	CellAssistanceInfoNRFullListAllServedCellsNR int64 = 0
)

var cellAssistanceInfoNRFullListConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CellAssistanceInfoNRFullList struct {
	Value int64
}

func (ie *CellAssistanceInfoNRFullList) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cellAssistanceInfoNRFullListConstraints)
}

func (ie *CellAssistanceInfoNRFullList) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cellAssistanceInfoNRFullListConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	CellAssistanceInfoNRChLimitedNRList   = 0
	CellAssistanceInfoNRChFullList        = 1
	CellAssistanceInfoNRChChoiceExtension = 2
)

var cellAssistanceInfoNRConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "limitedNR-List"},
		{Name: "full-List"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type CellAssistanceInfoNR struct {
	Choice          int
	LimitedNRList   []*NRCGI
	FullList        *CellAssistanceInfoNRFullList
	ChoiceExtension []byte
}

func (ie *CellAssistanceInfoNR) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(cellAssistanceInfoNRConstraints)
	switch ie.Choice {
	case 0: // limitedNR-List
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.LimitedNRList))); err != nil {
			return err
		}
		for _, item := range ie.LimitedNRList {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	case 1: // full-List
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.FullList.Encode(e); err != nil {
			return err
		}
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *CellAssistanceInfoNR) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(cellAssistanceInfoNRConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // limitedNR-List
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.LimitedNRList = make([]*NRCGI, n)
		for i := range ie.LimitedNRList {
			ie.LimitedNRList[i] = new(NRCGI)
			if err := ie.LimitedNRList[i].Decode(d); err != nil {
				return err
			}
		}
	case 1: // full-List
		ie.FullList = new(CellAssistanceInfoNRFullList)
		if err := ie.FullList.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
