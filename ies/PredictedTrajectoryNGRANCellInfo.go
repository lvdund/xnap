package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var predictedTrajectoryNGRANCellInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "globalNG-RANCell-ID"},
		{Name: "predictedTimeUEStaysInCell", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PredictedTrajectoryNGRANCellInfo struct {
	GlobalNGRANCellID          GlobalNGRANCellID
	PredictedTimeUEStaysInCell *int64
	IEExtensions               []byte
}

func (ie *PredictedTrajectoryNGRANCellInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(predictedTrajectoryNGRANCellInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PredictedTimeUEStaysInCell != nil, false}); err != nil {
		return err
	}
	if err := ie.GlobalNGRANCellID.Encode(e); err != nil {
		return err
	}
	if ie.PredictedTimeUEStaysInCell != nil {
		if err := e.EncodeInteger(*ie.PredictedTimeUEStaysInCell, per.IntegerConstraints{
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

func (ie *PredictedTrajectoryNGRANCellInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(predictedTrajectoryNGRANCellInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GlobalNGRANCellID.Decode(d); err != nil {
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
		ie.PredictedTimeUEStaysInCell = &val
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
