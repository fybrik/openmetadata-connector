package utils

import (
	"math/rand"
	"strings"

	client "github.com/fybrik/datacatalog-go-client"
	models "github.com/fybrik/datacatalog-go-models"
)

func AppendStrings(a string, b string) string {
	if strings.Contains(b, ".") {
		return a + ".\"" + b + "\""
	} else {
		return a + "." + b
	}
}

// If the tag is in the Fybrik tag category (e.g. "Fybrik.PII"), remove the
// "Fybrik." prefix. Otherwise, do nothing
func StripTag(tag string) string {
	return strings.TrimPrefix(tag, "Fybrik.")
}

func UpdateCustomProperty(customProperties map[string]interface{}, orig map[string]interface{}, key string, value *string) {
	// update customProperties only if there is a new value
	if value != nil && *value != "" {
		customProperties[key] = *value
		return
	}

	// if there is no new value, revert to original value
	if v, ok := orig[key]; ok && v != "" {
		customProperties[key] = v
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// create random string of letters of length n
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ExtractColumns(resourceColumns []models.ResourceColumn) []client.Column {
	var ret []client.Column
	for _, rc := range resourceColumns {
		column := *client.NewColumn("STRING", rc.Name)
		column.SetDataLength(0)
		ret = append(ret, column)
	}
	return ret
}
