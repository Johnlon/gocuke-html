package html

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"gitlab.com/rodrigoodhin/gocure/helpers"
	"gitlab.com/rodrigoodhin/gocure/models"
)

func (d *Data) scenariosResult() (passed, failed int) {

	for _, f := range d.Features {
		for _, s := range f.Scenarios {
			if s.Type == "scenario" {
				passed += 1
				for _, s := range s.Steps {
					if s.Result.Status != "passed" {
						passed -= 1
						failed += 1
						break
					}
				}
			}
		}
	}
	return
}

func (d *Data) scenarioResultByFeature(feature models.Feature) (passed, failed int) {

	for _, s := range feature.Scenarios {
		if s.Type == "scenario" {
			passed += 1
			for _, s := range s.Steps {
				if s.Result.Status != "passed" {
					passed -= 1
					failed += 1
					break
				}
			}
		}
	}
	return
}

func (d *Data) scenarioTags(scenario models.Scenario, featureLine int) (tags string) {

	for _, t := range scenario.Tags {
		if t.Line > featureLine {
			if tags != "" {
				tags += " "
			}
			tags += t.Name
		}
	}
	return
}

func (d *Data) printScenarios(feature models.Feature, featureIndex int) (scenarios string, duration int, err error) {

	for j, e := range feature.Scenarios {

		if e.Type == "scenario" {

			modalID := d.randomModalID("scenario")

			stepsPrint, stepsExecTime, err := d.printSteps(e)
			if err != nil {
				return "", 0, fmt.Errorf("error trying to print steps : %v", err)
			}

			d.Timer = d.Timer + stepsExecTime

			duration += stepsExecTime

			stepPassed, stepFailed, stepPending, stepSkipped := d.stepsResultByScenario(e)

			embedded := ""
			if d.ShowEmbeddedFiles {
				embedded, err = d.createEmbedModals(modalID, e.Keyword, e.Name, stepsExecTime, e.Embeddings)
				if err != nil {
					return "", 0, fmt.Errorf("error creating embed models")
				}
			}

			scenarioHTML := ScenarioHTML{
				ModalID:     modalID,
				Id:          "scenario-" + strconv.Itoa(featureIndex) + "-" + strconv.Itoa(j),
				Tags:        d.scenarioTags(e, feature.Line),
				Keyword:     e.Keyword,
				Name:        e.Name,
				Description: template.HTML(strings.ReplaceAll(e.Description, "\n", "<br>")),
				Skipped:     stepSkipped,
				Undefined:   stepPending,
				Failed:      stepFailed,
				Passed:      stepPassed,
				Steps:       template.HTML(stepsPrint),
				Embed:       template.HTML(embedded),
				Duration:    template.HTML(helpers.HumanizeExecution(stepsExecTime)),
			}

			content, err := d.tmplParse("tmpl/scenario.html", &scenarioHTML)
			if err != nil {
				return "", 0, fmt.Errorf("error trying to parse scenario file : %v", err)
			}

			scenarios += content

		}
	}

	return
}
