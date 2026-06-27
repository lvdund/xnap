package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qMCCoordinationRequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mN-to-SN-QMCCoordRequestList", Optional: true},
		{Name: "sN-to-MN-QMCCoordRequestList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QMCCoordinationRequest struct {
	MNToSNQMCCoordRequestList *MNToSNQMCCoordRequestList
	SNToMNQMCCoordRequestList *SNToMNQMCCoordRequestList
	IEExtensions              []byte
}

func (ie *QMCCoordinationRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qMCCoordinationRequestConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MNToSNQMCCoordRequestList != nil, ie.SNToMNQMCCoordRequestList != nil, false}); err != nil {
		return err
	}
	if ie.MNToSNQMCCoordRequestList != nil {
		if err := ie.MNToSNQMCCoordRequestList.Encode(e); err != nil {
			return err
		}
	}
	if ie.SNToMNQMCCoordRequestList != nil {
		if err := ie.SNToMNQMCCoordRequestList.Encode(e); err != nil {
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

func (ie *QMCCoordinationRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qMCCoordinationRequestConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.MNToSNQMCCoordRequestList = new(MNToSNQMCCoordRequestList)
		if err := ie.MNToSNQMCCoordRequestList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.SNToMNQMCCoordRequestList = new(SNToMNQMCCoordRequestList)
		if err := ie.SNToMNQMCCoordRequestList.Decode(d); err != nil {
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
