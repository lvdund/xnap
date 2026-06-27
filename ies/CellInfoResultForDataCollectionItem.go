package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellInfoResultForDataCollectionItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellID"},
		{Name: "predictedRadioResourceStatus", Optional: true},
		{Name: "predictedNumberofActiveUEs", Optional: true},
		{Name: "predictedRRCConnections", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellInfoResultForDataCollectionItem struct {
	CellID                       GlobalNGRANCellID
	PredictedRadioResourceStatus *RadioResourceStatus
	PredictedNumberofActiveUEs   *NumberofActiveUEs
	PredictedRRCConnections      *RRCConnections
	IEExtensions                 []byte
}

func (ie *CellInfoResultForDataCollectionItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellInfoResultForDataCollectionItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PredictedRadioResourceStatus != nil, ie.PredictedNumberofActiveUEs != nil, ie.PredictedRRCConnections != nil, false}); err != nil {
		return err
	}
	if err := ie.CellID.Encode(e); err != nil {
		return err
	}
	if ie.PredictedRadioResourceStatus != nil {
		if err := ie.PredictedRadioResourceStatus.Encode(e); err != nil {
			return err
		}
	}
	if ie.PredictedNumberofActiveUEs != nil {
		if err := ie.PredictedNumberofActiveUEs.Encode(e); err != nil {
			return err
		}
	}
	if ie.PredictedRRCConnections != nil {
		if err := ie.PredictedRRCConnections.Encode(e); err != nil {
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

func (ie *CellInfoResultForDataCollectionItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellInfoResultForDataCollectionItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CellID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PredictedRadioResourceStatus = new(RadioResourceStatus)
		if err := ie.PredictedRadioResourceStatus.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.PredictedNumberofActiveUEs = new(NumberofActiveUEs)
		if err := ie.PredictedNumberofActiveUEs.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.PredictedRRCConnections = new(RRCConnections)
		if err := ie.PredictedRRCConnections.Decode(d); err != nil {
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
