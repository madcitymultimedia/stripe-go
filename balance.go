package stripe

import "encoding/json"

// TransactionStatus is the list of allowed values for the transaction's status.
// Allowed values are "available", "pending".
type TransactionStatus string

// TransactionType is the list of allowed values for the transaction's type.
// Allowed values are "charge", "refund", "adjustment", "application_fee",
// "application_fee_refund", "transfer", "transfer_cancel", "transfer_failure".
type TransactionType string

const (
	TxAvailable TransactionStatus = "available"
	TxPending   TransactionStatus = "pending"

	TxCharge         TransactionType = "charge"
	TxRefund         TransactionType = "refund"
	TxAdjust         TransactionType = "adjustment"
	TxFee            TransactionType = "application_fee"
	TxFeeRefund      TransactionType = "application_fee_refund"
	TxTransfer       TransactionType = "transfer"
	TxTransferCancel TransactionType = "transfer_cancel"
	TxTransferFail   TransactionType = "transfer_failure"
)

// BalanceParams is the set of parameters that can be used when retrieving a balance.
// For more details see https://stripe.com/docs/api#balance.
type BalanceParams struct {
	Params
}

// TxParams is the set of parameters that can be used when retrieving a transaction.
// For more details see https://stripe.com/docs/api#retrieve_balance_transaction.
type TxParams struct {
	Params
}

// TxListParams is the set of parameters that can be used when listing balance transactions.
// For more details see https://stripe.com/docs/api/#balance_history.
type TxListParams struct {
	ListParams
	Created, Available      int64
	Currency, Src, Transfer string
	Type                    TransactionType
}

// Balance is the resource representing your Stripe balance.
// For more details see https://stripe.com/docs/api/#balance.
type Balance struct {
	// Live indicates the live mode.
	Live      bool     `json:"livemode"`
	Available []Amount `json:"available"`
	Pending   []Amount `json:"pending"`
}

// Transaction is the resource representing the balance transaction.
// For more details see https://stripe.com/docs/api/#balance.
type Transaction struct {
	Id         string            `json:"id"`
	Amount     int64             `json:"amount"`
	Currency   Currency          `json:"currency"`
	Available  int64             `json:"available_on"`
	Created    int64             `json:"created"`
	Fee        int64             `json:"fee"`
	FeeDetails []Fee             `json:"fee_details"`
	Net        int64             `json:"net"`
	Status     TransactionStatus `json:"status"`
	Type       TransactionType   `json:"type"`
	Desc       string            `json:"description"`
	Src        string            `json:"source"`
	Recipient  string            `json:"recipient"`
}

// Amount is a structure wrapping an amount value and its currency.
type Amount struct {
	Value    int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

// Fee is a structure that breaks down the fees in a transaction.
type Fee struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Type        string   `json:"type"`
	Desc        string   `json:"description"`
	Application string   `json:"application"`
}

// TransactionList is a list object for transactions.
type TransactionList struct {
	ListResponse
	Values []*Transaction `json:"data"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type transaction Transaction
	var tt transaction
	err := json.Unmarshal(data, &tt)
	if err == nil {
		*t = Transaction(tt)
	} else {
		// the id is surrounded by escaped \, so ignore those
		t.Id = string(data[1 : len(data)-1])
	}

	return nil
}