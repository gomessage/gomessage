package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

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

func createRelease(version string) (string, error) {
	url := "https://api.github.com/repos/gomessage/gomessage/releases"

	data := releaseBody{
		Name:                 version,
		TagName:              version,
		TargetCommitish:      "master",
		Body:                 "",
		Draft:                false,
		Prerelease:           false,
		GenerateReleaseNotes: false,
	}

	payload, _ := json.Marshal(data)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Authorization", "Bearer xxxxx")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
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

func upload(versionNum, releaseId, pkgName string) {

	url := fmt.Sprintf("https://uploads.github.com/repos/gomessage/gomessage/releases/%s/assets?pkgName=%s", releaseId, pkgName)

	filePath := fmt.Sprintf("build/%s/%s", versionNum, pkgName)

	fmt.Println(url)
	fmt.Println(filePath)
	payload, _ := ioutil.ReadFile(filePath)
	res, err := http.Post(url, "application/gzip", bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Header.Add("Authorization", "Bearer xxxxx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	versionNum := string("2.0.6")

	mac := fmt.Sprintf("gomessage-%s-mac-x64.tar.gz", versionNum)
	linux := fmt.Sprintf("gomessage-%s-linux-x64.tar.gz", versionNum)
	windows := fmt.Sprintf("gomessage-%s-windows-x64.tar.gz", versionNum)

	//创建release
	//version := fmt.Sprintf("v%s", versionNum)
	//releaseId, _ := createRelease(version)

	//上传程序包
	releaseId := "95593344"
	for _, pkgName := range []string{mac, linux, windows} {
		upload(versionNum, releaseId, pkgName)
	}
}
