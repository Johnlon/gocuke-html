package embedded

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.com/rodrigoodhin/gocure/helpers"
	"gitlab.com/rodrigoodhin/gocure/models"
)

func (d *Data) AddEmbeddedFileToStep() (err error) {

	features, err := helpers.ReadReport(false, d.InputJsonPath)
	if err != nil {
		return fmt.Errorf("error reading report")
	}

	if d.FeatureIndex > len(features)-1 || d.FeatureIndex < 0 {
		return fmt.Errorf("Feature index does not exist")
	} else if d.ScenarioIndex > len(features[d.FeatureIndex].Scenarios)-1 || d.ScenarioIndex < 0 {
		return fmt.Errorf("Scenario index does not exist")
	} else if d.StepIndex > len(features[d.FeatureIndex].Scenarios[d.ScenarioIndex].Steps)-1 || d.StepIndex < 0 {
		return fmt.Errorf("Scenario index does not exist")
	}

	d.Model = features[d.FeatureIndex].Scenarios[d.ScenarioIndex].Steps[d.StepIndex]
	newRepFile, err := d.AddEmbeddedFile()
	if err != nil {
		return fmt.Errorf("error adding files to step: %v", err)
	}

	features[d.FeatureIndex].Scenarios[d.ScenarioIndex].Steps[d.StepIndex] = newRepFile.(models.Step)

	newReport, err := json.Marshal(&features)
	if err != nil {
		return fmt.Errorf("error marshalling report: %v", err)
	}

	err = ioutil.WriteFile(d.OutputJsonPath, newReport, 0755)
	if err != nil {
		return fmt.Errorf("error writing new report file: %v", err)
	}

	return
}

func (d *Data) AddEmbeddedFileToScenario() (err error) {

	features, err := helpers.ReadReport(false, d.InputJsonPath)
	if err != nil {
		return fmt.Errorf("error reading report")
	}

	if d.FeatureIndex > len(features)-1 || d.FeatureIndex < 0 {
		return fmt.Errorf("Feature index does not exist")
	} else if d.ScenarioIndex > len(features[d.FeatureIndex].Scenarios)-1 || d.ScenarioIndex < 0 {
		return fmt.Errorf("Scenario index does not exist")
	}

	d.Model = features[d.FeatureIndex].Scenarios[d.ScenarioIndex]
	newRepFile, err := d.AddEmbeddedFile()
	if err != nil {
		return fmt.Errorf("error adding files to scenario: %v", err)
	}

	features[d.FeatureIndex].Scenarios[d.ScenarioIndex] = newRepFile.(models.Scenario)

	newReport, err := json.Marshal(&features)
	if err != nil {
		return fmt.Errorf("error marshalling report: %v", err)
	}

	err = ioutil.WriteFile(d.OutputJsonPath, newReport, 0755)
	if err != nil {
		return fmt.Errorf("error writing new report file: %v", err)
	}

	return
}

func (d *Data) AddEmbeddedFileToFeature() (err error) {

	features, err := helpers.ReadReport(false, d.InputJsonPath)
	if err != nil {
		return fmt.Errorf("error reading report")
	}

	if d.FeatureIndex > len(features)-1 || d.FeatureIndex < 0 {
		return fmt.Errorf("Feature index does not exist")
	}

	d.Model = features[d.FeatureIndex]
	newRepFile, err := d.AddEmbeddedFile()
	if err != nil {
		return fmt.Errorf("error adding files to feature: %v", err)
	}

	features[d.FeatureIndex] = newRepFile.(models.Feature)

	newReport, err := json.Marshal(&features)
	if err != nil {
		return fmt.Errorf("error marshalling report: %v", err)
	}

	err = ioutil.WriteFile(d.OutputJsonPath, newReport, 0755)
	if err != nil {
		return fmt.Errorf("error writing new report file: %v", err)
	}

	return
}

func (d *Data) AddEmbeddedFile() (embeddedModel interface{}, err error) {
	for _, f := range d.Files {
		if strings.HasPrefix(f, "http://") || strings.HasPrefix(f, "https://") {
			d.Model, err = d.addRemote(d.Model, f)
			if err != nil {
				return nil, fmt.Errorf("error adding remote file: %v", err)
			}
		} else {
			d.Model, err = d.addLocal(d.Model, f)
			if err != nil {
				return nil, fmt.Errorf("error adding local file: %v", err)
			}
		}
	}

	embeddedModel = d.Model

	return
}

func (d *Data) addLocal(model interface{}, filePath string) (embeddedModel interface{}, err error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading local file: %v", err)
	}

	return d.addFile(model, fileBytes)
}

func (d *Data) addRemote(model interface{}, fileUrl string) (embeddedModel interface{}, err error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting remotte file: %v", err)
	}

	defer resp.Body.Close()

	fileBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading remote file: %v", err)
	}

	return d.addFile(model, fileBytes)
}

func (d *Data) addFile(model interface{}, fileBytes []byte) (embeddedModel interface{}, err error) {
	fileData := helpers.ToBase64(fileBytes)
	fileType := http.DetectContentType(fileBytes)
	embed := models.Embed{
		Data: fileData,
		Media: models.Media{
			Type: fileType,
		},
	}

	switch m := model.(type) {
	case models.Feature:
		m.Embeddings = append(m.Embeddings, embed)
		embeddedModel = m
	case models.Scenario:
		m.Embeddings = append(m.Embeddings, embed)
		embeddedModel = m
	case models.Step:
		m.Embeddings = append(m.Embeddings, embed)
		embeddedModel = m
	default:
		return nil, fmt.Errorf("Unrecognized model")
	}

	return
}
