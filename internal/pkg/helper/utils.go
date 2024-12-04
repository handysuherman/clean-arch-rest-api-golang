package helper

import (
	"regexp"
	"strings"
)

// RedisPrefixes generates a prefixed key for Redis based on the input parameters.
// The function constructs the prefixed key by concatenating the prefixKey, expectedPrefixKey (if provided).
// and the main key, seperated by appropiate delimiters. The constructed key is then returned.
// If expectedPrefixKey is provided, the redisId is added as a prefix before it.
func RedisPrefixes(key, prefixKey, expectedPrefixKey, redisId string) string {
	prefix := prefixKey + ":" + key

	if expectedPrefixKey != "" {
		prefix = redisId + "_" + expectedPrefixKey + ":" + key
	}

	return prefix
}

func Slugify(text string) string {
	reSpace := regexp.MustCompile(`\s+`)
	text = reSpace.ReplaceAllString(text, "-")

	nonAlphaNum := regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	text = nonAlphaNum.ReplaceAllString(text, "-")

	consecutiveDash := regexp.MustCompile(`-{2,}`)
	text = consecutiveDash.ReplaceAllString(text, "-")

	text = strings.Trim(text, "-")
	// text = strings.ToLower(text)

	return text
}
