package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AvailableRVQoEMetricsApplicationLayerBufferLevelListTrue int64 = 0
)

var availableRVQoEMetricsApplicationLayerBufferLevelListConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type AvailableRVQoEMetricsApplicationLayerBufferLevelList struct {
	Value int64
}

func (ie *AvailableRVQoEMetricsApplicationLayerBufferLevelList) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, availableRVQoEMetricsApplicationLayerBufferLevelListConstraints)
}

func (ie *AvailableRVQoEMetricsApplicationLayerBufferLevelList) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(availableRVQoEMetricsApplicationLayerBufferLevelListConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	AvailableRVQoEMetricsPlayoutDelayForMediaStartupTrue int64 = 0
)

var availableRVQoEMetricsPlayoutDelayForMediaStartupConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type AvailableRVQoEMetricsPlayoutDelayForMediaStartup struct {
	Value int64
}

func (ie *AvailableRVQoEMetricsPlayoutDelayForMediaStartup) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, availableRVQoEMetricsPlayoutDelayForMediaStartupConstraints)
}

func (ie *AvailableRVQoEMetricsPlayoutDelayForMediaStartup) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(availableRVQoEMetricsPlayoutDelayForMediaStartupConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var availableRVQoEMetricsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "applicationLayerBufferLevelList", Optional: true},
		{Name: "playoutDelayForMediaStartup", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AvailableRVQoEMetrics struct {
	ApplicationLayerBufferLevelList *AvailableRVQoEMetricsApplicationLayerBufferLevelList
	PlayoutDelayForMediaStartup     *AvailableRVQoEMetricsPlayoutDelayForMediaStartup
	IEExtensions                    []byte
}

func (ie *AvailableRVQoEMetrics) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availableRVQoEMetricsConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ApplicationLayerBufferLevelList != nil, ie.PlayoutDelayForMediaStartup != nil, false}); err != nil {
		return err
	}
	if ie.ApplicationLayerBufferLevelList != nil {
		if err := ie.ApplicationLayerBufferLevelList.Encode(e); err != nil {
			return err
		}
	}
	if ie.PlayoutDelayForMediaStartup != nil {
		if err := ie.PlayoutDelayForMediaStartup.Encode(e); err != nil {
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

func (ie *AvailableRVQoEMetrics) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availableRVQoEMetricsConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ApplicationLayerBufferLevelList = new(AvailableRVQoEMetricsApplicationLayerBufferLevelList)
		if err := ie.ApplicationLayerBufferLevelList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.PlayoutDelayForMediaStartup = new(AvailableRVQoEMetricsPlayoutDelayForMediaStartup)
		if err := ie.PlayoutDelayForMediaStartup.Decode(d); err != nil {
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
