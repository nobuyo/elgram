package main

import (
	"testing"
)

func TestPickElements(t *testing.T) {
	t.Parallel()

	cases := []struct {
		desc        string
		args        []string
		expect      string
		shouldError bool
	}{
		{
			"Should return correct result from number",
			[]string{"", "3684736", "1356"},
			"3873",
			false,
		},
		{
			"Should return correct result from string",
			[]string{"", "teststring", "4733"},
			"trss",
			false,
		},
		{
			"When indices are comma-separeted",
			[]string{"", "longlongteststring", "5,3,6,10,12,3"},
			"lnoetn",
			false,
		},
		{
			"Should return error when illegal index range is given",
			[]string{"", "test", "5,3,1"},
			"",
			true,
		},
		{
			"Should return error when illegal character is given",
			[]string{"", "test", "14x"},
			"",
			true,
		},
		{
			"Should return error when too few argments",
			[]string{"", "test"},
			"",
			true,
		},
		{
			"Should return error when too much argments",
			[]string{"", "test", "1234", "3432"},
			"",
			true,
		},
	}

	for _, testCase := range cases {
		t.Log(testCase.desc)
		result, err := pickElements(testCase.args)

		if err != nil && !testCase.shouldError {
			t.Fatal("Got error", err)
		}

		if result != testCase.expect {
			t.Fatalf("Fatal: expected: %v, got: %v", testCase.expect, result)
		}
	}
}
