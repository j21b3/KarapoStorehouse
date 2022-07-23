package trans

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"KarapoStorehouse/model"
)

var RawPicURL = "raw/%s"

type Backend struct {
	Ip   string
	Port int

	Crypto string
	Key    []byte
}

func NewBackend(ip string, port int) *Backend {
	return &Backend{
		Ip:   ip,
		Port: port,
	}
}

// 获取原图片的接口封装
func (b *Backend) GetRawPic(idhex string) (*model.PicData, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/"+RawPicURL, b.Ip, b.Port, idhex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Bad response from server: %d", resp.StatusCode))
	}

	var buf []byte
	if _, err := resp.Body.Read(buf); err != nil {
		return nil, errors.New(fmt.Sprintf("Bad reply data from server (Body Read)"))
	}
	var picdata model.PicData
	if err := json.Unmarshal(buf, &picdata); err != nil {
		return nil, errors.New(fmt.Sprintf("Bad reply data from server (Json Unmarshal)"))
	}
	return &picdata, nil
}
