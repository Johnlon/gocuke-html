package html

import (
	"fmt"
	"html/template"
	"strings"

	"gitlab.com/rodrigoodhin/gocure/helpers"
	"gitlab.com/rodrigoodhin/gocure/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (d *Data) stepsResult() (passed, failed, undefined, skipped int) {

	for _, f := range d.Features {
		passed += 1
		for _, s := range f.Scenarios {
			if s.Type == "scenario" {
				for _, step := range s.Steps {
					switch step.Result.Status {
					case "passed":
						passed += 1
					case "failed":
						failed += 1
					case "skipped":
						skipped += 1
					case "undefined":
						undefined += 1
					}
				}
			}
		}
	}

	return
}

func (d *Data) stepsResultByScenario(scenario models.Scenario) (passed, failed, undefined, skipped int) {

	for _, s := range scenario.Steps {
		switch s.Result.Status {
		case "passed":
			passed += 1
		case "failed":
			failed += 1
		case "skipped":
			skipped += 1
		case "undefined":
			undefined += 1
		}
	}

	return
}

func (d *Data) printSteps(scenario models.Scenario) (steps string, duration int, err error) {

	for _, s := range scenario.Steps {

		modalID := d.randomModalID("step")

		duration += s.Result.Duration

		resultIcon := ""
		resultColor := ""
		switch s.Result.Status {
		case "passed":
			resultIcon = "done"
			resultColor = "#2E7D31"
		case "failed":
			resultIcon = "clear"
			resultColor = "#C62828"
		case "skipped":
			resultIcon = "remove"
			resultColor = "#ff6f00"
		case "undefined":
			resultIcon = "bolt"
			resultColor = "#01579b"
		}

		message := ""
		if s.Result.Status == "failed" {
			message = "<div class=\"card-step-message\">" + strings.ReplaceAll(s.Result.ErrorMessage, "\n", "<br>") + "</div>"
		}

		embedded := ""
		if d.ShowEmbeddedFiles {
			embedded, err = d.createEmbedModals(modalID, s.Keyword, s.Name, s.Result.Duration, s.Embeddings)
			if err != nil {
				return "", 0, fmt.Errorf("error creating embed models")
			}
		}

		table := ""
		if len(s.Rows) > 0 {
			table = "<div class=\"card-step-table\">"
			table += "<table>"
			for _, row := range s.Rows {
				table += "<tr>"
				for _, cell := range row.Cells {
					table += "<td>| " + cell + " </td>"
				}
				table += "<td>|</td>"
				table += "</tr>"
			}
			table += "</table></div>"
		}

		docString := ""
		if len(s.DocString.Value) > 0 {
			content, err := d.prettyString(s.DocString.Value)
			if err != nil {
				content = s.DocString.Value
			}
			
			docString = "<div class=\"card-step-table\">"
			docString += "<table>"
			docString += "<tr>"
			docString += "<td>\"\"\"</td>"
			docString += "</tr>"
			docString += "<tr>"
			docString += "<td><pre>" + content + "</pre></td>"
			docString += "</tr>"
			docString += "<tr>"
			docString += "<td>\"\"\"</td>"
			docString += "</tr>"
			docString += "</table></div>"
		}

		c := cases.Title(language.AmericanEnglish)
		stepHTML := StepHTML{
			ModalID:          modalID,
			Status:           s.Result.Status,
			StatusCapitalize: c.String(s.Result.Status),
			ResultIcon:       template.HTML("<span class='material-icons'>" + resultIcon + "</span>"),
			Keyword:          template.HTML("<span style='color: " + resultColor + ";'>" + s.Keyword + "</span>"),
			Name:             s.Name,
			Message:          template.HTML(message),
			Embed:            template.HTML(embedded),
			Table:            template.HTML(table),
			DocString:        template.HTML(docString),
			Duration:         template.HTML(helpers.HumanizeExecution(s.Result.Duration)),
		}

		content, err := d.tmplParse("tmpl/step.html", &stepHTML)
		if err != nil {
			return "", 0, fmt.Errorf("error trying to parse step file : %v", err)
		}

		steps += content

	}

	return
}
