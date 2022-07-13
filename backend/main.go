package main

import (
	"context"
	"net/http"
	"time"

	"backend.main/dblayer"

	"backend.main/model"

	"github.com/gin-gonic/gin"
)

var RPC *dblayer.RawPicDBController

func main() {
	Eng := gin.Default()

	RPC = dblayer.NewRawPicDBController("mongodb://localhost:27017")
	SetRoutes(Eng)
	Eng.Run(":25790")
}

func SetRoutes(Eng *gin.Engine) {
	rawpic := Eng.Group("/raw")
	{
		rawpic.GET("/:pid", GetRawPic)
	}
	Eng.POST("/upload", UploadRawPic)
}

// GET http://127.0.0.1:25790/raw/{picid} 获取原图片
func GetRawPic(c *gin.Context) {
	pidstr := c.Param("pid")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := RPC.FindPicById(ctx, pidstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ReturnData{
			Status: false,
			Err:    "ERROR Insert to Database",
			Data:   nil,
		})
		return
	}

	c.Writer.WriteString(string(data.Data))
}

// POST http://ip:25790/upload 上传原图片
func UploadRawPic(c *gin.Context) {
	data := model.PicData{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ReturnData{
			Status: false,
			Err:    "ERROR Upload",
			Data:   nil,
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := RPC.GenerateID()
	err := RPC.InsertPic(ctx, dblayer.RawPic{
		Id:         id,
		Title:      data.Title,
		FileName:   data.FileName,
		Data:       data.Data,
		Uploader:   data.Uploader,
		Message:    data.Message,
		CreateTime: time.Now(),
		Tags:       data.Tags,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ReturnData{
			Status: false,
			Err:    "ERROR Insert to Database",
			Data:   nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.ReturnData{
		Status: true,
	})
}
