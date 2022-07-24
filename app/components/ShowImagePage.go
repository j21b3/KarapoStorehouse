package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ImagePage struct {
	layout *fyne.Layout
}

func NewImagePage(showImage *ShowImage) fyne.CanvasObject {
	form := widget.NewForm()
	for key, val := range showImage.Lables {
		form.Append(key, val)
	}
	//adapt := container.NewAdaptiveGrid(5,showImage.Image )
	windowGroup := container.NewBorder(
		form,
		nil, nil, nil,
		showImage.Image,
	)
	return windowGroup
}
