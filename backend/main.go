package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	Eng := gin.Default()
	SetRoutes(Eng)
	Eng.Run(":25790")
}

func SetRoutes(Eng *gin.Engine) {
	rawpic := Eng.Group("/raw")
	{
		rawpic.GET("/:pid", GetRawPic)
	}
}

// /1111.jpg
func GetRawPic(c *gin.Context) {
	pid := c.Param("pid")
	{
		pidsplit := strings.Split(pid, ".")
		fmt.Println(pidsplit)
	}
	fmt.Println(pid)
	f, err := os.Open("./illust_63093148_20210101_095841.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	fmt.Println(string(data))
	c.Writer.WriteString(string(data))
}
