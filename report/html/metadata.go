package html

import (
	"fmt"
)

func (d *Data) printMetadata() (metadata string, err error) {

	metadata, err = d.tmplParse("tmpl/metadata.html", &d.Metadata)
	if err != nil {
		return "", fmt.Errorf("error trying to parse metadata file : %v", err)
	}

	return
}
