package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Categories []struct {
		Name        string   `yaml:"name"`
		Description string   `yaml:"desc"`
		Repos       []string `yaml:"repos"`
	}
}

func main() {
	f, err := os.Open("data.yaml")
	if err != nil {
		panic("open file: " + err.Error())
	}
	defer f.Close()

	var readme Data

	err = yaml.NewDecoder(f).Decode(&readme)
	if err != nil {
		panic("decoding yaml: " + err.Error())
	}

	myrepos, err := githubRepos()
	if err != nil {
		panic("loading github repos: " + err.Error())
	}

	fn := template.FuncMap{
		"repo": func(name string) repo {
			return search(myrepos, name)
		},
		"transform": transformString,
	}

	t, err := template.New("README.md").Funcs(fn).ParseFiles("README.md")
	if err != nil {
		panic("parsing template: " + err.Error())
	}

	rm, err := os.Create("../README.md")
	if err != nil {
		panic("create readme file: " + err.Error())
	}
	defer rm.Close()

	_, err = fmt.Fprintln(rm, "<!-- DO NOT EDIT. File is auto-generated -->")
	if err != nil {
		panic("writing to file: " + err.Error())
	}

	err = t.Execute(rm, readme)
	if err != nil {
		panic("executing template: " + err.Error())
	}
}

type repo struct {
	Name string `json:"name"`
	Link string `json:"html_url"`
	Desc string `json:"description"`

	Stars int `json:"stargazers_count"`
	Forks int `json:"forks_count"`
}

func (r repo) Title() string {
	switch {
	case r.Forks > 0 && r.Stars > 0:
		return more(r.Stars, "star") + ", " + more(r.Forks, "fork")
	case r.Forks == 0 && r.Stars > 0:
		return more(r.Stars, "star")
	case r.Stars == 0 && r.Forks > 0:
		return more(r.Forks, "fork")
	default:
		return ""
	}
}

func more(count int, s string) string {
	if count == 1 {
		return fmt.Sprintf("one %s", s)
	} else {
		return fmt.Sprintf("%d %ss", count, s)
	}
}

func githubRepos() (repos []repo, err error) {
	hc := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/xarantolus/repos?per_page=100", nil)
	if err != nil {
		return
	}

	resp, err := hc.Do(req)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&repos)

	return
}

func search(repos []repo, name string) repo {
	name = strings.ToLower(name)
	for _, r := range repos {
		if strings.ToLower(r.Name) == name {
			return r
		}
	}

	panic("supplied invalid repo name " + name)
}

var urlRegex = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

func transformString(s string) string {
	s = strings.TrimSpace(s)

	u := urlRegex.FindString(s)

	if u == "" {
		return s
	}

	var split = strings.Split(u, "/")
	if len(split) < 2 {
		return s
	}

	var linkText = split[len(split)-1]

	return strings.ReplaceAll(s, u, fmt.Sprintf("[%s](%s)", linkText, u))
}
