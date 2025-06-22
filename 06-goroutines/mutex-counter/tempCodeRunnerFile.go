c.mu.Lock() // Lock so only one goroutine at a time can access the map c.v.
	// defer c.mu.Unlock()