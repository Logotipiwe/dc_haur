package main

func ExistsInArr(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func GetValues[V any](inputMap map[string]V) []V {
	var values []V
	for _, value := range inputMap {
		values = append(values, value)
	}
	return values
}

func ChunkStrings(input []string, chunkSize int) [][]string {
	var result [][]string
	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}
