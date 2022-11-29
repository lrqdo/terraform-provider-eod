package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (c *Client) Delete(ID string) error {
	url := fmt.Sprintf("%s/%s", c.url, ID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = c.http.Do(req)
	return err
}

func (c *Client) Create(retentionDays int) (*DB, error) {
	body := fmt.Sprintf(`{"keep-days": "%d"}`, retentionDays)
	jsonBody := []byte(body)

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, errors.New("Unable to post:" + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	postRes, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	db, err := readBody(postRes.Body)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *Client) Read(ID string) (*DB, error) {
	response, err := c.http.Get(fmt.Sprintf("%s/%s", c.url, ID))
	if err != nil {
		return nil, errors.New("Unable to get" + err.Error())
	}
	if response.StatusCode == 404 {
		return nil, NotFound(ID)
	}
	db, err := readBody(response.Body)
	if err != nil {
		return nil, err
	}
	return db, err
}

func readBody(body io.ReadCloser) (*DB, error) {
	resBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New("Unable to readAll" + err.Error())
	}
	var db DB
	err = json.Unmarshal(resBody, &db)
	if err != nil {
		return nil, errors.New("Unable to Unmarshal: " + err.Error() + "[" + string(resBody) + "]")
	}
	return &db, nil
}

type DB struct {
	ID        string `json:"Id"`
	Name      string `json:"Name"`
	ExpiresAt string `json:"Expires_at"`
	Host      string `json:"Host"`
	Port      string `json:"Port"`
	Status    string `json:"Status"`
}

type NotFound string

func (db NotFound) Error() string {
	return fmt.Sprintf("Env %s not found", db)
}
