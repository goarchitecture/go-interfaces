// Package httpConfigSource - пакет хранящий источник конфигурирования для пакета configuration
// Этот пакет - пример того, что источником конфига может быть совершенно другой, сторонний пакет, который ничего "не знает"
// про пакет configuration
package httpConfigSource

import (
	"io/ioutil"
	"net/http"
)

const SourceTypeHttp = "http"

type HttpConfigSource struct {
	url string
}

func NewHttpConfigSource(url string) *HttpConfigSource {
	return &HttpConfigSource{url: url}
}

func (s *HttpConfigSource) Contents() ([]byte, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
