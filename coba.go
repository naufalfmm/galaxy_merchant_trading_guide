package main

import (
	"fmt"
	"io"
	"errors"
	// "bufio"
	// "os"
	// "strings"
)

func romanIntoNumbersConv(roman string) (error, int) {
	var romanValueConv map[byte]int
	var lastRomanDigit byte
	var romanStringLen, romansValue, lastRomansValue, lastRomanDigitCount int = len(roman), 0, 0, 0
	var romanValue, romanNextValue int
	var checkExist, checkExistNext bool

	romanValueConv = map[byte]int{
		73: 1, //I
		86: 5, //V
		88: 10, //X
		76: 50, //L
		67: 100, //C
		68: 500, //D
		77: 1000, //M
		105: 1, //i
		118: 5, //v
		120: 10, //x
		108: 50, //;
		99: 100, //c
		100: 500, //d
		109: 1000} //m
	
	for i := 0; i < romanStringLen; {
		romanValue, checkExist = romanValueConv[roman[i]]

		if !checkExist {
			return errors.New("Invalid: Digit not found"), 0
		}

		if romanValue > lastRomansValue && i > 0 {
			return errors.New("Invalid: The order of digits is invalid"), 0
		}

		if i < romanStringLen-1 {
			romanNextValue, checkExistNext = romanValueConv[roman[i+1]]

			if !checkExistNext {
				return errors.New("Invalid: Digit not found"), 0
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
					return errors.New("Invalid: Digits I, X, and C can subtract only the two greatest digits"), 0
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

// func intergalacticIntoRoman(m map[string]string, interUnits []string) string {
// 	// fmt.Println(len(interUnits))
// 	interUnitsLen := len(interUnits)

// 	for i := 0; i < interUnitsLen; i++ {
// 		// fmt.Println(interUnits[i], m[interUnits[i]])
// 		// fmt.Println(len(interUnits[i]))
// 		interUnits[i] = m[interUnits[i]]
// 	}
// 	// fmt.Println(interUnits)
// 	err := dummy()
// 	fmt.Println(err.Error())

// 	return strings.Join(interUnits, "")
// }

// func sodummy() interface{} {
// 	// fmt.Println("Something Happen")
// 	return 50
// }

// func dummy() error  {
// 	// panic("Error")
// 	// defer sodummy()
// 	return fmt.Errorf("Something happen")
// }

func main()  {
	var a string
	// a := 50
	// m := map[string]string{
	// 	"glob": "I",
	// 	"prok": "V",
	// 	"pish": "X",
	// 	"tegj": "L"}
	
	// fmt.Println(sodummy() + 1)
	// fmt.Println(strings.Split(a, " "))

	// n := map[string]int{
	// 	"glob": 1,
	// 	"prok": 2,
	// 	"pish": 3,
	// 	"tegj": 4}
	
	// val1, ok1 := m["globe"]
	// val2, ok2 := n["globe"] 

	// fmt.Println(val1, ok1, val2, ok2)

	// b := strings.Split("how much is pish tegj glob glob ?", " ")
	// fmt.Println(len(b))

	// in := bufio.NewReader(os.Stdin)

	// str, err := in.ReadString('\n')
	// for err != io.EOF {
	// 	// fmt.Println(romanIntoNumbersConv(a))
	// 	str = str[:len(str)-2]
	// 	// fmt.Println(len(str))
	// 	// str = strings.TrimSuffix(str, "\r\n")
	// 	fmt.Println(intergalacticIntoRoman(m, strings.Split(str, " ")))
	// 	str, err = in.ReadString('\n')
	// }

	_, err := fmt.Scanln(&a)
	for err != io.EOF {
		er, res := romanIntoNumbersConv(a)
		if er != nil {
			fmt.Println(er)
		} else {
			fmt.Println(res)
		}
		// fmt.Println(intergalacticIntoRoman(m, strings.Split(a, " ")))
		_, err = fmt.Scanln(&a)
	}
}