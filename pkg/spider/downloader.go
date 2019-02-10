package spider

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func getHttpResponse(url string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("request error")
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("response error")
	}
	defer response.Body.Close()
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New(fmt.Sprintf("status code error : %d", response.StatusCode))
	}
	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}
	return ioutil.ReadAll(response.Body)
}
