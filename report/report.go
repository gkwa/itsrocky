package report

import (
	"fmt"
	"html/template"
	"strings"

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

func RunReport2(repos daggerverse.CustomizedRepositoryInfoSlice) error {
	jsonBytes, err := repos.ToJson()
	if err != nil {
		return fmt.Errorf("error generating HTML report: %v", err)
	}

	fmt.Println(string(jsonBytes))

	return nil
}

func RunReport3(repos daggerverse.CustomizedRepositoryInfoSlice) error {
	jsonBytes, err := repos.ToJson()
	if err != nil {
		return fmt.Errorf("error generating HTML report: %v", err)
	}

	fmt.Println(string(jsonBytes))

	return nil
}

func RunReport4(repos daggerverse.CustomizedRepositoryInfoSlice) error {
	const repoInfoTemplate = `
#!/usr/bin/env bash
set -e

BASE_DIR=$(pwd)

{{range .}}
# {{.Path}} by {{.Author}}
if [ ! -d $BASE_DIR/{{.Path}} ]; then
    cd $BASE_DIR
    git submodule add --quiet {{.GitURL}} {{.Path}}
    git commit -am "Add submodule {{.Path}}"
fi

cd $BASE_DIR/{{.Path}}
git fetch --quiet

cd $BASE_DIR/{{.Path}}
if ! git checkout --quiet {{.GitCommit}}; then
    echo "Error checking out {{.GitCommit}} for {{.Path}}" >&2
fi

cd $BASE_DIR
if ! git diff-files --quiet; then
    git commit -am "Update repo {{.Path}} to {{.GitCommit}}"
fi
{{end}}
`
	tmpl, err := template.New("repositoryInfo").Parse(repoInfoTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	uRepos, err := daggerverse.MostRecentIndexed(repos)
	if err != nil {
		return fmt.Errorf("error getting most recent indexed: %v", err)
	}

	var output strings.Builder
	err = tmpl.Execute(&output, uRepos)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	trimmedOutput := strings.TrimSpace(output.String())
	output.Reset()
	output.WriteString(trimmedOutput)

	// trim leading and ending whitespace
	fmt.Println(output.String())

	return nil
}
