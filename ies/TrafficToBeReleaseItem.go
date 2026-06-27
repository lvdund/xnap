package ies

import (
	"github.com/lvdund/asn1go/per"
)

var trafficToBeReleaseItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "trafficIndex"},
		{Name: "bHInfoList", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TrafficToBeReleaseItem struct {
	TrafficIndex TrafficIndex
	BHInfoList   *BHInfoList
	IEExtensions []byte
}

func (ie *TrafficToBeReleaseItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(trafficToBeReleaseItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.BHInfoList != nil, false}); err != nil {
		return err
	}
	if err := ie.TrafficIndex.Encode(e); err != nil {
		return err
	}
	if ie.BHInfoList != nil {
		if err := ie.BHInfoList.Encode(e); err != nil {
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

func (ie *TrafficToBeReleaseItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(trafficToBeReleaseItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TrafficIndex.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.BHInfoList = new(BHInfoList)
		if err := ie.BHInfoList.Decode(d); err != nil {
			return err
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
