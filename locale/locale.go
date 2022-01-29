// Package locale contains definition of commonly used locales.
package locale

import (
	"strings"
	"sync"
	"time"
)

// Formatter is a function for formatting an output string for given time duration
type Formatter func(d time.Duration) string

// Formatters is a collection of formatter functions mapped to the closest time duration
type Formatters map[time.Duration]Formatter

// Locale represents an IETF BCP 47 formatted language tag.
type Locale string

// Split splits the locale into language and territory components.
// adapted from: https://github.com/jeandeaual/go-locale/blob/23669fb7cbc8f5714c70bc22594b8c0ff194ce5d/util.go#L13-L28
func (locale Locale) Split() (string, string) {
	// Remove the encoding, if present
	formattedLocale := strings.Split(string(locale), ".")[0]
	// Normalize by replacing the hyphens with underscores
	formattedLocale = strings.Replace(formattedLocale, "-", "_", -1)

	// Split at the underscore
	split := strings.Split(formattedLocale, "_")
	language := split[0]
	territory := ""
	if len(split) > 1 {
		territory = split[1]
	}

	return language, territory
}

// collection of registered range formatters keyed by locale
var (
	locales = make(map[string]Formatters)
	lock    = sync.RWMutex{}
)

// Register registers a new set of formatters for the given locale
func Register(locale Locale, formatters Formatters) {
	defer lock.Unlock()
	lock.Lock()
	locales[string(locale)] = formatters
}

// Lookup performs a lookup on the shared locales table and returns formatter
// for the given locale, else nil of none is found.
func Lookup(locale Locale) Formatters {
	defer lock.RUnlock()
	lock.RLock()

	formatters := locales[string(locale)]
	if formatters == nil {
		// try looking up a formatter defined at language level without territory classifier
		lang, _ := locale.Split()
		formatters = locales[lang]
	}

	return formatters
}
