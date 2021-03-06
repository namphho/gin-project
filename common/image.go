package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
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