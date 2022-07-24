package components

import (
	"KarapoStorehouse/app/trans"
	"fmt"
	"reflect"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ImagePage struct {
	layout *fyne.Layout
}

func NewForm(showImage *ShowImage) *widget.Form {
	form := widget.NewForm()
	/* tag_hbox := container.NewHBox() */

	myType := reflect.TypeOf(showImage).Elem()
	myValue := reflect.ValueOf(showImage).Elem()

	for i := 0; i < myType.NumField(); i++ {
		fld := myType.Field(i)
		tag := fld.Tag.Get("json")

		switch tag {
		case "":
			// 若是为空则为不显示内容
		case "tags":
			tags := []string{}
			for j := 0; j < myValue.Field(i).Len(); j++ {
				v := myValue.Field(i).Index(j)
				/* tag_hbox.Add(canvas.NewText(, color.RGBA{255, 255, 255, 255})) */
				tags = append(tags, v.String())
			}
			form.Append("tags:", widget.NewLabel(strings.Join(tags, ",")))

		default:
			form.Append(tag+":", widget.NewLabel(myValue.Field(i).String()))
		}
	}

	return form
}

func GetFromBackend(PicID string) *ShowImage {
	backend := trans.NewBackend("127.0.0.1", 25790)
	picdata, err := backend.GetRawPic(PicID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	showImage, err := NewShowImage(picdata)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return showImage
}

func NewImagePage(PicID string) fyne.CanvasObject {

	showImage := GetFromBackend(PicID)
	if showImage == nil {
		return nil
	}

	form := NewForm(showImage)
	//adapt := container.NewAdaptiveGrid(5,showImage.Image )

	windowGroup := container.NewBorder(
		nil, form, nil, nil,
		showImage.Image,
	)
	return windowGroup
}
