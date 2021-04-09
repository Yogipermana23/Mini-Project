package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

//InitRouter is used for initiate router
func InitRouter() *gin.Engine {
	router := gin.New()
	// router := gin.uuDefault()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://bri360.bri.co.id", "http://127.0.0.1:8240"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(apmgin.Middleware(router))

	//routers

	router.POST("/bri360/summary-dashboard/churn", handlerChurn)
	router.POST("/bri360/summary-dashboard/fasitiltas-rekening", FasilitasRekening)
	router.POST("/bri360/summary-dashboard/segmentasi", Segmentasi)
	router.POST("/bri360/summary-dashboard/prediksi-rata-saldo", Ratassaldo)
	router.POST("/bri360/summary-dashboard/top-frekuensi-merchant", TopFrekuensiMerchant)
	router.POST("/bri360/summary-dashboard/top-nominal-merchant", TopNominalFrekuensiMerchant)
	router.POST("/bri360/summary-dashboard/posisi-saldo", PosisiSaldo)
	router.POST("/bri360/summary-dashboard/top-product", TopProduct)
	return router
}
