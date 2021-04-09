package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Churn Handler
func handlerChurn(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, total_page, total_row, err := request.gethandlerChurn()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "total_page": total_page, "total_row": total_row})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}
}

//Filitas rekening
func FasilitasRekening(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getFasilitasRekening()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}
}

//segmentasi

func Segmentasi(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getSegmentasi()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}
}

func Ratassaldo(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getRatassaldo()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}
}

// Posisi Saldo

func PosisiSaldo(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getPosisiSaldo()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}

}

// Top Frekuensu Merchant
func TopFrekuensiMerchant(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getTopFrekuensiMerchant()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}

}

//Top Nominal Frekuensi Merchant
func TopNominalFrekuensiMerchant(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getTopNominalFrekuensiMerchant()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}
}

func TopProduct(c *gin.Context) {

	var request RequestData

	if err := c.BindJSON(&request); err == nil {
		message, period_date, modified_date, err := request.getTopProduct()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "data": message, "period_date": period_date, "modified_date": modified_date})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Error", "Error Message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}

}
