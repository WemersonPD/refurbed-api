package models

import "strings"

type Name string

func (n Name) ToString() string {
	return string(n)
}

func (n Name) Contains(substring string) bool {
	nameLower := strings.ToLower(n.ToString())
	substringLower := strings.ToLower(substring)

	return strings.Contains(nameLower, substringLower)
}

type Color string

const (
	ColorBlack  Color = "black"
	ColorWhite  Color = "white"
	ColorSilver Color = "silver"
	ColorGray   Color = "gray"
	ColorBlue   Color = "blue"
	ColorRed    Color = "red"
	ColorGreen  Color = "green"
	ColorPink   Color = "pink"
	ColorOrange Color = "orange"
	ColorViolet Color = "violet"
	ColorYellow Color = "yellow"
	ColorBeige  Color = "beige"
	ColorGold   Color = "gold"
)

type Colors []Color

func (colors Colors) Contains(color string) bool {
	for _, c := range colors {
		if strings.EqualFold(string(c), string(color)) {
			return true
		}
	}

	return false
}

type ProductMetadata struct {
	ID        string  `json:"id"`
	Name      Name    `json:"name"`
	BasePrice float64 `json:"base_price"`
	ImageURL  string  `json:"image_url"`
}

type Category string

const (
	CategorySmartphones Category = "smartphones"
	CategoryTablets     Category = "tablets"
	CategoryLaptops     Category = "laptops"
	CategoryAccessories Category = "accessories"
)

func (c Category) Contains(category []string) bool {
	for _, cat := range category {
		if strings.EqualFold(string(c), string(cat)) {
			return true
		}
	}

	return false
}

type Brand string

const (
	BrandApple   Brand = "apple"
	BrandSamsung Brand = "samsung"
	BrandGoogle  Brand = "google"
	BrandXiaomi  Brand = "xiaomi"
)

func (b Brand) Contains(brand []string) bool {
	for _, br := range brand {
		if strings.EqualFold(string(b), string(br)) {
			return true
		}
	}

	return false
}

type Condition string

const (
	ConditionNew         Condition = "new"
	ConditionRefurbished Condition = "refurbished"
	ConditionUsed        Condition = "used"
)

func (c Condition) Contains(condition []string) bool {
	for _, cond := range condition {
		if strings.EqualFold(string(c), string(cond)) {
			return true
		}
	}

	return false
}

type ProductDetail struct {
	ID              string    `json:"id"`
	DiscountPercent int       `json:"discount_percent"`
	Bestseller      bool      `json:"bestseller"`
	Colors          Colors    `json:"colors"`
	Stock           int       `json:"stock"`
	Category        Category  `json:"category"`
	Brand           Brand     `json:"brand"`
	Condition       Condition `json:"condition"`
}

type Product struct {
	ID              string  `json:"id"`
	DiscountedPrice float64 `json:"discounted_price"`

	// Metadata fields
	Name      Name    `json:"name"`
	BasePrice float64 `json:"base_price"`
	ImageURL  string  `json:"image_url"`

	// Detail fields
	DiscountPercent int       `json:"discount_percent"`
	Bestseller      bool      `json:"bestseller"`
	Colors          Colors    `json:"colors"`
	Stock           int       `json:"stock"`
	Category        Category  `json:"category"`
	Brand           Brand     `json:"brand"`
	Condition       Condition `json:"condition"`
}

type ProductsResponse struct {
	Products []*Product `json:"products"`
	Count    int        `json:"count"`
}
