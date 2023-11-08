
package model

import "errors"

// SimpleModel represents a basic linear model for demonstration
type SimpleModel struct {
	Weights []float64
	Bias    float64
}

// NewSimpleModel creates a new SimpleModel instance
func NewSimpleModel(weights []float64, bias float64) *SimpleModel {
	return &SimpleModel{
		Weights: weights,
		Bias:    bias,
	}
}

// Predict performs a simple linear prediction
func (m *SimpleModel) Predict(input []float64) (float64, error) {
	if len(input) != len(m.Weights) {
		return 0, errors.New("input dimension mismatch")
	}

	var sum float64
	for i, val := range input {
		sum += val * m.Weights[i]
	}
	return sum + m.Bias, nil
}
