package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rRCReestabInitiatedReportingWoUERLFReportConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failureCellPCI"},
		{Name: "reestabCellCGI"},
		{Name: "c-RNTI"},
		{Name: "shortMAC-I"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RRCReestabInitiatedReportingWoUERLFReport struct {
	FailureCellPCI NGRANCellPCI
	ReestabCellCGI GlobalNGRANCellID
	CRNTI          CRNTI
	ShortMACI      MACI
	IEExtensions   []byte
}

func (ie *RRCReestabInitiatedReportingWoUERLFReport) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestabInitiatedReportingWoUERLFReportConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.FailureCellPCI.Encode(e); err != nil {
		return err
	}
	if err := ie.ReestabCellCGI.Encode(e); err != nil {
		return err
	}
	if err := ie.CRNTI.Encode(e); err != nil {
		return err
	}
	if err := ie.ShortMACI.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RRCReestabInitiatedReportingWoUERLFReport) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestabInitiatedReportingWoUERLFReportConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FailureCellPCI.Decode(d); err != nil {
		return err
	}
	if err := ie.ReestabCellCGI.Decode(d); err != nil {
		return err
	}
	if err := ie.CRNTI.Decode(d); err != nil {
		return err
	}
	if err := ie.ShortMACI.Decode(d); err != nil {
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
