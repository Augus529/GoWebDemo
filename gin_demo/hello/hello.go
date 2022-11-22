package main

import "github.com/gin-gonic/gin"

/**
 * @Project GoWebDemo
 * @File    hello.go
 * @Author  Augus Lee
 * @Date    2022/11/9 15:51
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在http://localhost:8080启动服务
	r.Run()
}
