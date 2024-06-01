package html

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"gitlab.com/rodrigoodhin/gocure/helpers"
	"gitlab.com/rodrigoodhin/gocure/models"
)

func (d *Data) featuresResult() (passed, failed int) {

	for _, feature := range d.Features {
		passed += 1
	out:
		for _, element := range feature.Scenarios {
			if element.Type == "scenario" {
				for _, step := range element.Steps {
					if step.Result.Status != "passed" {
						passed -= 1
						failed += 1
						break out
					}
				}
			}
		}
	}
	return
}

func (d *Data) featureTags(feature models.Feature) (tags string) {

	for _, tag := range feature.Tags {
		if tags != "" {
			tags += " "
		}
		tags += tag.Name
	}
	return
}

func (d *Data) printFeatures() (features string, err error) {

	for i, f := range d.Features {

		modalID := d.randomModalID("feature")

		scenarioPrint, scenarioExecTime, err := d.printScenarios(f, i)
		if err != nil {
			return "", fmt.Errorf("error trying to print scenarios : %v", err)
		}

		scenarioPassed, scenarioFailed := d.scenarioResultByFeature(f)

		result := "passed"
		if scenarioFailed > 0 {
			result = "failed"
		}

		embedded := ""
		if d.ShowEmbeddedFiles {
			embedded, err = d.createEmbedModals(modalID, f.Keyword, f.Name, scenarioExecTime, f.Embeddings)
			if err != nil {
				return "", fmt.Errorf("error reading report")
			}
		}

		featureHTML := FeatureHTML{
			Id:          "feature-" + strconv.Itoa(i),
			Tags:        d.featureTags(f),
			Name:        f.Name,
			Description: template.HTML(strings.ReplaceAll(f.Description, "\n", "<br>")),
			Failed:      scenarioFailed,
			Passed:      scenarioPassed,
			Scenarios:   template.HTML(scenarioPrint),
			Embed:       template.HTML(embedded),
			Duration:    template.HTML(helpers.HumanizeExecution(scenarioExecTime)),
			Result:      result,
		}

		content, err := d.tmplParse("tmpl/feature.html", &featureHTML)
		if err != nil {
			return "", fmt.Errorf("error trying to parse feature file : %v", err)
		}

		features += content

	}

	return
}
