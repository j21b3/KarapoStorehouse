package main

import (
	"KarapoStorehouse/app/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myapp := app.New()
	// a.SetIcon()

	w := myapp.NewWindow("图片显示测试")

	/* f, err := os.Open("D:/MyProject/gopath/src/KarapoStorehouse/backend/illust_63093148_20210101_095841.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	readbuf := bytes.NewBuffer(data)
	imgd, _, err := image.Decode(readbuf)
	if err != nil {
		fmt.Println("Bad reply data from server (image.Decode)")
	}

	img := canvas.NewImageFromImage(imgd)
	img.FillMode = canvas.ImageFillOriginal
	*/

	page := components.NewImagePage("62dbd0d283d5641480925801")
	w.Resize(fyne.Size{Width: 1000, Height: 1000})
	w.SetContent(page)

	w.ShowAndRun()
}
