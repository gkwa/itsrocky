package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/taylormonacelli/itsrocky/daggerverse"
)

func Main() error {
	repos, err := LoadFromFile()
	if err != nil {
		return fmt.Errorf("error loading from file: %v", err)
	}

	repoList, err := BuildCustomizedRepositoryInfoSlice(repos)
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

func BuildCustomizedRepositoryInfoSlice(repos []daggerverse.RepositoryInfo) (daggerverse.CustomizedRepositoryInfoSlice, error) {
	var err error

	var reposSlice daggerverse.CustomizedRepositoryInfoSlice
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

		cr.ProjectDir, err = daggerverse.GetProjectDir(cr.Path)
		if err != nil {
			return nil, fmt.Errorf("error getting project dir: %v", err)
		}
		reposSlice = append(reposSlice, cr)
	}

	uniqueRepos, err := daggerverse.MostRecentIndexed(reposSlice)
	if err != nil {
		return nil, fmt.Errorf("error uniquifying: %v", err)
	}

	sort.Slice(uniqueRepos, func(i, j int) bool {
		return uniqueRepos[i].CreatedAt.After(uniqueRepos[j].CreatedAt)
	})

	return uniqueRepos, nil
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
