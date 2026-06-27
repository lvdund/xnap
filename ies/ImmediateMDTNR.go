package ies

import (
	"github.com/lvdund/asn1go/per"
)

var immediateMDTNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementsToActivate"},
		{Name: "m1Configuration", Optional: true},
		{Name: "m4Configuration", Optional: true},
		{Name: "m5Configuration", Optional: true},
		{Name: "mDT-Location-Info", Optional: true},
		{Name: "m6Configuration", Optional: true},
		{Name: "m7Configuration", Optional: true},
		{Name: "bluetoothMeasurementConfiguration", Optional: true},
		{Name: "wLANMeasurementConfiguration", Optional: true},
		{Name: "sensorMeasurementConfiguration", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ImmediateMDTNR struct {
	MeasurementsToActivate            MeasurementsToActivate
	M1Configuration                   *M1Configuration
	M4Configuration                   *M4Configuration
	M5Configuration                   *M5Configuration
	MDTLocationInfo                   *MDTLocationInfo
	M6Configuration                   *M6Configuration
	M7Configuration                   *M7Configuration
	BluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration
	WLANMeasurementConfiguration      *WLANMeasurementConfiguration
	SensorMeasurementConfiguration    *SensorMeasurementConfiguration
	IEExtensions                      []byte
}

func (ie *ImmediateMDTNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(immediateMDTNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.M1Configuration != nil, ie.M4Configuration != nil, ie.M5Configuration != nil, ie.MDTLocationInfo != nil, ie.M6Configuration != nil, ie.M7Configuration != nil, ie.BluetoothMeasurementConfiguration != nil, ie.WLANMeasurementConfiguration != nil, ie.SensorMeasurementConfiguration != nil, false}); err != nil {
		return err
	}
	if err := ie.MeasurementsToActivate.Encode(e); err != nil {
		return err
	}
	if ie.M1Configuration != nil {
		if err := ie.M1Configuration.Encode(e); err != nil {
			return err
		}
	}
	if ie.M4Configuration != nil {
		if err := ie.M4Configuration.Encode(e); err != nil {
			return err
		}
	}
	if ie.M5Configuration != nil {
		if err := ie.M5Configuration.Encode(e); err != nil {
			return err
		}
	}
	if ie.MDTLocationInfo != nil {
		if err := ie.MDTLocationInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.M6Configuration != nil {
		if err := ie.M6Configuration.Encode(e); err != nil {
			return err
		}
	}
	if ie.M7Configuration != nil {
		if err := ie.M7Configuration.Encode(e); err != nil {
			return err
		}
	}
	if ie.BluetoothMeasurementConfiguration != nil {
		if err := ie.BluetoothMeasurementConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if ie.WLANMeasurementConfiguration != nil {
		if err := ie.WLANMeasurementConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if ie.SensorMeasurementConfiguration != nil {
		if err := ie.SensorMeasurementConfiguration.Encode(e); err != nil {
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

func (ie *ImmediateMDTNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(immediateMDTNRConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MeasurementsToActivate.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.M1Configuration = new(M1Configuration)
		if err := ie.M1Configuration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.M4Configuration = new(M4Configuration)
		if err := ie.M4Configuration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.M5Configuration = new(M5Configuration)
		if err := ie.M5Configuration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.MDTLocationInfo = new(MDTLocationInfo)
		if err := ie.MDTLocationInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.M6Configuration = new(M6Configuration)
		if err := ie.M6Configuration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.M7Configuration = new(M7Configuration)
		if err := ie.M7Configuration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.BluetoothMeasurementConfiguration = new(BluetoothMeasurementConfiguration)
		if err := ie.BluetoothMeasurementConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(8) {
		ie.WLANMeasurementConfiguration = new(WLANMeasurementConfiguration)
		if err := ie.WLANMeasurementConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(9) {
		ie.SensorMeasurementConfiguration = new(SensorMeasurementConfiguration)
		if err := ie.SensorMeasurementConfiguration.Decode(d); err != nil {
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
