package models

type CET struct {
	Outcome string `json:"outcome"`
	Payout  int    `json:"payout"`
}

type Offer struct {
	Network         string   `json:"network"`
	ID              string   `json:"id"`
	OracleInfo      string   `json:"oracleInfo"`
	MakerPubkey     string   `json:"makerPubkey"`
	MakerAddress    string   `json:"makerAddress"`
	MakerCollateral int      `json:"makerCollateral"`
	ProposedCETs    []CET    `json:"proposedCETs"`
	ChangeAddress   string   `json:"changeAddress"`
	ExpiresBy       int64    `json:"expiresBy"`
	RefundTxn       string   `json:"refundTxn"`
	Topic           string   `json:"topic"`
	TopicID         string   `json:"topicID"`
	MakerSignature  string   `json:"makerSignature"`
}