package main

import (
	"slices"
	"testing"
)

// Функция, которую мы тестируем.
func findMaxUniqueInK(K, L []int) int {
	maxVal := -1
	for _, val := range K {
		if !slices.Contains(L, val) && val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func TestFindMaxUniqueInK(t *testing.T) {
	tests := []struct {
		name     string
		K        []int
		L        []int
		expected int
	}{
		{
			name:     "Нет пересечений",
			K:        []int{1, 2, 3, 4, 5},
			L:        []int{6, 7, 8},
			expected: 5,
		},
		{
			name:     "Есть пересечения",
			K:        []int{10, 20, 30},
			L:        []int{20, 30, 40},
			expected: 10,
		},
		{
			name:     "Все элементы K есть в L",
			K:        []int{1, 2, 3},
			L:        []int{1, 2, 3, 4},
			expected: -1,
		},
		{
			name:     "Отрицательные числа",
			K:        []int{-1, -2, -3},
			L:        []int{-2, -3},
			expected: -1,
		},
		{
			name:     "Пустой массив K",
			K:        []int{},
			L:        []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "Пустой массив L",
			K:        []int{1, 2, 3},
			L:        []int{},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxUniqueInK(tt.K, tt.L)
			if result != tt.expected {
				t.Errorf("findMaxUniqueInK() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
