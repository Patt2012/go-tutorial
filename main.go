package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	//"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	initTimeZone()
	initConfig()
	db := initDatabase()
	

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)
	_ = customerRepositoryDB
	_ = customerRepositoryMock

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	//log.Printf("Banking service started at port %v", viper.GetInt("app.port"))
	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

	/* customer, err := customerService.GetCustomer(2001)
	if err != nil {
		panic(err)
	} */

	//fmt.Println(customer)

	/* customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers) */

	/* _ = customerRepository

	customer, err := customerRepository.GetById(2000)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)  */

/* 	customers, err := customerRepository.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers) */

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}