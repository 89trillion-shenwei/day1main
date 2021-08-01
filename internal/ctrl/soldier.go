package ctrl

import (
	"day1/internal"
	"day1/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindSoldierByRaUnCvApi(c *gin.Context) (string, error) {
	ra := c.Query("Rarity")
	if ra == "" {
		return "", internal.IsEmptyError("稀有度不能为空")
	}
	un := c.Query("UnlockArena")
	if un == "" {
		return "", internal.IsEmptyError("未解锁阶段不能为空")
	}
	cv := c.Query("Cvc")
	if cv == "" {
		return "", internal.IsEmptyError("cvc不能为空")
	}
	//c.String(http.StatusOK, handler.FindSoldierByRaUnCv(ra, un, cv))\
	return handler.FindSoldierByRaUnCv(ra, un, cv)
}

func FindSoldierRaByIdApi(c *gin.Context) (string, error) {
	id := c.Query("Id")
	if id == "" {
		return "", internal.IsEmptyError("id不能为空")
	}
	//c.String(http.StatusOK, "稀有度:"+handler.FindSoldierRaById(id))
	return handler.FindSoldierRaById(id), nil
}

func FindSoldierCoByIdApi(c *gin.Context) (string, error) {
	id := c.Query("Id")
	if id == "" {
		return "", internal.IsEmptyError("id不能为空")
	}
	//c.String(http.StatusOK, "战力："+handler.FindSoldierCoById(id))
	return handler.FindSoldierCoById(id), nil
}

func FindSoldierByCvApi(c *gin.Context) (string, error) {
	cvc := c.Query("Cvc")
	if cvc == "" {
		return "", internal.IsEmptyError("cvc不能为空")
	}
	//c.String(http.StatusOK, handler.FindSoldierByCv(cvc))
	return handler.FindSoldierByCv(cvc)
}

func FindSoldierByUnApi(c *gin.Context) (string, error) {
	//c.String(http.StatusOK, handler.FindSoldierByUn())
	return handler.FindSoldierByUn()
}

type Api func(c *gin.Context) (string, error)

func ReturnData(api Api) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := api(c)
		if err != nil {
			globalError := err.(internal.GlobalError)
			c.JSON(globalError.Status, globalError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"查询结果": data,
		})
	}
}
