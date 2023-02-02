package route

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, c *Controller) {

	r.GET("/blocks", c.getLatestNBlocks)
	r.GET("/blocks/:id", c.getBlockByNumber)
	r.GET("/transaction/:txHash", c.getTxByHash)

}
