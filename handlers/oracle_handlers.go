package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/backpacker69/storasd/db"
	"github.com/backpacker69/storasd/models"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func GetOracles(c *gin.Context) {
	var oracles []models.Oracle
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("oracles"))
		return b.ForEach(func(k, v []byte) error {
			var oracle models.Oracle
			if err := json.Unmarshal(v, &oracle); err != nil {
				return err
			}
			oracles = append(oracles, oracle)
			return nil
		})
	})
	c.JSON(http.StatusOK, oracles)
}

func GetOracle(c *gin.Context) {
	id := c.Param("id")
	var oracle models.Oracle
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("oracles"))
		v := b.Get([]byte(id))
		if v == nil {
			return nil // Oracle not found
		}
		return json.Unmarshal(v, &oracle)
	})
	if oracle.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Oracle not found"})
		return
	}
	c.JSON(http.StatusOK, oracle)
}

func CreateOracle(c *gin.Context) {
	var oracle models.Oracle
	if err := c.BindJSON(&oracle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("oracles"))
		encoded, err := json.Marshal(oracle)
		if err != nil {
			return err
		}
		return b.Put([]byte(oracle.ID), encoded)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, oracle)
}

func UpdateOracle(c *gin.Context) {
	id := c.Param("id")
	var oracle models.Oracle
	if err := c.BindJSON(&oracle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oracle.ID = id
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("oracles"))
		encoded, err := json.Marshal(oracle)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), encoded)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, oracle)
}

func DeleteOracle(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("oracles"))
		return b.Delete([]byte(id))
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
