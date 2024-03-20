package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	dsn := "user=postgres password=super dbname=soa host=localhost port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Person{})
	database.AutoMigrate(&model.Student{})
	database.AutoMigrate(&model.TravelTimeAndMethod{})
	database.AutoMigrate(&model.Checkpoint{})
	database.AutoMigrate(&model.Tour{})
	database.AutoMigrate(&model.Equipment{})
	database.AutoMigrate(&model.Coupon{})

	err = database.AutoMigrate(&model.Person{}, &model.Student{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startServer(
	handler *handler.StudentHandler,
	tourHandler *handler.TourHandler,
	checkpointHandler *handler.CheckpointHandler,
	equipmentHandler *handler.AuthorEquipmentHandler,
	couponHandler *handler.CouponHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/students", handler.Create).Methods("POST")

	router.HandleFunc("/tour/{id}", tourHandler.Get).Methods("GET")
	router.HandleFunc("/tour/authortours/{authorId}", tourHandler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/tour", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tour/updatetour", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tour/publish", tourHandler.PublishTour).Methods("PUT")
	router.HandleFunc("/tour/archive", tourHandler.ArchiveTour).Methods("PUT")

	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Get).Methods("GET")
	router.HandleFunc("/checkpoint/tour/{tourId}", checkpointHandler.GetByTourId).Methods("GET")
	router.HandleFunc("/checkpoint", checkpointHandler.Create).Methods("POST")
	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Delete).Methods("DELETE")

	router.HandleFunc("/author/equipment", equipmentHandler.GetAll).Methods("GET")

	router.HandleFunc("/coupon/add-coupon", couponHandler.CreateCoupon).Methods("POST")
	router.HandleFunc("/coupon/delete-coupon", couponHandler.DeleteCoupon).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8090",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router)))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	studentService := &service.StudentService{StudentRepo: studentRepo}
	studentHandler := &handler.StudentHandler{StudentService: studentService}

	checkpointRepo := &repo.CheckpointRepository{DatabaseConnection: database}
	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo}
	checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourhandler := &handler.TourHandler{TourService: tourService, CheckpointService: checkpointService}

	equipmentRepo := &repo.AuthorEquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.AuthorEquipmentService{AuthorEquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.AuthorEquipmentHandler{AuthorEquipmentService: equipmentService}

	couponRepo := &repo.CouponRepository{DatabaseConnection: database}
	couponService := &service.CouponService{CouponRepo: couponRepo}
	couponHandler := &handler.CouponHandler{CouponService: couponService}

	startServer(studentHandler, tourhandler, checkpointHandler, equipmentHandler, couponHandler)
}
