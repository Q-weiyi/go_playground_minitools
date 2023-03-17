package helper

import "time"

type DirectDepositToCallback struct {
	ID                           int64     `json:"-"`
	MerchantCallbackID           int64     `json:"-"`
	PaymentAddress               string    `json:"payment_address"`
	CompletedAt                  time.Time `json:"completed_at"`
	DepositAmount                float64   `json:"deposit_amount"`
	FeeAmount                    float64   `json:"fee_amount"`
	MerchantUserID               string    `json:"merchant_user_id"`
	TransactionHash              string    `json:"transaction_hash"`
	MerchantID                   string    `json:"merchant_id"`
	IntegrationSystemID          int       `json:"-"`
	CryptoCurrency               string    `json:"crypto_currency"`
	CryptoCurrencyNetwork        string    `json:"crypto_currency_network"`
	CallbackSettingUrl           string    `json:"-"`
	CallbackSettingSigningSecret string    `json:"-"`
	DirectDepositID              *int64    `json:"direct_deposit_id,omitempty" gorm:"-"`
}

type DepositToCallback struct {
	ID                 int64 `json:"-"`
	DepositStatusID    int   `json:"-"`
	MerchantCallbackID int64 `json:"-"`

	MerchantDepositID   string    `json:"merchant_deposit_id"`
	PaymentAddress      string    `json:"payment_address"`
	CompletedAt         time.Time `json:"completed_at"`
	ActualDepositAmount float64   `json:"actual_deposit_amount"`
	ActualPaymentAmount float64   `json:"actual_payment_amount"`
	FeeAmount           float64   `json:"fee_amount"`
	DetectedAmount      float64   `json:"detected_amount"`
	Status              *string   `json:"status,omitempty" gorm:"-"`

	MerchantID          string `json:"merchant_id"`
	IntegrationSystemID int    `json:"-"`

	CryptoCurrency        string `json:"crypto_currency"`
	CryptoCurrencyNetwork string `json:"crypto_currency_network"`

	CallbackSettingUrl           string `json:"-"`
	CallbackSettingSigningSecret string `json:"-"`

	Transactions []TransactionToCallback `json:"transactions" gorm:"foreignKey:DepositID"`

	InvoiceDepositID *int64 `json:"invoice_deposit_id,omitempty" gorm:"-"`
}

type TransactionToCallback struct {
	DepositID       int64   `json:"-"`
	TransactionHash string  `json:"transaction_hash"`
	Amount          float64 `json:"amount"`
}

type RefundToCallback struct {
	ID                 int64 `json:"-"`
	RefundStatusID     int   `json:"-"`
	MerchantCallbackID int64 `json:"-"`

	MerchantRefundID string    `json:"merchant_refund_id"`
	RefundAddress    string    `json:"refund_address"`
	CompletedAt      time.Time `json:"completed_at"`
	RequestedAmount  float64   `json:"requested_amount"`
	NetworkFeeAmount float64   `json:"network_fee_amount"`
	TransactionHash  string    `json:"transaction_hash"`
	Status           *string   `json:"status,omitempty" gorm:"-"`

	MerchantID string `json:"merchant_id"`

	CryptoCurrency        string `json:"crypto_currency"`
	CryptoCurrencyNetwork string `json:"crypto_currency_network"`
	NetworkFeeCurrency    string `json:"network_fee_currency"`

	CallbackSettingUrl           string `json:"-"`
	CallbackSettingSigningSecret string `json:"-"`

	RefundID int64 `json:"refund_id" gorm:"-"`
}

type SettlementToCallback struct {
	ID                 int64 `json:"-"`
	SettlementStatusID int   `json:"-"`
	MerchantCallbackID int64 `json:"-"`

	MerchantSettlementID string    `json:"merchant_settlement_id"`
	SettlementAddress    string    `json:"settlement_address"`
	CompletedAt          time.Time `json:"completed_at"`
	RequestedAmount      float64   `json:"requested_amount"`
	FeeAmount            float64   `json:"fee_amount"`
	NetworkFeeAmount     float64   `json:"network_fee_amount"`
	TransactionHash      string    `json:"transaction_hash"`
	Status               *string   `json:"status,omitempty" gorm:"-"`

	MerchantID string `json:"merchant_id"`

	CryptoCurrency        string `json:"crypto_currency"`
	CryptoCurrencyNetwork string `json:"crypto_currency_network"`
	NetworkFeeCurrency    string `json:"network_fee_currency"`

	CallbackSettingUrl           string `json:"-"`
	CallbackSettingSigningSecret string `json:"-"`

	SettlementID int64 `json:"settlement_id" gorm:"-"`
}

type TransactionToCallback2 struct {
	ID                 int64 `json:"-"`
	TransactionTypeID  int   `json:"-"`
	MerchantID_        int64 `json:"-" gorm:"merchant_id_"`
	MerchantCallbackID int64 `json:"-"`

	PaymentAddress string    `json:"payment_address"`
	CompletedAt    time.Time `json:"completed_at"`

	Amount    float64  `json:"amount"`
	FeeAmount *float64 `json:"fee_amount,omitempty"`

	MerchantID string `json:"merchant_id"`

	CryptoCurrency        string `json:"crypto_currency"`
	CryptoCurrencyNetwork string `json:"crypto_currency_network"`

	CallbackSettingUrl           string `json:"-"`
	CallbackSettingSigningSecret string `json:"-"`

	TransactionHash string `json:"transaction_hash" gorm:"primaryKey"`

	FeeTransactions []FeeTransaction `json:"-" gorm:"foreignKey:TransactionHash"`

	TransactionID int64 `json:"transaction_id" gorm:"-"`
}

type FeeTransaction struct {
	TransactionHash string  `json:"transaction_hash"`
	Amount          float64 `json:"amount"`
	MerchantID      int64   `json:"merchant_id"`
}

type WithdrawalToCallback struct {
	ID                 int64 `json:"-"`
	WithdrawalStatusID int   `json:"-"`
	MerchantCallbackID int64 `json:"-"`

	MerchantWithdrawalID string    `json:"merchant_withdrawal_id"`
	WithdrawalAddress    string    `json:"withdrawal_address"`
	CompletedAt          time.Time `json:"completed_at"`
	RequestedAmount      float64   `json:"requested_amount"`
	FeeAmount            float64   `json:"fee_amount"`
	NetworkFeeAmount     float64   `json:"network_fee_amount"`
	TransactionHash      string    `json:"transaction_hash"`
	Status               *string   `json:"status,omitempty" gorm:"-"`

	MerchantID          string `json:"merchant_id"`
	IntegrationSystemID int    `json:"-"`

	CryptoCurrency        string `json:"crypto_currency"`
	CryptoCurrencyNetwork string `json:"crypto_currency_network"`
	NetworkFeeCurrency    string `json:"network_fee_currency"`

	CallbackSettingUrl           string `json:"-"`
	CallbackSettingSigningSecret string `json:"-"`

	WithdrawalID *int64 `json:"withdrawal_id,omitempty" gorm:"-"`
}
