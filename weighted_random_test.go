package goutils

import "testing"

func TestWeightedChoice(t *testing.T) {
	// sum of weight = 0
	// 1. no choices
	i, _ := WeightedChoice(nil)
	if i != -1 {
		t.Errorf("expected -1 but get %v", i)
	}

	// 2. 3 choices, select second one
	SetRandomIntn(func(int) int {
		return 1
	})
	_, choice := WeightedChoice([]Choice{
		{Item: "1"},
		{Item: "2"},
		{Item: "3"},
	})
	if item := choice.Item.(string); item != "2" {
		t.Errorf("expected item 2 but get %v", item)
	}

	// sum of weight > 0
	SetRandomIntn(func(int) int {
		return 45
	})
	_, choice = WeightedChoice([]Choice{
		{Item: "1", Weight: 10},
		{Item: "2", Weight: 10},
		{Item: "3", Weight: 35},
		{Item: "4", Weight: 25},
		{Item: "5", Weight: 10},
		{Item: "6", Weight: 10},
	})
	if item := choice.Item.(string); item != "3" {
		t.Errorf("expected item 3 but get %v", item)
	}
}
