package cache

import (
	"L2/http/models"
	"log"
	"sync"
	"time"
)

type Cache struct {
	m    sync.Mutex
	pool map[string]*models.Event
}

func NewCache() *Cache {
	return &Cache{
		m:    sync.Mutex{},
		pool: map[string]*models.Event{},
	}
}

func (c *Cache) Create(e models.Event) {
	c.m.Lock()
	defer c.m.Unlock()
	c.pool[e.Date] = &e
}

func (c *Cache) ReadDay(date string) (models.Event, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	val, ok := c.pool[date]
	if ok {
		return *val, true
	}
	log.Println("Err at getting", date, ": no events detected")
	return models.Event{}, false
}

func (c *Cache) ReadWeek(date string) ([]models.Event, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	result := []models.Event{}

	today, errParseToday := time.Parse("2006-01-02", date)

	if errParseToday != nil {
		log.Println("Error at parsing date: ", errParseToday)
		return nil, false
	}

	for _, event := range c.pool {

		evDate, errParseEv := time.Parse("2006-01-02", event.Date)
		if errParseEv != nil {
			log.Println("Error at parsing date: ", errParseEv)
			return nil, false
		}

		if evDate.After(today) && evDate.Before(today.Add(time.Hour*168)) {
			result = append(result, *event)
		}
	}

	return result, true
}

func (c *Cache) ReadMonth(date string) []models.Event {
	c.m.Lock()
	defer c.m.Unlock()

	result := make([]models.Event, 1)

	today, errParseToday := time.Parse("2006-01-02", date)

	if errParseToday != nil {
		log.Println("Error at parsing date: ", errParseToday)
		return nil
	}

	for _, event := range c.pool {

		evDate, errParseEv := time.Parse("2006-01-02", event.Date)
		if errParseEv != nil {
			log.Println("Error at parsing date: ", errParseEv)
			return nil
		}

		if evDate.After(today) && evDate.Before(today.Add(time.Hour*24*30)) {
			result = append(result, *event)
		}
	}

	return result
}

func (c *Cache) Update(e models.Event) {
	c.m.Lock()
	defer c.m.Unlock()
	c.pool[e.Date] = &e
}
