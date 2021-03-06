package lib

import (
	"testing"
)

func TestMyName(t *testing.T) {
	if MyName() != "pstop" {
		t.Errorf("MyName() expected to be %v but actually was %v", "pstop", MyName())
	}
}

func Test_sec_to_time(t *testing.T) {
	type stuff struct {
		input  uint64
		output string
	}
	test_data := []stuff{
		{0, "00:00:00"},
		{1, "00:00:01"},
		{61, "00:01:01"},
		{3601, "01:00:01"},
		{7384, "02:03:04"},
		// add more values here
	}
	for i := range test_data {
		if sec_to_time(test_data[i].input) != test_data[i].output {
			t.Errorf("sec_to_time(%v) expected to be %v but actually was %v", test_data[i].input, test_data[i].output, sec_to_time(test_data[i].input))
		}
	}
}

func TestFormatTime(t *testing.T) {
	type stuff struct {
		input  uint64
		output string
	}
	test_data := []stuff{
		{0, ""},
		{1, "1 ps"},
		{1000, "   1.00 ns"},
		{1000000, "   1.00 us"},
		{1000000000, "   1.00 ms"},
		{1000000000000, "    1.00 s"},
		// add more values here
	}
	for i := range test_data {
		if FormatTime(test_data[i].input) != test_data[i].output {
			t.Errorf("FormatTime(%v) expected to be %v but actually was %v", test_data[i].input, test_data[i].output, FormatTime(test_data[i].input))
		}
	}
}
