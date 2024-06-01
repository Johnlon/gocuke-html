package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/rodrigoodhin/gocure/api/handlers"
	"gitlab.com/rodrigoodhin/gocure/api/models"
	"gitlab.com/rodrigoodhin/gocure/embedded"
	md "gitlab.com/rodrigoodhin/gocure/models"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
	"gitlab.com/rodrigoodhin/gocure/report/html"
)

func recoverRoutes() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
    }
}

// EmbedToFeature
// @Summary embed files to feature
// @Description embed files to feature of a cucumber report
// @Tags embed
// @Accept json
// @Produce json
// @Param message body models.Embed true "Embed Info"
// @Success 200 {object} handlers.ResponseData
// @Failure 400 {object} handlers.ResponseData
// @Router /embed/toFeature [post]
func EmbedToFeature(c *fiber.Ctx) (err error) {
	defer recoverRoutes()

	var embed = models.Embed{}

	if err = c.BodyParser(&embed); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to parse request body",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	// Set embedded options
	emb := gocure.Embedded{
		Config: embedded.Data{
			InputJsonPath:  embed.InputJsonPath,
			OutputJsonPath: embed.OutputJsonPath,
			Files:          embed.Files,
			FeatureIndex:   embed.FeatureIndex,
			ScenarioIndex:  embed.ScenarioIndex,
			StepIndex:      embed.StepIndex,
		},
	}

	if err = emb.AddToFeature(); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to embed files to feature",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	return handlers.Response{
		Ctx:    c,
		Status: fiber.StatusOK,
	}.Send()
}

// EmbedToScenario
// @Summary embed files to scenario
// @Description embed files to scenario of a cucumber report
// @Tags embed
// @Accept json
// @Produce json
// @Param message body models.Embed true "Embed Info"
// @Success 200 {object} handlers.ResponseData
// @Failure 400 {object} handlers.ResponseData
// @Router /embed/toScenario [post]
func EmbedToScenario(c *fiber.Ctx) (err error) {
	defer recoverRoutes()

	var embed = models.Embed{}

	if err = c.BodyParser(&embed); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to parse request body",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	// Set embedded options
	emb := gocure.Embedded{
		Config: embedded.Data{
			InputJsonPath:  embed.InputJsonPath,
			OutputJsonPath: embed.OutputJsonPath,
			Files:          embed.Files,
			FeatureIndex:   embed.FeatureIndex,
			ScenarioIndex:  embed.ScenarioIndex,
			StepIndex:      embed.StepIndex,
		},
	}

	if err = emb.AddToScenario(); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to embed files to scenario",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	return handlers.Response{
		Ctx:    c,
		Status: fiber.StatusOK,
	}.Send()
}

// EmbedToStep
// @Summary embed files to step
// @Description embed files to step of a cucumber report
// @Tags embed
// @Accept json
// @Produce json
// @Param message body models.Embed true "Embed Info"
// @Success 200 {object} handlers.ResponseData
// @Failure 400 {object} handlers.ResponseData
// @Router /embed/toStep [post]
func EmbedToStep(c *fiber.Ctx) (err error) {
	defer recoverRoutes()

	var embed = models.Embed{}

	if err = c.BodyParser(&embed); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to parse request body",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	// Set embedded options
	emb := gocure.Embedded{
		Config: embedded.Data{
			InputJsonPath:  embed.InputJsonPath,
			OutputJsonPath: embed.OutputJsonPath,
			Files:          embed.Files,
			FeatureIndex:   embed.FeatureIndex,
			ScenarioIndex:  embed.ScenarioIndex,
			StepIndex:      embed.StepIndex,
		},
	}

	if err = emb.AddToStep(); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to embed files to step",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	return handlers.Response{
		Ctx:    c,
		Status: fiber.StatusOK,
	}.Send()
}

// HTMLGenerate
// @Summary generate html reports
// @Description generate html reports from cucumber report
// @Tags html
// @Accept json
// @Produce json
// @Param message body models.HTML true "HTML Info"
// @Success 200 {object} handlers.ResponseData
// @Failure 400 {object} handlers.ResponseData
// @Router /html/generate [post]
func HTMLGenerate(c *fiber.Ctx) (err error) {
	defer recoverRoutes()

	var rep = models.HTML{}

	if err = c.BodyParser(&rep); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to parse request body",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	// Set html options
	r := gocure.HTML{
		Config: html.Data{
			InputJsonPath:      rep.InputJsonPath,
			InputFolderPath:    rep.InputFolderPath,
			MergeFiles:         rep.MergeFiles,
			IgnoreBadJsonFiles: rep.IgnoreBadJsonFiles,
			OutputHtmlFolder:   rep.OutputHtmlFolder,
			Title:              rep.Title,
			ShowEmbeddedFiles:  rep.ShowEmbeddedFiles,
			Metadata: md.Metadata{
				AppVersion:      rep.Metadata.AppVersion,
				TestEnvironment: rep.Metadata.TestEnvironment,
				Browser:         rep.Metadata.Browser,
				Platform:        rep.Metadata.Platform,
				Parallel:        rep.Metadata.Parallel,
				Executed:        rep.Metadata.Executed,
			},
		},
	}

	if err = r.Generate(); err != nil {
		return handlers.Response{
			Ctx:    c,
			Status: fiber.StatusBadRequest,
			Msg:    "Error to generate html report",
			Data:   fiber.Map{"error": err.Error()},
		}.Send()
	}

	return handlers.Response{
		Ctx:    c,
		Status: fiber.StatusOK,
	}.Send()
}
