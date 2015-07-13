package entrapped

// centeral hub to manage all connections
type trap struct {
	// registered connections
	troopers map[string]*connection
	// register requests
	enter chan *trooper
	// unregister requests
	dead chan *trooper
	// errors
	errors chan string
}

var hub = &trap{
	troopers: make(map[string]*connection),
	enter:    make(chan *trooper),
	dead:     make(chan *trooper),
	errors:   make(chan string),
}

func (t *trap) run() {
	for {
		select {
		case c := <-t.enter:
			if _, ok := t.troopers[c.nickname]; ok {
				t.errors <- "user exists"
			} else {
				t.troopers[c.nickname] = c.connection
			}
		case c := <-t.dead:
			if _, ok := t.troopers[c.nickname]; ok {
				delete(t.troopers, c.nickname)
				close(c.connection.data)
			}
		}
	}
}

func (t *trap) add(name string, conn *connection) {
	t.enter <- &trooper{name, conn}

	err := <-t.errors
	if len(err) != 0 {
		conn.ws.WriteJSON(err)
		conn.ws.Close()
	}

	for {
	}
}
