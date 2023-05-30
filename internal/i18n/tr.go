package i18n

import (
	"strconv"
	"strings"
)

// Tr represents translationsations map for specific locale.
type Tr struct {
	translations map[string]interface{}
}

// Value traverses the translationsations map and finds translationsation for
// given key. If no translationsation is found, returns value of given key.
func (t Tr) Value(key string, args ...string) string {
	if t.exists(key) {
		res, ok := t.translations[key].(string)
		if ok {
			return t.parseArgs(res, args)
		}
	}

	// nested key: recursive lookup
	tuple := strings.Split(key, ".")
	for i := range tuple {
		k1 := strings.Join(tuple[0:i], ".")
		k2 := strings.Join(tuple[i:], ".")
		if t.exists(k1) {
			newt := &Tr{
				translations: t.translations[k1].(map[string]interface{}),
			}
			return newt.Value(k2, args...)
		}
	}

	// unable to translationsate, return the key itself
	return key
}

// parseArgs replaces the argument placeholders with given arguments
func (t Tr) parseArgs(value string, args []string) string {
	res := value
	for i := 0; i < len(args); i++ {
		tok := "{" + strconv.Itoa(i) + "}"
		res = strings.Replace(res, tok, args[i], -1)
	}
	return res
}

// exists checks if value exists for given key
func (t Tr) exists(key string) bool {
	_, exists := t.translations[key]
	return exists
}
