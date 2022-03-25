package storage

import (
	todo "WB_GO_L0"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"sync"
)

type Cache struct {
	mx sync.RWMutex
	m  map[int]todo.OrderItems
}

//Get out the value by key
func (c *Cache) Get(key int) (todo.OrderItems, bool) {
	c.mx.RLock()
	val, ok := c.m[key]
	c.mx.RUnlock()
	return val, ok
}

//Add new value in Cache
func (c *Cache) Add(key int, value todo.OrderItems) {
	c.mx.Lock()
	c.m[key] = value
	c.mx.Unlock()
}

//NewCache creating
func NewCache() *Cache {
	return &Cache{
		m:  make(map[int]todo.OrderItems),
		mx: sync.RWMutex{},
	}
}

func GetCache(db *sqlx.DB) (*Cache, error) {

	cache := NewCache()

	psqlQ := fmt.Sprintf("SELECT * FROM %s ;", OrdersTable)
	rows, err := db.Queryx(psqlQ)
	if err != nil {
		return &Cache{}, err
	}

	for rows.Next() {
		order := todo.Order{}
		items := make([]todo.Item, 0)

		err = rows.StructScan(&order)
		if err != nil {
			return &Cache{}, err
		}
		psqlQ = fmt.Sprintf("SELECT * FROM %s WHERE %s = %s ;", ItemsTable, orderIdColumns, strconv.Itoa(order.OrderId))
		rowsI, err := db.Queryx(psqlQ)
		if err != nil {
			return &Cache{}, err
		}
		for rowsI.Next() {
			item := todo.Item{}
			err := rowsI.StructScan(&item)
			if err != nil {
				return &Cache{}, err
			}
			items = append(items, item)
		}
		cache.Add(order.OrderId, todo.OrderItems{
			Order: order,
			Items: items,
		})
	}

	return cache, nil
}
