package api

type Locale string

func (l Locale) String() string {
	return string(l)
}

const (
	LocaleEnglish         Locale = "en"
	LocaleEnglishCanadian Locale = "en_CA"
	LocaleRussian         Locale = "ru"
	LocalePortuguese      Locale = "pt"
	LocaleBrazilian       Locale = "pt_BR"
	LocaleChinese         Locale = "zh"
	LocaleFrench          Locale = "fr"
	LocaleFrenchCanadian  Locale = "fr_CA"
	LocaleGerman          Locale = "de"
	LocaleJapanese        Locale = "ja"
	LocalePolish          Locale = "pl"
	LocaleRomanian        Locale = "ro"
	LocaleSpanish         Locale = "es"
	LocaleUkrainian       Locale = "uk"
	LocaleHungarian       Locale = "hu"
	LocalePhilippines     Locale = "tl"
	LocaleVietnam         Locale = "vi"
	LocaleThailand        Locale = "th"
	LocaleHindi           Locale = "hi"
)
