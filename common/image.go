package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"-"`
	Url       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	CloudName string `json:"cloud_name,omitempty"`
	Extension string `json:"extension,omitempty"`
}

func (j *Image) Scan(value interface{})  error{
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var image Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}
	*j = image
	return nil
}

//value return json value, implement driver.value interface
func (j *Image) Value() (driver.Value, error){
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{})  error{
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var images Images
	if err := json.Unmarshal(bytes, &images); err != nil {
		return err
	}
	*j = images
	return nil
}

func (j *Images) Value() (driver.Value, error){
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}