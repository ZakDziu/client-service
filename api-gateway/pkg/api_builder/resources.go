package api_builder

import (
	"api-gateway/pkg/logger"
	"api-gateway/pkg/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type resources struct {
	builder *apiBuilder
}

func NewResources(builder *apiBuilder) Resources {
	return &resources{builder: builder}
}

func (a *resources) Users(requestParams map[string]string) (*map[int]model.User, error) {
	resp, err := a.builder.getRequest(a.builder.configPaths.ResourcesApiUtl+resourcesUsers, requestParams)
	if err != nil {
		logger.Errorf("Users.getRequest", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Users.ReadAll", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var response map[int]model.User

	if err := json.Unmarshal(body, &response); err != nil {
		logger.Errorf("Users.Unmarshal", err)
		return nil, err
	}

	return &response, nil
}

func (a *resources) Books(requestParams map[string]string) (*map[int]model.Book, error) {
	resp, err := a.builder.getRequest(a.builder.configPaths.ResourcesApiUtl+resourcesBooks, requestParams)
	if err != nil {
		logger.Errorf("Books.getRequest", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Books.ReadAll", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var response map[int]model.Book

	if err := json.Unmarshal(body, &response); err != nil {
		logger.Errorf("Books.Unmarshal", err)
		return nil, err
	}

	return &response, nil
}
