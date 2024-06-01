package main

import (
	"fmt"

	"gitlab.com/rodrigoodhin/gocure/embedded"
	"gitlab.com/rodrigoodhin/gocure/models"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
	"gitlab.com/rodrigoodhin/gocure/report/html"
)

func main() {

	// Set embedded options
	emb := gocure.Embedded{
		Config: embedded.Data{
			InputJsonPath:  "./example/newreport.json",
			OutputJsonPath: "./example/newReport.json",
			Files: []string{
				"./example/OpenGocureRepository.mp4",
				"./example/empty.txt",
				"./example/text.txt",
				"https://i.ibb.co/LpRkTqf/gocure-small.png",
			},
			FeatureIndex: 0,
			ScenarioIndex: 0,
			StepIndex: 2,
		},
	}

	err := emb.AddToFeature()
	if err != nil {
		fmt.Printf("error adding files to feature: %v", err)
	}

	err = emb.AddToScenario()
	if err != nil {
		fmt.Printf("error adding files to scenario: %v", err)
	}

	err = emb.AddToStep()
	if err != nil {
		fmt.Printf("error adding files to step: %v", err)
	}

	// Set html options
	html := gocure.HTML{
		Config: html.Data{
			InputJsonPath:     "./example/newReport.json",
			OutputHtmlFolder:  "./example/",
			ShowEmbeddedFiles: true,
			Metadata: models.Metadata{
				AppVersion:      "0.8.7",
				TestEnvironment: "development",
				Browser:         "Google Chrome",
				Platform:        "Linux",
				Parallel:        "Scenarios",
				Executed:        "Remote",
			},
		},
	}

	// Generate HTML report
	err = html.Generate()
	if err != nil {
		fmt.Printf("error generatig html report: %v", err)
	}

}
