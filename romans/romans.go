package romans

import (
	"errors"
)

/*
	Roman Package
	Roman digits are ranged I into M.
*/

var romanValueConv = map[byte]int{
		73: 1, //I
		86: 5, //V
		88: 10, //X
		76: 50, //L
		67: 100, //C
		68: 500, //D
		77: 1000} //M

// 
func checkGetRomanDigit(romanDigit byte) (error, int) {
	if romanDigit >= 97 && romanDigit <= 122 {
		romanDigit -= 32
	}

	romanValue, checkExist := romanValueConv[romanDigit]
	if !checkExist {
		return errors.New("Invalid: Digit not found"), 0
	}

	return nil, romanValue
}

// Convert betweem roman digits into number.
func RomanIntoNumbersConv(roman string) (error, int) {
	var lastRomanDigit byte
	var romanStringLen, romansValue, lastRomansValue, lastRomanDigitCount int = len(roman), 0, 0, 0
	var romanValue, romanNextValue int
	var checkError, checkErrorNext error
	
	for i := 0; i < romanStringLen; {
		checkError, romanValue = checkGetRomanDigit(roman[i])
		if checkError != nil {
			return checkError, 0
		}

		if romanValue > lastRomansValue && i > 0 {
			return errors.New("Invalid: The order of digits is invalid"), 0
		}

		if i < romanStringLen-1 {
			checkErrorNext, romanNextValue = checkGetRomanDigit(roman[i+1])

			if checkErrorNext != nil {
				return checkErrorNext, 0
			}

			if romanValue >= romanNextValue {
				romansValue += romanValue
				i++
				lastRomansValue = romanValue
			} else {
				if romanValue == 5 || romanValue == 50 || romanValue == 500 {
					return errors.New("Invalid: Digits L, V, and D cannot subtract"), 0
				}

				if romanValue < romanNextValue/10 || romanValue > romanNextValue {
					return errors.New("Invalid: Digits I, X, and C can subtract only the second greater digit"), 0
				}

				if romanNextValue > romanValueConv[lastRomanDigit] && i > 0 {
					return errors.New("Invalid: Subtracting digit cannot subtract more than one time"), 0
				}

				romansValue += (romanNextValue - romanValue)
				lastRomansValue = (romanNextValue - romanValue)
				i += 2
			}
		} else {
			romansValue += romanValue
			i++
		}

		if roman[i-1] == lastRomanDigit {
			if roman[i-1] == 86 || roman[i-1] == 118 || roman[i-1] == 76 || roman[i-1] == 108 || roman[i-1] == 68 || roman[i-1] == 100 {
				return errors.New("Invalid: Digits L, V, and D cannot be repeated"), 0
			}
			lastRomanDigitCount++
		} else {
			lastRomanDigit = roman[i-1]
			lastRomanDigitCount = 1
		}

		if lastRomanDigitCount == 4 {
			return errors.New("Invalid: Digit cannot be repeated more than thrice"), 0
		}
	}

	return nil, romansValue
}