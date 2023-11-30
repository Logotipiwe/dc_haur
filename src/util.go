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
