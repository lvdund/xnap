package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABCellInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nRCGI"},
		{Name: "iAB-DU-Cell-Resource-Configuration-Mode-Info", Optional: true},
		{Name: "iAB-STC-Info", Optional: true},
		{Name: "rACH-Config-Common", Optional: true},
		{Name: "rACH-Config-Common-IAB", Optional: true},
		{Name: "cSI-RS-Configuration", Optional: true},
		{Name: "sR-Configuration", Optional: true},
		{Name: "pDCCH-ConfigSIB1", Optional: true},
		{Name: "sCS-Common", Optional: true},
		{Name: "multiplexingInfo", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABCellInformation struct {
	NRCGI                                  NRCGI
	IABDUCellResourceConfigurationModeInfo *IABDUCellResourceConfigurationModeInfo
	IABSTCInfo                             *IABSTCInfo
	RACHConfigCommon                       *RACHConfigCommon
	RACHConfigCommonIAB                    *RACHConfigCommonIAB
	CSIRSConfiguration                     []byte
	SRConfiguration                        []byte
	PDCCHConfigSIB1                        []byte
	SCSCommon                              []byte
	MultiplexingInfo                       *MultiplexingInfo
	IEExtensions                           []byte
}

func (ie *IABCellInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABCellInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.IABDUCellResourceConfigurationModeInfo != nil, ie.IABSTCInfo != nil, ie.RACHConfigCommon != nil, ie.RACHConfigCommonIAB != nil, len(ie.CSIRSConfiguration) > 0, len(ie.SRConfiguration) > 0, len(ie.PDCCHConfigSIB1) > 0, len(ie.SCSCommon) > 0, ie.MultiplexingInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.NRCGI.Encode(e); err != nil {
		return err
	}
	if ie.IABDUCellResourceConfigurationModeInfo != nil {
		if err := ie.IABDUCellResourceConfigurationModeInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.IABSTCInfo != nil {
		if err := ie.IABSTCInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.RACHConfigCommon != nil {
		if err := ie.RACHConfigCommon.Encode(e); err != nil {
			return err
		}
	}
	if ie.RACHConfigCommonIAB != nil {
		if err := ie.RACHConfigCommonIAB.Encode(e); err != nil {
			return err
		}
	}
	if len(ie.CSIRSConfiguration) > 0 {
		if err := e.EncodeOctetString(ie.CSIRSConfiguration, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.SRConfiguration) > 0 {
		if err := e.EncodeOctetString(ie.SRConfiguration, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.PDCCHConfigSIB1) > 0 {
		if err := e.EncodeOctetString(ie.PDCCHConfigSIB1, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.SCSCommon) > 0 {
		if err := e.EncodeOctetString(ie.SCSCommon, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if ie.MultiplexingInfo != nil {
		if err := ie.MultiplexingInfo.Encode(e); err != nil {
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

func (ie *IABCellInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABCellInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRCGI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.IABDUCellResourceConfigurationModeInfo = new(IABDUCellResourceConfigurationModeInfo)
		if err := ie.IABDUCellResourceConfigurationModeInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.IABSTCInfo = new(IABSTCInfo)
		if err := ie.IABSTCInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.RACHConfigCommon = new(RACHConfigCommon)
		if err := ie.RACHConfigCommon.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.RACHConfigCommonIAB = new(RACHConfigCommonIAB)
		if err := ie.RACHConfigCommonIAB.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.CSIRSConfiguration = val
	}
	if seq.IsComponentPresent(6) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.SRConfiguration = val
	}
	if seq.IsComponentPresent(7) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.PDCCHConfigSIB1 = val
	}
	if seq.IsComponentPresent(8) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.SCSCommon = val
	}
	if seq.IsComponentPresent(9) {
		ie.MultiplexingInfo = new(MultiplexingInfo)
		if err := ie.MultiplexingInfo.Decode(d); err != nil {
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
