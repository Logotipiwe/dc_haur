package utils

import (
	"bytes"
	config "github.com/logotipiwe/dc_go_config_lib"
	"image"
	"image/png"
)

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

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func ToMap[K string, V any](arr []V, groupingFunc func(val V) K) map[K]V {
	result := make(map[K]V)
	for _, v := range arr {
		key := groupingFunc(v)
		result[key] = v
	}
	return result
}

func EncodeImageToBytes(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetOwnerChatID() string {
	return config.GetConfig("OWNER_TG_CHAT_ID")
}

func FindIndex(array []string, target string) int {
	for i, str := range array {
		if str == target {
			return i
		}
	}
	return -1
}
