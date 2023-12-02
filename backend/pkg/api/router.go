package api

import (
	"niya-voyage/backend/docs"
	"niya-voyage/backend/pkg/api/blogs"
	"niya-voyage/backend/pkg/api/books"
	"niya-voyage/backend/pkg/api/customers"
	"niya-voyage/backend/pkg/api/destinations"
	"niya-voyage/backend/pkg/api/hotels"
	"niya-voyage/backend/pkg/api/orderDetails"
	"niya-voyage/backend/pkg/api/orders"
	"niya-voyage/backend/pkg/api/payments"
	"niya-voyage/backend/pkg/api/products"
	"niya-voyage/backend/pkg/api/reviewRatings"
	"niya-voyage/backend/pkg/api/testimonials"
	"niya-voyage/backend/pkg/api/tourPackages"
	"niya-voyage/backend/pkg/auth"
	"niya-voyage/backend/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
		//r.Use(middleware.Xss())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{

		v1.GET("/books", middleware.APIKeyAuth(), books.FindBooks)
		v1.POST("/books", middleware.APIKeyAuth(), middleware.JWTAuth(), books.CreateBook)
		v1.GET("/books/:id", middleware.APIKeyAuth(), books.FindBook)
		v1.PUT("/books/:id", middleware.APIKeyAuth(), books.UpdateBook)
		v1.DELETE("/books/:id", middleware.APIKeyAuth(), books.DeleteBook)

		// Hotels routes
		v1.GET("/hotels", middleware.APIKeyAuth(), hotels.FindHotels)
		v1.POST("/hotels", middleware.APIKeyAuth(), middleware.JWTAuth(), hotels.CreateHotel)
		v1.GET("/hotels/:id", middleware.APIKeyAuth(), hotels.FindHotel)
		v1.PUT("/hotels/:id", middleware.APIKeyAuth(), hotels.UpdateHotel)
		v1.DELETE("/hotels/:id", middleware.APIKeyAuth(), hotels.DeleteHotel)

		// Destinations routes
		v1.GET("/destinations", middleware.APIKeyAuth(), destinations.FindDestinations)
		v1.POST("/destinations", middleware.APIKeyAuth(), middleware.JWTAuth(), destinations.CreateDestination)
		v1.GET("/destinations/:id", middleware.APIKeyAuth(), destinations.FindDestination)
		v1.PUT("/destinations/:id", middleware.APIKeyAuth(), destinations.UpdateDestination)
		v1.DELETE("/destinations/:id", middleware.APIKeyAuth(), destinations.DeleteDestination)

		// Testimonials routes
		v1.GET("/testimonials", middleware.APIKeyAuth(), testimonials.FindTestimonials)
		v1.POST("/testimonials", middleware.APIKeyAuth(), middleware.JWTAuth(), testimonials.CreateTestimonial)
		v1.GET("/testimonials/:id", middleware.APIKeyAuth(), testimonials.FindTestimonial)
		v1.PUT("/testimonials/:id", middleware.APIKeyAuth(), testimonials.UpdateTestimonial)
		v1.DELETE("/testimonials/:id", middleware.APIKeyAuth(), testimonials.DeleteTestimonial)

		// Blogs routes
		v1.GET("/blogs", middleware.APIKeyAuth(), blogs.FindBlogs)
		v1.POST("/blogs", middleware.APIKeyAuth(), middleware.JWTAuth(), blogs.CreateBlog)
		v1.GET("/blogs/:id", middleware.APIKeyAuth(), blogs.FindBlog)
		v1.PUT("/blogs/:id", middleware.APIKeyAuth(), blogs.UpdateBlog)
		v1.DELETE("/blogs/:id", middleware.APIKeyAuth(), blogs.DeleteBlog)

		//Login routes
		v1.POST("/login", middleware.APIKeyAuth(), auth.LoginHandler)
		v1.POST("/register", middleware.APIKeyAuth(), auth.RegisterHandler)

		//Cutomer routes
		v1.GET("/", customers.Healthcheck)
		v1.GET("/customers", middleware.APIKeyAuth(), customers.FindCustomers)
		v1.POST("/customers", middleware.APIKeyAuth(), middleware.JWTAuth(), customers.CreateCustomer)
		v1.GET("/customers/:id", middleware.APIKeyAuth(), customers.FindCustomer)
		v1.PUT("/customers/:id", middleware.APIKeyAuth(), customers.UpdateCustomer)
		v1.DELETE("/customers/:id", middleware.APIKeyAuth(), customers.DeleteCustomer)

		//product routes
		v1.GET("/products", middleware.APIKeyAuth(), products.FindProducts)
		v1.POST("/products", middleware.APIKeyAuth(), middleware.JWTAuth(), products.CreateProduct)
		v1.GET("/products/:id", middleware.APIKeyAuth(), products.FindProduct)
		v1.PUT("/products/:id", middleware.APIKeyAuth(), products.UpdateProduct)
		v1.DELETE("/products/:id", middleware.APIKeyAuth(), products.DeleteProduct)

		//order routes
		v1.GET("/orders", middleware.APIKeyAuth(), orders.FindOrders)
		v1.POST("/orders", middleware.APIKeyAuth(), middleware.JWTAuth(), orders.CreateOrder)
		v1.GET("/orders/:id", middleware.APIKeyAuth(), orders.FindOrder)
		v1.PUT("/orders/:id", middleware.APIKeyAuth(), orders.UpdateOrder)
		v1.DELETE("/orders/:id", middleware.APIKeyAuth(), orders.DeleteOrder)

		//orderDetail routes
		v1.GET("/orderDetails", middleware.APIKeyAuth(), orderDetails.FindOrderDetails)
		v1.POST("/orderDetails", middleware.APIKeyAuth(), middleware.JWTAuth(), orderDetails.CreateOrderDetail)
		v1.GET("/orderDetails/:id", middleware.APIKeyAuth(), orderDetails.FindOrderDetail)
		v1.PUT("/orderDetails/:id", middleware.APIKeyAuth(), orderDetails.UpdateOrderDetail)
		v1.DELETE("/orderDetails/:id", middleware.APIKeyAuth(), orderDetails.DeleteOrderDetail)

		// tourPackage routes
		v1.GET("/tourPackages", middleware.APIKeyAuth(), tourPackages.FindTourPackages)
		v1.POST("/tourPackages", middleware.APIKeyAuth(), middleware.JWTAuth(), tourPackages.CreateTourPackage)
		v1.GET("/tourPackages/:id", middleware.APIKeyAuth(), tourPackages.FindTourPackage)
		v1.PUT("/tourPackages/:id", middleware.APIKeyAuth(), tourPackages.UpdateTourPackage)
		v1.DELETE("/tourPackages/:id", middleware.APIKeyAuth(), tourPackages.DeleteTourPackage)

		// payment routes
		v1.GET("/payments", middleware.APIKeyAuth(), payments.FindPayments)
		v1.POST("/payments", middleware.APIKeyAuth(), middleware.JWTAuth(), payments.CreatePayment)
		v1.GET("/payments/:id", middleware.APIKeyAuth(), payments.FindPayment)
		v1.PUT("/payments/:id", middleware.APIKeyAuth(), payments.UpdatePayment)
		v1.DELETE("/payments/:id", middleware.APIKeyAuth(), payments.DeletePayment)

		// reviewRating routes
		v1.GET("/reviewRatings", middleware.APIKeyAuth(), reviewRatings.FindReviewRatings)
		v1.POST("/reviewRatings", middleware.APIKeyAuth(), middleware.JWTAuth(), reviewRatings.CreateReviewRating)
		v1.GET("/reviewRatings/:id", middleware.APIKeyAuth(), reviewRatings.FindReviewRating)
		v1.PUT("/reviewRatings/:id", middleware.APIKeyAuth(), reviewRatings.UpdateReviewRating)
		v1.DELETE("/reviewRatings/:id", middleware.APIKeyAuth(), reviewRatings.DeleteReviewRating)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
