package product

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/charlesguo404/example_go/database"
)

// 这里实际上是一个匿名结构体，同时初始化给了值，所以后面可以直接使用productMap.m

var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{
	m: make(map[int]Product),
}

func getProduct(productID int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result := database.DbConn.QueryRowContext(ctx, `select productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products
	WHERE productId = ?`, productID)
	product := &Product{}
	err := result.Scan(&product.ProductID,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return product, nil
}

func removeProduct(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM products WHERE productId = ?`, productID)
	if err != nil {
		return err
	}
	return nil
}

func getProductList() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `select productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products`)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	products := make([]Product, 0)

	for results.Next() {
		var product Product
		results.Scan(
			&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)

	}
	return products, nil
}

func updateProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `UPDATE products SET
	manufacturer=?,
	sku=?,
	upc=?,
	pricePerUnit=CAST(? AS DECIMAL(13,2)),
	quantityOnHand=?,
	productName=?
	WHERE productId=?`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)
	if err != nil {
		return err
	}
	return nil
}

func insertProduct(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO products 
	(manufacturer,
		sku,
		upc,
		pricePerUnit,
		quantityOnHand,
		productName) VALUES (?,?,?,?,?,?)`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName)

	if err != nil {
		return 0, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}