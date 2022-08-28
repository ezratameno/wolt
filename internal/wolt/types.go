package wolt

type Orders []struct {
	AutomaticRejectionTime struct {
		Date int64 `json:"$date"`
	} `json:"automatic_rejection_time"`
	ClientPreEstimate string `json:"client_pre_estimate"`
	Credits           int    `json:"credits"`
	Currency          string `json:"currency"`
	DeliveryBasePrice int    `json:"delivery_base_price"`
	DeliveryComment   string `json:"delivery_comment"`
	DeliveryDistance  int    `json:"delivery_distance"`
	DeliveryEta       struct {
		Date int64 `json:"$date"`
	} `json:"delivery_eta,omitempty"`
	DeliveryLocation struct {
		Address     string `json:"address"`
		Alias       string `json:"alias"`
		Apartment   string `json:"apartment"`
		City        string `json:"city"`
		Coordinates struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"coordinates"`
		Street string `json:"street"`
	} `json:"delivery_location"`
	DeliveryMethod     string `json:"delivery_method"`
	DeliveryPrice      int    `json:"delivery_price"`
	DeliveryPriceShare int    `json:"delivery_price_share"`
	DeliveryTime       struct {
		Date int64 `json:"$date"`
	} `json:"delivery_time,omitempty"`
	DriverType      string        `json:"driver_type"`
	IsHostPaying    bool          `json:"is_host_paying"`
	IsMarketplaceV2 bool          `json:"is_marketplace_v2"`
	ItemChangeLog   []interface{} `json:"item_change_log"`
	Items           []struct {
		Count     int    `json:"count"`
		EndAmount int    `json:"end_amount"`
		ID        string `json:"id"`
		Name      string `json:"name"`
		Options   []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Values []struct {
				Count int    `json:"count"`
				ID    string `json:"id"`
				Name  string `json:"name"`
				Price int    `json:"price"`
			} `json:"values"`
		} `json:"options"`
		Price                int  `json:"price"`
		RowNumber            int  `json:"row_number"`
		SkipOnRefill         bool `json:"skip_on_refill"`
		SubstitutionSettings struct {
			AllowedItems []interface{} `json:"allowed_items"`
			IsAllowed    bool          `json:"is_allowed"`
		} `json:"substitution_settings"`
	} `json:"items"`
	ItemsPrice          int           `json:"items_price"`
	ListImage           string        `json:"list_image"`
	ListImageBlurhash   string        `json:"list_image_blurhash"`
	MainImage           string        `json:"main_image"`
	MainImageBlurhash   string        `json:"main_image_blurhash"`
	OrderAdjustmentRows []interface{} `json:"order_adjustment_rows"`
	OrderID             string        `json:"order_id"`
	OrderNumber         string        `json:"order_number"`
	PaymentAmount       int           `json:"payment_amount"`
	PaymentMethod       struct {
		ID       string `json:"id"`
		Provider string `json:"provider"`
		Type     string `json:"type"`
	} `json:"payment_method"`
	PaymentName string `json:"payment_name"`
	PaymentTime struct {
		Date int64 `json:"$date"`
	} `json:"payment_time"`
	Status                         string    `json:"status"`
	Subscribed                     bool      `json:"subscribed"`
	Subtotal                       int       `json:"subtotal"`
	Tip                            int       `json:"tip"`
	TipShare                       int       `json:"tip_share"`
	Tokens                         int       `json:"tokens"`
	TotalPrice                     int       `json:"total_price"`
	TotalPriceShare                int       `json:"total_price_share"`
	VenueAddress                   string    `json:"venue_address"`
	VenueCoordinates               []float64 `json:"venue_coordinates"`
	VenueCountry                   string    `json:"venue_country"`
	VenueFullAddress               string    `json:"venue_full_address"`
	VenueID                        string    `json:"venue_id"`
	VenueName                      string    `json:"venue_name"`
	VenueOpen                      bool      `json:"venue_open"`
	VenueOpenOnPurchase            bool      `json:"venue_open_on_purchase"`
	VenuePhone                     string    `json:"venue_phone"`
	VenueProductLine               string    `json:"venue_product_line"`
	VenueTimezone                  string    `json:"venue_timezone"`
	VenueURL                       string    `json:"venue_url"`
	DeliveryDistanceSurcharge      int       `json:"delivery_distance_surcharge,omitempty"`
	DeliverySizeSurcharge          int       `json:"delivery_size_surcharge,omitempty"`
	DeliverySizeSurchargeWithoutSf int       `json:"delivery_size_surcharge_without_sf,omitempty"`
	RejectionReason                struct {
		Message       string `json:"message"`
		RejectionKey  int    `json:"rejection_key"`
		RejectionType string `json:"rejection_type"`
	} `json:"rejection_reason,omitempty"`
	PreorderStatus string `json:"preorder_status,omitempty"`
	PreorderTime   struct {
		Date int64 `json:"$date"`
	} `json:"preorder_time,omitempty"`
	CancellableStatus struct {
		Reason    string `json:"reason"`
		ShowTimer bool   `json:"show_timer"`
		Start     struct {
			Date int64 `json:"$date"`
		} `json:"start"`
		Until struct {
			Date int64 `json:"$date"`
		} `json:"until"`
	} `json:"cancellable_status,omitempty"`
}
