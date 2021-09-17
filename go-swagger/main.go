package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go-swagger/docs"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Order struct{
	OrderId			string		`json:"order_id" example:"1"`
	CustomerName 	string		`json:"customer_name" example:"Ali""`
	OrderedAt		time.Time	`json:"ordered_at"`
	Items			Item		`json:"items"`
}

type Item struct {
	ItemId			string		`json:"item_id example:"ABC"`
	Description		string		`json:"description"`
	Quantity		int			`json:"quantity"`
}

// CreateOrder godoc
// @Summary Create an order
// @Description creater new order with input on payload
// @Tags orders
// @Produce json
// @Accept json
// @Param order body Order true "create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderId = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// @Summary get all order
// CreateOrder godoc
// @Description creater new order with input on payload
// @Tags orders
// @Produce json
// @Accept json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// @Summary get detail order
// CreateOrder godoc
// @Description creater new order with input on payload
// @Tags orders
// @Produce json
// @Accept json
// @Success 200 {object} Order
// @Router /orders/{orderID} [get]
// @Param orderId path int true "id order"
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderId == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}
//@Tags orders
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderId == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}
//@Tags orders
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderId == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

var orders []Order
var prevOrderID = 0



//@title kampus merdeka API Docs
//@version1.0
//@description
//@termsofService
func main(){
	router := mux.NewRouter()

	router.HandleFunc("/", func( w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")

	})

	router.HandleFunc("/orders", createOrder).Methods("POST")
	//router.HandleFunc("/orders", getOrder).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)


	log.Println("server running on part 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	//http.ListenAndServe(":8080", router)

}



