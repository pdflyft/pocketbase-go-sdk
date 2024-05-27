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
		Meta *MetaConfig `json:"meta"`
		Smtp *SmtpConfig `json:"smtp"`
	}

	SmtpConfig struct {
		Enabled    bool   `json:"enabled"`
		Host       string `host:"host"`
		Port       int    `json:"port"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Tls        bool   `json:"tls"`
		AuthMethod string `json:"authMethod"`
		LocalName  string `json:"localName"`
	}

	MetaConfig struct {
		AppName                    string        `json:"appName"`
		AppUrl                     string        `json:"appUrl"`
		SenderName                 string        `json:"senderName"`
		SenderAddress              string        `json:"senderAddress"`
		VerificationTemplate       EmailTemplate `json:"verificationTemplate"`
		ResetPasswordTemplate      EmailTemplate `json:"resetPasswordTemplate"`
		ConfirmEmailChangeTemplate EmailTemplate `json:"confirmEmailChangeTemplate"`
	}

	EmailTemplate struct {
		Body    string `json:"body"`
		Subject string `json:"subject"`
	}
)

// All returns all settings.
func (s Settings) All() (ResponseSettingsAll, error) {
	var response ResponseSettingsAll
	if err := s.Authorize(); err != nil {
		return response, err
	}

	request := s.client.R().
		SetHeader("Content-Type", "application/json")

	resp, err := request.Get(s.url + "/api/settings")
	if err != nil {
		return response, fmt.Errorf("[settings] can't send get settings request to pocketbase, err %w", err)
	}

	if resp.IsError() {
		return response, fmt.Errorf("[settings] pocketbase returned status at getting all settings: %d, msg: %s, err %w",
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

func (s Settings) Update(body any) (ResponseSettingsAll, error) {
	var response ResponseSettingsAll

	if err := s.Authorize(); err != nil {
		return response, err
	}

	request := s.client.R().
		SetHeader("Content-Type", "application/json")
	request = request.SetBody(body)

	resp, err := request.Patch(s.url + "/api/settings")
	if err != nil {
		return response, fmt.Errorf("[settings] can't send update settings request to pocketbase, err %w", err)
	}

	if resp.IsError() {
		return response, fmt.Errorf("[settings] pocketbase returned status at updating settings: %d, msg: %s, err %w",
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
