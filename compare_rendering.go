package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Result struct {
	URL              string
	Slug             string
	RodHTMLPath      string
	CurlLen          int
	RodLen           int
	CurlTitle        string
	RodTitle         string
	CurlHasH1        bool
	RodHasH1         bool
	CurlTextWords    int
	RodTextWords     int
	RenderingVerdict string
	CurlErr          string
	RodErr           string
}

var titleRe = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
var h1Re = regexp.MustCompile(`(?is)<h1[^>]*>(.*?)</h1>`)
var tagRe = regexp.MustCompile(`(?s)<[^>]*>`)
var wsRe = regexp.MustCompile(`\s+`)
var nonAlphaNumRe = regexp.MustCompile(`[^a-z0-9]+`)

func getTitle(html string) string {
	m := titleRe.FindStringSubmatch(html)
	if len(m) < 2 {
		return ""
	}
	return strings.TrimSpace(tagRe.ReplaceAllString(m[1], ""))
}

func textWordCount(html string) int {
	text := tagRe.ReplaceAllString(html, " ")
	text = wsRe.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)
	if text == "" {
		return 0
	}
	return len(strings.Split(text, " "))
}

func slugify(url string) string {
	cleaned := strings.ToLower(url)
	cleaned = strings.TrimPrefix(cleaned, "https://")
	cleaned = strings.TrimPrefix(cleaned, "http://")
	cleaned = nonAlphaNumRe.ReplaceAllString(cleaned, "-")
	cleaned = strings.Trim(cleaned, "-")
	if cleaned == "" {
		return "article"
	}
	return cleaned
}

func fetchCurl(url string) (string, error) {
	cmd := exec.Command("curl", "-L", "-sS", "--max-time", "45", url)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func fetchRod(browser *rod.Browser, url string) (string, error) {
	page := browser.MustPage()
	defer page.MustClose()
	page = page.Timeout(60 * time.Second)
	if err := page.Navigate(url); err != nil {
		return "", err
	}
	page.MustWaitLoad()
	time.Sleep(2 * time.Second)
	return page.HTML()
}

func classify(curlHTML, rodHTML string) string {
	curlWords := textWordCount(curlHTML)
	rodWords := textWordCount(rodHTML)
	if rodWords > curlWords*2 && rodWords-curlWords > 300 {
		return "Likely client-side rendered"
	}
	return "Likely server-side rendered"
}

func analyze(browser *rod.Browser, url, htmlDir string) Result {
	res := Result{URL: url, Slug: slugify(url)}

	curlHTML, curlErr := fetchCurl(url)
	if curlErr != nil {
		res.CurlErr = curlErr.Error()
	} else {
		res.CurlLen = len(curlHTML)
		res.CurlTitle = getTitle(curlHTML)
		res.CurlHasH1 = h1Re.MatchString(curlHTML)
		res.CurlTextWords = textWordCount(curlHTML)
	}

	rodHTML, rodErr := fetchRod(browser, url)
	if rodErr != nil {
		res.RodErr = rodErr.Error()
	} else {
		res.RodLen = len(rodHTML)
		res.RodTitle = getTitle(rodHTML)
		res.RodHasH1 = h1Re.MatchString(rodHTML)
		res.RodTextWords = textWordCount(rodHTML)
		res.RodHTMLPath = filepath.Join(htmlDir, res.Slug+".html")
		_ = os.WriteFile(res.RodHTMLPath, []byte(rodHTML), 0o644)
	}

	switch {
	case res.CurlErr != "" && res.RodErr != "":
		res.RenderingVerdict = "Comparison unavailable: both methods failed"
	case res.CurlErr != "" || res.RodErr != "":
		res.RenderingVerdict = "Comparison unavailable: one method failed"
	default:
		res.RenderingVerdict = classify(curlHTML, rodHTML)
	}

	return res
}

func main() {
	urls := []string{
		"https://builders.ramp.com/post/meet-ramp-research",
		"https://modal.com/blog/how-ramp-built-a-full-context-background-coding-agent-on-modal",
		"https://simonwillison.net/2026/Feb/20/beats/",
		"https://stripe.dev/blog/minions-stripes-one-shot-end-to-end-coding-agents",
	}

	htmlDir := "rendered_html"
	if err := os.MkdirAll(htmlDir, 0o755); err != nil {
		panic(err)
	}

	l := launcher.New().Headless(true).Set("ignore-certificate-errors")
	browserURL, err := l.Launch()
	if err != nil {
		panic(err)
	}

	browser := rod.New().ControlURL(browserURL).MustConnect()
	defer browser.MustClose()

	fmt.Println("# Rendering Comparison: go-rod vs curl")
	for _, url := range urls {
		r := analyze(browser, url, htmlDir)
		fmt.Printf("\n## %s\n", r.URL)
		fmt.Printf("- slug: %s\n", r.Slug)
		fmt.Printf("- curl html length: %d\n", r.CurlLen)
		fmt.Printf("- rod html length: %d\n", r.RodLen)
		fmt.Printf("- curl title: %q\n", r.CurlTitle)
		fmt.Printf("- rod title: %q\n", r.RodTitle)
		fmt.Printf("- curl has <h1>: %v\n", r.CurlHasH1)
		fmt.Printf("- rod has <h1>: %v\n", r.RodHasH1)
		fmt.Printf("- curl visible-text word count: %d\n", r.CurlTextWords)
		fmt.Printf("- rod visible-text word count: %d\n", r.RodTextWords)
		if r.RodHTMLPath != "" {
			fmt.Printf("- rod rendered html path: %s\n", r.RodHTMLPath)
		}
		if r.CurlErr != "" {
			fmt.Printf("- curl error: %s\n", r.CurlErr)
		}
		if r.RodErr != "" {
			fmt.Printf("- rod error: %s\n", r.RodErr)
		}
		fmt.Printf("- verdict: %s\n", r.RenderingVerdict)
	}
}
