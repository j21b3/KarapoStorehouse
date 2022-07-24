package components

import (
	"KarapoStorehouse/model"
	"bytes"
	"fmt"

	"fyne.io/fyne/v2/canvas"
)

type ShowImage struct {
	Title    string `json:"title"`
	FileName string
	Uploader string `json:"uploader"`
	Message  string `json:"message"`

	Image  *canvas.Image
	Lables []string `json:"tags"`
}

type RetPic struct {
	Data   model.PicData `json:"data"`
	Status bool          `json:"status"`
	Err    string        `json:"err"`
}

// 将接口返回的image结构体转换为显示需要的ShowImage结构体
func NewShowImage(picdata *model.PicData) (*ShowImage, error) {

	showimage := ShowImage{
		Title:    picdata.Title,
		Message:  picdata.Message,
		FileName: picdata.FileName,
		Uploader: picdata.Uploader,
		Lables:   picdata.Tags,
	}

	readbuf := bytes.NewBuffer(picdata.Data)

	showimage.Image = canvas.NewImageFromReader(readbuf, picdata.FileName)
	if showimage.Image == nil {
		return nil, fmt.Errorf("bad reply data from server (newimagefromimage)")
	}
	showimage.Image.FillMode = canvas.ImageFillContain

	return &showimage, nil
}
