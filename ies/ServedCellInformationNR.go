package ies

import (
	"github.com/lvdund/asn1go/per"
)

var servedCellInformationNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrPCI"},
		{Name: "cellID"},
		{Name: "tac"},
		{Name: "ranac", Optional: true},
		{Name: "broadcastPLMN"},
		{Name: "nrModeInfo"},
		{Name: "measurementTimingConfiguration"},
		{Name: "connectivitySupport"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellInformationNR struct {
	NrPCI                          NRPCI
	CellID                         NRCGI
	Tac                            TAC
	Ranac                          *RANAC
	BroadcastPLMN                  BroadcastPLMNs
	NrModeInfo                     NRModeInfo
	MeasurementTimingConfiguration []byte
	ConnectivitySupport            ConnectivitySupport
	IEExtensions                   []byte
}

func (ie *ServedCellInformationNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellInformationNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Ranac != nil, false}); err != nil {
		return err
	}
	if err := ie.NrPCI.Encode(e); err != nil {
		return err
	}
	if err := ie.CellID.Encode(e); err != nil {
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
	if err := ie.BroadcastPLMN.Encode(e); err != nil {
		return err
	}
	if err := ie.NrModeInfo.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.MeasurementTimingConfiguration, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if err := ie.ConnectivitySupport.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ServedCellInformationNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellInformationNRConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrPCI.Decode(d); err != nil {
		return err
	}
	if err := ie.CellID.Decode(d); err != nil {
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
	if err := ie.BroadcastPLMN.Decode(d); err != nil {
		return err
	}
	if err := ie.NrModeInfo.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.MeasurementTimingConfiguration = val
	}
	if err := ie.ConnectivitySupport.Decode(d); err != nil {
		return err
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
