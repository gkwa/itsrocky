package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

func RunFetch() error {
	timeout := 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stale, age := isFileOlderThan(DataFilename, maxAgeDuration)

	if !stale && !viper.GetBool("no-cache") {
		fmt.Fprintf(os.Stderr, "data age %s<5h, skipping re-fetch\n", age.Truncate(time.Second))
		return nil
	}

	err := os.MkdirAll(filepath.Dir(DataFilename), 0o755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %v", filepath.Dir(DataFilename), err)
	}

	data, err := fetchDataWithContext(ctx, endpointURL)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	err = writeJsonTofile(DataFilename, data)
	if err != nil {
		return fmt.Errorf("error writing to %s: %v", DataFilename, err)
	}

	fmt.Fprintf(os.Stderr, "Data fetched and written to %s successfully.\n", DataFilename)

	return nil
}

func fetchDataWithContext(ctx context.Context, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error performing HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func writeJsonTofile(filename string, data []byte) error {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	prettyPrintedData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data to pretty-printed JSON: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filename, err)
	}
	defer file.Close()

	_, err = file.Write(prettyPrintedData)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filename, err)
	}

	return nil
}

func isFileOlderThan(filename string, maxAge time.Duration) (bool, time.Duration) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return true, 0
	}

	age := time.Since(fileInfo.ModTime())

	return age > maxAge, age
}
