package ies

import (
	"github.com/lvdund/asn1go/per"
)

var criticalityDiagnosticsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "procedureCode", Optional: true},
		{Name: "triggeringMessage", Optional: true},
		{Name: "procedureCriticality", Optional: true},
		{Name: "iEsCriticalityDiagnostics", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode
	TriggeringMessage         *TriggeringMessage
	ProcedureCriticality      *Criticality
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList
	IEExtensions              []byte
}

func (ie *CriticalityDiagnostics) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(criticalityDiagnosticsConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ProcedureCode != nil, ie.TriggeringMessage != nil, ie.ProcedureCriticality != nil, ie.IEsCriticalityDiagnostics != nil, false}); err != nil {
		return err
	}
	if ie.ProcedureCode != nil {
		if err := ie.ProcedureCode.Encode(e); err != nil {
			return err
		}
	}
	if ie.TriggeringMessage != nil {
		if err := ie.TriggeringMessage.Encode(e); err != nil {
			return err
		}
	}
	if ie.ProcedureCriticality != nil {
		if err := ie.ProcedureCriticality.Encode(e); err != nil {
			return err
		}
	}
	if ie.IEsCriticalityDiagnostics != nil {
		if err := ie.IEsCriticalityDiagnostics.Encode(e); err != nil {
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

func (ie *CriticalityDiagnostics) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(criticalityDiagnosticsConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ProcedureCode = new(ProcedureCode)
		if err := ie.ProcedureCode.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.TriggeringMessage = new(TriggeringMessage)
		if err := ie.TriggeringMessage.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.ProcedureCriticality = new(Criticality)
		if err := ie.ProcedureCriticality.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.IEsCriticalityDiagnostics = new(CriticalityDiagnosticsIEList)
		if err := ie.IEsCriticalityDiagnostics.Decode(d); err != nil {
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
