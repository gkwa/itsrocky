package report

import (
	"fmt"
	"html/template"
	"math/rand"
	"strings"
	"time"

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

func RunReport5(repos daggerverse.CustomizedRepositoryInfoSlice) error {
	if len(repos) == 0 {
		return fmt.Errorf("empty slice, cannot get a random object")
	}

	const repoInfoTemplate = `
#!/usr/bin/env bash
set -e

BASE_DIR=$(pwd)
scratch=$BASE_DIR/scratch

rm -rf $scratch
mkdir -p $scratch

git clone --quiet --depth 1 {{.GitURL}} $scratch/{{.Path}}
cd $scratch/{{.Path}}
dir=$(
	rg --files-with-matches --glob=dagger.json '"sdk": "go"' $scratch/{{.Path}} |
	xargs dirname |
	xargs -d '\n' -a - rg -l dag.Container |
	grep -i main.go |
	xargs dirname |
	sort -R |
	head -1
)

if [ -z "$dir" ]; then
	echo "No go sdk found in {{.Path}}"
	exit 1
fi

cd $dir
docker ps --format {{"{{"}}.Names{{"}}"}} --filter name='^/dagger-engine-*' | xargs --no-run-if-empty -I"{}" docker rm --force {}
code $dir
echo "cd $dir && dagger mod install {{ .ModInstallPath }}"
`

	tmpl, err := template.New("repositoryInfo").Parse(repoInfoTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	uRepos, err := daggerverse.MostRecentIndexed(repos)
	if err != nil {
		return fmt.Errorf("error getting most recent indexed: %v", err)
	}

	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	var output strings.Builder
	err = tmpl.Execute(&output, uRepos[rng.Intn(len(repos))])
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
