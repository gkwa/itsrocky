package daggerverse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func convertToURL(path string) (*url.URL, error) {
	if path == "" {
		return nil, nil
	}

	path = strings.ToLower(path)

	x := strings.TrimPrefix(path, "https://")

	fullURL := "https://" + x

	u, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %v", err)
	}
	return u, nil
}

func GetFinalPathSegment(u *url.URL) string {
	segments := strings.Split(u.Path, "/")
	return segments[len(segments)-1]
}

func GetProjectDir(path string) (string, error) {
	u, err := convertToURL(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	final := GetFinalPathSegment(u)

	return final, nil
}

func GetAuthor(path string) (string, error) {
	u, err := convertToURL(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	p := strings.TrimPrefix(u.Path, "/")
	z := strings.Index(p, "/")
	author := p[:z]

	return author, nil
}

func getHost(path string) (string, error) {
	u, err := convertToURL(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	return u.Host, nil
}

func GetAuthorRepoURL(path string) (string, error) {
	author, err := GetAuthor(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	host, err := getHost(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	y := "https://" + host + "/" + author

	return y, nil
}

func (c CustomizedRepositoryInfo) String() (string, error) {
	jsonBytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}
	return string(jsonBytes), nil
}

func (l CustomizedRepositoryInfoSlice) String() string {
	jsonBytes, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v\n", err)
	}
	return string(jsonBytes)
}

func (cr CustomizedRepositoryInfos) String() string {
	// Use json.MarshalIndent for the entire slice
	jsonBytes, err := json.MarshalIndent(cr, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v\n", err)
	}
	return string(jsonBytes)
}

func MostRecentIndexed(records CustomizedRepositoryInfoSlice) (CustomizedRepositoryInfoSlice, error) {
	uniqueRecords := make(map[string]CustomizedRepositoryInfo)

	for _, record := range records {
		existingRecord, exists := uniqueRecords[record.Path]

		// Check if the key doesn't exist in the map OR the current record has a younger 'IndexedAt' timestamp
		if !exists || record.IndexedAt.After(existingRecord.IndexedAt) {
			// If either condition is true, update 'uniqueRecords' with the current record using its 'Path' as the key
			uniqueRecords[record.Path] = record
		}
	}
	// Convert the map to a slice for sorting
	var sortedRecords []CustomizedRepositoryInfo
	for _, record := range uniqueRecords {
		sortedRecords = append(sortedRecords, record)
	}

	// Sort the records based on the indexed_at field
	sort.Slice(sortedRecords, func(i, j int) bool {
		return sortedRecords[i].IndexedAt.Before(sortedRecords[j].IndexedAt)
	})

	return sortedRecords, nil
}
