package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceChangeRequiredInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataforwardinginfofromSource", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceChangeRequiredInfoSNterminated struct {
	DataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource
	IEExtensions                 []byte
}

func (ie *PDUSessionResourceChangeRequiredInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceChangeRequiredInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DataforwardinginfofromSource != nil, false}); err != nil {
		return err
	}
	if ie.DataforwardinginfofromSource != nil {
		if err := ie.DataforwardinginfofromSource.Encode(e); err != nil {
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

func (ie *PDUSessionResourceChangeRequiredInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceChangeRequiredInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DataforwardinginfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardinginfofromSource.Decode(d); err != nil {
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
