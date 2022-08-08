package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"KarapoStorehouse/backend/dblayer"

	"KarapoStorehouse/model"

	"KarapoStorehouse/tools"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 现版本先只用一个会话连接数据库，后续更新用户管理后每个用户都分配一个
var RPC *dblayer.RawPicDBController

func main() {
	Eng := gin.Default()

	RPC = dblayer.NewRawPicDBController("mongodb://localhost:27017")
	SetMiddleware(Eng)
	SetRoutes(Eng)
	Eng.Run(":25790")
}

func SetRoutes(Eng *gin.Engine) {
	rawpic := Eng.Group("/raw")
	{
		rawpic.GET("/:pid", GetRawPic)
		rawpic.GET("/thumbnail", GetThumbnailPic)
		rawpic.GET("/detail", GetPicDetail)
	}
	Eng.POST("/upload", UploadRawPic)
	Eng.GET("/timeline/:page", GetTimeline)
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
	// c.JSON(
	// 	http.StatusOK,
	// 	model.ReturnData{
	// 		Status: true,
	// 		Data:   data.PicData,
	// 	})
}

// 获取图片详细信息接口 GET http://127.0.0.1:25790/raw/detail?pid=123
func GetPicDetail(c *gin.Context) {
	form := struct {
		Pid string `form:"pid" binding:"required"`
	}{}
	if err := c.Bind(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ReturnData{
			Status: false,
			Err:    "ERROR input param for raw detail",
			Data:   nil,
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := RPC.FindDetailById(ctx, form.Pid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ReturnData{
			Status: false,
			Err:    "ERROR find Database",
			Data:   nil,
		})
		return
	}
	ret_data := struct {
		Title      string    `json:"title" bson:"title"`
		FileName   string    `json:"file_name" binding:"required" bson:"file_name"`
		Uploader   string    `json:"uploader" binding:"required" bson:"uploader"`
		Message    string    `json:"message" bson:"message"`
		Tags       []string  `json:"tags" bson:"tags"`
		CreateTime time.Time `json:"update_time" bson:"create_time"`
	}{
		Title:      data.Title,
		FileName:   data.FileName,
		Uploader:   data.Uploader,
		Message:    data.Message,
		Tags:       data.Tags,
		CreateTime: data.CreateTime,
	}

	c.JSON(
		http.StatusOK,
		model.ReturnData{
			Status: true,
			Data:   ret_data,
		})
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
		CreateTime: time.Now(),
		PicData: model.PicData{
			Title:    data.Title,
			FileName: data.FileName,
			Data:     data.Data,
			Uploader: data.Uploader,
			Message:  data.Message,
			Tags:     data.Tags,
		},
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

// GET http://127.0.0.1:25790/raw/thumbnail?pid=111&width=222 获取原图片
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
		//默认为360*360的正方形
		thumbnailData, err = tools.GenerateThumbnail(data.Data, 360, 360)
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
	//c.Writer.WriteString(string(thumbnailData))
	c.JSON(http.StatusOK, model.ReturnData{
		Status: true,
		Data:   thumbnailData,
	})
}

// GET http://127.0.0.1:25790/timeline/(0,int64max)
func GetTimeline(c *gin.Context) {
	pageStr := c.Param("page")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ReturnData{
			Status: false,
			Err:    "ERROR input param for timeline",
			Data:   nil,
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ret, err := RPC.GetTimelineID(ctx, page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ReturnData{
			Status: false,
			Err:    "ERROR generate timeline",
			Data:   nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.ReturnData{
		Status: true,
		Data:   ret,
	})
}

func SetMiddleware(Eng *gin.Engine) {
	// 解决跨域访问，default允许所有的origins，后续视情况修改
	// Eng.Use(cors.Default())

	webOrigin := []string{"http://127.0.0.1:3000", "http://127.0.0.1:8848", "http://localhost:3000"}

	conf := cors.Config{
		AllowOrigins: webOrigin,
		AllowMethods: []string{"GET", "POST", "UPDATE", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Access-Token"},
		MaxAge:       6 * time.Hour,
	}

	Eng.Use(cors.New(conf))
}
