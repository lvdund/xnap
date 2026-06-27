package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	CellAssistanceInfoEUTRAFullListAllServedCellsEUTRA int64 = 0
)

var cellAssistanceInfoEUTRAFullListConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CellAssistanceInfoEUTRAFullList struct {
	Value int64
}

func (ie *CellAssistanceInfoEUTRAFullList) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cellAssistanceInfoEUTRAFullListConstraints)
}

func (ie *CellAssistanceInfoEUTRAFullList) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cellAssistanceInfoEUTRAFullListConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	CellAssistanceInfoEUTRAChLimitedEUTRAList = 0
	CellAssistanceInfoEUTRAChFullList         = 1
	CellAssistanceInfoEUTRAChChoiceExtension  = 2
)

var cellAssistanceInfoEUTRAConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "limitedEUTRA-List"},
		{Name: "full-List"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type CellAssistanceInfoEUTRA struct {
	Choice           int
	LimitedEUTRAList []*EUTRACGI
	FullList         *CellAssistanceInfoEUTRAFullList
	ChoiceExtension  []byte
}

func (ie *CellAssistanceInfoEUTRA) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(cellAssistanceInfoEUTRAConstraints)
	switch ie.Choice {
	case 0: // limitedEUTRA-List
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.LimitedEUTRAList))); err != nil {
			return err
		}
		for _, item := range ie.LimitedEUTRAList {
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

func (ie *CellAssistanceInfoEUTRA) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(cellAssistanceInfoEUTRAConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // limitedEUTRA-List
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.LimitedEUTRAList = make([]*EUTRACGI, n)
		for i := range ie.LimitedEUTRAList {
			ie.LimitedEUTRAList[i] = new(EUTRACGI)
			if err := ie.LimitedEUTRAList[i].Decode(d); err != nil {
				return err
			}
		}
	case 1: // full-List
		ie.FullList = new(CellAssistanceInfoEUTRAFullList)
		if err := ie.FullList.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
