package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	WLANMeasurementConfigurationWlanRssiTrue int64 = 0
)

var wLANMeasurementConfigurationWlanRssiConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type WLANMeasurementConfigurationWlanRssi struct {
	Value int64
}

func (ie *WLANMeasurementConfigurationWlanRssi) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, wLANMeasurementConfigurationWlanRssiConstraints)
}

func (ie *WLANMeasurementConfigurationWlanRssi) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(wLANMeasurementConfigurationWlanRssiConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	WLANMeasurementConfigurationWlanRttTrue int64 = 0
)

var wLANMeasurementConfigurationWlanRttConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type WLANMeasurementConfigurationWlanRtt struct {
	Value int64
}

func (ie *WLANMeasurementConfigurationWlanRtt) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, wLANMeasurementConfigurationWlanRttConstraints)
}

func (ie *WLANMeasurementConfigurationWlanRtt) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(wLANMeasurementConfigurationWlanRttConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var wLANMeasurementConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "wlanMeasConfig"},
		{Name: "wlanMeasConfigNameList", Optional: true},
		{Name: "wlan-rssi", Optional: true},
		{Name: "wlan-rtt", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type WLANMeasurementConfiguration struct {
	WlanMeasConfig         WLANMeasConfig
	WlanMeasConfigNameList *WLANMeasConfigNameList
	WlanRssi               *WLANMeasurementConfigurationWlanRssi
	WlanRtt                *WLANMeasurementConfigurationWlanRtt
	IEExtensions           []byte
}

func (ie *WLANMeasurementConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(wLANMeasurementConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.WlanMeasConfigNameList != nil, ie.WlanRssi != nil, ie.WlanRtt != nil, false}); err != nil {
		return err
	}
	if err := ie.WlanMeasConfig.Encode(e); err != nil {
		return err
	}
	if ie.WlanMeasConfigNameList != nil {
		if err := ie.WlanMeasConfigNameList.Encode(e); err != nil {
			return err
		}
	}
	if ie.WlanRssi != nil {
		if err := ie.WlanRssi.Encode(e); err != nil {
			return err
		}
	}
	if ie.WlanRtt != nil {
		if err := ie.WlanRtt.Encode(e); err != nil {
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

func (ie *WLANMeasurementConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(wLANMeasurementConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.WlanMeasConfig.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.WlanMeasConfigNameList = new(WLANMeasConfigNameList)
		if err := ie.WlanMeasConfigNameList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.WlanRssi = new(WLANMeasurementConfigurationWlanRssi)
		if err := ie.WlanRssi.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.WlanRtt = new(WLANMeasurementConfigurationWlanRtt)
		if err := ie.WlanRtt.Decode(d); err != nil {
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
