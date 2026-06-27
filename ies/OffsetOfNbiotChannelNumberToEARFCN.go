package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	OffsetOfNbiotChannelNumberToEARFCNMinusTen          int64 = 0
	OffsetOfNbiotChannelNumberToEARFCNMinusNine         int64 = 1
	OffsetOfNbiotChannelNumberToEARFCNMinusEightDotFive int64 = 2
	OffsetOfNbiotChannelNumberToEARFCNMinusEight        int64 = 3
	OffsetOfNbiotChannelNumberToEARFCNMinusSeven        int64 = 4
	OffsetOfNbiotChannelNumberToEARFCNMinusSix          int64 = 5
	OffsetOfNbiotChannelNumberToEARFCNMinusFive         int64 = 6
	OffsetOfNbiotChannelNumberToEARFCNMinusFourDotFive  int64 = 7
	OffsetOfNbiotChannelNumberToEARFCNMinusFour         int64 = 8
	OffsetOfNbiotChannelNumberToEARFCNMinusThree        int64 = 9
	OffsetOfNbiotChannelNumberToEARFCNMinusTwo          int64 = 10
	OffsetOfNbiotChannelNumberToEARFCNMinusOne          int64 = 11
	OffsetOfNbiotChannelNumberToEARFCNMinusZeroDotFive  int64 = 12
	OffsetOfNbiotChannelNumberToEARFCNZero              int64 = 13
	OffsetOfNbiotChannelNumberToEARFCNOne               int64 = 14
	OffsetOfNbiotChannelNumberToEARFCNTwo               int64 = 15
	OffsetOfNbiotChannelNumberToEARFCNThree             int64 = 16
	OffsetOfNbiotChannelNumberToEARFCNThreeDotFive      int64 = 17
	OffsetOfNbiotChannelNumberToEARFCNFour              int64 = 18
	OffsetOfNbiotChannelNumberToEARFCNFive              int64 = 19
	OffsetOfNbiotChannelNumberToEARFCNSix               int64 = 20
	OffsetOfNbiotChannelNumberToEARFCNSeven             int64 = 21
	OffsetOfNbiotChannelNumberToEARFCNSevenDotFive      int64 = 22
	OffsetOfNbiotChannelNumberToEARFCNEight             int64 = 23
	OffsetOfNbiotChannelNumberToEARFCNNine              int64 = 24
)

var offsetOfNbiotChannelNumberToEARFCNConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
	ExtValues:  nil,
}

type OffsetOfNbiotChannelNumberToEARFCN struct {
	Value int64
}

func (ie *OffsetOfNbiotChannelNumberToEARFCN) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, offsetOfNbiotChannelNumberToEARFCNConstraints)
}

func (ie *OffsetOfNbiotChannelNumberToEARFCN) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(offsetOfNbiotChannelNumberToEARFCNConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
