package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/taylormonacelli/itsrocky/daggerverse"
)

func Main() error {
	repos, err := LoadFromFile()
	if err != nil {
		return fmt.Errorf("error loading from file: %v", err)
	}

	repoList, err := buildCustomizedRepositoryInfoList(repos)
	if err != nil {
		return fmt.Errorf("error building customized repository info list: %v", err)
	}

	uniqueRepos, err := daggerverse.MostRecentIndexed(repoList)
	if err != nil {
		return fmt.Errorf("error uniquifying: %v", err)
	}

	fmt.Println(uniqueRepos)

	return nil
}

func buildCustomizedRepositoryInfoList(repos []daggerverse.RepositoryInfo) (daggerverse.CustomizedRepositoryInfoList, error) {
	var err error

	var repoList daggerverse.CustomizedRepositoryInfoList
	for _, repo := range repos {
		cr := daggerverse.CustomizedRepositoryInfo{RepositoryInfo: repo}

		cr.Author, err = daggerverse.GetAuthor(cr.Path)
		if err != nil {
			return nil, fmt.Errorf("error getting author: %v", err)
		}

		cr.AuthorRepoURL, err = daggerverse.GetAuthorRepoURL(cr.Path)
		if err != nil {
			return nil, fmt.Errorf("error getting author repo URL: %v", err)
		}

		repoList = append(repoList, cr)
	}

	return repoList, nil
}

func LoadFromFile() ([]daggerverse.RepositoryInfo, error) {
	file, err := os.Open(DataFilename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var repos []daggerverse.RepositoryInfo
	err = json.Unmarshal(data, &repos)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %v", err)
	}
	return repos, nil
}
