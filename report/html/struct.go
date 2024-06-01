package html

import (
	"html/template"

	"gitlab.com/rodrigoodhin/gocure/models"
)

type Data struct {
	InputJsonPath      string `json:"inputJsonPath,omitempty"`
	InputFolderPath    string `json:"inputFolderPath,omitempty"`
	MergeFiles         bool   `json:"mergeFiles,omitempty"`
	IgnoreBadJsonFiles bool   `json:"ignoreBadJsonFiles,omitempty"`
	OutputHtmlFolder   string `json:"outputHtmlFolder,omitempty"`
	Title              string `json:"title,omitempty"`
	ShowEmbeddedFiles  bool   `json:"showEmbeddedFiles,omitempty"`
	Timer              int
	Features           []models.Feature
	JSONContent        []byte
	HTMLContent        string
	Metadata           models.Metadata `json:"metadata,omitempty"`
	GeneratedAt        string
}

type BaseHTML struct {
	Title          string
	Timer          template.HTML
	Chart          template.HTML
	ChartJsLibrary template.HTML
	ChartJs        template.HTML
	Content        template.HTML
	CountAll       int
	CountPassed    int
	CountFailed    int
	CSS            template.HTML
	JS             template.HTML
	Metadata       template.HTML
	GeneratedAt    string
}

type ChartJsHTML struct {
	PassedFeatureChart  int
	FailedFeatureChart  int
	PassedScenarioChart int
	FailedScenarioChart int
	PassedStepChart     int
	FailedStepChart     int
	UndefinedStepChart  int
	SkippedStepChart    int
}

type ChartHTML struct {
	Id string
}

type FeatureHTML struct {
	ModalID     string
	Id          string
	Tags        string
	Name        string
	Description template.HTML
	Failed      int
	Passed      int
	Scenarios   template.HTML
	Embed       template.HTML
	Duration    template.HTML
	Result      string
}

type ScenarioHTML struct {
	ModalID     string
	Id          string
	Tags        string
	Keyword     string
	Name        string
	Description template.HTML
	Skipped     int
	Undefined   int
	Failed      int
	Passed      int
	Steps       template.HTML
	Embed       template.HTML
	Duration    template.HTML
}

type StepHTML struct {
	ModalID          string
	Status           string
	StatusCapitalize string
	ResultIcon       template.HTML
	Keyword          template.HTML
	Name             string
	Message          template.HTML
	Embed            template.HTML
	Table            template.HTML
	DocString        template.HTML
	Duration         template.HTML
}

type ModalHTML struct {
	ModalID   string
	Keyword   string
	Name      string
	EmbedHTML template.HTML
	Duration  template.HTML
}

type ModalFileHTML struct {
	RowClass       string
	MediaType      string
	MediaSize      string
	MediaData      string
	MediaExtension string
	CopyOption     template.HTML
	PrintOption    template.HTML
	MediaContent   template.HTML
	ContentID      string
}
