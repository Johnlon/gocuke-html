package html

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
	"time"

	"gitlab.com/rodrigoodhin/gocure/helpers"
	"gitlab.com/rodrigoodhin/gocure/models"
)

//go:embed tmpl/*
var htmlTmpl embed.FS

func (d *Data) CreateReport() (err error) {

	d.validateDefaultValues()

	if d.InputJsonPath != "" {
		d.JSONContent, err = ioutil.ReadFile(d.InputJsonPath)
		if err != nil {
			return fmt.Errorf("error reading json file : %v", err)
		}

		d.Features, err = helpers.ReportParse(d.IgnoreBadJsonFiles, d.JSONContent)
		if err != nil {
			return fmt.Errorf("error parsing json file : %v", err)
		}

		report, err := d.generateHTML()
		if err != nil {
			return fmt.Errorf("error creating report : %v", err)
		}

		d.HTMLContent = report
		err = d.writeReport()
		if err != nil {
			return fmt.Errorf("error writing report : %v", err)
		}

	} else if d.InputFolderPath != "" {
		var fullOutputJsonReport []models.Feature
		fileList, err := helpers.ReadFolders(d.InputFolderPath)
		if err != nil {
			return fmt.Errorf("error reading folders")
		}

		if d.MergeFiles {
			for _, file := range fileList {

				if strings.Contains(file, ".json") {
					reportFileContent, err := ioutil.ReadFile(file)
					if err != nil {
						return fmt.Errorf("error reading json (%v) file : %v", file, err)
					}

					var outputJsonReport []models.Feature
					err = json.Unmarshal(reportFileContent, &outputJsonReport)
					if err != nil && !d.IgnoreBadJsonFiles {
						return fmt.Errorf("error to unmarshall json (%v) file : %v", file, err)
					} else if err != nil && d.IgnoreBadJsonFiles {
						err = nil
					} else {
						fullOutputJsonReport = append(fullOutputJsonReport, outputJsonReport...)
					}
				}
			}

			d.JSONContent, err = json.MarshalIndent(fullOutputJsonReport, "", " ")
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			d.Features, err = helpers.ReportParse(d.IgnoreBadJsonFiles, d.JSONContent)
			if err != nil {
				return fmt.Errorf("error parsing json file : %v", err)
			}

			report, err := d.generateHTML()
			if err != nil {
				return fmt.Errorf("error creating report : %v", err)
			}

			d.HTMLContent = report
			err = d.writeReport()
			if err != nil {
				return fmt.Errorf("error writing report : %v", err)
			}
		} else {
			for _, file := range fileList {

				if strings.Contains(file, ".json") {
					d.Timer = 0

					d.JSONContent, err = ioutil.ReadFile(file)
					if err != nil {
						return fmt.Errorf("error reading json (%v) file : %v", file, err)
					}

					d.Features, err = helpers.ReportParse(d.IgnoreBadJsonFiles, d.JSONContent)
					if err != nil && !d.IgnoreBadJsonFiles {
						return fmt.Errorf("error parsing json (%v) file : %v", file, err)
					} else if err != nil && d.IgnoreBadJsonFiles {
						err = nil
					} else {

						d.HTMLContent, err = d.generateHTML()
						if err != nil && !d.IgnoreBadJsonFiles {
							return fmt.Errorf("error creating report : %v", err)
						}

						err = d.writeReport()
						if err != nil && !d.IgnoreBadJsonFiles {
							return fmt.Errorf("error writing report : %v", err)
						}

					}
				}
			}
		}
	}

	return
}

func (d *Data) generateHTML() (htmlReport string, err error) {
	htmlReport += ""

	content, err := d.printFeatures()
	if err != nil {
		return "", fmt.Errorf("error trying to print feature file : %v", err)
	}

	chartPrint := ""
	chartJsPrint, err := d.printChartJs()
	if err != nil {
		return "", fmt.Errorf("error trying to print chartJs file : %v", err)
	}

	metadataInfo, err := d.printMetadata()
	if err != nil {
		return "", fmt.Errorf("error trying to print metadata file : %v", err)
	}

	for _, chart := range d.getChartsName() {

		chartBlock, err := d.printChart(chart)

		if err != nil {
			return "", fmt.Errorf("error trying to print chart file : %v", err)
		}

		chartPrint += chartBlock
	}

	featurePassed, featureFailed := d.featuresResult()

	minifiedCSS, err := d.minifyCSS()
	if err != nil {
		return "", fmt.Errorf("error on call minify CSS function %v", err)
	}

	minifiedJS, err := d.minifyJS()
	if err != nil {
		return "", fmt.Errorf("error on call minify JS function %v", err)
	}

	minifiedChartJSLibrary, err := d.minifyChartJSLibrary()
	if err != nil {
		return "", fmt.Errorf("error on call minify chart JS library function %v", err)
	}

	d.GeneratedAt = time.Now().Format("Jan 02, 2006 - 15:04:05")

	baseHTML := BaseHTML{
		Title:          d.Title,
		Timer:          template.HTML(helpers.HumanizeExecution(d.Timer)),
		Chart:          template.HTML(chartPrint),
		ChartJsLibrary: template.HTML(minifiedChartJSLibrary),
		ChartJs:        template.HTML(chartJsPrint),
		Content:        template.HTML(content),
		CountAll:       (featurePassed + featureFailed),
		CountPassed:    featurePassed,
		CountFailed:    featureFailed,
		CSS:            template.HTML(minifiedCSS),
		JS:             template.HTML(minifiedJS),
		Metadata:       template.HTML(metadataInfo),
		GeneratedAt:    d.GeneratedAt,
	}

	tmpl, err := template.ParseFS(htmlTmpl, "tmpl/base.html")
	if err != nil {
		return "", fmt.Errorf("error trying to load html embedded : %v", err)
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, &baseHTML); err != nil {
		return "", fmt.Errorf("error trying to parse base file : %v", err)
	}

	htmlReport, err = d.minifyHTML(buf.String())
	if err != nil {
		return "", fmt.Errorf("error on call minify HTML function %v", err)
	}

	return
}
