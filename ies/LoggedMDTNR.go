package ies

import (
	"github.com/lvdund/asn1go/per"
)

var loggedMDTNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "loggingInterval"},
		{Name: "loggingDuration"},
		{Name: "reportType"},
		{Name: "bluetoothMeasurementConfiguration", Optional: true},
		{Name: "wLANMeasurementConfiguration", Optional: true},
		{Name: "sensorMeasurementConfiguration", Optional: true},
		{Name: "areaScopeOfNeighCellsList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type LoggedMDTNR struct {
	LoggingInterval                   LoggingInterval
	LoggingDuration                   LoggingDuration
	ReportType                        ReportType
	BluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration
	WLANMeasurementConfiguration      *WLANMeasurementConfiguration
	SensorMeasurementConfiguration    *SensorMeasurementConfiguration
	AreaScopeOfNeighCellsList         *AreaScopeOfNeighCellsList
	IEExtensions                      []byte
}

func (ie *LoggedMDTNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedMDTNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.BluetoothMeasurementConfiguration != nil, ie.WLANMeasurementConfiguration != nil, ie.SensorMeasurementConfiguration != nil, ie.AreaScopeOfNeighCellsList != nil, false}); err != nil {
		return err
	}
	if err := ie.LoggingInterval.Encode(e); err != nil {
		return err
	}
	if err := ie.LoggingDuration.Encode(e); err != nil {
		return err
	}
	if err := ie.ReportType.Encode(e); err != nil {
		return err
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
	if ie.AreaScopeOfNeighCellsList != nil {
		if err := ie.AreaScopeOfNeighCellsList.Encode(e); err != nil {
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

func (ie *LoggedMDTNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedMDTNRConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.LoggingInterval.Decode(d); err != nil {
		return err
	}
	if err := ie.LoggingDuration.Decode(d); err != nil {
		return err
	}
	if err := ie.ReportType.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.BluetoothMeasurementConfiguration = new(BluetoothMeasurementConfiguration)
		if err := ie.BluetoothMeasurementConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.WLANMeasurementConfiguration = new(WLANMeasurementConfiguration)
		if err := ie.WLANMeasurementConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.SensorMeasurementConfiguration = new(SensorMeasurementConfiguration)
		if err := ie.SensorMeasurementConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.AreaScopeOfNeighCellsList = new(AreaScopeOfNeighCellsList)
		if err := ie.AreaScopeOfNeighCellsList.Decode(d); err != nil {
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
