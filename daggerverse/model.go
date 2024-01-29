package daggerverse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type RepositoryInfo struct {
	BrowseURL string    `json:"browse_url"`
	CreatedAt time.Time `json:"created_at"`
	GitCommit string    `json:"git_commit"`
	GitURL    string    `json:"git_url"`
	IndexedAt time.Time `json:"indexed_at"`
	Path      string    `json:"path"`
	Release   string    `json:"release"`
	Subpath   string    `json:"subpath"`
	Version   string    `json:"version"`
}

type CustomizedRepositoryInfo struct {
	RepositoryInfo
	MyAuthor string
}

func (c CustomizedRepositoryInfo) String() (string, error) {
	jsonBytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}
	return string(jsonBytes), nil
}

type CustomizedRepositoryInfos []*CustomizedRepositoryInfo

func (cr CustomizedRepositoryInfos) String() string {
	// Use json.MarshalIndent for the entire slice
	jsonBytes, err := json.MarshalIndent(cr, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v\n", err)
	}
	return string(jsonBytes)
}

func (cr CustomizedRepositoryInfos) AddRepo(repo *CustomizedRepositoryInfo) CustomizedRepositoryInfos {
	return append(cr, repo)
}

type CustomizedRepositoryInfoBuilder struct {
	info *CustomizedRepositoryInfo
}

func NewCustomizedRepositoryInfoBuilder() *CustomizedRepositoryInfoBuilder {
	return &CustomizedRepositoryInfoBuilder{
		info: &CustomizedRepositoryInfo{
			RepositoryInfo: RepositoryInfo{}, // Initialize embedded struct
		},
	}
}

func (b *CustomizedRepositoryInfoBuilder) Build() *CustomizedRepositoryInfo {
	return b.info
}

func (b *CustomizedRepositoryInfoBuilder) BrowseURL(browseURL string) *CustomizedRepositoryInfoBuilder {
	b.info.BrowseURL = browseURL
	return b
}

func getAuthor(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	x := strings.TrimPrefix(path, "https://")

	fullURL := "https://" + x

	u, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	p := strings.TrimPrefix(u.Path, "/")
	z := strings.Index(p, "/")
	author := p[:z]

	return author, nil
}

func (b *CustomizedRepositoryInfoBuilder) Author(author string) *CustomizedRepositoryInfoBuilder {
	a, err := getAuthor(author)
	if err != nil {
		panic(err)
	}
	b.info.MyAuthor = a
	return b
}

func (b *CustomizedRepositoryInfoBuilder) CreatedAt(createdAt time.Time) *CustomizedRepositoryInfoBuilder {
	b.info.CreatedAt = createdAt
	return b
}

func (b *CustomizedRepositoryInfoBuilder) GitCommit(gitCommit string) *CustomizedRepositoryInfoBuilder {
	b.info.GitCommit = gitCommit
	return b
}

func (b *CustomizedRepositoryInfoBuilder) GitURL(gitURL string) *CustomizedRepositoryInfoBuilder {
	b.info.GitURL = gitURL
	return b
}

func (b *CustomizedRepositoryInfoBuilder) IndexedAt(indexedAt time.Time) *CustomizedRepositoryInfoBuilder {
	b.info.IndexedAt = indexedAt
	return b
}

func (b *CustomizedRepositoryInfoBuilder) Path(path string) *CustomizedRepositoryInfoBuilder {
	b.info.Path = path
	return b
}

func (b *CustomizedRepositoryInfoBuilder) Release(release string) *CustomizedRepositoryInfoBuilder {
	b.info.Release = release
	return b
}

func (b *CustomizedRepositoryInfoBuilder) Subpath(subpath string) *CustomizedRepositoryInfoBuilder {
	b.info.Subpath = subpath
	return b
}

func (b *CustomizedRepositoryInfoBuilder) Version(version string) *CustomizedRepositoryInfoBuilder {
	b.info.Version = version
	return b
}
