package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRUChannelItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nR-U-ChannelID"},
		{Name: "channelOccupancyTimePercentageDL"},
		{Name: "energyDetectionThresholdDL"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRUChannelItem struct {
	NRUChannelID                     NRUChannelID
	ChannelOccupancyTimePercentageDL ChannelOccupancyTimePercentage
	EnergyDetectionThresholdDL       EnergyDetectionThreshold
	IEExtensions                     []byte
}

func (ie *NRUChannelItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRUChannelItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NRUChannelID.Encode(e); err != nil {
		return err
	}
	if err := ie.ChannelOccupancyTimePercentageDL.Encode(e); err != nil {
		return err
	}
	if err := ie.EnergyDetectionThresholdDL.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRUChannelItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRUChannelItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRUChannelID.Decode(d); err != nil {
		return err
	}
	if err := ie.ChannelOccupancyTimePercentageDL.Decode(d); err != nil {
		return err
	}
	if err := ie.EnergyDetectionThresholdDL.Decode(d); err != nil {
		return err
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
