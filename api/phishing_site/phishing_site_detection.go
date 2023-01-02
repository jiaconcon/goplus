package phishing_site

import (
	"encoding/json"
	"fmt"
	"github.com/jiaconcon/goplus/conf"
	"net/http"
	"time"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type PhishingSiteDetection struct {
	AccessToken string
	Config      *Config
}

func NewPhishingSiteDetection(accessToken string, config *Config) *PhishingSiteDetection {
	return &PhishingSiteDetection{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

func (s *PhishingSiteDetection) Run(url string) (*Result, error) {

	fullUrl := fmt.Sprintf(conf.Domain+"/api/v1/phishing_site?url=%s", url)

	client := &http.Client{}
	request, err := http.NewRequest("GET", fullUrl, nil)
	if s.AccessToken != "" {
		request.Header.Add("Authorization", s.AccessToken)
	}

	if err != nil {
		return nil, err
	}

	if s.Config != nil && s.Config.Timeout > 0 {
		client.Timeout = time.Duration(s.Config.Timeout) * time.Second
	} else {
		client.Timeout = time.Duration(conf.Timeout) * time.Second
	}

	response, err := client.Do(request)

	var res Result

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
