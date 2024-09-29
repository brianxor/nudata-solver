package internal

import (
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

func GenerateUuid(isUpper bool) string {
	uuidV4 := uuid.NewString()

	if isUpper {
		return strings.ToUpper(uuidV4)
	}

	return uuidV4
}

func GenerateRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GenerateRandomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func Rot13(input string) string {
	result := make([]rune, len(input))

	for i, char := range input {
		switch {
		case char >= 'A' && char <= 'Z':
			result[i] = 'A' + (char-'A'+13)%26
		case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a'+13)%26
		default:
			result[i] = char
		}
	}

	return string(result)
}

func GetRandomItem[T any](items []T) T {
	randomIndex := rand.Intn(len(items))
	return items[randomIndex]
}
