package i18n

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
)

var (
	// B is the singleton i18n bundle.
	B *L

	// Tr is the translation map with the same locale as B's default locale.
	tr Tr
)

// Offer provides a singleton L instance.
func Offer(defaultLocale string) *L {
	translations := AssetNames()
	l := &L{
		bundle:        make(map[string]Tr, len(translations)),
		supported:     make([]Locale, 0, len(translations)),
		defaultLocale: defaultLocale,
	}
	for _, fileName := range translations {
		if !strings.HasSuffix(fileName, ".json") {
			continue
		}

		b, err := Asset(fileName)
		if err != nil {
			panic(err)
		}

		t := Tr{
			translations: make(map[string]interface{}),
		}
		if err = json.Unmarshal(b, &t.translations); err != nil {
			panic(fmt.Errorf("%v %s", err, string(b)))
		}

		locale := strings.Split(path.Base(fileName), ".")[0]
		l.supported = append(l.supported, ParseLocale(locale))
		l.bundle[locale] = t
	}

	tr = l.TranslationsForLocale(defaultLocale)

	return l
}

// T translates a msg by key with default locale.
func T(key string, args ...string) string {
	return tr.Value(key, args...)
}
