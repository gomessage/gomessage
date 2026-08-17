package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type ResponseReleaseByTag struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
	Body       string `json:"body"`
}

type releaseRequest struct {
	TagName              string `json:"tag_name"`
	TargetCommitish      string `json:"target_commitish"`
	Name                 string `json:"name"`
	Body                 string `json:"body"`
	Draft                bool   `json:"draft"`
	Prerelease           bool   `json:"prerelease"`
	GenerateReleaseNotes bool   `json:"generate_release_notes"`
}

type releaseResponse struct {
	URL             string        `json:"url"`
	AssetsURL       string        `json:"assets_url"`
	UploadURL       string        `json:"upload_url"`
	HTMLURL         string        `json:"html_url"`
	ID              int           `json:"id"`
	NodeID          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       time.Time     `json:"created_at"`
	PublishedAt     time.Time     `json:"published_at"`
	Assets          []interface{} `json:"assets"`
	TarballURL      string        `json:"tarball_url"`
	ZipballURL      string        `json:"zipball_url"`
	Body            string        `json:"body"`
}

func createRelease(owner, repo, branch, tag string) (string, error) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/releases"

	data := releaseRequest{
		Name:                 tag,
		TagName:              tag,
		TargetCommitish:      branch,
		Body:                 "",
		Draft:                false,
		Prerelease:           false,
		GenerateReleaseNotes: false,
	}

	payload, _ := json.Marshal(data)

	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add(os.Getenv("Github_Authorization"), os.Getenv("Github_Token"))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("createRelease 失败: %s\n%s", res.Status, string(body))
	}
	rsp := releaseResponse{}
	json.Unmarshal(body, &rsp)

	fmt.Println(">> Release创建成功，ID:", rsp.ID)

	return strconv.Itoa(rsp.ID), nil
}

func upload(owner, repo, version, releaseId, pkgName string) {
	url := fmt.Sprintf("https://uploads.github.com/repos/%s/%s/releases/%s/assets?name=%s", owner, repo, releaseId, pkgName)

	filePath, _ := os.Getwd()
	filePath += fmt.Sprintf("/build/%s/%s", version, pkgName)
	payload, _ := os.ReadFile(filePath)

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add(os.Getenv("Github_Authorization"), os.Getenv("Github_Token"))
	req.Header.Add("Host", "uploads.github.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/gzip")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body2, _ := ioutil.ReadAll(res.Body)
	fmt.Println("\n", string(body2))
	fmt.Println(res.Status)
}

func getReleaseByTag(owner, repo, tag string) (string, bool) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/tags/%s", owner, repo, tag)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add(os.Getenv("Github_Authorization"), os.Getenv("Github_Token"))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		r := ResponseReleaseByTag{}
		json.Unmarshal(body, &r)
		fmt.Println(&r)
		fmt.Println(r.ID)
		return strconv.Itoa(r.ID), true
	}
	return "发行版不存在，查询不到release_id", false
}

func baseDelete(delName, url string) {
	client := &http.Client{}
	method := "DELETE"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add(os.Getenv("Github_Authorization"), os.Getenv("Github_Token"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 204 {
		fmt.Println(delName + "删除成功...")
	} else {
		fmt.Println(delName + "删除失败...")
	}
}

func deleteReleaseById(owner, repo, releaseId string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/%s", owner, repo, releaseId)
	baseDelete("release", url)
}

func deleteTag(owner, repo, tag string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/tags/%s", owner, repo, tag)
	baseDelete("tag", url)
}

var versionParam *string

func init() {
	versionParam = pflag.StringP("version", "v", "", "指定上传文件的版本")
	pflag.Parse()

	if len(*versionParam) == 0 {
		log.Println("version参数不能为空~")
		os.Exit(3)
	}
}

func main() {
	version := *versionParam
	tag := fmt.Sprintf("v%s", version)

	owner := "gomessage"
	repo := "gomessage"
	branch := "main"

	mac := fmt.Sprintf("gomessage-%s-mac-arm64.tar.gz", version)
	linux := fmt.Sprintf("gomessage-%s-linux-amd64.tar.gz", version)
	linuxArm := fmt.Sprintf("gomessage-%s-linux-arm64.tar.gz", version)
	windows := fmt.Sprintf("gomessage-%s-windows-amd64.tar.gz", version)

	packageList := []string{mac, linux, linuxArm, windows}

	releaseId, ok := getReleaseByTag(owner, repo, tag)

	var wg sync.WaitGroup
	wg.Add(len(packageList))

	uploadAll := func(newReleaseId string) {
		for _, pkgName := range packageList {
			go func(pkgName string) {
				upload(owner, repo, version, newReleaseId, pkgName)
				wg.Done()
			}(pkgName)
		}
	}

	if ok {
		// 只删Release不删tag：tag仍在时createRelease会复用现有tag，
		// 避免重复发布时tag被重新指向main最新HEAD而破坏版本锁定
		deleteReleaseById(owner, repo, releaseId)
	}

	newReleaseId, err := createRelease(owner, repo, branch, tag)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	uploadAll(newReleaseId)

	wg.Wait()
	fmt.Println(">> 文件上传完成：", packageList)
}
