package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/backpacker69/storasd/db"
	"github.com/backpacker69/storasd/models"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetOffers(t *testing.T) {
	// Initialize test database
	db.InitDB()
	defer db.CloseDB()

	r := SetUpRouter()
	r.GET("/offers", GetOffers)

	// Create a test offer
	testOffer := models.Offer{
		ID:      "test-id",
		Network: "testnet",
		Topic:   "Test Topic",
	}
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		encoded, _ := json.Marshal(testOffer)
		return b.Put([]byte(testOffer.ID), encoded)
	})

	// Make a GET request to /offers
	req, _ := http.NewRequest("GET", "/offers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var offers []models.Offer
	err := json.Unmarshal(w.Body.Bytes(), &offers)
	assert.NoError(t, err)
	/*
		assert.Len(t, offers, 1)
		assert.Equal(t, testOffer.ID, offers[0].ID)
		assert.Equal(t, testOffer.Network, offers[0].Network)
		assert.Equal(t, testOffer.Topic, offers[0].Topic)
	*/
}

func TestCreateOffer(t *testing.T) {
	// Initialize test database
	db.InitDB()
	defer db.CloseDB()

	r := SetUpRouter()
	r.POST("/offers", CreateOffer)

	// Create a test offer
	testOffer := models.Offer{
		ID:      "new-test-id",
		Network: "testnet",
		Topic:   "New Test Topic",
	}

	jsonValue, _ := json.Marshal(testOffer)
	req, _ := http.NewRequest("POST", "/offers", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)

	var createdOffer models.Offer
	err := json.Unmarshal(w.Body.Bytes(), &createdOffer)
	assert.NoError(t, err)
	assert.Equal(t, testOffer.ID, createdOffer.ID)
	assert.Equal(t, testOffer.Network, createdOffer.Network)
	assert.Equal(t, testOffer.Topic, createdOffer.Topic)

	// Check if the offer was actually saved in the database
	var savedOffer models.Offer
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		v := b.Get([]byte(testOffer.ID))
		return json.Unmarshal(v, &savedOffer)
	})
	assert.Equal(t, testOffer.ID, savedOffer.ID)
	assert.Equal(t, testOffer.Network, savedOffer.Network)
	assert.Equal(t, testOffer.Topic, savedOffer.Topic)
}

func TestGetOffer(t *testing.T) {
	// Initialize test database
	db.InitDB()
	defer db.CloseDB()

	r := SetUpRouter()
	r.GET("/offers/:id", GetOffer)

	// Create a test offer
	testOffer := models.Offer{
		ID:      "get-test-id",
		Network: "testnet",
		Topic:   "Get Test Topic",
	}
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		encoded, _ := json.Marshal(testOffer)
		return b.Put([]byte(testOffer.ID), encoded)
	})

	// Make a GET request to /offers/get-test-id
	req, _ := http.NewRequest("GET", "/offers/get-test-id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedOffer models.Offer
	err := json.Unmarshal(w.Body.Bytes(), &retrievedOffer)
	assert.NoError(t, err)
	assert.Equal(t, testOffer.ID, retrievedOffer.ID)
	assert.Equal(t, testOffer.Network, retrievedOffer.Network)
	assert.Equal(t, testOffer.Topic, retrievedOffer.Topic)
}

func TestUpdateOffer(t *testing.T) {
	// Initialize test database
	db.InitDB()
	defer db.CloseDB()

	r := SetUpRouter()
	r.PUT("/offers/:id", UpdateOffer)

	// Create a test offer
	testOffer := models.Offer{
		ID:      "update-test-id",
		Network: "testnet",
		Topic:   "Update Test Topic",
	}
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		encoded, _ := json.Marshal(testOffer)
		return b.Put([]byte(testOffer.ID), encoded)
	})

	// Update the offer
	updatedOffer := testOffer
	updatedOffer.Topic = "Updated Test Topic"

	jsonValue, _ := json.Marshal(updatedOffer)
	req, _ := http.NewRequest("PUT", "/offers/update-test-id", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedOffer models.Offer
	err := json.Unmarshal(w.Body.Bytes(), &retrievedOffer)
	assert.NoError(t, err)
	assert.Equal(t, updatedOffer.ID, retrievedOffer.ID)
	assert.Equal(t, updatedOffer.Network, retrievedOffer.Network)
	assert.Equal(t, updatedOffer.Topic, retrievedOffer.Topic)

	// Check if the offer was actually updated in the database
	var savedOffer models.Offer
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		v := b.Get([]byte(testOffer.ID))
		return json.Unmarshal(v, &savedOffer)
	})
	assert.Equal(t, updatedOffer.ID, savedOffer.ID)
	assert.Equal(t, updatedOffer.Network, savedOffer.Network)
	assert.Equal(t, updatedOffer.Topic, savedOffer.Topic)
}

func TestDeleteOffer(t *testing.T) {
	// Initialize test database
	db.InitDB()
	defer db.CloseDB()

	r := SetUpRouter()
	r.DELETE("/offers/:id", DeleteOffer)

	// Create a test offer
	testOffer := models.Offer{
		ID:      "delete-test-id",
		Network: "testnet",
		Topic:   "Delete Test Topic",
	}
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		encoded, _ := json.Marshal(testOffer)
		return b.Put([]byte(testOffer.ID), encoded)
	})

	// Make a DELETE request to /offers/delete-test-id
	req, _ := http.NewRequest("DELETE", "/offers/delete-test-id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Check if the offer was actually deleted from the database
	var deletedOffer models.Offer
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		v := b.Get([]byte(testOffer.ID))
		if v == nil {
			return nil
		}
		return json.Unmarshal(v, &deletedOffer)
	})
	assert.Empty(t, deletedOffer.ID)
}
