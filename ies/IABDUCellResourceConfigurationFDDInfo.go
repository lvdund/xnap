package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABDUCellResourceConfigurationFDDInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gNB-DU-Cell-Resource-Configuration-FDD-UL"},
		{Name: "gNB-DU-Cell-Resource-Configuration-FDD-DL"},
		{Name: "uLFrequencyInfo", Optional: true},
		{Name: "dLFrequencyInfo", Optional: true},
		{Name: "uLTransmissionBandwidth", Optional: true},
		{Name: "dlTransmissionBandwidth", Optional: true},
		{Name: "uLCarrierList", Optional: true},
		{Name: "dlCarrierList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABDUCellResourceConfigurationFDDInfo struct {
	GNBDUCellResourceConfigurationFDDUL GNBDUCellResourceConfiguration
	GNBDUCellResourceConfigurationFDDDL GNBDUCellResourceConfiguration
	ULFrequencyInfo                     *NRFrequencyInfo
	DLFrequencyInfo                     *NRFrequencyInfo
	ULTransmissionBandwidth             *NRTransmissionBandwidth
	DlTransmissionBandwidth             *NRTransmissionBandwidth
	ULCarrierList                       *NRCarrierList
	DlCarrierList                       *NRCarrierList
	IEExtensions                        []byte
}

func (ie *IABDUCellResourceConfigurationFDDInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABDUCellResourceConfigurationFDDInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ULFrequencyInfo != nil, ie.DLFrequencyInfo != nil, ie.ULTransmissionBandwidth != nil, ie.DlTransmissionBandwidth != nil, ie.ULCarrierList != nil, ie.DlCarrierList != nil, false}); err != nil {
		return err
	}
	if err := ie.GNBDUCellResourceConfigurationFDDUL.Encode(e); err != nil {
		return err
	}
	if err := ie.GNBDUCellResourceConfigurationFDDDL.Encode(e); err != nil {
		return err
	}
	if ie.ULFrequencyInfo != nil {
		if err := ie.ULFrequencyInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.DLFrequencyInfo != nil {
		if err := ie.DLFrequencyInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.ULTransmissionBandwidth != nil {
		if err := ie.ULTransmissionBandwidth.Encode(e); err != nil {
			return err
		}
	}
	if ie.DlTransmissionBandwidth != nil {
		if err := ie.DlTransmissionBandwidth.Encode(e); err != nil {
			return err
		}
	}
	if ie.ULCarrierList != nil {
		if err := ie.ULCarrierList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DlCarrierList != nil {
		if err := ie.DlCarrierList.Encode(e); err != nil {
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

func (ie *IABDUCellResourceConfigurationFDDInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABDUCellResourceConfigurationFDDInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GNBDUCellResourceConfigurationFDDUL.Decode(d); err != nil {
		return err
	}
	if err := ie.GNBDUCellResourceConfigurationFDDDL.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.ULFrequencyInfo = new(NRFrequencyInfo)
		if err := ie.ULFrequencyInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DLFrequencyInfo = new(NRFrequencyInfo)
		if err := ie.DLFrequencyInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.ULTransmissionBandwidth = new(NRTransmissionBandwidth)
		if err := ie.ULTransmissionBandwidth.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.DlTransmissionBandwidth = new(NRTransmissionBandwidth)
		if err := ie.DlTransmissionBandwidth.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.ULCarrierList = new(NRCarrierList)
		if err := ie.ULCarrierList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.DlCarrierList = new(NRCarrierList)
		if err := ie.DlCarrierList.Decode(d); err != nil {
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
