package html

import (
	"fmt"
)

func (d *Data) printChartJs() (chartJS string, err error) {

	featurePassed, featureFailed := d.featuresResult()
	scenarioPassed, scenarioFailed := d.scenariosResult()
	stepPassed, stepFailed, stepPending, stepSkipped := d.stepsResult()

	chartJsHTML := ChartJsHTML{
		PassedFeatureChart:  featurePassed,
		FailedFeatureChart:  featureFailed,
		PassedScenarioChart: scenarioPassed,
		FailedScenarioChart: scenarioFailed,
		PassedStepChart:     stepPassed,
		FailedStepChart:     stepFailed,
		UndefinedStepChart:  stepPending,
		SkippedStepChart:    stepSkipped,
	}

	content, err := d.tmplParse("tmpl/chartJs.html", &chartJsHTML)
	if err != nil {
		return "", fmt.Errorf("error trying to parse chart JS file : %v", err)
	}

	chartJS, err = d.minifyChartJS(content)
	if err != nil {
		return "", fmt.Errorf("error trying to minify chart JS file %v", err)
	}

	return
}

func (d *Data) printChart(chartName string) (chart string, err error) {

	chartHTML := ChartHTML{
		Id: "chart_" + chartName,
	}

	content, err := d.tmplParse("tmpl/chart.html", &chartHTML)
	if err != nil {
		return "", fmt.Errorf("error trying to parse chart file : %v", err)
	}

	chart += content

	return
}

func (d *Data) getChartsName() []string {
	return []string{"features", "scenarios", "steps"}
}
