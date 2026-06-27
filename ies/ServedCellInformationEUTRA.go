package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ServedCellInformationEUTRAFreqBandIndicatorPriorityNotBroadcast int64 = 0
	ServedCellInformationEUTRAFreqBandIndicatorPriorityBroadcast    int64 = 1
)

var servedCellInformationEUTRAFreqBandIndicatorPriorityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type ServedCellInformationEUTRAFreqBandIndicatorPriority struct {
	Value int64
}

func (ie *ServedCellInformationEUTRAFreqBandIndicatorPriority) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, servedCellInformationEUTRAFreqBandIndicatorPriorityConstraints)
}

func (ie *ServedCellInformationEUTRAFreqBandIndicatorPriority) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(servedCellInformationEUTRAFreqBandIndicatorPriorityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	ServedCellInformationEUTRABandwidthReducedSIScheduled int64 = 0
)

var servedCellInformationEUTRABandwidthReducedSIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ServedCellInformationEUTRABandwidthReducedSI struct {
	Value int64
}

func (ie *ServedCellInformationEUTRABandwidthReducedSI) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, servedCellInformationEUTRABandwidthReducedSIConstraints)
}

func (ie *ServedCellInformationEUTRABandwidthReducedSI) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(servedCellInformationEUTRABandwidthReducedSIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var servedCellInformationEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "e-utra-pci"},
		{Name: "e-utra-cgi"},
		{Name: "tac"},
		{Name: "ranac", Optional: true},
		{Name: "broadcastPLMNs"},
		{Name: "e-utra-mode-info"},
		{Name: "numberofAntennaPorts", Optional: true},
		{Name: "prach-configuration", Optional: true},
		{Name: "mBSFNsubframeInfo", Optional: true},
		{Name: "multibandInfo", Optional: true},
		{Name: "freqBandIndicatorPriority", Optional: true},
		{Name: "bandwidthReducedSI", Optional: true},
		{Name: "protectedE-UTRAResourceIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellInformationEUTRA struct {
	EUtraPci                         EUTRAPCI
	EUtraCgi                         EUTRACGI
	Tac                              TAC
	Ranac                            *RANAC
	BroadcastPLMNs                   []*ServedCellInformationEUTRAPerBPLMN
	EUtraModeInfo                    ServedCellInformationEUTRAModeInfo
	NumberofAntennaPorts             *NumberOfAntennaPortsEUTRA
	PrachConfiguration               *EUTRAPRACHConfiguration
	MBSFNsubframeInfo                *MBSFNSubframeInfoEUTRA
	MultibandInfo                    *EUTRAMultibandInfoList
	FreqBandIndicatorPriority        *ServedCellInformationEUTRAFreqBandIndicatorPriority
	BandwidthReducedSI               *ServedCellInformationEUTRABandwidthReducedSI
	ProtectedEUTRAResourceIndication *ProtectedEUTRAResourceIndication
	IEExtensions                     []byte
}

func (ie *ServedCellInformationEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellInformationEUTRAConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Ranac != nil, ie.NumberofAntennaPorts != nil, ie.PrachConfiguration != nil, ie.MBSFNsubframeInfo != nil, ie.MultibandInfo != nil, ie.FreqBandIndicatorPriority != nil, ie.BandwidthReducedSI != nil, ie.ProtectedEUTRAResourceIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.EUtraPci.Encode(e); err != nil {
		return err
	}
	if err := ie.EUtraCgi.Encode(e); err != nil {
		return err
	}
	if err := ie.Tac.Encode(e); err != nil {
		return err
	}
	if ie.Ranac != nil {
		if err := ie.Ranac.Encode(e); err != nil {
			return err
		}
	}
	soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(1)),
		Max:        common.Ptr(int64(common.MaxnoofBPLMNs)),
	})
	if err := soEnc.EncodeLength(int64(len(ie.BroadcastPLMNs))); err != nil {
		return err
	}
	for _, item := range ie.BroadcastPLMNs {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.EUtraModeInfo.Encode(e); err != nil {
		return err
	}
	if ie.NumberofAntennaPorts != nil {
		if err := ie.NumberofAntennaPorts.Encode(e); err != nil {
			return err
		}
	}
	if ie.PrachConfiguration != nil {
		if err := ie.PrachConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if ie.MBSFNsubframeInfo != nil {
		if err := ie.MBSFNsubframeInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.MultibandInfo != nil {
		if err := ie.MultibandInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.FreqBandIndicatorPriority != nil {
		if err := ie.FreqBandIndicatorPriority.Encode(e); err != nil {
			return err
		}
	}
	if ie.BandwidthReducedSI != nil {
		if err := ie.BandwidthReducedSI.Encode(e); err != nil {
			return err
		}
	}
	if ie.ProtectedEUTRAResourceIndication != nil {
		if err := ie.ProtectedEUTRAResourceIndication.Encode(e); err != nil {
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

func (ie *ServedCellInformationEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellInformationEUTRAConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EUtraPci.Decode(d); err != nil {
		return err
	}
	if err := ie.EUtraCgi.Decode(d); err != nil {
		return err
	}
	if err := ie.Tac.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.Ranac = new(RANAC)
		if err := ie.Ranac.Decode(d); err != nil {
			return err
		}
	}
	{
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofBPLMNs)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.BroadcastPLMNs = make([]*ServedCellInformationEUTRAPerBPLMN, n)
		for i := range ie.BroadcastPLMNs {
			ie.BroadcastPLMNs[i] = new(ServedCellInformationEUTRAPerBPLMN)
			if err := ie.BroadcastPLMNs[i].Decode(d); err != nil {
				return err
			}
		}
	}
	if err := ie.EUtraModeInfo.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(6) {
		ie.NumberofAntennaPorts = new(NumberOfAntennaPortsEUTRA)
		if err := ie.NumberofAntennaPorts.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.PrachConfiguration = new(EUTRAPRACHConfiguration)
		if err := ie.PrachConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(8) {
		ie.MBSFNsubframeInfo = new(MBSFNSubframeInfoEUTRA)
		if err := ie.MBSFNsubframeInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(9) {
		ie.MultibandInfo = new(EUTRAMultibandInfoList)
		if err := ie.MultibandInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(10) {
		ie.FreqBandIndicatorPriority = new(ServedCellInformationEUTRAFreqBandIndicatorPriority)
		if err := ie.FreqBandIndicatorPriority.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(11) {
		ie.BandwidthReducedSI = new(ServedCellInformationEUTRABandwidthReducedSI)
		if err := ie.BandwidthReducedSI.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(12) {
		ie.ProtectedEUTRAResourceIndication = new(ProtectedEUTRAResourceIndication)
		if err := ie.ProtectedEUTRAResourceIndication.Decode(d); err != nil {
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
