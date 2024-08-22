package api_builder

import (
	"bytes"
	"net/http"
	"net/url"

	"api-gateway/pkg/config"
	"api-gateway/pkg/logger"
)

type apiBuilder struct {
	configPaths config.ServicesPath

	auth      Auth
	resources Resources
}

func New(configPaths config.ServicesPath) InternalAPI {
	return &apiBuilder{configPaths: configPaths}
}

func (a *apiBuilder) postRequest(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		logger.Errorf("postRequest.NewRequest", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("postRequest.Do", err)
		return nil, err
	}

	return resp, nil
}

func (a *apiBuilder) getRequest(link string, requestParams map[string]string) (*http.Response, error) {
	params := url.Values{}
	for key, val := range requestParams {
		params.Add(key, val)
	}
	req, err := http.NewRequest("GET", link+"?"+params.Encode(), bytes.NewBuffer([]byte{}))
	if err != nil {
		logger.Errorf("getRequest.NewRequest", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("getRequest.Do", err)
		return nil, err
	}

	return resp, nil
}

func (a *apiBuilder) Auth() Auth {
	if a.auth == nil {
		a.auth = NewAuth(a)
	}

	return a.auth
}

func (a *apiBuilder) Resources() Resources {
	if a.resources == nil {
		a.resources = NewResources(a)
	}

	return a.resources
}
