package models

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	PublicKey string `json:"publicKey"`
	Address   string `json:"address"`
}