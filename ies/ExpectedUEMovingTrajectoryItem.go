package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var expectedUEMovingTrajectoryItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nGRAN-CGI"},
		{Name: "timeStayedInCell", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         GlobalNGRANCellID
	TimeStayedInCell *int64
	IEExtensions     []byte
}

func (ie *ExpectedUEMovingTrajectoryItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(expectedUEMovingTrajectoryItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.TimeStayedInCell != nil, false}); err != nil {
		return err
	}
	if err := ie.NGRANCGI.Encode(e); err != nil {
		return err
	}
	if ie.TimeStayedInCell != nil {
		if err := e.EncodeInteger(*ie.TimeStayedInCell, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(4095)),
		}); err != nil {
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

func (ie *ExpectedUEMovingTrajectoryItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(expectedUEMovingTrajectoryItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NGRANCGI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(4095)),
		})
		if err != nil {
			return err
		}
		ie.TimeStayedInCell = &val
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
