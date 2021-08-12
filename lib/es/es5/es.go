package es5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

type hit struct {
	Source Metadata `json:"_source"`
}

type searchResult struct {
	Hits struct {
		Total int
		Hits  []hit
	}
}

func getMetadata(name string, versionId int) (meta Metadata, err error) {
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d/_source", os.Getenv("ES_SERVER"), name, versionId)
	result, err := http.Get(url)
	if err != nil {
		return
	}
	if result.StatusCode != http.StatusOK {
		err = fmt.Errorf("fail to get %s_%d:%d", name, versionId, result.StatusCode)
		return
	}
	result2, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(result2, &meta)
	return
}

func SearchLatestVersion(name string) (meta Metadata, err error) {
	url := fmt.Sprintf("http://%s//metadata/_search?q=name:%s&size=1&sort=version:desc", os.Getenv("ES_SERVER"), url2.PathEscape(name))
	result, err := http.Get(url)
	if err != nil {
		return
	}
	if result.StatusCode != http.StatusOK {
		err = fmt.Errorf("fail to search latest metadata:%s", result.StatusCode)
		return
	}
	result2, _ := ioutil.ReadAll(result.Body)
	var sr searchResult
	json.Unmarshal(result2, &sr)
	if len(sr.Hits.Hits) != 0 {
		meta = sr.Hits.Hits[0].Source
	}
	return
}

func GetMetadata(name string, version int) (Metadata, error) {
	if version == 0 {
		return SearchLatestVersion(name)
	}
	return getMetadata(name, version)
}

func PutMetadata(name string, version int, size int64, hash string) error {
	document := fmt.Sprintf(`{"name":"%s","version":%d,"size":%d,"hash":"%s"}`, name, version, size, hash)
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d?op_type=create", os.Getenv("ES_SERVER"), name, version)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(document))
	request.Header.Set("Content-Type", "application/json")
	result, err := client.Do(request)
	if err != nil {
		return err
	}
	if result.StatusCode == http.StatusConflict {
		return PutMetadata(name, version+1, size, hash)
	}
	if result.StatusCode != http.StatusCreated {
		result2, _ := ioutil.ReadAll(result.Body)
		return fmt.Errorf("fail to put metadata:%d %s", result.StatusCode, string(result2))
	}
	return nil
}

func AddVersion(name, hash string, size int64) error {
	version, err := SearchLatestVersion(name)
	if err != nil {
		return err
	}
	return PutMetadata(name, version.Version+1, size, hash)
}

func SearchAllVersions(name string, from, size int) ([]Metadata, error) {
	url := fmt.Sprintf("http://%s/metadata/_search?sort=name,version&from=%d&size=%d", os.Getenv("ES_SERVER"), from, size)
	if name != "" {
		url += "&q=name:" + name
	}
	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	metas := make([]Metadata, 0)
	result2, _ := ioutil.ReadAll(result.Body)
	var sr searchResult
	json.Unmarshal(result2, &sr)
	for i := range sr.Hits.Hits {
		metas = append(metas, sr.Hits.Hits[i].Source)
	}
	return metas, nil
}

func DelMetadata(name string, version int) {
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d", os.Getenv("ES_SERVER"), name, version)
	client := http.Client{}
	request, _ := http.NewRequest("DELETE", url, nil)
	client.Do(request)
}

type Bucket struct {
	Key        string
	DocCount   int
	MinVersion struct {
		Value float32
	}
}

type aggregateResult struct {
	Aggregations struct {
		GroupByName struct {
			Buckets []Bucket
		}
	}
}

func SearchVersionStatus(minDocCount int) ([]Bucket, error) {
	url := fmt.Sprintf("http://%s/metadata/_search", os.Getenv("ES_SERVER"))
	body := fmt.Sprintf(`
		{
			"size": 0,
			"aggs": {
				"group_by_name": {
					"terms": {
						"field": "name",
						"min_doc_count": %d
					},
					"aggs": {
						"min_version": {
							"min": {
								"field": "version"
							}
						}
					}
				}
			}
		}`, minDocCount)
	client := http.Client{}
	request, _ := http.NewRequest("GET", url, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	result, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, _ := ioutil.ReadAll(result.Body)
	var ar aggregateResult
	json.Unmarshal(responseBody, &ar)
	return ar.Aggregations.GroupByName.Buckets, nil
}

func HasHash(hash string) (bool, error) {
	url := fmt.Sprintf("http://%s/metadata/_search?q=hash:%s&size=0", os.Getenv("ES_SERVER"), hash)
	result, err := http.Get(url)
	if err != nil {
		return false, err
	}
	body, _ := ioutil.ReadAll(result.Body)
	var sr searchResult
	json.Unmarshal(body, &sr)
	return sr.Hits.Total != 0, nil
}

func SearchHashSize(hash string) (size int64, err error) {
	url := fmt.Sprintf("http://%s/metadata/_search?q=hash:%s&size=1", os.Getenv("ES_SERVER"), hash)
	result, err := http.Get(url)
	if err != nil {
		return
	}
	if result.StatusCode != http.StatusOK {
		err = fmt.Errorf("fail to search hash size:%d", result.StatusCode)
		return
	}
	body, _ := ioutil.ReadAll(result.Body)
	var sr searchResult
	json.Unmarshal(body, &sr)
	if len(sr.Hits.Hits) != 0 {
		size = sr.Hits.Hits[0].Source.Size
	}
	return
}
