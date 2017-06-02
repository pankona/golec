// fuzzbuzz_test.go

package main

import "testing"

type testData struct {
	input  int
	expect string
}

func TestFizzBuzz(t *testing.T) {
	testDataTable := []testData{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
	}

	for _, v := range testDataTable {
		if fizzbuzz(v.input) != v.expect {
			t.Errorf("error. input = %d, expect = %s, actual return = %s", v.input, v.expect, fizzbuzz(v.input))
		}
	}
}
