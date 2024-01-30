package report

import (
	"fmt"

	"github.com/taylormonacelli/itsrocky/daggerverse"
)

func RunReport1(repos daggerverse.CustomizedRepositoryInfoSlice) error {
	htmlReport, err := repos.GenerateHTMLReport()
	if err != nil {
		return fmt.Errorf("error generating HTML report: %v", err)
	}

	fmt.Println(htmlReport)

	return nil
}
