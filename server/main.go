package main

import (
	"time"
	"udfiles/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	//cors "github.com/rs/cors/wrapper/gin"
)



func main() {
	router := gin.Default()

	//router.Use(cors.Default())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	//router.MaxMultipartMemory = 8 << 20  // 8 MiB
	//router.Use(cors.Default())

	router.Use(cors.New(cors.Config{
       			AllowOrigins:     []string{"*"},
        		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
        		AllowHeaders:     []string{"access-control-allow-headers","authorization", "expires", "pragma", "dnt","user-agent","x-requested-with","if-modified-since","cache-control","content-type","Range"},
			//AllowCredentials: true,
        		//AllowOriginFunc: func(origin string) bool {
            		//return origin == "*"
       	 		//},
        		MaxAge: 12 * time.Hour,
    	}))
	//Cargar archivo

	router.GET("/user/:name", handlers.Probando)

	router.POST("/upload/:name/:challenge", handlers.UploadSIF)

	//Descargar archivo
	router.GET("/download/:filename", handlers.DownloadSolv)

	router.Run("152.74.52.188:8080")
	//router.Run(":8080")

}
