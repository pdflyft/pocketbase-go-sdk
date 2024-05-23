package pocketbase

import (
	"encoding/json"
	"fmt"
)

type (
	Settings struct {
		*Client
	}

	ResponseSettingsAll struct {
		Smtp *Smtp `json:"smtp"`
	}

	Smtp struct {
		Enabled    bool   `json:"enabled"`
		Host       string `host:"host"`
		Port       int    `json:"port"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		TLS        bool   `json:"tls"`
		AuthMethod string `json:"authMethod"`
		LocalName  string `json:"localName"`
	}
)

// All returns all settings.
func (b Settings) All() (ResponseSettingsAll, error) {
	var response ResponseSettingsAll
	if err := b.Authorize(); err != nil {
		return response, err
	}

	request := b.client.R().
		SetHeader("Content-Type", "application/json")

	resp, err := request.Get(b.url + "/api/settings")
	if err != nil {
		return response, fmt.Errorf("[settings] can't send all request to pocketbase, err %w", err)
	}

	if resp.IsError() {
		return response, fmt.Errorf("[settings] pocketbase returned status: %d, msg: %s, err %w",
			resp.StatusCode(),
			resp.String(),
			ErrInvalidResponse,
		)
	}

	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return response, fmt.Errorf("[settings] can't unmarshal response, err %w", err)
	}

	return response, nil
}
