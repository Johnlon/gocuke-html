package gocure

import (
	"fmt"
)

func (h *HTML) Generate() (err error) {

	err = h.Config.CreateReport()
	if err != nil {
		return fmt.Errorf("error generating report : %v", err)
	}

	return
}
