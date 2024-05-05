package entity

import "gorm.io/gorm"

type OpenDonation struct {
	gorm.Model
	Title       string  `json:"title"`
	Photo       string  `json:"photo"`
	Description string  `json:"description"`
	Target      float64 `json:"target"`
	IsDisaster  bool    `json:"is_disaster"`
}

type OpenDonationResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Photo       string  `json:"photo"`
	Description string  `json:"description"`
	Total       float64 `json:"total"`
	Target      float64 `json:"target"`
	IsDisaster  bool    `json:"is_disaster"`
	TotalDonors int64   `json:"total_donors"`
}
