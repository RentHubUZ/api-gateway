package models

type CreateProperties struct {
	Address        string   `json:"address"`
	Price          float64  `json:"price"`
	Property_type  string   `json:"property_type"`
	Bedrooms       int32    `json:"bedrooms"`
	Bathrooms      int32    `json:"bathrooms"`
	Square_footage float64  `json:"square_footage"`
	Listing_status string   `json:"listing_status"`
	Description    string   `json:"description"`
	Roommate_count int32    `json:"roommate_count"`
	Lease_terms    string   `json:"lease_terms"`
	Lease_duration int32    `json:"lease_duration"`
	Top_status     bool     `json:"top_status"`
	Image_url      []string `json:"image_url"`
	Latitude       float64  `json:"latitude"`
	Longitude      float64  `json:"longitude"`
}

type UpdateProperties struct {
	Id             string   `json:"id"`
	Address        string   `json:"address"`
	Price          float64  `json:"price"`
	Property_type  string   `json:"property_type"`
	Bedrooms       int32    `json:"bedrooms"`
	Bathrooms      int32    `json:"bathrooms"`
	Square_footage float64  `json:"square_footage"`
	Listing_status string   `json:"listing_status"`
	Description    string   `json:"description"`
	Roommate_count int32    `json:"roommate_count"`
	Lease_terms    string   `json:"lease_terms"`
	Lease_duration int32    `json:"lease_duration"`
	Top_status     bool     `json:"top_status"`
	Image_url      []string `json:"image_url"`
	Latitude       float64  `json:"latitude"`
	Longitude      float64  `json:"longitude"`
}
