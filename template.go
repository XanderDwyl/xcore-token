package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	//"github.com/a-fis/urhodling.com/helpers"
	//"github.com/a-fis/urhodling.com/libs/currency"
	//"github.com/a-fis/urhodling.com/libs/translate"
	"github.com/labstack/echo"
)

/*
view files should be below
- app/views/layouts/base.html
- app/views/layouts/amp.html
- app/views/layouts/admin.html
- app/views/partials/XXXX.html
- app/views/XXX/YYY.html
*/

// Template ...
type Template struct {
	templates map[string]*template.Template
}

// Add ...
func (t Template) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	t.templates[name] = tmpl
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if _, ok := t.templates[name]; ok == false {
		// not such view
		return fmt.Errorf("no such view. (%s)", name)
	}
	return t.templates[name].Execute(w, data)
}

// NewTemplate creates a new template
func NewTemplate(templatesDir string) *Template {
	ext := ".html"
	ins := Template{
		templates: map[string]*template.Template{},
	}

	layout := templatesDir + "layouts/base" + ext
	adminLayout := templatesDir + "layouts/admin" + ext
	ampLayout := templatesDir + "layouts/amp" + ext

	_, err := os.Stat(layout)
	if err != nil {
		log.Panicf("cannot find %s", layout)
		os.Exit(1)
	}
	_, err = os.Stat(adminLayout)
	if err != nil {
		log.Printf("cannot find %s", adminLayout)
		os.Exit(1)
	}
	_, err = os.Stat(ampLayout)
	if err != nil {
		log.Printf("cannot find %s", ampLayout)
		os.Exit(1)
	}

	partials, err := filepath.Glob(templatesDir + "partials/" + "*" + ext)
	if err != nil {
		log.Print("cannot find " + templatesDir + "partials/" + "*" + ext)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML {
			return template.HTML(text)
		},
		//"translate": func(str, locale string) template.HTML {
		//	return template.HTML(translate.Translate(str, locale))
		//},
		//"humanize": func(qty decimal.Decimal) template.HTML {
		//	formattedNumber := helpers.FormatNumber(qty.String())
		//
		//	return template.HTML(formattedNumber)
		//},
		//"convert": func(qty decimal.Decimal, from, to string) template.HTML {
		//	formattedNumber := helpers.FormatNumber(currency.Convert(qty, from, to).StringFixed(2))
		//
		//	return template.HTML(formattedNumber)
		//},
		//"percent": func(a, b decimal.Decimal) template.HTML {
		//	z := decimal.NewFromFloat(0)
		//	if a.Equal(z) || b.Equal(z) {
		//		return template.HTML(z.String())
		//	}
		//
		//	return template.HTML(a.Div(b).Mul(decimal.NewFromFloat(100)).StringFixed(2))
		//},
		"unixToDate": func(value int64) string {
			tm := time.Unix(value, 0)

			return fmt.Sprintf("%s", tm.Format("2006-01-02"))
		},
		"unixToDatetime": func(value int64) string {
			tm := time.Unix(value, 0)

			return fmt.Sprintf("%s", tm.Format("2006-01-02-03:04:05"))
		},
		//"displayTimezone": func(timezone string) string {
		//	return helpers.DisplayTimezone(timezone)
		//},
		//"spaceToUnderscore": func(word string) string {
		//	return strings.Replace(word, " ", "_", -1)
		//},
		//"isExchangeActive": func(isActive bool) string {
		//	if isActive == true {
		//		return "check_circle"
		//	}
		//
		//	return "not_interested"
		//},
		//"textSuccessFailed": func(isActive bool) string {
		//	if isActive == true {
		//		return "green-text"
		//	}
		//
		//	return "red-text"
		//},
		//"sanitizeLocale": func(locale string) string {
		//	properLocale := map[string]string{
		//		"en": "English",
		//		"ja": "Japanese",
		//	}
		//	return properLocale[locale]
		//},
		//"displayCoinIcon": func(coinSymbol string) string {
		//	imagePath := "app/data/assets/img/svg/color/" + strings.ToLower(coinSymbol) + ".svg"
		//	coinIconPath := "/assets/img/svg/color/" + strings.ToLower(coinSymbol) + ".svg"
		//	exists, _ := helpers.FileExists(imagePath)
		//	if !exists {
		//		coinIconPath = "/assets/img/default-coin.png"
		//	}
		//	return coinIconPath
		//},
	}

	views, _ := filepath.Glob(templatesDir + "**/*" + ext)
	for _, view := range views {
		dir, file := filepath.Split(view)
		dir = strings.Replace(dir, templatesDir, "", 1)
		file = strings.TrimSuffix(file, ext)
		renderName := dir + file

		if strings.Contains(renderName, "_amp") {
			tmplfiles := append([]string{ampLayout, view}, partials...)
			tmpl := template.Must(template.New(filepath.Base(ampLayout)).Funcs(funcMap).ParseFiles(tmplfiles...))
			ins.Add(renderName, tmpl)
			log.Printf("renderName: %s, layout: %s", renderName, ampLayout)
		} else if strings.Contains(renderName, "admin") {
			tmplfiles := append([]string{adminLayout, view}, partials...)
			tmpl := template.Must(template.New(filepath.Base(adminLayout)).Funcs(funcMap).ParseFiles(tmplfiles...))
			ins.Add(renderName, tmpl)
			log.Printf("renderName: %s, layout: %s", renderName, adminLayout)
		} else {
			tmplfiles := append([]string{layout, view}, partials...)
			tmpl := template.Must(template.New(filepath.Base(layout)).Funcs(funcMap).ParseFiles(tmplfiles...))
			ins.Add(renderName, tmpl)
			log.Printf("renderName: %s, layout: %s", renderName, layout)
		}
	}
	return &ins
}
