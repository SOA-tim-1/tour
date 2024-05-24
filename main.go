package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/proto/checkpoint"
	"database-example/proto/coupon"
	"database-example/proto/equipment"
	"database-example/proto/tour"
	"database-example/repo"
	"database-example/service"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func startServerGRPC(
	tourHandler *handler.TourHandlergRPC,
	checkpointHandler *handler.CheckpointHandlergRPC,
	equpmentHandler *handler.EquipmentHandlergRPC,
	couponHandler *handler.CouponHandlergRPC,
) {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	// Bootstrap gRPC server.
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Bootstrap gRPC service server and respond to request.
	//userHandler := handlers.UserHandler{}
	tour.RegisterTourServiceServer(grpcServer, tourHandler)
	checkpoint.RegisterCheckpointServiceServer(grpcServer, checkpointHandler)
	equipment.RegisterEquipmentServiceServer(grpcServer, equpmentHandler)
	coupon.RegisterCouponServiceServer(grpcServer, couponHandler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
		println("Server starting")
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()

}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	checkpointRepo := &repo.CheckpointRepository{DatabaseConnection: database}
	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}

	equipmentRepo := &repo.AuthorEquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.AuthorEquipmentService{AuthorEquipmentRepo: equipmentRepo}

	couponRepo := &repo.CouponRepository{DatabaseConnection: database}
	couponService := &service.CouponService{CouponRepo: couponRepo}

	//studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	//studentService := &service.StudentService{StudentRepo: studentRepo}
	// studentHandler := &handler.StudentHandler{StudentService: studentService}
	// checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}
	// tourHandler := &handler.TourHandler{TourService: tourService, CheckpointService: checkpointService}
	// equipmentHandler := &handler.AuthorEquipmentHandler{AuthorEquipmentService: equipmentService}
	// couponHandler := &handler.CouponHandler{CouponService: couponService}

	// startServer(studentHandler, tourHandler, checkpointHandler, equipmentHandler, couponHandler)

	tourHandlerGRPC := &handler.TourHandlergRPC{TourService: tourService, CheckpointService: checkpointService}
	checkpointHandlerGRPC := &handler.CheckpointHandlergRPC{CheckpointService: checkpointService}
	equipmentHandlerGRPC := &handler.EquipmentHandlergRPC{AuthorEquipmentService: equipmentService}
	couponHandlerGRPC := &handler.CouponHandlergRPC{CouponService: couponService}

	startServerGRPC(tourHandlerGRPC, checkpointHandlerGRPC, equipmentHandlerGRPC, couponHandlerGRPC)

}
