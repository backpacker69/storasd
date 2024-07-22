package models

type Oracle struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	URL       string `json:"url"`
}