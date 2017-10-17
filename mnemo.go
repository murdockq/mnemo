/*
MIT License

Copyright (c) 2017 github.com/murdockq

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package mnemo provides generation of
// human readible names for ids.
package mnemo

import (
	"fmt"
	"strings"
)

type mnemo struct {
	name  string
	index int
}

var syllables = []string{
	"ba", "bi", "bu", "be", "bo",
	"cha", "chi", "chu", "che", "cho",
	"da", "di", "du", "de", "do",
	"fa", "fi", "fu", "fe", "fo",
	"ga", "gi", "gu", "ge", "go",
	"ha", "hi", "hu", "he", "ho",
	"ja", "ji", "ju", "je", "jo",
	"ka", "ki", "ku", "ke", "ko",
	"la", "li", "lu", "le", "lo",
	"ma", "mi", "mu", "me", "mo",
	"na", "ni", "nu", "ne", "no",
	"pa", "pi", "pu", "pe", "po",
	"ra", "ri", "ru", "re", "ro",
	"sa", "si", "su", "se", "so",
	"sha", "shi", "shu", "she", "sho",
	"ta", "ti", "tu", "te", "to",
	"tsa", "tsi", "tsu", "tse", "tso",
	"wa", "wi", "wu", "we", "wo",
	"ya", "yi", "yu", "ye", "yo",
	"za", "zi", "zu", "ze", "zo",
}

var negative = "xa"

// toString recursively generate mnemonic from int
func toString(i int, s *string) {
	mod := i % len(syllables)
	rst := i / len(syllables)
	if rst > 0 {
		toString(rst, s)
	}
	*s += syllables[mod]
}

// ToString provides a way to convert integer to mnemoic
func ToString(i int) (mnemonic string) {
	if i < 0 {
		mnemonic += negative
		i = -i
	}

	toString(i, &mnemonic)

	return
}

// ToInt provides a way to convert mnemoic to integer
func ToInt(s string) (value int, err error) {
	sign := 1

	if strings.HasPrefix(s, negative) {
		sign = -1
		s = strings.TrimPrefix(s, negative)
	}

	for len(s) > 0 && err == nil {
		for i, syllable := range syllables {
			if strings.HasPrefix(s, syllable) {
				//TODO: Can be sped up by not making copies and just doing substring
				s = strings.TrimPrefix(s, syllable)
				value = len(syllables)*value + i
				break
			}

			//if it reaches the last element of the array then it is not valid
			if i == len(syllables)-1 {
				err = fmt.Errorf("mnemo: invalid mnemonic %s", s)
				value = 0
				break
			}
		}
	}

	value = sign * value

	return
}

// IsValid checks if the passed in string is a valid mnemonic
func IsValid(s string) bool {
	if _, err := ToInt(s); err != nil {
		return false
	}
	return true
}

// Convert is a convience method to convert mnemoic to and from an integer
func Convert(input ...interface{}) (mnemonic string, value int) {
	for _, it := range input {
		switch it.(type) {
		case string:
			mnemonic = it.(string)
			value, _ = ToInt(it.(string))
		case int:
			mnemonic = ToString(it.(int))
			value = it.(int)
		}
	}
	return
}
