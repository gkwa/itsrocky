package fetch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func RunFetch() error {
	endpointURL := "https://daggerverse.dev/api/refs"

	timeout := 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	data, err := fetchDataWithContext(ctx, endpointURL)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	err = writeJsonTofile("daggervers.json", data)
	if err != nil {
		return fmt.Errorf("error writing to daggervers.json: %v", err)
	}

	fmt.Println("Data fetched and written to daggervers.json successfully.")

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
