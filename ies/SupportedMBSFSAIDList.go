package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var supportedMBSFSAIDListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSFSAs)),
}

type SupportedMBSFSAIDList struct {
	Value []*MBSFrequencySelectionAreaIdentity
}

func (ie *SupportedMBSFSAIDList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(supportedMBSFSAIDListConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SupportedMBSFSAIDList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(supportedMBSFSAIDListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSFrequencySelectionAreaIdentity, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSFrequencySelectionAreaIdentity)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
