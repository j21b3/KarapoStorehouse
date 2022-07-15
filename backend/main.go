package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"backend.main/dblayer"

	"backend.main/model"
	"backend.main/tools"

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
		rawpic.GET("/thumbnail", GetThumbnailPic)
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
			Err:    "ERROR find Database",
			Data:   nil,
		})
		return
	}

	//TODO:后续需要调整为协议内容
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
	//TODO:后续需要校验data内容
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

// GET http://127.0.0.1:25790/thumbnail?pid=111&width=222 获取原图片
func GetThumbnailPic(c *gin.Context) {
	//固定生成的略缩图都为正方形
	form := struct {
		Pid   string `form:"pid" binding:"required"`
		Width string `form:"width"`
	}{}
	if err := c.Bind(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ReturnData{
			Status: false,
			Err:    "ERROR input param for thumbnail",
			Data:   nil,
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := RPC.FindPicById(ctx, form.Pid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ReturnData{
			Status: false,
			Err:    "ERROR find Database",
			Data:   nil,
		})
		return
	}

	var thumbnailData []byte
	if form.Width == "" {
		//默认为1000*1000的正方形
		thumbnailData, err = tools.GenerateThumbnail(data.Data, 1000, 1000)
	} else {
		var wid int64
		if wid, err = strconv.ParseInt(form.Width, 10, 32); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ReturnData{
				Status: false,
				Err:    "ERROR input param for thumbnail",
				Data:   nil,
			})
			return
		}
		thumbnailData, err = tools.GenerateThumbnail(data.Data, int(wid), int(wid))
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ReturnData{
			Status: false,
			Err:    "ERROR generate thumbnail picture",
			Data:   nil,
		})
		return
	}
	//TODO:后续需要调整为协议内容
	c.Writer.WriteString(string(thumbnailData))
}
