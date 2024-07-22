package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/backpacker69/storasd/db"
	"github.com/backpacker69/storasd/models"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	var messages []models.Message
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		return b.ForEach(func(k, v []byte) error {
			var message models.Message
			if err := json.Unmarshal(v, &message); err != nil {
				return err
			}
			messages = append(messages, message)
			return nil
		})
	})
	c.JSON(http.StatusOK, messages)
}

func GetMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		v := b.Get([]byte(id))
		if v == nil {
			return nil // Message not found
		}
		return json.Unmarshal(v, &message)
	})
	if message.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}
	c.JSON(http.StatusOK, message)
}

func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		encoded, err := json.Marshal(message)
		if err != nil {
			return err
		}
		return b.Put([]byte(message.ID), encoded)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, message)
}

func UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message.ID = id
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		encoded, err := json.Marshal(message)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), encoded)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, message)
}

func DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		return b.Delete([]byte(id))
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
