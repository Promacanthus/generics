package generic

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestMax(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		args []T
		want T
	}

	case1 := testCase[int32]{"int32", []int32{1, 2, 3}, 3}
	case2 := testCase[float64]{"float64", []float64{1.2, 3.4, 5.6}, 5.6}

	t.Run(case1.name, func(t *testing.T) {
		if got := Max(case1.args); !reflect.DeepEqual(got, case1.want) {
			t.Errorf("Max() = %v, want %v", got, case1)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Max(case2.args); !reflect.DeepEqual(got, case2.want) {
			t.Errorf("Max() = %v, want %v", got, case2)
		}
	})
}

func TestMin(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		args []T
		want T
	}

	case1 := testCase[int32]{"int32", []int32{1, 2, 3}, 1}
	case2 := testCase[float64]{"float64", []float64{1.2, 3.4, 5.6}, 1.2}

	t.Run(case1.name, func(t *testing.T) {
		if got := Min(case1.args); !reflect.DeepEqual(got, case1.want) {
			t.Errorf("Min() = %v, want %v", got, case1)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Min(case2.args); !reflect.DeepEqual(got, case2.want) {
			t.Errorf("Min() = %v, want %v", got, case2)
		}
	})
}

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		name   string
		args   []T
		target T
		want   bool
	}

	case1 := testCase[string]{"string", []string{"a", "b", "c"}, "b", true}
	case2 := testCase[int]{"int", []int{1, 2, 3}, 4, false}

	t.Run(case1.name, func(t *testing.T) {
		if got := Contains(case1.args, case1.target); got != case1.want {
			t.Errorf("Contains() = %v, want %v", got, case1)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Contains(case2.args, case2.target); got != case2.want {
			t.Errorf("Contains() = %v, want %v", got, case2)
		}
	})
}
