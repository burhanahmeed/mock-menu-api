package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Menu represents the entire menu structure
type Menu struct {
	Menu struct {
		Categories []Category `json:"categories"`
	} `json:"menu"`
}

// Category represents a menu category
type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
}

// Item represents a menu item
type Item struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Price          float64        `json:"price"`
	Image          string         `json:"image"`
	Ingredients    []string       `json:"ingredients"`
	Customizations []Customization `json:"customizations"`
}

// Customization represents item customization options
type Customization struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Required bool     `json:"required"`
	Options  []Option `json:"options"`
}

// Option represents a customization option
type Option struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	ID             string                  `json:"id"`
	ItemID         string                  `json:"item_id"`
	Name           string                  `json:"name"`
	Price          float64                 `json:"price"`
	Quantity       int                     `json:"quantity"`
	Customizations []SelectedCustomization `json:"customizations"`
	TotalPrice     float64                 `json:"total_price"`
}

// SelectedCustomization represents a customization selected by the user
type SelectedCustomization struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	SelectedOptions []SelectedOption `json:"selected_options"`
}

// SelectedOption represents a customization option selected by the user
type SelectedOption struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Cart represents the shopping cart
type Cart struct {
	Items      []CartItem `json:"items"`
	TotalItems int        `json:"total_items"`
	TotalPrice float64    `json:"total_price"`
}

var menuData Menu
var cart Cart

func main() {
	// Load menu data
	err := loadMenuData()
	if err != nil {
		log.Fatalf("Failed to load menu data: %v", err)
	}

	// Create router
	r := mux.NewRouter()

	// Define routes - all using query parameters
	// Menu routes
	r.HandleFunc("/categories", listCategories).Methods("GET")
	r.HandleFunc("/category", getCategory).Methods("GET") // ?id=appetizers
	r.HandleFunc("/items", listItemsByCategory).Methods("GET") // ?category_id=appetizers
	r.HandleFunc("/all-items", listAllItems).Methods("GET") // Show all items across all categories
	r.HandleFunc("/item", getItem).Methods("GET") // ?id=app1
	r.HandleFunc("/customizations", listCustomizationsByItem).Methods("GET") // ?item_id=app1
	r.HandleFunc("/customization", getCustomization).Methods("GET") // ?id=spice_level
	
	// Cart routes
	r.HandleFunc("/cart", getCart).Methods("GET")
	r.HandleFunc("/cart/add", addToCart).Methods("POST")
	r.HandleFunc("/cart/remove", removeFromCart).Methods("POST")
	r.HandleFunc("/cart/clear", clearCart).Methods("POST")

	// Start server
	port := 8080
	fmt.Printf("Server starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func loadMenuData() error {
	// Read the menu_data.json file
	file, err := os.Open("menu_data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse the JSON data
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the menuData variable
	err = json.Unmarshal(byteValue, &menuData)
	if err != nil {
		return err
	}

	return nil
}

// listCategories returns all categories
func listCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Create a simplified response with just category info, not all items
	type SimplifiedCategory struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	
	simplifiedCategories := make([]SimplifiedCategory, 0)
	for _, category := range menuData.Menu.Categories {
		simplifiedCategories = append(simplifiedCategories, SimplifiedCategory{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	
	json.NewEncoder(w).Encode(simplifiedCategories)
}

// getCategory returns a specific category by ID using query parameter
func getCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get category ID from query parameter
	categoryID := r.URL.Query().Get("id")
	if categoryID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing category id parameter"})
		return
	}
	
	for _, category := range menuData.Menu.Categories {
		if category.ID == categoryID {
			json.NewEncoder(w).Encode(category)
			return
		}
	}
	
	// If category not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
}

// listItemsByCategory returns all items in a specific category using query parameter
func listItemsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get category ID from query parameter
	categoryID := r.URL.Query().Get("category_id")
	if categoryID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing category_id parameter"})
		return
	}
	
	for _, category := range menuData.Menu.Categories {
		if category.ID == categoryID {
			json.NewEncoder(w).Encode(category.Items)
			return
		}
	}
	
	// If category not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
}

// getItem returns a specific item by ID using query parameter
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get item ID from query parameter
	itemID := r.URL.Query().Get("id")
	if itemID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing item id parameter"})
		return
	}
	
	for _, category := range menuData.Menu.Categories {
		for _, item := range category.Items {
			if item.ID == itemID {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}
	
	// If item not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
}

// listCustomizationsByItem returns all customizations for a specific item using query parameter
func listCustomizationsByItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get item ID from query parameter
	itemID := r.URL.Query().Get("item_id")
	if itemID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing item_id parameter"})
		return
	}
	
	for _, category := range menuData.Menu.Categories {
		for _, item := range category.Items {
			if item.ID == itemID {
				json.NewEncoder(w).Encode(item.Customizations)
				return
			}
		}
	}
	
	// If item not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
}

// getCustomization returns a specific customization by ID using query parameter
func getCustomization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get customization ID from query parameter
	customizationID := r.URL.Query().Get("id")
	if customizationID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing customization id parameter"})
		return
	}
	
	for _, category := range menuData.Menu.Categories {
		for _, item := range category.Items {
			for _, customization := range item.Customizations {
				if customization.ID == customizationID {
					json.NewEncoder(w).Encode(customization)
					return
				}
			}
		}
	}
	
	// If customization not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Customization not found"})
}

// listAllItems returns all menu items across all categories
func listAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Collect all items from all categories
	allItems := []Item{}
	
	for _, category := range menuData.Menu.Categories {
		for _, item := range category.Items {
			// Add category information to each item for reference
			itemWithCategory := item
			allItems = append(allItems, itemWithCategory)
		}
	}
	
	// Return all items
	json.NewEncoder(w).Encode(allItems)
}

// getCart returns the current cart
func getCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Return the cart
	json.NewEncoder(w).Encode(cart)
}

// addToCart adds an item to the cart
func addToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Parse the request body
	var requestBody struct {
		ItemID         string                  `json:"item_id"`
		Quantity       int                     `json:"quantity"`
		Customizations []SelectedCustomization `json:"customizations"`
	}
	
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	
	// Validate the item ID
	if requestBody.ItemID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing item_id"})
		return
	}
	
	// Validate the quantity
	if requestBody.Quantity <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Quantity must be greater than 0"})
		return
	}
	
	// Find the item in the menu
	var item Item
	itemFound := false
	
	for _, category := range menuData.Menu.Categories {
		for _, menuItem := range category.Items {
			if menuItem.ID == requestBody.ItemID {
				item = menuItem
				itemFound = true
				break
			}
		}
		if itemFound {
			break
		}
	}
	
	if !itemFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
		return
	}
	
	// Calculate the total price including customizations
	totalPrice := item.Price
	for _, customization := range requestBody.Customizations {
		for _, option := range customization.SelectedOptions {
			totalPrice += option.Price
		}
	}
	totalPrice *= float64(requestBody.Quantity)
	
	// Create a new cart item
	cartItem := CartItem{
		ID:             fmt.Sprintf("%s_%d", requestBody.ItemID, len(cart.Items)),
		ItemID:         requestBody.ItemID,
		Name:           item.Name,
		Price:          item.Price,
		Quantity:       requestBody.Quantity,
		Customizations: requestBody.Customizations,
		TotalPrice:     totalPrice,
	}
	
	// Add the item to the cart
	cart.Items = append(cart.Items, cartItem)
	
	// Update the cart totals
	cart.TotalItems += requestBody.Quantity
	cart.TotalPrice += totalPrice
	
	// Return the updated cart
	json.NewEncoder(w).Encode(cart)
}

// removeFromCart removes an item from the cart
func removeFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Parse the request body
	var requestBody struct {
		CartItemID string `json:"cart_item_id"`
	}
	
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	
	// Validate the cart item ID
	if requestBody.CartItemID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing cart_item_id"})
		return
	}
	
	// Find the item in the cart
	itemIndex := -1
	for i, item := range cart.Items {
		if item.ID == requestBody.CartItemID {
			itemIndex = i
			break
		}
	}
	
	if itemIndex == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Cart item not found"})
		return
	}
	
	// Update the cart totals
	cart.TotalItems -= cart.Items[itemIndex].Quantity
	cart.TotalPrice -= cart.Items[itemIndex].TotalPrice
	
	// Remove the item from the cart
	cart.Items = append(cart.Items[:itemIndex], cart.Items[itemIndex+1:]...)
	
	// Return the updated cart
	json.NewEncoder(w).Encode(cart)
}

// clearCart removes all items from the cart
func clearCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Clear the cart
	cart.Items = []CartItem{}
	cart.TotalItems = 0
	cart.TotalPrice = 0
	
	// Return the empty cart
	json.NewEncoder(w).Encode(cart)
}
