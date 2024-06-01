package html

import (
	"fmt"
	"regexp"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"gitlab.com/rodrigoodhin/gocure/report/html/assets"
)

func (d *Data) minifyCSS() (minifiedCSS string, err error) {

	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	minifiedCSS, err = m.String("text/css", assets.CSS)
	if err != nil {
		return "", fmt.Errorf("error trying to minify CSS file %v", err)
	}

	minifiedCSS = "<style>" + minifiedCSS + "</style>"

	return
}

func (d *Data) minifyJS() (minifiedJS string, err error) {

	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	minifiedJS, err = m.String("text/javascript", assets.JS)
	if err != nil {
		return "", fmt.Errorf("error trying to minify JS file %v", err)
	}

	minifiedJS = "<script>" + minifiedJS + "</script>"

	return
}

func (d *Data) minifyChartJSLibrary() (minifiedJS string, err error) {

	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	minifiedJS, err = m.String("text/javascript", assets.ChartJSlibrary)
	if err != nil {
		return "", fmt.Errorf("error trying to minify chart JS library file %v", err)
	}

	minifiedJS = "<script>" + minifiedJS + "</script>"

	return
}

func (d *Data) minifyChartJS(content string) (minifiedJS string, err error) {

	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	minifiedJS, err = m.String("text/javascript", content)
	if err != nil {
		return "", fmt.Errorf("error trying to minify chart JS file %v", err)
	}

	minifiedJS = "<script>" + minifiedJS + "</script>"

	return
}

func (d *Data) minifyHTML(htmlContent string) (minifiedHMTL string, err error) {

	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	minifiedHMTL, err = m.String("text/html", htmlContent)
	if err != nil {
		return "", fmt.Errorf("error trying to minify HTML file %v", err)
	}

	return
}
