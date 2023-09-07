package kattis

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/albe2669/kattissues/internal"
)

func createUrlBase(host string) (*url.URL, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("error parsing host: %w", err)
	}

	u.Scheme = "https"

	return u, nil
}

func makeContestUrl(baseUrl *url.URL, contestId string) *url.URL {
	return baseUrl.JoinPath("contests", contestId)
}

func makeRequest(url *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "kattissues")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return res, nil
}

func parseHTML(htmlReader io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %w", err)
	}

	return doc, nil
}

type KattisProblem struct {
	ProblemId string
	Name      string
	Label     string
	Url       string
}

type KattisContestStatus int

const (
	ContestNotStarted KattisContestStatus = iota
	ContestRunning
	ContestEnded
)

type KattisContest struct {
	ContestId string
	StartTime time.Time
	EndTime   time.Time
	Status    KattisContestStatus
	Problems  []*KattisProblem
}

func parseTimes(htmlBody *goquery.Document) (time.Time, time.Time, error) {
	var times []time.Time
	var err error

	htmlBody.Find(".table2").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		// Value is in the second child
		raw := s.Find("td").Eq(1).Text()
		raw = strings.Trim(raw, " \n")

		// Example time: 2023-09-07 01:00 CEST
		t, e := time.Parse("2006-01-02 15:04 MST", raw)
		if e != nil {
			fmt.Println("error parsing time: ", err)
			err = e
		}
		times = append(times, t)
	})

	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("error parsing times: %w", err)
	}

	return times[0], times[1], nil
}

func getContestStatus(startTime time.Time, endTime time.Time) KattisContestStatus {
	now := time.Now()

	if now.Before(startTime) {
		return ContestNotStarted
	} else if now.After(endTime) {
		return ContestEnded
	} else {
		return ContestRunning
	}
}

func getProblemId(problemUrl string) string {
	return path.Base(problemUrl)
}

func parseProblem(baseUrl *url.URL, tableRow *goquery.Selection) *KattisProblem {
	problem := &KattisProblem{}

	tableRow.Find("td").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			problem.Label = s.Text()
		case 1:
			problem.Name = s.Text()
			problem.Url, _ = s.Find("a").Attr("href")
		}
	})

	problem.ProblemId = getProblemId(problem.Url)
	problem.Url = baseUrl.JoinPath(problem.Url).String()

	return problem
}

func parseProblems(baseUrl *url.URL, htmlBody *goquery.Document) []*KattisProblem {
	var problems []*KattisProblem

	htmlBody.Find(".table2 > tbody").Eq(1).Find("tr").Each(func(i int, s *goquery.Selection) {
		problems = append(problems, parseProblem(baseUrl, s))
	})

	return problems
}

func GetContest(contestId string) (*KattisContest, error) {
	contest := &KattisContest{}
	baseUrl, err := createUrlBase(internal.Config.Kattis.Host)
	if err != nil {
		return nil, fmt.Errorf("error creating base url: %w", err)
	}
	contestUrl := makeContestUrl(baseUrl, contestId)
	res, err := makeRequest(contestUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	doc, err := parseHTML(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %w", err)
	}

	contest.ContestId = contestId
	contest.StartTime, contest.EndTime, err = parseTimes(doc)
	if err != nil {
		return nil, fmt.Errorf("error parsing times: %w", err)
	}

	contest.Status = getContestStatus(contest.StartTime, contest.EndTime)
	if contest.Status == ContestNotStarted {
		return contest, nil
	}

	contest.Problems = parseProblems(baseUrl, doc)

	return contest, nil
}
