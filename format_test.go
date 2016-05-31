package fnum

import (
	"testing"
)

func TestFormatInt(t *testing.T) {
	testList{
		{"1", FormatInt(1, ",")},
		{"10", FormatInt(10, ",")},
		{"100", FormatInt(100, ",")},
		{"1,000", FormatInt(1000, ",")},
		{"1.000", FormatInt(1000, ".")},
		{"98.665", FormatInt(98665, ".")},
		{"5,555", FormatInt(5555, ",")},
		{"108.654", FormatInt(108654, ".")},
		{"500,123", FormatInt(500123, ",")},
		{"1.345.099", FormatInt(1345099, ".")},
		{"10,876,890", FormatInt(10876890, ",")},
		{"890,765", FormatInt(890765, ",")},
	}.validate(t)
}

func TestFormatFloat(t *testing.T) {
	testList{
		{"1", FormatFloat(1.00, ".", ",")},
		{"10", FormatFloat(10.00, ".", ",")},
		{"100,05", FormatFloat(100.05, ".", ",")},
		{"1.000,1", FormatFloat(1000.10, ".", ",")},
		{"986.65", FormatFloat(986.65, ",", ".")},
		{"5,555.55", FormatFloat(5555.55, ",", ".")},
		{"108.654,3", FormatFloat(108654.30, ".", ",")},
		{"500,123.32", FormatFloat(500123.32, ",", ".")},
		{"1.345.099,25", FormatFloat(1345099.25, ".", ",")},
		{"10,876,890.88", FormatFloat(10876890.88, ",", ".")},
		{"890.765,74", FormatFloat(890765.74, ".", ",")},
	}.validate(t)
}
