package spinwriter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var baseUrl = "https://www.spinwriter.com/action/api"

type Status string

const (
	OK    Status = "OK"
	ERROR Status = "ERROR"
)

type Resp struct {
	Status            Status     `json:"status"`
	RequestsMade      int        `json:"api_requests_made"`
	RequestsAvailable int        `json:"api_requests_available"`
	ProtectedTerms    []string   `json:"protected_terms"`
	ConfidenceLevel   Confidence `json:"confidence_level"`
	Data              string     `json:"response"`
}

type Client struct {
	HTTP   http.Client
	email  string
	apiKey string
}

func New(email string, apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		email:  email,
		HTTP:   http.Client{},
	}
}

func (c *Client) makeRequest(action string, options []Option) (string, error) {
	body := url.Values{}
	body.Set("email_address", c.email)
	body.Set("api_key", c.apiKey)
	body.Set("action", action)

	if options != nil {
		for _, option := range options {
			body.Set(option.Key, option.Value)
		}
	}

	resp, err := c.HTTP.Post(baseUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data Resp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Status != OK {
		return "", errors.New(data.Data)
	}

	return data.Data, nil
}

// returns the number of made and remaining API calls for the 24-hour period
func (c *Client) Quota() (string, error) {
	return c.makeRequest("api_quota", nil)
}

// returns the processed spun text with spintax
func (c *Client) Spintax(text string, options ...Option) (*Spintax, error) {
	raw, err := c.makeRequest("text_with_spintax", append(options, WithText(text)))
	if err != nil {
		return nil, err
	}

	return NewSpintax(raw, SpintaxFormat(optionValue(options, "spintax_format", SpintaxFormatPipeBraces.String()))), nil
}

// returns a unique variation of processed given text
func (c *Client) UniqueVariation(text string, options ...Option) (string, error) {
	return c.makeRequest("unique_variation", append(options, WithText(text)))
}

// returns a unique variation of already spun text
func (c *Client) UniqueSpintaxVariation(text string, options ...Option) (string, error) {
	return c.makeRequest("unique_variation_from_spintax", append(options, WithText(text)))
}
