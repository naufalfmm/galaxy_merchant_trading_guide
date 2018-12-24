
Galaxy Merchant Trading Guide

===========================
This is repository of the Galaxy Merchant Trading Guide problem solution created by [Naufal Farras][1]. The solution of problem is typed in [Go Programming Language][2].

#  Requirement

The solution requires Go >= 1.8.

#  Problem Statement

You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot).Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you.The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate translation between them.Roman numerals are based on seven symbols.

####  Symbol Value
* I 1
* V 5
* X 10
* L 250
* C 100
* D 500
* M 1,000

Numbers are formed by combining symbols together and adding the values. For example, MMVI is 1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.) "D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately. Inthe above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

[Source: Wikipedia Roman Numbers][3]

Input to your program consists of lines of text detailing your notes on the conversion between intergalactic units and roman numerals. **You are expected to handle invalid queries appropriately**.

###  Test input:
```
glob is I
prok is V
pish is X
tegj is L
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
how much is pish tegj glob glob ?
how many Credits is glob prok Silver ?
how many Credits is glob prok Gold ?
how many Credits is glob prok Iron ?
how much wood could a woodchuck chuck if a woodchuck could chuck wood ?
```
###  Test Output:
```
pish tegj glob glob is 42
glob prok Silver is 68 Credits
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
I have no idea what you are talking about
```
#  Solution Approach
The flowchart of solution looks like below

![Solution Flowchart][4]

Input is divided into 4 types (categorized by using Regex):

###  Type 1

Input categorized into type 1 looks like *glob is I* with regex
```
"[a-zA-Z]+ is [IVXLCDMivxlcdm]{1}"
```
The type of input tells convention between intergalactic numerals and roman numerals. For example
* glob is I means glob equals to I
* prok is V means prok equals t V
* etc

The name of function for the task of type is *interRomanInput()*.

###  Type 2

Input categorized into type 2 looks like *glob glob Silver is 34 Credits* with regex
```
"[a-zA-Z ]+ is \\d+ [cC]redits"
```
The type of input tells credits value of the metal. For example
* glob glob Silver is 34 Credits means 2 (glob glob is II as following convention) silver equals 34 Credits
* glob prok Gold is 57800 Credits means 4 (glob prok is IV as following convention) gold equals 57800 Credits
* etc

The name of function for the task of type is *metalsCreditValueInput()*.

###  Type 3
Input categorized into type 3 looks like *how much is pish tegj glob glob ?* with regex
```
"how much is [a-zA-Z ]+ \\?"
```
The type of input is a question of intergalactic numerals value (through roman numeral). It will return answer the value of intergalactic numerals asked. For example
* how much is pish tegj glob glob? will return pish tegj glob glob is 42
* etc

The name of function for the task of type is *interRomanQuestionInput()*.

###  Type 4

Input categorized into type 4 looks like *how many Credits is glob prok Silver ?* with regex
```
"how many [cC]redits is [a-zA-Z ]+ \\?"
```
The type of input is a question of some of metals credits value. It will return answer the value of some of metals. For example
* how many Credits is glob prok Silver ? will return glob prok Silver is 68 Credits
* how many Credits is glob prok Gold ? will return glob prok Gold is 57800 Credits
* etc

The name of function for the task of type is *interCreditsValueQuestionInput()*.

If input can't be categorized into the following types, program will return "I have no idea what you are talking about".

#  Tests

The tests are done manually not programmatically. The tests are divided into

##  Unit Tests

####  RomanIntoNumberConv() in './romans/romans.go'

Convert roman digits into numbers (int)

| Test Case | Expected | Actual |
| ------------ | -------- | ------- |
| VIII | 8 | 8 |
| IX | 9 | 9 |
| VX | *Error* | *Error* |
| IIX | *Error* | *Error* |
| IM | *Error* | *Error* |
| LCM | *Error* | *Error* |

####  intergalacticIntoRoman()

Convert intergalactic digits into roman digits

**Known glob is I, prok is V, pish is X, tegj is L**

| Test Case | Expected | Actual |
| -------------- | -------- | ------- |
| glob glob | II | II |
| pish glob prok | XIV | XIV |
| tegj | L | L |
| mush pish | *Error* | *Error* |

####  inputClassification()

Classify input into 4 types

| Test Case | Expected | Actual |
| --------- | -------- | ------- |
| glob is I | Type 1 | Type 1 |
| glob glob Silver is 34 Credits | Type 2 | Type 2 |
| how much is pish tegj glob glob ? | Type 3 | Type 3 |
| how many Credits is glob prok Iron ? | Type 4 | Type 4 |
| how much wood could a woodchuck chuck if a woodchuck could chuck wood ? | *Error* | *Error* |
| aaaaaaaegegwebrwhwr | *Error* | *Error* |

[1]: https://gitlab.com/naufalfmm

[2]: https://golang.org/

[3]: http://en.wikipedia.org/wiki/Roman_numerals

[4]: https://github.com/naufalfmm/galaxy_merchant_trading_guide/blob/master/docs/solution_flow.png