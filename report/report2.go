package report

import (
	"fmt"

	"github.com/taylormonacelli/itsrocky/daggerverse"
)

func RunReport2(repos daggerverse.CustomizedRepositoryInfoSlice) error {
	jsonBytes, err := repos.ToJson()
	if err != nil {
		return fmt.Errorf("error generating HTML report: %v", err)
	}

	fmt.Println(string(jsonBytes))

	return nil
}
