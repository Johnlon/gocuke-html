package models

type Feature struct {
	URI         string     `json:"uri"`
	ID          string     `json:"id"`
	Keyword     string     `json:"keyword"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Line        int        `json:"line"`
	Comments    []Comments `json:"comments"`
	Tags        []Tags     `json:"tags"`
	Scenarios   []Scenario `json:"elements"`
	Embeddings  []Embed    `json:"embeddings"`
}

type Comments struct {
	Value string `json:"value"`
	Line  int    `json:"line"`
}

type Tags struct {
	Name string `json:"name"`
	Line int    `json:"line"`
}

type Match struct {
	Location string `json:"location"`
}

type Result struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
	Duration     int    `json:"duration"`
}

type Step struct {
	Keyword    string  `json:"keyword"`
	Name       string  `json:"name"`
	Line       int     `json:"line"`
	Match      Match   `json:"match"`
	Result     Result  `json:"result"`
	Rows       []Row   `json:"rows"`
	Embeddings []Embed `json:"embeddings"`
	DocString  DocString  `json:"doc_string"`
}

type Embed struct {
	Media Media  `json:"media"`
	Data  string `json:"data"`
}

type Media struct {
	Type string `json:"type"`
}

type Row struct {
	Cells []string `json:"cells"`
}

type DocString struct {
	Value       string `json:"value"`
	ContentType string `json:"content_type"`
	Line        int    `json:"line"`
}

type Scenario struct {
	ID          string  `json:"id"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Line        int     `json:"line"`
	Type        string  `json:"type"`
	Tags        []Tags  `json:"tags"`
	Steps       []Step  `json:"steps"`
	Embeddings  []Embed `json:"embeddings"`
}

type Metadata struct {
	AppVersion      string `json:"appVersion,omitempty"`
	TestEnvironment string `json:"testEnvironment,omitempty"`
	Browser         string `json:"browser,omitempty"`
	Platform        string `json:"platform,omitempty"`
	Parallel        string `json:"parallel,omitempty"`
	Executed        string `json:"executed,omitempty"`
}
