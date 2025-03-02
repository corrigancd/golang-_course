package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for index, stringVal := range strings {
		float, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("There was a problem converting string to float64: " + err.Error())
		}

		floats[index] = float
	}

	return floats, nil
}
