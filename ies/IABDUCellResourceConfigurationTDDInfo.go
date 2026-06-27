package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABDUCellResourceConfigurationTDDInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gNB-DU-Cell-Resource-Configuration-TDD"},
		{Name: "frequencyInfo", Optional: true},
		{Name: "transmissionBandwidth", Optional: true},
		{Name: "carrierList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABDUCellResourceConfigurationTDDInfo struct {
	GNBDUCellResourceConfigurationTDD GNBDUCellResourceConfiguration
	FrequencyInfo                     *NRFrequencyInfo
	TransmissionBandwidth             *NRTransmissionBandwidth
	CarrierList                       *NRCarrierList
	IEExtensions                      []byte
}

func (ie *IABDUCellResourceConfigurationTDDInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABDUCellResourceConfigurationTDDInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.FrequencyInfo != nil, ie.TransmissionBandwidth != nil, ie.CarrierList != nil, false}); err != nil {
		return err
	}
	if err := ie.GNBDUCellResourceConfigurationTDD.Encode(e); err != nil {
		return err
	}
	if ie.FrequencyInfo != nil {
		if err := ie.FrequencyInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.TransmissionBandwidth != nil {
		if err := ie.TransmissionBandwidth.Encode(e); err != nil {
			return err
		}
	}
	if ie.CarrierList != nil {
		if err := ie.CarrierList.Encode(e); err != nil {
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

func (ie *IABDUCellResourceConfigurationTDDInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABDUCellResourceConfigurationTDDInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GNBDUCellResourceConfigurationTDD.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.FrequencyInfo = new(NRFrequencyInfo)
		if err := ie.FrequencyInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.TransmissionBandwidth = new(NRTransmissionBandwidth)
		if err := ie.TransmissionBandwidth.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.CarrierList = new(NRCarrierList)
		if err := ie.CarrierList.Decode(d); err != nil {
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
