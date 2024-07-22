package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/backpacker69/storasd/db"
	"github.com/backpacker69/storasd/models"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func GetOffers(c *gin.Context) {
	var offers []models.Offer
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		return b.ForEach(func(k, v []byte) error {
			var offer models.Offer
			if err := json.Unmarshal(v, &offer); err != nil {
				return err
			}
			offers = append(offers, offer)
			return nil
		})
	})
	c.JSON(http.StatusOK, offers)
}

func GetOffer(c *gin.Context) {
	id := c.Param("id")
	var offer models.Offer
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		v := b.Get([]byte(id))
		if v == nil {
			return nil // Offer not found
		}
		return json.Unmarshal(v, &offer)
	})
	if offer.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}
	c.JSON(http.StatusOK, offer)
}

func CreateOffer(c *gin.Context) {
	var offer models.Offer
	if err := c.BindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		encoded, err := json.Marshal(offer)
		if err != nil {
			return err
		}
		return b.Put([]byte(offer.ID), encoded)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, offer)
}

func UpdateOffer(c *gin.Context) {
	id := c.Param("id")
	var offer models.Offer
	if err := c.BindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offer.ID = id
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		encoded, err := json.Marshal(offer)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), encoded)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, offer)
}

func DeleteOffer(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		return b.Delete([]byte(id))
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
