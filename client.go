package spinrewriter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var baseUrl = "https://www.spinrewriter.com/action/api"

type Status string

const (
	OK    Status = "OK"
	ERROR Status = "ERROR"
)

type Resp struct {
	Status          Status     `json:"status"`
	ReqsMade        int        `json:"api_requests_made"`
	ReqsAvailable   int        `json:"api_requests_available"`
	ProtectedTerms  string     `json:"protected_terms"`
	ConfidenceLevel Confidence `json:"confidence_level"`
	Data            string     `json:"response"`
}

type Quota struct {
	ReqsMade      int
	ReqsAvailable int
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

func (c *Client) makeRequest(action string, options []Option) (*Resp, error) {
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
		return nil, err
	}
	defer resp.Body.Close()

	var data Resp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Status != OK {
		return nil, errors.New(data.Data)
	}

	return &data, nil
}

// returns the number of made and remaining API calls for the 24-hour period
func (c *Client) Quota() (*Quota, error) {
	resp, err := c.makeRequest("api_quota", nil)
	if err != nil {
		return nil, err
	}

	return &Quota{
		ReqsMade:      resp.ReqsMade,
		ReqsAvailable: resp.ReqsAvailable,
	}, nil
}

// returns the processed spun text with spintax
func (c *Client) Spintax(text string, options ...Option) (*Spintax, error) {
	resp, err := c.makeRequest("text_with_spintax", append(options, withText(text)))
	if err != nil {
		return nil, err
	}

	return NewSpintax(resp.Data, SpintaxFormat(optionValue(options, "spintax_format", FormatPipeBraces.String())), resp.ReqsMade, resp.ReqsAvailable), nil
}

// returns a unique variation of processed given text
func (c *Client) UniqueVariation(text string, options ...Option) (string, error) {
	resp, err := c.makeRequest("unique_variation", append(options, withText(text)))
	if err != nil {
		return "", err
	}

	return resp.Data, nil
}

// returns a unique variation of already spun text
func (c *Client) UniqueSpintaxVariation(spintax *Spintax, options ...Option) (string, error) {
	resp, err := c.makeRequest("unique_variation_from_spintax", append(options, withText(spintax.String())))
	if err != nil {
		return "", err
	}

	return resp.Data, nil
}
