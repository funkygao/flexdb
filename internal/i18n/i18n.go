package i18n

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// L represents i18n bundle.
type L struct {
	bundle    map[string]Tr
	supported []Locale

	// defaultLocale is used when requested locale is not found
	defaultLocale string
}

// New creates the i18n bundle with specified translations dir.
func New(defaultLocale, path string) *L {
	files, _ := ioutil.ReadDir(path)
	l := &L{
		bundle:        make(map[string]Tr),
		defaultLocale: defaultLocale,
		supported:     make([]Locale, 0),
	}
	for _, f := range files {
		fileName := f.Name()
		b, err := ioutil.ReadFile(path + "/" + fileName)
		if err != nil {
			log.Printf("Cannot read file %s, file corrupt.", fileName)
			log.Printf("Error: %s", err)
			continue
		}
		t := Tr{
			translations: make(map[string]interface{}),
		}
		if err = json.Unmarshal(b, &t.translations); err != nil {
			panic(err)
		}

		locale := strings.Split(fileName, ".")[0]
		l.supported = append(l.supported, ParseLocale(locale))
		l.bundle[locale] = t
	}
	return l
}

func (l *L) exists(locale string) bool {
	_, exists := l.bundle[locale]
	return exists
}

// TranslationsForRequest will get the best matched Tr for given
// Request. If no Tr is found, returns default Tr
func (l *L) TranslationsForRequest(r *http.Request) Tr {
	locales := GetLocales(r)
	for _, locale := range locales {
		t, exists := l.bundle[locales[0].Name()]
		if exists {
			return t
		}
		for _, sup := range l.supported {
			if locale.Lang == sup.Lang {
				return l.bundle[sup.Name()]
			}
		}
	}

	// not found, use default
	return l.bundle[l.defaultLocale]
}

// TranslationsForLocale will get the Tr for specific locale.
// If no locale is found, returns default Tr
func (l *L) TranslationsForLocale(locale string) Tr {
	t, exists := l.bundle[locale]
	if exists {
		return t
	}

	return l.bundle[l.defaultLocale]
}
