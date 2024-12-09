package Parser

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// Mock handler for testing
func mockHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Test Page</title>
			</head>
			<body>
				<h1 class="mw-page-title-main">Mock Title</h1>
			</body>
		</html>
	`
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(html))
}

func TestParseArticle(t *testing.T) {
	// Start a test server
	server := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer server.Close()

	// Call ParseArticle with the mock server URL
	parseArticle(server.URL)
}

func TestGetTitle(t *testing.T) {
	// Create a sample document
	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Test Page</title>
			</head>
			<body>
				<h1 class="mw-page-title-main">Mock Title</h1>
			</body>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	title := getTitle(doc)
	expectedTitle := "Mock Title"

	if title != expectedTitle {
		t.Errorf("Expected title %q but got %q", expectedTitle, title)
	}
}
