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

type releaseBody struct {
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

// 创建一个发行版
func createRelease(tag string) (string, error) {
	url := "https://api.github.com/repos/gomessage/gomessage/releases"

	data := releaseBody{
		Name:                 tag,
		TagName:              tag,
		TargetCommitish:      "master",
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
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	rsp := releaseResponse{}
	json.Unmarshal(body, &rsp)

	fmt.Println(&rsp)

	return strconv.Itoa(rsp.ID), nil
}

// 上传软件包
func upload(version, releaseId, pkgName string) {
	url := fmt.Sprintf("https://uploads.github.com/repos/gomessage/gomessage/releases/%s/assets?name=%s", releaseId, pkgName)
	//这里是相对路径，后面如有需要，记得变动
	filePath, _ := os.Getwd()
	filePath += fmt.Sprintf("/build/%s/%s", version, pkgName)

	fmt.Println(filePath)
	payload, _ := os.ReadFile(filePath)
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Host", "uploads.github.com")
	req.Header.Add(os.Getenv("Github_Authorization"), os.Getenv("Github_Token"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body2, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body2))

}

// 按照标签查询发行版
func getReleaseByTag(tag string) (string, bool) {
	url := fmt.Sprintf("https://api.github.com/repos/gomessage/gomessage/releases/tags/%s", tag)

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
	} else {
		return "发行版不存在，查询不到release_id", false
	}

}

func baseDelete(url string) {
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
		fmt.Println("删除成功...")
	} else {
		fmt.Println("删除失败...")
	}
}

// 按照id删除发行版
func deleteReleaseById(releaseId string) {
	url := fmt.Sprintf("https://api.github.com/repos/gomessage/gomessage/releases/%s", releaseId)
	baseDelete(url)
}

// 删除指定tag
func deleteTag(tag string) {
	url := fmt.Sprintf("https://api.github.com/repos/gomessage/gomessage/git/refs/tags/%s", tag)
	baseDelete(url)
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
	version := *versionParam           //纯数字格式：2.0.1
	tag := fmt.Sprintf("v%s", version) //格式：v2.0.1

	mac := fmt.Sprintf("gomessage-%s-mac-x64.tar.gz", version)
	linux := fmt.Sprintf("gomessage-%s-linux-x64.tar.gz", version)
	windows := fmt.Sprintf("gomessage-%s-windows-x64.tar.gz", version)

	releaseId, ok := getReleaseByTag(tag)
	if !ok {
		//如果发行版不存在，则直接创建一个新的release并上传包
		newReleaseId, _ := createRelease(tag)
		fmt.Println(newReleaseId)
		for _, pkgName := range []string{mac, linux, windows} {
			upload(version, newReleaseId, pkgName)
		}

	} else {
		//如果发行版存在，则先删除原来的，再重新创建
		deleteReleaseById(releaseId)
		deleteTag(tag)

		newReleaseId, _ := createRelease(tag)
		fmt.Println(newReleaseId)
		for _, pkgName := range []string{mac, linux, windows} {
			upload(version, newReleaseId, pkgName)
		}
	}

	//releaseId := "95650756"
	//for _, pkgName := range []string{mac} {
	//	upload(version, releaseId, pkgName)
	//}

}
