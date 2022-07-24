package components

import (
	"KarapoStorehouse/model"
	"bytes"
	"fmt"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type ShowImage struct {
	Title    string
	FileName string
	Uploader int
	Message  string

	Image  *canvas.Image
	Lables map[string]*widget.Label
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
