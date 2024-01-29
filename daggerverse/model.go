package daggerverse

import "time"

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

type ItemsWrapper struct {
	Items []RepositoryInfo `json:"items"`
}

type CustomizedRepositoryInfo struct {
	RepositoryInfo
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
