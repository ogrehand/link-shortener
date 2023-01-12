package middleware

// func CorsHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		startTime := time.Now()

// 		// before request

// 		c.Next()
// 		var listErr []string
// 		for _, ginErr := range c.Errors {
// 			listErr = append(listErr, ginErr.Error())
// 		}
// 		fmt.Println(c.Get("key1"))
// 		model.SaveLog(c, startTime, time.Now(), "bebas", listErr)
// 	}
// }

// r.Use(cors.New(cors.Config{
// 	AllowOrigins:     []string{"https://foo.com"},
// 	AllowMethods:     []string{"*"},
// 	AllowHeaders:     []string{"Origin"},
// 	ExposeHeaders:    []string{"Content-Length"},
// 	AllowCredentials: true,
// 	AllowOriginFunc: func(origin string) bool {
// 		return origin == "https://github.com"
// 	},
// 	MaxAge: 12 * time.Hour,
// }))
