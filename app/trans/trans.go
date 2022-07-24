package trans

import (
	"KarapoStorehouse/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var RawPicURL = "raw/%s"

type Backend struct {
	Ip   string
	Port int

	Crypto string
	Key    []byte
}

type ShowImage struct {
	Title    string
	FileName string
	Uploader int
	Message  string

	Image  *canvas.Image
	Lables map[string]*widget.Label
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

// 将接口返回的结构体转换为显示需要的ShowImage结构体
func NewShowImage(retpic *RetPic) (*ShowImage, error) {
	picdata := retpic.Data

	showimage := ShowImage{
		Title:    picdata.Title,
		Message:  picdata.Message,
		FileName: picdata.FileName,
		Uploader: picdata.Uploader,
		Lables:   make(map[string]*widget.Label),
	}

	readbuf := bytes.NewBuffer(picdata.Data)

	showimage.Image = canvas.NewImageFromReader(readbuf, picdata.FileName)
	if showimage.Image == nil {
		return nil, fmt.Errorf("bad reply data from server (newimagefromimage)")
	}
	showimage.Image.FillMode = canvas.ImageFillOriginal

	for _, each := range picdata.Tags {
		showimage.Lables[each] = widget.NewLabel(each)
	}
	return &showimage, nil
}

// 获取原图片的接口封装
func (b *Backend) GetRawPic(idhex string) (*ShowImage, error) {
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

	return NewShowImage(&retpic)
}
