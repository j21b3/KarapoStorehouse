package main

import (
	"KarapoStorehouse/app/trans"
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {
	myapp := app.New()
	// a.SetIcon()

	w := myapp.NewWindow("图片显示测试")

	backend := trans.NewBackend("127.0.0.1", 25790)
	img, err := backend.GetRawPic("62dbd0d283d5641480925801")
	if err != nil {
		fmt.Println(err)
		return
	}

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
	w.SetContent(img.Image)

	w.ShowAndRun()
}
