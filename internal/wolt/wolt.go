// wolt handles the wolt rest api.
package wolt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	h "github.com/ezratameno/wolt/pkg/http"
)

type Wolt struct {
	do    *http.Client
	token string
}

func New(token string) *Wolt {
	return &Wolt{
		do:    http.DefaultClient,
		token: token,
	}
}

// GetOrders returns the orders.
func (c Wolt) GetOrders() (Orders, error) {
	var res Orders
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + c.token
	// go through every 100 orders
	for i := 0; ; i += 100 {
		url := fmt.Sprintf("https://restaurant-api.wolt.com/v2/order_details/?limit=100&skip=%s",
			strconv.Itoa(i))
		resp, err := h.Request(http.MethodGet, url, nil, headers)
		if err != nil {
			return nil, err
		}
		var data Orders
		// break when there are no more orders.
		// 3 - the resp is empty
		if len(resp) == 3 {
			break
		}
		err = json.Unmarshal(resp, &data)
		if err != nil {
			return nil, err
		}
		res = append(res, data...)
	}

	return res, nil
}

// GetSummary return the summry of every month
func (c Wolt) GetSummary() error {
	orders, err := c.GetOrders()
	if err != nil {
		return err
	}
	var counter int
	for _, order := range orders {
		paymentTime := time.UnixMilli(order.PaymentTime.Date)
		fmt.Println("Time:", paymentTime.Format("02-01-2006 15:04:05"))
		fmt.Println("Food Price:", order.ItemsPrice/100)
		fmt.Println("Delivery Price:", order.PaymentAmount/100-order.ItemsPrice/100)
		fmt.Println("Total Price:", order.PaymentAmount/100)
		counter += order.PaymentAmount / 100
	}
	fmt.Println("Total Orders value: ", counter)
	return nil
}
