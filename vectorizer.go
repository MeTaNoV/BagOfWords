package Vectorizer

import (
	"github.com/MeTaNoV/CloudForest"
)

type Vectorizer struct {
	maxFeatures  int // max number of features to keep
	maxSizeGroup int // number of contiguous words that can be grouped together to form a feature
}

func New(maxFeatures, maxSizeGroup int) *Vectorizer {
	v := new(Vectorizer)
	v.maxFeatures = maxFeatures
	v.maxSizeGroup = maxSizeGroup
	return &v
}

// FitTranform take a training set of string to learn and set up a features set
// and return the associated feature set
func (v *Vectorizer) FitTranform() *FeatureMatrix {

}

// Tranform take a test set of string and return the associated feature set taking into account its own vocabulary
func (v *Vectorizer) Tranform() *FeatureMatrix {

}
