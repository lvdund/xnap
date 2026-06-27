package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qMCCoordinationResponseConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mN-to-SN-QMCCoordResponseList", Optional: true},
		{Name: "sN-to-MN-QMCCoordResponseList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QMCCoordinationResponse struct {
	MNToSNQMCCoordResponseList *MNToSNQMCCoordResponseList
	SNToMNQMCCoordResponseList *SNToMNQMCCoordResponseList
	IEExtensions               []byte
}

func (ie *QMCCoordinationResponse) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qMCCoordinationResponseConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MNToSNQMCCoordResponseList != nil, ie.SNToMNQMCCoordResponseList != nil, false}); err != nil {
		return err
	}
	if ie.MNToSNQMCCoordResponseList != nil {
		if err := ie.MNToSNQMCCoordResponseList.Encode(e); err != nil {
			return err
		}
	}
	if ie.SNToMNQMCCoordResponseList != nil {
		if err := ie.SNToMNQMCCoordResponseList.Encode(e); err != nil {
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

func (ie *QMCCoordinationResponse) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qMCCoordinationResponseConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.MNToSNQMCCoordResponseList = new(MNToSNQMCCoordResponseList)
		if err := ie.MNToSNQMCCoordResponseList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.SNToMNQMCCoordResponseList = new(SNToMNQMCCoordResponseList)
		if err := ie.SNToMNQMCCoordResponseList.Decode(d); err != nil {
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
