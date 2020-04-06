package expo

type PushReceipt struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

type PushReceiptResponse struct {
	Data   map[string]PushReceipt `json:"data"`
	Errors []map[string]string    `json:"errors"`
}
