package goshopify

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const inventoryItemsBasePath = "inventory_items"

// InventoryItemService is an interface for interacting with the
// inventory items endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventoryitem
type InventoryItemService interface {
	List(interface{}) ([]InventoryItem, error)
	ListBatch([]int64, interface{}) ([]InventoryItem, error)
	Get(int64, interface{}) (*InventoryItem, error)
	Update(InventoryItem) (*InventoryItem, error)
}

// InventoryItemServiceOp is the default implementation of the InventoryItemService interface
type InventoryItemServiceOp struct {
	client *Client
}

// InventoryItem represents a Shopify inventory item
type InventoryItem struct {
	ID                int64            `json:"id,omitempty"`
	SKU               string           `json:"sku,omitempty"`
	CreatedAt         *time.Time       `json:"created_at,omitempty"`
	UpdatedAt         *time.Time       `json:"updated_at,omitempty"`
	Cost              *decimal.Decimal `json:"cost,omitempty"`
	Tracked           *bool            `json:"tracked,omitempty"`
	AdminGraphqlAPIID string           `json:"admin_graphql_api_id,omitempty"`
}

// InventoryItemResource is used for handling single item requests and responses
type InventoryItemResource struct {
	InventoryItem *InventoryItem `json:"inventory_item"`
}

// InventoryItemsResource is used for handling multiple item responsees
type InventoryItemsResource struct {
	InventoryItems []InventoryItem `json:"inventory_items"`
}

// List inventory items
func (s *InventoryItemServiceOp) List(options interface{}) ([]InventoryItem, error) {
	path := fmt.Sprintf("%s.json", inventoryItemsBasePath)
	resource := new(InventoryItemsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItems, err
}

// List inventory items for variant IDs
func (s *InventoryItemServiceOp) ListBatch(variantIDs []int64, options interface{}) ([]InventoryItem, error) {
	var idStrs []string
	for _, id := range variantIDs {
		idStrs = append(idStrs, fmt.Sprintf("%d", id))
	}
	path := fmt.Sprintf("%s.json?ids=%s", inventoryItemsBasePath, strings.Join(idStrs, ","))
	resource := new(InventoryItemsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItems, err
}

// Get a inventory item
func (s *InventoryItemServiceOp) Get(id int64, options interface{}) (*InventoryItem, error) {
	path := fmt.Sprintf("%s/%d.json", inventoryItemsBasePath, id)
	resource := new(InventoryItemResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItem, err
}

// Update a inventory item
func (s *InventoryItemServiceOp) Update(item InventoryItem) (*InventoryItem, error) {
	path := fmt.Sprintf("%s/%d.json", inventoryItemsBasePath, item.ID)
	wrappedData := InventoryItemResource{InventoryItem: &item}
	resource := new(InventoryItemResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.InventoryItem, err
}
