package api_builder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"api-gateway/pkg/authmiddleware"
	"api-gateway/pkg/logger"
	"api-gateway/pkg/model"
)

type auth struct {
	builder *apiBuilder
}

func NewAuth(builder *apiBuilder) Auth {
	return &auth{builder: builder}
}

func (a *auth) Token(request model.AuthUser) (*authmiddleware.Tokens, error) {

	req, err := json.Marshal(request)
	if err != nil {
		logger.Errorf("Token.Marshal", err)
		return nil, err
	}

	resp, err := a.builder.postRequest(fmt.Sprint(a.builder.configPaths.AuthApiUrl, authToken), req)
	if err != nil {
		logger.Errorf("Token.postRequest", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Token.ReadAll", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var response authmiddleware.Tokens

	if err := json.Unmarshal(body, &response); err != nil {
		logger.Errorf("Token.Unmarshal", err)
		return nil, err
	}

	return &response, nil
}

func (a *auth) CheckToken(request model.Token) error {
	req, err := json.Marshal(request)
	if err != nil {
		logger.Errorf("CheckToken.Marshal", err)
		return err
	}

	resp, err := a.builder.postRequest(fmt.Sprint(a.builder.configPaths.AuthApiUrl, authCheckToken), req)
	if err != nil {
		logger.Errorf("CheckToken.postRequest", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("CheckToken.ReadAll", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	return nil
}
