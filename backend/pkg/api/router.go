package api

import (
	"niya-voyage/backend/docs"
	"niya-voyage/backend/pkg/api/activities"
	"niya-voyage/backend/pkg/api/addBanners"
	"niya-voyage/backend/pkg/api/blogs"
	"niya-voyage/backend/pkg/api/cars"
	"niya-voyage/backend/pkg/api/changePasses"
	"niya-voyage/backend/pkg/api/customers"
	"niya-voyage/backend/pkg/api/destinations"
	"niya-voyage/backend/pkg/api/flights"
	"niya-voyage/backend/pkg/api/golfs"
	"niya-voyage/backend/pkg/api/hotels"
	"niya-voyage/backend/pkg/api/locationInfos"
	"niya-voyage/backend/pkg/api/merchants"
	"niya-voyage/backend/pkg/api/orderDetails"
	"niya-voyage/backend/pkg/api/orders"
	"niya-voyage/backend/pkg/api/payments"
	"niya-voyage/backend/pkg/api/products"
	"niya-voyage/backend/pkg/api/rentals"
	"niya-voyage/backend/pkg/api/reviewRatings"
	"niya-voyage/backend/pkg/api/testimonials"
	"niya-voyage/backend/pkg/api/tours"
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
	// r.Use(middleware.Cors())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 600)) // 60 requests per minute

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
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

		//merchant routes
		v1.GET("/merchants", middleware.APIKeyAuth(), merchants.FindMerchants)
		v1.POST("/merchants", middleware.APIKeyAuth(), middleware.JWTAuth(), merchants.CreateMerchant)
		v1.GET("/merchants/:id", middleware.APIKeyAuth(), merchants.FindMerchant)
		v1.PUT("/merchants/:id", middleware.APIKeyAuth(), merchants.UpdateMerchant)
		v1.DELETE("/merchants/:id", middleware.APIKeyAuth(), merchants.DeleteMerchant)

		//location routes
		v1.GET("/locationInfos", middleware.APIKeyAuth(), locationInfos.FindLocations)
		v1.POST("/locationInfos", middleware.APIKeyAuth(), middleware.JWTAuth(), locationInfos.CreateLocation)
		v1.GET("/locationInfos/:id", middleware.APIKeyAuth(), locationInfos.FindLocation)
		v1.PUT("/locationInfos/:id", middleware.APIKeyAuth(), locationInfos.UpdateLocation)
		v1.DELETE("/locationInfos/:id", middleware.APIKeyAuth(), locationInfos.DeleteLocation)

		//password routes
		v1.GET("/changePasses", middleware.APIKeyAuth(), changePasses.FindPasswords)
		v1.POST("/changePasses", middleware.APIKeyAuth(), middleware.JWTAuth(), changePasses.CreatePassword)
		v1.GET("/changePasses/:id", middleware.APIKeyAuth(), changePasses.FindPassword)
		v1.PUT("/changePasses/:id", middleware.APIKeyAuth(), changePasses.UpdatePassword)
		v1.DELETE("/changePasses/:id", middleware.APIKeyAuth(), changePasses.DeletePassword)

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

		// activities routes
		v1.GET("/activities", middleware.APIKeyAuth(), activities.FindActivities)
		v1.POST("/activities", middleware.APIKeyAuth(), middleware.JWTAuth(), activities.CreateActivity)
		v1.GET("/activities/:id", middleware.APIKeyAuth(), activities.FindActivity)
		v1.PUT("/activities/:id", middleware.APIKeyAuth(), activities.UpdateActivity)
		v1.DELETE("/activities/:id", middleware.APIKeyAuth(), activities.DeleteActivity)

		// tour routes
		v1.GET("/tours", middleware.APIKeyAuth(), tours.FindTours)
		v1.POST("/tours", middleware.APIKeyAuth(), middleware.JWTAuth(), tours.CreateTour)
		v1.GET("/tours/:id", middleware.APIKeyAuth(), tours.FindTour)
		v1.PUT("/tours/:id", middleware.APIKeyAuth(), tours.UpdateTour)
		v1.DELETE("/tours/:id", middleware.APIKeyAuth(), tours.DeleteTour)

		// rental routes
		v1.GET("/rentals", middleware.APIKeyAuth(), rentals.FindRentals)
		v1.POST("/rentals", middleware.APIKeyAuth(), middleware.JWTAuth(), rentals.CreateRental)
		v1.GET("/rentals/:id", middleware.APIKeyAuth(), rentals.FindRental)
		v1.PUT("/rentals/:id", middleware.APIKeyAuth(), rentals.UpdateRental)
		v1.DELETE("/rentals/:id", middleware.APIKeyAuth(), rentals.DeleteRental)

		// car routes
		v1.GET("/cars", middleware.APIKeyAuth(), cars.FindCars)
		v1.POST("/cars", middleware.APIKeyAuth(), middleware.JWTAuth(), cars.CreateCar)
		v1.GET("/cars/:id", middleware.APIKeyAuth(), cars.FindCar)
		v1.PUT("/cars/:id", middleware.APIKeyAuth(), cars.UpdateCar)
		v1.DELETE("/cars/:id", middleware.APIKeyAuth(), cars.DeleteCar)

		// golf routes
		v1.GET("/golfs", middleware.APIKeyAuth(), golfs.FindGolfs)
		v1.POST("/golfs", middleware.APIKeyAuth(), middleware.JWTAuth(), golfs.CreateGolf)
		v1.GET("/golfs/:id", middleware.APIKeyAuth(), golfs.FindGolf)
		v1.PUT("/golfs/:id", middleware.APIKeyAuth(), golfs.UpdateGolf)
		v1.DELETE("/golfs/:id", middleware.APIKeyAuth(), golfs.DeleteGolf)

		// flight routes
		v1.GET("/flights", middleware.APIKeyAuth(), flights.FindFlights)
		v1.POST("/flights", middleware.APIKeyAuth(), middleware.JWTAuth(), flights.CreateFlight)
		v1.GET("/flights/:id", middleware.APIKeyAuth(), flights.FindFlight)
		v1.PUT("/flights/:id", middleware.APIKeyAuth(), flights.UpdateFlight)
		v1.DELETE("/flights/:id", middleware.APIKeyAuth(), flights.DeleteFlight)

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

		// addBanner routes
		v1.GET("/addBanners", middleware.APIKeyAuth(), addBanners.FindAddBanners)
		v1.POST("/addBanners", middleware.APIKeyAuth(), middleware.JWTAuth(), addBanners.CreateAddBanner)
		v1.GET("/addBanners/:id", middleware.APIKeyAuth(), addBanners.FindAddBanner)
		v1.PUT("/addBanners/:id", middleware.APIKeyAuth(), addBanners.UpdateAddBanner)
		v1.DELETE("/addBanners/:id", middleware.APIKeyAuth(), addBanners.DeleteAddBanner)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
