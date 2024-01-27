package pkg

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestChunkStrings_SmallChunk(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	chunkSize := 2
	expectedResult := [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	result := ChunkStrings(input, chunkSize)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestChunkStrings_EqualChunk(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	chunkSize := 6
	expectedResult := [][]string{{"a", "b", "c", "d", "e", "f"}}
	result := ChunkStrings(input, chunkSize)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestChunkStrings_LargeChunk(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	chunkSize := 8
	expectedResult := [][]string{{"a", "b", "c", "d", "e", "f"}}
	result := ChunkStrings(input, chunkSize)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestChunkStrings_EmptyInput(t *testing.T) {
	input := []string{}
	chunkSize := 3
	expectedResult := [][]string{}
	result := ChunkStrings(input, chunkSize)
	if !assert.Equal(t, 0, len(result)) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name   string
		array  []string
		target string
		want   int
	}{
		{
			name:   "TargetPresent",
			array:  []string{"apple", "banana", "cherry"},
			target: "banana",
			want:   1,
		},
		{
			name:   "TargetLast",
			array:  []string{"apple", "banana", "cherry"},
			target: "cherry",
			want:   2,
		},
		{
			name:   "TargetAbsent",
			array:  []string{"apple", "banana", "cherry"},
			target: "mango",
			want:   -1,
		},
		{
			name:   "TargetInEmptyArray",
			array:  []string{},
			target: "mango",
			want:   -1,
		},
		{
			name:   "TargetInSingleElementArray",
			array:  []string{"apple"},
			target: "apple",
			want:   0,
		},
		{
			name:   "TargetAbsentInSingleElementArray",
			array:  []string{"apple"},
			target: "mango",
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndex(tt.array, tt.target); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
