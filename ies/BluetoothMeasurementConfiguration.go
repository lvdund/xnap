package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	BluetoothMeasurementConfigurationBtRssiTrue int64 = 0
)

var bluetoothMeasurementConfigurationBtRssiConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type BluetoothMeasurementConfigurationBtRssi struct {
	Value int64
}

func (ie *BluetoothMeasurementConfigurationBtRssi) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, bluetoothMeasurementConfigurationBtRssiConstraints)
}

func (ie *BluetoothMeasurementConfigurationBtRssi) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(bluetoothMeasurementConfigurationBtRssiConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var bluetoothMeasurementConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bluetoothMeasConfig"},
		{Name: "bluetoothMeasConfigNameList", Optional: true},
		{Name: "bt-rssi", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BluetoothMeasurementConfiguration struct {
	BluetoothMeasConfig         BluetoothMeasConfig
	BluetoothMeasConfigNameList *BluetoothMeasConfigNameList
	BtRssi                      *BluetoothMeasurementConfigurationBtRssi
	IEExtensions                []byte
}

func (ie *BluetoothMeasurementConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bluetoothMeasurementConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.BluetoothMeasConfigNameList != nil, ie.BtRssi != nil, false}); err != nil {
		return err
	}
	if err := ie.BluetoothMeasConfig.Encode(e); err != nil {
		return err
	}
	if ie.BluetoothMeasConfigNameList != nil {
		if err := ie.BluetoothMeasConfigNameList.Encode(e); err != nil {
			return err
		}
	}
	if ie.BtRssi != nil {
		if err := ie.BtRssi.Encode(e); err != nil {
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

func (ie *BluetoothMeasurementConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bluetoothMeasurementConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BluetoothMeasConfig.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.BluetoothMeasConfigNameList = new(BluetoothMeasConfigNameList)
		if err := ie.BluetoothMeasConfigNameList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.BtRssi = new(BluetoothMeasurementConfigurationBtRssi)
		if err := ie.BtRssi.Decode(d); err != nil {
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
