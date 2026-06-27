package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	QoEMeasurementResultsAppLayerSessionStatusStarted int64 = 0
	QoEMeasurementResultsAppLayerSessionStatusStopped int64 = 1
)

var qoEMeasurementResultsAppLayerSessionStatusConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type QoEMeasurementResultsAppLayerSessionStatus struct {
	Value int64
}

func (ie *QoEMeasurementResultsAppLayerSessionStatus) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoEMeasurementResultsAppLayerSessionStatusConstraints)
}

func (ie *QoEMeasurementResultsAppLayerSessionStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoEMeasurementResultsAppLayerSessionStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var qoEMeasurementResultsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qOEReference"},
		{Name: "rrcContainerForRVQoEReport", Optional: true},
		{Name: "rrcContainerForQoEReport", Optional: true},
		{Name: "appLayerSessionStatus", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoEMeasurementResults struct {
	QOEReference               []byte
	RrcContainerForRVQoEReport []byte
	RrcContainerForQoEReport   []byte
	AppLayerSessionStatus      *QoEMeasurementResultsAppLayerSessionStatus
	IEExtensions               []byte
}

func (ie *QoEMeasurementResults) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoEMeasurementResultsConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.RrcContainerForRVQoEReport) > 0, len(ie.RrcContainerForQoEReport) > 0, ie.AppLayerSessionStatus != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.QOEReference, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(6)),
		Max:        common.Ptr(int64(6)),
	}); err != nil {
		return err
	}
	if len(ie.RrcContainerForRVQoEReport) > 0 {
		if err := e.EncodeOctetString(ie.RrcContainerForRVQoEReport, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.RrcContainerForQoEReport) > 0 {
		if err := e.EncodeOctetString(ie.RrcContainerForQoEReport, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if ie.AppLayerSessionStatus != nil {
		if err := ie.AppLayerSessionStatus.Encode(e); err != nil {
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

func (ie *QoEMeasurementResults) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoEMeasurementResultsConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(6)),
		})
		if err != nil {
			return err
		}
		ie.QOEReference = val
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.RrcContainerForRVQoEReport = val
	}
	if seq.IsComponentPresent(2) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.RrcContainerForQoEReport = val
	}
	if seq.IsComponentPresent(3) {
		ie.AppLayerSessionStatus = new(QoEMeasurementResultsAppLayerSessionStatus)
		if err := ie.AppLayerSessionStatus.Decode(d); err != nil {
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
