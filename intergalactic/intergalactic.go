package intergalactic

import (
	"errors"
	"strings"
)

// Convert intergalactic digits into roman digits
func IntergalacticIntoRoman(m map[string]string, interUnits []string) (error, string) {
	var interUnitsLen int
	var intergalacticValue string
	var checkExist bool
	var interUnitsAnswer []string

	interUnitsLen = len(interUnits)

	for i := 0; i < interUnitsLen; i++ {
		intergalacticValue, checkExist = m[interUnits[i]]
		if !checkExist {
			return errors.New("Invalid: Digit not found"), ""
		}
		interUnitsAnswer = append(interUnitsAnswer, intergalacticValue)
	}

	return nil, strings.Join(interUnitsAnswer, "")
}