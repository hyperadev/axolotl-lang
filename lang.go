package lang

import (
	"embed"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.toml
var localeFS embed.FS

type Data map[string]any

var (
	supportedLocales = []discordgo.Locale{
		discordgo.EnglishGB,
		discordgo.EnglishUS,
		discordgo.Japanese,
	}

	bundle  *i18n.Bundle
	locales = make(map[discordgo.Locale]*Localizer)
)

func LoadLocales() error {
	bundle = i18n.NewBundle(language.MustParse(string(supportedLocales[0])))
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, locale := range supportedLocales {
		_, err := bundle.LoadMessageFileFS(localeFS, fmt.Sprintf("locales/%s.toml", string(locale)))
		if err != nil {
			return fmt.Errorf("load locale file for %q: %w", locale, err)
		}
		locales[locale] = newLocalizer(locale)
	}

	return nil
}

func Bundle() *i18n.Bundle {
	return bundle
}

func GetLocalizer(locale discordgo.Locale) (*Localizer, bool) {
	if l, ok := locales[locale]; ok {
		return l, true
	}

	// Use default locale.
	return locales[supportedLocales[0]], false
}

func Localize(locale discordgo.Locale, id string) string {
	l, _ := GetLocalizer(locale)
	return l.Localize(id)
}

func LocalizeWithData(locale discordgo.Locale, id string, data Data) string {
	l, _ := GetLocalizer(locale)
	return l.LocalizeWithData(id, data)
}

func LocalizePlural(locale discordgo.Locale, id string, count any) string {
	l, _ := GetLocalizer(locale)
	return l.LocalizePlural(id, count)
}

func LocalizePluralWithData(locale discordgo.Locale, id string, count any, data Data) string {
	l, _ := GetLocalizer(locale)
	return l.LocalizePluralWithData(id, count, data)
}

type Localizer struct {
	*i18n.Localizer
}

func newLocalizer(locale discordgo.Locale) *Localizer {
	return &Localizer{Localizer: i18n.NewLocalizer(bundle, string(locale))}
}

func (l *Localizer) Localize(id string) string {
	s, err := l.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: id,
	})
	if s == "" && err != nil {
		return id
	}
	return s
}

func (l *Localizer) LocalizeWithData(id string, data Data) string {
	s, err := l.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if s == "" && err != nil {
		return id
	}
	return s
}

func (l *Localizer) LocalizePlural(id string, count any) string {
	s, err := l.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:   id,
		PluralCount: count,
	})
	if s == "" && err != nil {
		return id
	}
	return s
}

func (l *Localizer) LocalizePluralWithData(id string, count any, data Data) string {
	s, err := l.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
		PluralCount:  count,
	})
	if s == "" && err != nil {
		return id
	}
	return s
}
