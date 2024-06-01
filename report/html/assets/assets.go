package assets

import (
	_ "embed"
)

//go:embed css/style.css
var CSS string

//go:embed js/script.js
var JS string

//go:embed js/chart.js
var ChartJSlibrary string