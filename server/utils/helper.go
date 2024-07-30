package utils

import (
	"unicode"

	"cloud.google.com/go/firestore"
)

func CapitalizeKey(key string) string {
	if len(key) == 0 {
		return ""
	}
	runes := []rune(key)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func FirestoreUpdate(data map[string]interface{}) []firestore.Update {
	updates := []firestore.Update{}
	for key, value := range data {
		updates = append(updates, firestore.Update{
			Path:  CapitalizeKey(key),
			Value: value,
		})
	}
	return updates
}
