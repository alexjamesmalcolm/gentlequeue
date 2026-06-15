package backend

import (
	"context"
	"time"

	"github.com/alexjamesmalcolm/gentlequeue/uuid"
)

type Importance uint16

type (
	ItemID uuid.UUID
	Item   struct {
		ID          ItemID
		Title       string
		Description string
		Importance  Importance
	}
)

type (
	CompletionEventID uuid.UUID
	CompletionEvent   struct {
		ID     CompletionEventID
		ItemID ItemID
		Time   time.Time
	}
)

type Backend interface {
	GetAllItems(context.Context) ([]Item, error)
	GetItemByID(context.Context, ItemID) (Item, error)
	DeleteItemByID(context.Context, ItemID) error
	UpdateItem(context.Context, Item) error
	CreateItem(context.Context, Item) error

	GetCompletionEventsByItemID(ctx context.Context, itemID ItemID, before time.Time, after time.Time) ([]CompletionEvent, error)
	DeleteCompletionEventByID(context.Context, CompletionEventID) error
	UpdateCompletionEvent(context.Context, CompletionEvent) error
	CreateCompletionEvent(context.Context, CompletionEvent) error
}
