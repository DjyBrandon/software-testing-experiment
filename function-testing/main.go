package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"time"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	fmt.Println(db)
	fmt.Println(err)

	sqLDB, err := db.DB()
	// SetMaxIdLeConns 设置空闲连接池中连接的最大数量
	sqLDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqLDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqLDB.SetConnMaxLifetime(10 * time.Second) //10秒钟

	// 结构体
	type List struct {
		gorm.Model
		Name    string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
		State   string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
		Phone   string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
		Email   string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
		Address string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
	}

	db.AutoMigrate(&List{})

	// 接口
	r := gin.Default()

	// 测试
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Request Success",
		})
	})

	// 增
	r.POST("/user/add", func(c *gin.Context) {
		var data List

		err := c.ShouldBindJSON(&data)

		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "Add Fail",
				"data": gin.H{},
				"code": 400,
			})
		} else {
			db.Create(&data)

			c.JSON(200, gin.H{
				"msg":  "Add Success",
				"data": data,
				"code": 200,
			})
		}
	})

	// 删
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		var data []List

		id := c.Param("id")

		db.Where("id = ?", id).Find(&data)

		if len(data) == 0 {
			c.JSON(200, gin.H{
				"msg":  "ID Not Found, Delete Fail",
				"code": 400,
			})
		} else {
			db.Where("id = ?", id).Delete(&data)

			c.JSON(200, gin.H{
				"msg":  "Delete Success",
				"code": 200,
			})
		}
	})

	// 改
	r.PUT("/user/update/:id", func(c *gin.Context) {

		var data List

		id := c.Param("id")

		db.Select("id").Where("id = ?", id).Find(&data)

		if data.ID == 0 {
			c.JSON(200, gin.H{
				"msg":  "User id Not Found",
				"code": 400,
			})
		} else {
			err := c.ShouldBindJSON(&data)

			if err != nil {
				c.JSON(200, gin.H{
					"msg":  "Update Fail",
					"code": 400,
				})
			} else {
				db.Where("id = ?", id).Updates(&data)

				c.JSON(200, gin.H{
					"msg":  "Update Success",
					"code": 200,
				})
			}
		}

	})

	// 查

	// 条件查询
	r.GET("/user/list/:name", func(c *gin.Context) {

		name := c.Param("name")

		var dataList []List

		db.Where("name = ?", name).Find(&dataList)

		if len(dataList) == 0 {
			c.JSON(200, gin.H{
				"msg":  "Not found data",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "Query Success",
				"code": 400,
				"data": dataList,
			})
		}
	})

	// 全部查询
	r.GET("/user/list", func(c *gin.Context) {

		var dataList []List

		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		pageNum, _ := strconv.Atoi(c.Query("pageNum"))

		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = -1
		}

		offsetVal := (pageNum - 1) * pageSize
		if pageNum == -1 && pageSize == -1 {
			offsetVal = -1
		}

		var total int64

		db.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)

		if len(dataList) == 0 {
			c.JSON(200, gin.H{
				"msg":  "Not Found Data",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "Query Success",
				"code": 200,
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
			})
		}
	})

	PORT := "3001"
	r.Run(":" + PORT)
}
