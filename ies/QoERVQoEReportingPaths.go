package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoERVQoEReportingPathsQoEReportingPathSrb4 int64 = 0
	QoERVQoEReportingPathsQoEReportingPathSrb5 int64 = 1
)

var qoERVQoEReportingPathsQoEReportingPathConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type QoERVQoEReportingPathsQoEReportingPath struct {
	Value int64
}

func (ie *QoERVQoEReportingPathsQoEReportingPath) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoERVQoEReportingPathsQoEReportingPathConstraints)
}

func (ie *QoERVQoEReportingPathsQoEReportingPath) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoERVQoEReportingPathsQoEReportingPathConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	QoERVQoEReportingPathsRVQoEReportingPathSrb4 int64 = 0
	QoERVQoEReportingPathsRVQoEReportingPathSrb5 int64 = 1
)

var qoERVQoEReportingPathsRVQoEReportingPathConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type QoERVQoEReportingPathsRVQoEReportingPath struct {
	Value int64
}

func (ie *QoERVQoEReportingPathsRVQoEReportingPath) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoERVQoEReportingPathsRVQoEReportingPathConstraints)
}

func (ie *QoERVQoEReportingPathsRVQoEReportingPath) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoERVQoEReportingPathsRVQoEReportingPathConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var qoERVQoEReportingPathsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qoEReportingPath", Optional: true},
		{Name: "rVQoEReportingPath", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoERVQoEReportingPaths struct {
	QoEReportingPath   *QoERVQoEReportingPathsQoEReportingPath
	RVQoEReportingPath *QoERVQoEReportingPathsRVQoEReportingPath
	IEExtensions       []byte
}

func (ie *QoERVQoEReportingPaths) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoERVQoEReportingPathsConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QoEReportingPath != nil, ie.RVQoEReportingPath != nil, false}); err != nil {
		return err
	}
	if ie.QoEReportingPath != nil {
		if err := ie.QoEReportingPath.Encode(e); err != nil {
			return err
		}
	}
	if ie.RVQoEReportingPath != nil {
		if err := ie.RVQoEReportingPath.Encode(e); err != nil {
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

func (ie *QoERVQoEReportingPaths) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoERVQoEReportingPathsConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.QoEReportingPath = new(QoERVQoEReportingPathsQoEReportingPath)
		if err := ie.QoEReportingPath.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.RVQoEReportingPath = new(QoERVQoEReportingPathsRVQoEReportingPath)
		if err := ie.RVQoEReportingPath.Decode(d); err != nil {
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
