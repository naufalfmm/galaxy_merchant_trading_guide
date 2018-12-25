package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
	"regexp"
	"errors"
	"strconv"

	"./romans"
	"./intergalactic"
)

// Input type 1
func interRomanInput(m map[string]string, inputStr string) error {
	spl := strings.Split(inputStr, " is ")
	m[spl[0]] = spl[1]
	return nil
}

// Input type 2
func metalsCreditValueInput(m map[string]string, n map[string]float32, inputStr string) error {
	var checkError error
	var metalValueRoman string
	var metalValue int

	spl := strings.Split(inputStr, " is ")
	metalValueType := strings.Split(spl[0], " ")
	metalCreditsValue, checkErr := strconv.Atoi(strings.Split(spl[1], " ")[0])
	if checkErr != nil {
		return checkErr
	}

	checkError, metalValueRoman = intergalactic.IntergalacticIntoRoman(m, metalValueType[:len(metalValueType)-1])
	if checkError != nil {
		return checkError
	}

	checkError, metalValue = romans.RomanIntoNumbersConv(metalValueRoman)
	if checkError != nil {
		return checkError
	}

	// The credit values of metal is stored in format the credit value of metal per metal
	n[metalValueType[len(metalValueType)-1]] = float32(metalCreditsValue)/float32(metalValue)

	return nil
}

// Input type 3
func interRomanQuestionInput(m map[string]string, inputStr string) (error, string) {
	var checkError error
	var romanizedIntergal string
	var romansValue int

	spl := strings.Split(inputStr, " is ")
	intergalDigitsMark := strings.Split(spl[1], " ")
	intergalDigits := intergalDigitsMark[:len(intergalDigitsMark)-1]
	
	checkError, romanizedIntergal = intergalactic.IntergalacticIntoRoman(m, intergalDigits)
	if checkError != nil {
		return checkError, ""
	}
	
	checkError, romansValue = romans.RomanIntoNumbersConv(romanizedIntergal)
	if checkError != nil {
		return checkError, ""
	}

	return nil, strings.Join(intergalDigits, " ") + " is " + strconv.Itoa(romansValue)
}

// Input type 4
func interCreditsValueQuestionInput(m map[string]string, n map[string]float32, inputStr string) (error, string)  {
	var checkError error
	var romanizedIntergal string
	var romansValue int
	var creditsValue float32

	spl := strings.Split(inputStr, " is ")
	intergalDigitsMark := strings.Split(spl[1], " ")
	metalType := intergalDigitsMark[len(intergalDigitsMark) - 2]
	intergalDigits := intergalDigitsMark[:len(intergalDigitsMark)-2]

	checkError, romanizedIntergal = intergalactic.IntergalacticIntoRoman(m, intergalDigits)
	if checkError != nil {
		return checkError, ""
	}

	checkError, romansValue = romans.RomanIntoNumbersConv(romanizedIntergal)
	if checkError != nil {
		return checkError, ""
	}
	
	creditsValue = n[metalType] * float32(romansValue)

	return nil, strings.Join(intergalDigits, " ") + " " + metalType + " is " + strconv.FormatFloat(float64(creditsValue), 'f', -1, 32) + " Credits"
}

// Classify input into 4 types
func inputClassification(m map[string]string, n map[string]float32, inputStr string) (error, string) {
	var checkError error
	var checkMatch bool
	var returnAnswer string
	var invalidInputMessage = "I have no idea what you are talking about"

	inputStr = strings.TrimRight(inputStr, "\r\n")
	// Check whether inputStr is type 1
	if checkMatch, checkError = regexp.MatchString("[a-zA-Z]+ is [IVXLCDMivxlcdm]{1}", inputStr); checkMatch {
		if checkError != nil {
			return checkError, ""
		}

		checkError = interRomanInput(m, inputStr)
		if checkError != nil {
			return checkError, ""
		}
		return nil, ""
	// Check whether inputStr is type 2
	} else if checkMatch, checkError = regexp.MatchString("[a-zA-Z ]+ is \\d+ [cC]redits", inputStr); checkMatch {
		if checkError != nil {
			return checkError, ""
		}

		checkError = metalsCreditValueInput(m, n, inputStr)
		if checkError != nil {
			return checkError, ""
		}
		return nil, ""
	// Check whether inputStr is type 3
	} else if checkMatch, checkError = regexp.MatchString("how much is [a-zA-Z ]+ \\?", inputStr); checkMatch {
		if checkError != nil {
			return checkError, ""
		}

		checkError, returnAnswer = interRomanQuestionInput(m, inputStr)
		if checkError != nil {
			return checkError, ""
		}
		return nil, returnAnswer
	// Check whether inputStr is type 4
	} else if checkMatch, checkError = regexp.MatchString("how many [cC]redits is [a-zA-Z ]+ \\?", inputStr); checkMatch {
		if checkError != nil {
			return checkError, ""
		}
		checkError, returnAnswer = interCreditsValueQuestionInput(m, n, inputStr)
		if checkError != nil {
			return checkError, ""
		}
		return nil, returnAnswer
	// inputStr is not in the following types, it will return invalidInputMessage
	} else {
		return errors.New(invalidInputMessage), ""
	}
}

func main()  {
	m := make(map[string]string)
	n := make(map[string]float32)
	in := bufio.NewReader(os.Stdin)

	str, err := in.ReadString('\n')
	for err != io.EOF {
		er, ans := inputClassification(m, n, str)
		if er == nil {
			if len(ans) > 0 {
				fmt.Println(ans)
			}
		} else {
			fmt.Println(er.Error())
		}
		str, err = in.ReadString('\n')
	}
}