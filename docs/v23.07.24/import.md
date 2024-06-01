# Importing into a project

## Import

To import Gocure in a project to generate html reports, use this code:

```go
import (
	"gitlab.com/rodrigoodhin/gocure/models"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
	"gitlab.com/rodrigoodhin/gocure/report/html"
)
```

To import Gocure in a project to embed files to Ccucumber json reports, use this code:

```go
import (
	"gitlab.com/rodrigoodhin/gocure/embedded"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
)
```

To import Gocure in a project to generate html reports and to embed files to Ccucumber json reports, use this code:

```go
import (
	"gitlab.com/rodrigoodhin/gocure/embedded"
	"gitlab.com/rodrigoodhin/gocure/models"
	"gitlab.com/rodrigoodhin/gocure/report/html"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
)
```

## Configure

After the import you have to configure the gocure options.

To embed files, use this code:

```go
emb := gocure.Embedded{
    Config: embedded.Data{
        InputJsonPath:  "./example/report.json",
        OutputJsonPath: "./example/newReport.json",
        Files: []string{
            "./example/video.mp4",
            "./example/text.txt",
            "https://i.ibb.co/LpRkTqf/gocure-small.png",
        },
        FeatureIndex: 0,
        ScenarioIndex: 0,
        StepIndex: 2,
    },
}
```

To generate html reports, use this code:

```golang
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
```

### Options

To embed files, these are the options:

| Name | Type | Description | Default |
| --- | --- | --- | --- |
| InputJsonPath | String | Path of the Cucumber JSON report file | |
| OutputJsonPath | String | Output json path | |
| Files | []String | Paths of the files to be embedded | |
| FeatureIndex | Int | Index of the Feature | |
| ScenarioIndex | Int | Index of the Scenario | |
| StepIndex | Int | Index of the Step | |

To generate html reports, these are the options:

| Option | Type | Description | Default |
| --- | --- | --- | --- |
| InputJsonPath | String | Path of the Cucumber JSON report file | |
| InputFolderPath | String | Path of a directory with more than one Cucumber JSON report file<br>This option will be ignored if option "InputJsonPath" were used | |
| MergeFiles | Bool | Merge all Cucumber JSON report files into only one HTML report<br>This option is only used if option "InputJsonPath" were used | false |
| IgnoreBadJsonFiles | Bool | Ignore bad Cucumber JSON report files<br>This option is only used if option "InputJsonPath" were used | false |
| OutputHtmlFolder | String | Path of the HTML report output folder | |
| Title | String | Title of your report | GO Cucumber HTML Report |
| ShowEmbeddedFiles | Bool | Show embedded files | false |
| Metadata.AppVersion | String | Metadata - App Version | - |
| Metadata.TestEnvironment | String | Metadata - Test Environment | - |
| Metadata.Browser | String | Metadata - Browser | - |
| Metadata.Platform | String | Metadata - Platform | - |
| Metadata.Parallel | String | Metadata - Parallel | - |
| Metadata.Executed | String | Metadata - Executed | - |

## Run 

After setting up gocure you can finally run it.

Gocure can embed files into Features, Scenarios or Steps. 

> [!NOTE]
> These three options are used to determinate into which Feature/Scenario/Step the files should be embedded.

If you want to embed files to ```Step```, the options ```FeatureIndex```, ```ScenarioIndex``` and ```StepIndex``` must be set. 
To embed files to Step, use this code:

```go
err = emb.AddToStep()
if err != nil {
    log.Fatal(err)
}
```

If you want to embed files to ```Scenario```, the options ```FeatureIndex``` and ```ScenarioIndex``` must be set. 
To embed files to Scenario, use this code:

```go
err = emb.AddToScenario()
if err != nil {
    log.Fatal(err)
}
```

If you want to embed files to ```Feature```, the options ```FeatureIndex``` must be set. 
To embed files to Feature, use this code:

```go
err = emb.AddToFeature()
if err != nil {
    log.Fatal(err)
}
```

If you want to generate html reports, use this code:

```go
err = html.Generate()
if err != nil {
    log.Fatal(err)
}
```

## Full example

Full example with all codes:

```go
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
			InputJsonPath:  "./example/report.json",
			OutputJsonPath: "./example/newReport.json",
			Files: []string{
				"./example/video.mp4",
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
```