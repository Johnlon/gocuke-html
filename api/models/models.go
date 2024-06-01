package models

type Embed struct {
	InputJsonPath  string   `json:"inputJsonPath,omitempty"`
	OutputJsonPath string   `json:"outputJsonPath,omitempty"`
	Files          []string `json:"files,omitempty"`
	FeatureIndex   int      `json:"featureIndex,omitempty"`
	ScenarioIndex  int      `json:"scenarioIndex,omitempty"`
	StepIndex      int      `json:"stepIndex,omitempty"`
}

type HTML struct {
	InputJsonPath      string   `json:"inputJsonPath,omitempty"`
	InputFolderPath    string   `json:"inputFolderPath,omitempty"`
	MergeFiles         bool     `json:"mergeFiles,omitempty"`
	IgnoreBadJsonFiles bool     `json:"ignoreBadJsonFiles,omitempty"`
	OutputHtmlFolder   string   `json:"outputHtmlFolder,omitempty"`
	Title              string   `json:"title,omitempty"`
	ShowEmbeddedFiles  bool     `json:"showEmbeddedFiles,omitempty"`
	Metadata           Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	AppVersion      string `json:"appVersion,omitempty"`
	TestEnvironment string `json:"testEnvironment,omitempty"`
	Browser         string `json:"browser,omitempty"`
	Platform        string `json:"platform,omitempty"`
	Parallel        string `json:"parallel,omitempty"`
	Executed        string `json:"executed,omitempty"`
}
