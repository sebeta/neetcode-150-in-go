package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"nil slice", nil, false},
		{"empty slice", []int{}, false},
		{"single element", []int{1}, false},
		{"duplicate present", []int{1, 2, 3, 1}, true},
		{"all distinct", []int{1, 2, 3, 4}, false},
		{"many duplicates", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"adjacent duplicates", []int{5, 5}, true},
		{"negative numbers duplicate", []int{-1, -2, -3, -1}, true},
		{"negative numbers distinct", []int{-1, -2, -3}, false},
		{"zero duplicate", []int{0, 1, 0}, true},
		{"values equal to indices but distinct", []int{0, 1, 2, 3}, false},
		{"large values distinct", []int{1000000000, -1000000000, 7}, false},
		{"large values duplicate", []int{1000000000, 2, 1000000000}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.nums); got != tt.want {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
