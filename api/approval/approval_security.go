package approval

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

type ApprovalSecurity struct {
	AccessToken string
	Config      *Config
}

func NewApprovalSecurity(accessToken string, config *Config) *ApprovalSecurity {
	return &ApprovalSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

func (a *ApprovalSecurity) Run(chainId, address string) (*Result, error) {

	url := fmt.Sprintf(conf.Domain+"/api/v1/approval_security/%s?contract_addresses=%s", chainId, address)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if a.AccessToken != "" {
		request.Header.Add("Authorization", a.AccessToken)
	}

	if err != nil {
		return nil, err
	}

	if a.Config != nil && a.Config.Timeout > 0 {
		client.Timeout = time.Duration(a.Config.Timeout) * time.Second
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
