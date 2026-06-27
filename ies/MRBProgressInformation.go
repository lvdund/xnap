package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	MRBProgressInformationChPdcpSN12        = 0
	MRBProgressInformationChPdcpSN18        = 1
	MRBProgressInformationChChoiceExtension = 2
)

var mRBProgressInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pdcp-SN12"},
		{Name: "pdcp-SN18"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type MRBProgressInformation struct {
	Choice          int
	PdcpSN12        *int64
	PdcpSN18        *int64
	ChoiceExtension []byte
}

func (ie *MRBProgressInformation) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(mRBProgressInformationConstraints)
	switch ie.Choice {
	case 0: // pdcp-SN12
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeInteger(*ie.PdcpSN12, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(4095)),
		}); err != nil {
			return err
		}
	case 1: // pdcp-SN18
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeInteger(*ie.PdcpSN18, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(262143)),
		}); err != nil {
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

func (ie *MRBProgressInformation) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(mRBProgressInformationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // pdcp-SN12
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(4095)),
		})
		if err != nil {
			return err
		}
		ie.PdcpSN12 = &val
	case 1: // pdcp-SN18
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(262143)),
		})
		if err != nil {
			return err
		}
		ie.PdcpSN18 = &val
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
