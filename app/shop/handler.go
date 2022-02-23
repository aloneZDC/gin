package shop

import (
	_ "gin/config"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func postHandler(e *gin.Context) {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	var (
		result gin.H
	)
	if sum == 0 {
		result = gin.H{
			"data": nil,
			"msg":  -1,
		}
	} else {
		result = gin.H{
			"data": sum,
			"msg":  1,
		}
	}
	e.JSON(http.StatusOK, result)
}

// 求元素和
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func getValue(e *gin.Context) {
	b := [5]int{1, 3, 5, 8, 7}

	var (
		result gin.H
		value  bool
	)
	value = myTest(b, 8)
	if value == false {
		result = gin.H{
			"data": nil,
			"msg":  -1,
		}
	} else {
		result = gin.H{
			"data": value,
			"msg":  1,
		}
	}
	e.JSON(http.StatusOK, result)
}

// 求元素和，是给定的值
func myTest(a [5]int, target int) bool {
	var result = false
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				result = true
				return result
			}
		}
	}
	return result
}
