package embedded

type Data struct {
	InputJsonPath  string      `json:"inputJsonPath,omitempty"`
	OutputJsonPath string      `json:"outputJsonPath,omitempty"`
	Files          []string    `json:"files,omitempty"`
	FeatureIndex   int         `json:"featureIndex,omitempty"`
	ScenarioIndex  int         `json:"scenarioIndex,omitempty"`
	StepIndex      int         `json:"stepIndex,omitempty"`
	Model          interface{} 
}
