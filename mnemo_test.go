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

package mnemo_test

import (
	"testing"

	"github.com/murdockq/mnemo"
)

var gold_answers = []struct {
	key   int
	value string
}{
	{0, "ba"},
	{1, "bi"},
	{34, "jo"},
	{42, "lu"},
	{99, "zo"},
	{100, "biba"},
	{101, "bibi"},
	{41112, "bodidu"},
	{412122, "ligigu"},

	{392406, "kogochi"},
	{25437225, "haleshuha"},
	{392406392406392406, "kogochikogochikogochi"},
	{2543722525437225254, "bunokugununokugununo"},

	{-1, "xabi"},
	{-31, "xaji"},
	{-99, "xazo"},
	{-100, "xabiba"},
	{947110, "yoshida"},
	{-947110, "xayoshida"},
	{79523582, "tonukatsu"},
}

func Test_id_to_string_list(t *testing.T) {
	for _, answer := range gold_answers {

		value := mnemo.ToString(answer.key)
		expected := answer.value
		if value != expected {
			t.Errorf("answer:%v expected:{%v %v}", answer, value, expected)
		}
	}
}

func Test_string_to_id_list(t *testing.T) {
	for _, answer := range gold_answers {

		value, _ := mnemo.ToInt(answer.value)
		expected := answer.key
		if value != expected {
			t.Errorf("answer:%v expected:{%v %v}", answer, value, expected)
		}
	}
}

func Test_negative_id_to_string(t *testing.T) {
	value := mnemo.ToString(-1)
	expected := "xabi"
	if value != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_string_to_id(t *testing.T) {
	value, err := mnemo.ToInt("xabi")
	expected := -1
	if err != nil && value != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_convert_string_to_id(t *testing.T) {
	_, i := mnemo.Convert("xabi")
	expected := -1
	if i != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_invalid_convert_string_to_id(t *testing.T) {
	_, i := mnemo.Convert("invalid")
	expected := 0
	if i != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_convert_id_to_string(t *testing.T) {
	s, _ := mnemo.Convert(-1)
	expected := "xabi"
	if s != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_convert_nil(t *testing.T) {
	s, i := mnemo.Convert(nil)
	expected := 0
	if s != "" && i != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_is_valid_fail(t *testing.T) {
	value := mnemo.IsValid("invalid")
	expected := false
	if value != expected {
		t.Errorf("expected %v", expected)
	}
}

func Test_is_valid_succeed(t *testing.T) {
	value := mnemo.IsValid("yoshida")
	expected := true
	if value != expected {
		t.Errorf("expected %v", expected)
	}
}

func Benchmark_to_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mnemo.ToString(i)
	}
}

func Benchmark_convert_to_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mnemo.Convert(i)
	}
}

func Benchmark_to_int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, answer := range gold_answers {
			mnemo.ToInt(answer.value)
		}
	}
}

func Benchmark_convert_to_int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, answer := range gold_answers {
			mnemo.Convert(answer.value)
		}
	}
}
