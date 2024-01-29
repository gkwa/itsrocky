package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/taylormonacelli/itsrocky/daggerverse"
)

func LoadFromFile() error {
	file, err := os.Open(DataFilename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var repos []daggerverse.RepositoryInfo
	err = json.Unmarshal(data, &repos)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	var customizedRepos daggerverse.CustomizedRepositoryInfos
	for _, repo := range repos {
		customizedRepo := daggerverse.CustomizedRepositoryInfo{RepositoryInfo: repo}
		cr := daggerverse.NewCustomizedRepositoryInfoBuilder().
			Author(customizedRepo.Path).
			Build()
		customizedRepos = append(customizedRepos, cr)

	}

	fmt.Println(customizedRepos)

	return nil
}
