package routes

import (
	"server_go/api_warehouse_system/controllers"

	"github.com/gin-gonic/gin"
)

func Coredata(routes *gin.RouterGroup) {
	coredata := routes.Group("/coredata")
	{
		coredata.GET("/", controllers.Hello)                          // localhost:xxxx/coredata/
		coredata.GET("/positions/*act", controllers.Positions)        // localhost:xxxx/coredata/position
		coredata.GET("/getposition/:id_enc", controllers.GetPosition) // localhost:xxxx/coredata/getposition/xx

		coredata.GET("/branchs/*act", controllers.Branchs)        // localhost:xxxx/coredata/branchs
		coredata.GET("/getbranch/:id_enc", controllers.GetBranch) // localhost:xxxx/coredata/getbranch/xx

		coredata.GET("/categorys/*act", controllers.Categorys)        // localhost:xxxx/coredata/categorys
		coredata.GET("/getcategory/:id_enc", controllers.GetCategory) // localhost:xxxx/coredata/getcategory/xx

		coredata.GET("/suppliers/*act", controllers.Suppliers)        // localhost:xxxx/coredata/suppliers
		coredata.GET("/getsupplier/:id_enc", controllers.Getsupplier) // localhost:xxxx/coredata/getsupplier/xx

		coredata.GET("/products/*act", controllers.Products)        // localhost:xxxx/coredata/products
		coredata.GET("/getproduct/:id_enc", controllers.Getproduct) // localhost:xxxx/coredata/getproduct/xx

		// coredata.GET("/branchs/*catch-all", func(c *gin.Context) {
		// 	path := strings.Split(c.Request.RequestURI, "/")

		// 	fmt.Println(c.Request.RequestURI)
		// 	fmt.Println(path)
		// 	c.IndentedJSON(http.StatusOK, gin.H{
		// 		"message": "api sub folder routes",
		// 	})
		// })
	}
}
