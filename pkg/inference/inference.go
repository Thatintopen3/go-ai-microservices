
package inference

import (
	"errors"
	"go-ai-microservices/pkg/model"
)

// Predict uses the provided model to make a prediction on the input data
func Predict(m *model.SimpleModel, data []float64) (float64, error) {
	if m == nil {
		return 0, errors.New("model is nil")
	}

	// In a real scenario, this would involve more complex inference logic
	// For this example, we're using the simple model's predict method
	prediction, err := m.Predict(data)
	if err != nil {
		return 0, fmt.Errorf("model prediction failed: %w", err)
	}

	return prediction, nil
}
