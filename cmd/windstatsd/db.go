package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

// DBClient ...
type DBClient struct {
	db     string
	client client.Client
}

// NewDBClient ...
func NewDBClient(addr string, db string, user string, pass string) *DBClient {
	client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: user,
		Password: pass,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	return &DBClient{db: db, client: client}
}

// Query ...
func (c *DBClient) Query(measurement string) ([]map[string]interface{}, error) {

	return nil, nil
}

// Insert ...
func (c *DBClient) Insert(measurement string, tags map[string]string, fields map[string]interface{}) error {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  c.db,
		Precision: "s",
	})
	if err != nil {
		return fmt.Errorf("Error creating batch points %w", err)
	}
	pt, err := client.NewPoint(measurement, tags, fields, time.Now())
	if err != nil {
		return fmt.Errorf("Error creating new point %w", err)
	}
	bp.AddPoint(pt)
	if err := c.client.Write(bp); err != nil {
		return fmt.Errorf("Error writing data %w", err)
	}
	return nil
}

// Close ...
func (c *DBClient) Close() {
	c.client.Close()
}
