package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJson := `[
		{
			"productId":1,
   		"manufacturer":"Johns-Jenkins",
   		"sku":"p5z343vdS",
   		"upc":"939581000000",
   		"pricePerUnit":"497.45",
   		"quantityOnHand":9703,
   		"productName":"sticky note"
		},
		{
			"productId":2,
   		"manufacturer":"Hessel, Schimmel and Feeney",
   		"sku":"i7v300kmx",
   		"upc":"740979000000",
   		"pricePerUnit":"282.29",
   		"quantityOnHand":9217,
   		"productName":"leg warmers"
		},
		{
			"productId":3,
   		"manufacturer":"Swaniawski, Bartoletti and Bruen",
   		"sku":"q0L657ys7",
   		"upc":"11173000000",
   		"pricePerUnit":"436.26",
   		"quantityOnHand":5905,
   		"productName":"lamp shade"
		}
	]`

	err := json.Unmarshal([]byte(productsJson), &productList)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(&productList)
}

func findProductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product, listItemIndex := findProductByID(productID)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJson)

	case http.MethodPut:
		var updateProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updateProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updateProduct.ProductID != productID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		product = &updateProduct
		productList[listItemIndex] = *product
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)

	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished; %s", time.Since(start))
	})
}

func main() {
	productListHandler := http.HandlerFunc(productsHandler)
	productItemHandler := http.HandlerFunc(productHandler)
	http.Handle("/products", middlewareHandler(productListHandler))
	http.Handle("/products/", middlewareHandler(productItemHandler))
	http.ListenAndServe(":5000", nil)
}

// func main() {

// 	product := &Product{
// 		ProductID:      123,
// 		Manufacturer:   "Big box company",
// 		Sku:            "4561Hxjy",
// 		Upc:            "788979798546",
// 		PricePerUnit:   "12.889",
// 		QuantityOnHand: 28,
// 		ProductName:    "Gizmo",
// 	}

// 	productJson, err := json.Marshal(product)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(productJson))
// }

// func main() {
// 	productJson := `{
// 		"productId":123,
// 		"manufacturer":"Big box company",
// 		"sku":"4561Hxjy",
// 		"upc":"788979798546",
// 		"pricePerUnit":"12.889",
// 		"quantityOnHand":28,
// 		"productName":"Gizmo"
// 	}`

// 	product := Product{}
// 	err := json.Unmarshal([]byte(productJson), &product)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(product.ProductName)
// }
