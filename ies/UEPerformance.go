package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEPerformanceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dL-UE-AverageThroughput", Optional: true},
		{Name: "uL-UE-AverageThroughput", Optional: true},
		{Name: "uE-AveragePacketDelay", Optional: true},
		{Name: "uE-AveragePacketLossDL", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEPerformance struct {
	DLUEAverageThroughput *BitRate
	ULUEAverageThroughput *BitRate
	UEAveragePacketDelay  *AveragePacketDelay
	UEAveragePacketLossDL *PacketLossRate
	IEExtensions          []byte
}

func (ie *UEPerformance) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEPerformanceConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DLUEAverageThroughput != nil, ie.ULUEAverageThroughput != nil, ie.UEAveragePacketDelay != nil, ie.UEAveragePacketLossDL != nil, false}); err != nil {
		return err
	}
	if ie.DLUEAverageThroughput != nil {
		if err := ie.DLUEAverageThroughput.Encode(e); err != nil {
			return err
		}
	}
	if ie.ULUEAverageThroughput != nil {
		if err := ie.ULUEAverageThroughput.Encode(e); err != nil {
			return err
		}
	}
	if ie.UEAveragePacketDelay != nil {
		if err := ie.UEAveragePacketDelay.Encode(e); err != nil {
			return err
		}
	}
	if ie.UEAveragePacketLossDL != nil {
		if err := ie.UEAveragePacketLossDL.Encode(e); err != nil {
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

func (ie *UEPerformance) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEPerformanceConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DLUEAverageThroughput = new(BitRate)
		if err := ie.DLUEAverageThroughput.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ULUEAverageThroughput = new(BitRate)
		if err := ie.ULUEAverageThroughput.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.UEAveragePacketDelay = new(AveragePacketDelay)
		if err := ie.UEAveragePacketDelay.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.UEAveragePacketLossDL = new(PacketLossRate)
		if err := ie.UEAveragePacketLossDL.Decode(d); err != nil {
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
