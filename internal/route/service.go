package route

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrBlockNumberShouldBeInteger = errors.New("block number should be an integer")

func (ctrl *Controller) getLatestNBlocks(c *gin.Context) {
	limit := c.Query("limit")
	if limit == "" {
		limit = "1"
	}

	n, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatal(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})

	}

	res, err := ctrl.RDS.GetLatestNBlocks(uint64(n))

	if err != nil {
		log.Fatal(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

func (ctrl *Controller) getBlockByNumber(c *gin.Context) {
	id := c.Param("id")

	number, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})

	}

	block, err := ctrl.RDS.GetBlockByNumber(uint64(number))

	if err != nil {
		log.Fatal(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"res": block,
	})
}

func (ctrl *Controller) getTxByHash(c *gin.Context) {
	txHash := c.Param("txHash")

	transaction, err := ctrl.RDS.GetTransactionByHash(txHash)

	if err != nil {
		log.Fatal(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"res": transaction,
	})
}
