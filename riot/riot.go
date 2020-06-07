package riot

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type riot struct {
	client *http.Client
	apiKey string
}

func New() (*riot, error) {
	client := &http.Client{}

	key, err := getDotEnvVariable("api_key")
	if err != nil {
		return nil, err
	}

	return &riot{client, key}, nil
}

func (r *riot) sendRequest(url string, data Parser) (int, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return http.StatusNotFound, err
	}
	req.Header.Add("X-riot-Token", r.apiKey)

	resp, err := r.client.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	if err := data.parse(respBody); err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, err
}

func getDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
