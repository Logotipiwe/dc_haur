package utils

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
