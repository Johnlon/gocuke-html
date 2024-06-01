package html

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gitlab.com/rodrigoodhin/gocure/data"
)

func (d *Data) writeReport() (err error) {

	timeNow := time.Now().Format("20060102_150405.000")

	timeNow = strings.ReplaceAll(timeNow, ".", "_")

	if _, err = os.Stat(d.OutputHtmlFolder); os.IsNotExist(err) {

		folders := strings.Split(d.OutputHtmlFolder, "/")

		createFolder := ""

		for i := 0; i < len(folders)-1; i++ {

			createFolder += folders[i] + "/"

			errDir := os.Mkdir(createFolder, 0777)

			if errDir != nil {
				return fmt.Errorf(errDir.Error())
			}
		}
	}

	err = ioutil.WriteFile(d.OutputHtmlFolder+"report_"+timeNow+".html", []byte(d.HTMLContent), 0644)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return
}

func (d *Data) randomModalID(prefix string) (dateTime string) {
	dateNow := time.Now().Format("20060102")
	timeNow := time.Now().Format("150405.000000")
	timeNow = strings.ReplaceAll(timeNow, ".", "")
	dateTime = prefix + dateNow[2:8] + timeNow[0:12]
	return
}

func (d *Data) validateDefaultValues() {
	if d.Title == "" {
		d.Title = data.DefaultTitle
	}

	if d.Metadata.AppVersion == "" {
		d.Metadata.AppVersion = data.DefaultAppVersion
	}

	if d.Metadata.TestEnvironment == "" {
		d.Metadata.TestEnvironment = data.DefaultTestEnvironment
	}

	if d.Metadata.Browser == "" {
		d.Metadata.Browser = data.DefaultBrowser
	}

	if d.Metadata.Platform == "" {
		d.Metadata.Platform = data.DefaultPlatform
	}

	if d.Metadata.Parallel == "" {
		d.Metadata.Parallel = data.DefaultParallel
	}

	if d.Metadata.Executed == "" {
		d.Metadata.Executed = data.DefaultExecuted
	}
}

func (d *Data) tmplParse(name string, data interface{}) (content string, err error) {
	tmpl, err := template.ParseFS(htmlTmpl, name)
	if err != nil {
		return "", fmt.Errorf("error trying to load %v file : %v", name, err)
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, &data); err != nil {
		return "", fmt.Errorf("error trying to parse %v file : %v", name, err)
	}

	content += buf.String()

	return
}

func (d *Data) prettyString(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}
