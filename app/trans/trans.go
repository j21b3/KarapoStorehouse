package trans

import (
	"KarapoStorehouse/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

type RetPic struct {
	Data   model.PicData `json:"data"`
	Status bool          `json:"status"`
	Err    string        `json:"err"`
}

// 获取原图片的接口封装
func (b *Backend) GetRawPic(idhex string) (*model.PicData, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/"+RawPicURL, b.Ip, b.Port, idhex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response from server: %d", resp.StatusCode)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var retpic RetPic

	// 检查json转换、返回数据状态是否正确
	if err := json.Unmarshal(buf, &retpic); err != nil {
		return nil, err
	}

	if !retpic.Status {
		return nil, fmt.Errorf("return error, status is false")
	}

	return &retpic.Data, nil
}
