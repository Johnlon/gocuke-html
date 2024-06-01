package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/rodrigoodhin/gocure/data"
	"gitlab.com/rodrigoodhin/gocure/embedded"
	"gitlab.com/rodrigoodhin/gocure/models"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
	"gitlab.com/rodrigoodhin/gocure/report/html"
)

func main() {

	htmlOption := flag.Bool("h", false, "HTML option")
	flag.BoolVar(htmlOption, "html", false, "HTML option")

	embeddedOption := flag.Bool("e", false, "Embedded option")
	flag.BoolVar(embeddedOption, "embedded", false, "Embedded option")

	// common tags
	inputJsonPath := flag.String("j", "", "Input json path")

	// html tags
	inputFolderPath := flag.String("f", "", "Input folder path")
	mergeFiles := flag.Bool("m", false, "Merge cucumber json files into one report")
	ignoreBadJsonFiles := flag.Bool("i", false, "Ignore bad json files")
	outputFolder := flag.String("o", "", "Output folder")
	title := flag.String("t", data.DefaultTitle, "Report title")
	showEmbeddedFiles := flag.Bool("s", false, "Show embedded files")
	metaAppVersion := flag.String("AppVersion", data.DefaultAppVersion, "Metadata - App version")
	metaTestEnvironment := flag.String("TestEnvironment", data.DefaultTestEnvironment, "Metadata - Test environment")
	metaBrowser := flag.String("Browser", data.DefaultBrowser, "Metadata - Browser")
	metaPlatform := flag.String("Platform", data.DefaultPlatform, "Metadata - Platform")
	metaParallel := flag.String("Parallel", data.DefaultParallel, "Metadata - Parallel")
	metaExecuted := flag.String("Executed", data.DefaultExecuted, "Metadata - Executed")

	// embedded tags
	outputJsonPath := flag.String("u", "", "Output json path")
	var embeddedFilesPath models.ParamList
	flag.Var(&embeddedFilesPath, "l", "Embedded file path")
	featureIndex := flag.Int("a", -1, "Feature index")
	scenarioIndex := flag.Int("c", -1, "Scenario index")
	stepIndex := flag.Int("p", -1, "Step index")

	flag.Parse()

	var err error
	errorsMsg := ""

	if !*htmlOption && !*embeddedOption {
		errorsMsg += "You must set an option\n"
	}
	if *htmlOption && (*inputJsonPath == "" && *inputFolderPath == "") {
		errorsMsg += "You must set an input method path for the generate HTML Reports\n"
	}

	if *embeddedOption && (*inputJsonPath == "") {
		errorsMsg += "You must set an input report path for the Embed files to report\n"
	}

	if *embeddedOption && *outputJsonPath == "" {
		errorsMsg += "You must set an output file path for the Embed files to report\n"
	}

	if *embeddedOption && len(embeddedFilesPath) == 0 {
		errorsMsg += "You must set at least one file path for the Embed files to report\n"
	}

	if *embeddedOption && (*featureIndex == -1 && *scenarioIndex == -1 && *stepIndex == -1) {
		errorsMsg += "You must set an index for the Embed files to report\n"
	}

	if errorsMsg != "" {
		fmt.Println("errors:")
		fmt.Print(errorsMsg)
		usage()
		return
	}

	if *embeddedOption && errorsMsg == "" {
		emb := gocure.Embedded{
			Config: embedded.Data{
				InputJsonPath:  *inputJsonPath,
				OutputJsonPath: *outputJsonPath,
				Files:          embeddedFilesPath,
				FeatureIndex:   *featureIndex,
				ScenarioIndex:  *scenarioIndex,
				StepIndex:      *stepIndex,
			},
		}

		if emb.Config.FeatureIndex >= 0 && emb.Config.ScenarioIndex >= 0 && emb.Config.StepIndex >= 0 {
			err = emb.AddToStep()

		} else if emb.Config.FeatureIndex >= 0 && emb.Config.ScenarioIndex >= 0 && emb.Config.StepIndex == -1 {
			err = emb.AddToScenario()

		} else if emb.Config.FeatureIndex >= 0 && emb.Config.ScenarioIndex == -1 && emb.Config.StepIndex == -1 {
			err = emb.AddToFeature()

		} else {
			err = fmt.Errorf("no index was chosen")
		}

		if err != nil {
			fmt.Println(err)
		}
	}

	if *htmlOption && errorsMsg == "" {
		html := gocure.HTML{
			Config: html.Data{
				InputJsonPath:      *inputJsonPath,
				InputFolderPath:    *inputFolderPath,
				MergeFiles:         *mergeFiles,
				IgnoreBadJsonFiles: *ignoreBadJsonFiles,
				OutputHtmlFolder:   *outputFolder,
				Title:              *title,
				ShowEmbeddedFiles:  *showEmbeddedFiles,
				Metadata: models.Metadata{
					AppVersion:      *metaAppVersion,
					TestEnvironment: *metaTestEnvironment,
					Browser:         *metaBrowser,
					Platform:        *metaPlatform,
					Parallel:        *metaParallel,
					Executed:        *metaExecuted,
				},
			},
		}

		err = html.Generate()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func usage() {
	fmt.Printf("\nExample:\n%s [TAG] [VALUE]\n", os.Args[0])
	fmt.Printf("-h, --html string\n   HTML option\n")
	fmt.Printf("-e, --embedded string\n   Embedded option\n")
	fmt.Printf("\nAvailable common tags:\n")
	fmt.Printf("-j string\n   Input json path\n")
	fmt.Printf("\nAvailable tags for 'html' option:\n")
	fmt.Printf("-f string\n   Input folder path\n")
	fmt.Printf("-m bool\n   Merge cucumber json files into one report (default false)\n")
	fmt.Printf("-i bool\n   Ignore bad json files (default false)\n")
	fmt.Printf("-o string\n   Output folder\n")
	fmt.Printf("-t string\n   Report title (default '" + data.DefaultTitle + "')\n")
	fmt.Printf("-s string\n   Show embedded files (default false)\n")
	fmt.Printf("--AppVersion string\n   Metadata - App version (default '" + data.DefaultAppVersion + "')\n")
	fmt.Printf("--TestEnvironment string\n   Metadata - Test environment (default '" + data.DefaultTestEnvironment + "')\n")
	fmt.Printf("--Browser string\n   Metadata - Browser (default '" + data.DefaultBrowser + "')\n")
	fmt.Printf("--Platform string\n   Metadata - Platform (default '" + data.DefaultPlatform + "')\n")
	fmt.Printf("--Parallel string\n   Metadata - Parallel (default '" + data.DefaultParallel + "')\n")
	fmt.Printf("--Executed string\n   Metadata - Executed (default '" + data.DefaultExecuted + "')\n")
	fmt.Printf("\nAvailable tags for 'embedded' option:\n")
	fmt.Printf("-u string\n   Output json path\n")
	fmt.Printf("-l string\n   Embedded file path\n")
	fmt.Printf("-a int\n   Feature index (default -1)\n")
	fmt.Printf("-c int\n   Scenario index (default -1)\n")
	fmt.Printf("-p int\n   Step index (default -1)\n")
}
