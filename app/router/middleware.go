package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)
//获取响应处理时间间隔
func RunTime(c *gin.Context) {
	currentTime := time.Now().Nanosecond()
	log.Println(fmt.Sprintf("开始%v#", currentTime))
	c.Next()
	currentTime2 := time.Now().Nanosecond()
	log.Println(fmt.Sprintf("结束%v#", currentTime2))
}
