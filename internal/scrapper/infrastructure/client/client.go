package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)


type Client struct {
	baseURL string
	http    *http.Client
}


func New(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client {
		baseURL: strings.TrimRight(baseURL, "/"),
		http: httpClient,
	}
} 


func (c *Client) RegisterChat(ctx context.Context, chatID int64) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/tg-chat/%d", c.baseURL, chatID),
		nil,
	)

	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return parseErrorResponse(resp)
}


func (c *Client) DeleteChat(ctx context.Context, chatID int64) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/tg-chat/%d", c.baseURL, chatID),
		nil,
	)

	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return parseErrorResponse(resp)
}




func (c *Client) GetLinks(ctx context.Context, chatID int64) ([]domain.Link, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/links", c.baseURL),
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Tg-Chat-Id", strconv.FormatInt(chatID, 10))
	
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, parseErrorResponse(resp)
	}

	var dto ListLinksResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return nil, err
	}

	return toDomainLinks(dto.Links), nil
}


func (c *Client) AddLink(ctx context.Context, chatID int64, link string, tags, filters []string) (domain.Link, error) {
	bodyDTO := AddLinkRequestDTO {
		Link: link,
		Tags: tags,
		Filters: filters,
	}

	body, err := json.Marshal(bodyDTO) //превращаем в []byte формат для считывания из JSON
	if err != nil {
		return domain.Link{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/links", c.baseURL),
		bytes.NewReader(body),
	)
	if err != nil {
		return domain.Link{}, err 
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tg-Chat-Id", strconv.FormatInt(chatID, 10))

	resp, err := c.http.Do(req)
	if err != nil {
		return domain.Link{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Link{}, parseErrorResponse(resp)
	}

	var dto LinkResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return domain.Link{}, err
	}

	return toDomainLink(dto), nil
}


func (c *Client) RemoveLink(ctx context.Context, chatID int64, link string) (domain.Link, error) {
	bodyDTO := RemoveLinkRequestDTO{
		Link: link,
	}

	body, err := json.Marshal(bodyDTO)
	if err != nil {
		return domain.Link{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/links", c.baseURL),
		bytes.NewReader(body),
	)
	if err != nil {
		return domain.Link{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tg-Chat-Id", strconv.FormatInt(chatID, 10))

	resp, err := c.http.Do(req)
	if err != nil {
		return domain.Link{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Link{}, parseErrorResponse(resp)
	}

	var dto LinkResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return domain.Link{}, err
	}

	return toDomainLink(dto), nil
}