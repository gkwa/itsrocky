package daggerverse

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func DateToAge(t time.Time) string {
	return humanize.Time(t)
}

func (repos CustomizedRepositoryInfoSlice) GenerateHTMLReport() (string, error) {
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Repository Report</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
		<style>
			body {
				padding: 20px;
			}

			table {
				width: 100%;
				border-collapse: collapse;
				margin-top: 20px;
			}

			th, td {
				border: 1px solid #ddd;
				padding: 8px;
				text-align: left;
			}

			th {
				background-color: #f2f2f2;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1 class="mt-4 mb-4">Repository Report</h1>
			<table class="table table-bordered">
				<thead>
					<tr>
						<th>Author</th>
						<th>Recent Commit Age</th>
						<th>Browser URL</th>
						<th>Git URL</th>
					</tr>
				</thead>
				<tbody>
					{{range .Repos}}
					<tr>
						<td><a href="{{.AuthorRepoURL}}">{{.Author}}</a></td>
						<td>{{DateToAge .CreatedAt }}</td>
					<td><a target="_blank" href="{{.BrowseURL}}">{{.ProjectDir}}</a></td>
						<td><a target="_blank" href="{{.GitURL}}">{{.GitURL}}</a></td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</div>
		<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
	</body>
	</html>
	`

	tmpl, err := template.New("report").Funcs(template.FuncMap{
		"DateToAge": DateToAge,
	}).Parse(htmlTemplate)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %v", err)
	}

	currentTime := time.Now()
	outputBuffer := &strings.Builder{}

	err = tmpl.Execute(outputBuffer, struct {
		Repos       CustomizedRepositoryInfoSlice
		CurrentTime time.Time
	}{
		Repos:       repos,
		CurrentTime: currentTime,
	})
	if err != nil {
		return "", fmt.Errorf("error executing template: %v", err)
	}

	return outputBuffer.String(), nil
}
