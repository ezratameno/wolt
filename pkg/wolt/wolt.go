package wolt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
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

//GetOrders returns the orders.
func (c Wolt) GetOrders() (Orders, error) {
	var res Orders
	// go through every 100 orders
	for i := 0; ; i += 100 {
		url := fmt.Sprintf("https://restaurant-api.wolt.com/v2/order_details/?limit=100&skip=%s",
			strconv.Itoa(i))
		resp, err := c.request(http.MethodGet, url, nil)
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
		// TODO: fix the year format
		time := time.Unix(order.PaymentTime.Date, 0)
		fmt.Println("Time:", time.UTC())
		fmt.Println("Food Price:", order.ItemsPrice/100)
		fmt.Println("Delivery Price:", order.PaymentAmount/100-order.ItemsPrice/100)
		fmt.Println("Total Price:", order.PaymentAmount/100)
		counter += order.PaymentAmount / 100
	}
	fmt.Println("Total Orders value: ", counter)
	return nil
}

// request make a http request and returns the result.
func (c Wolt) request(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.do.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
