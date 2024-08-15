package main
import (
    "net/http"
    "encoding/json"
    "fmt"
    "strconv"
    "time"
)

// Helping functions for events 
func parseEvent(r *http.Request) (Event, error) {
    ev := Event{}
    err := json.NewDecoder(r.Body).Decode(&ev)
    if err != nil {
	return ev, err
    }
    return ev, nil
}

type DelEvent struct {
    Id		uint64	`json:"id"`
    UserId	uint64	`json:"user_id"`
}

func parseDelete(r *http.Request) (DelEvent, error) {
    dev := DelEvent{}
    err := json.NewDecoder(r.Body).Decode(&dev)
    if err != nil {
	return dev, err
    }
    return dev, nil
}

func validateEvent(ev Event) bool {
    if ev.UserId == 0 || ev.Name == "" {
	return false
    }
    return true
}

func (s *Server) findEventUserId(id uint64) uint64 {
    s.cache.RWMutex.RLock()
    defer s.cache.RWMutex.RUnlock()
    ev, ok := s.cache.Cache[id]
    if ok {
	return ev.UserId
    }
    return 0
}


// Handlers and their help functions

//=======CREATE=======
func (s *Server) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
	respError(w, fmt.Errorf("Wrong HTTP Method"), http.StatusBadRequest)
	return
    }
    newEvent, err := parseEvent(r)
    if err != nil {
	respError(w, err, http.StatusBadRequest)
	return
    }
    if !validateEvent(newEvent) {
	respError(w, fmt.Errorf("Invalid Event"), http.StatusBadRequest)
	return
    }
    err = s.createEvent(newEvent)
    if err != nil {
	respError(w, err, http.StatusServiceUnavailable)
    } else {
	var evs []Event
	evs = append(evs, newEvent)
	respResult(w, fmt.Sprintf("Event with ID %v was created", newEvent.Id), evs, http.StatusOK)
    }
}

func (s *Server) createEvent(ev Event) error {
    uid := s.findEventUserId(ev.Id) 
    if uid != 0 {
	return fmt.Errorf("Event with this ID already exists")
    }

    s.cache.RWMutex.Lock()
    s.cache.Cache[ev.Id] = ev
    s.cache.RWMutex.Unlock()
    return nil
}

//=======UPDATE========
func (s *Server) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
	respError(w, fmt.Errorf("Wrong HTTP Method"), http.StatusBadRequest)
	return
    }
    updEvent, err := parseEvent(r)
    if err != nil {
	respError(w, fmt.Errorf("Invalid Event"), http.StatusBadRequest)
	return
    }

    if !validateEvent(updEvent) {
	respError(w, fmt.Errorf("Invalid Event"), http.StatusBadRequest)
	return
    }

    err = s.updateEvent(updEvent)
    if err != nil {
	respError(w, err, http.StatusServiceUnavailable)
    } else {	
	var evs []Event
	evs = append(evs, updEvent)
	respResult(w, fmt.Sprintf("Event with ID %v was updated", updEvent.Id), evs, http.StatusOK)
    }
}

func (s *Server) updateEvent(ev Event) error {
    uid := s.findEventUserId(ev.Id) 
    if ev.UserId == uid {
	s.cache.RWMutex.Lock()
	s.cache.Cache[ev.Id] = ev
	s.cache.RWMutex.Unlock()
	return nil
    } else if uid == 0 {
	return fmt.Errorf("Event with this ID not found")
    } else {
	return fmt.Errorf("Wrong UserID")
    }
}

//=======DELETE=======
func (s *Server) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
	respError(w, fmt.Errorf("Wrong HTTP Method"), http.StatusBadRequest)
	return
    }
    delEvent, err := parseDelete(r)
    if err != nil {
	respError(w, fmt.Errorf("Invalid delete command parameters"), http.StatusBadRequest)
	return
    }
    err = s.deleteEvent(delEvent)
    if err != nil {
	respError(w, err, http.StatusServiceUnavailable)
    } else {
	respResult(w, fmt.Sprintf("Event with ID %v was deleted", delEvent.Id), nil, http.StatusOK)
    }
 }

func (s *Server) deleteEvent(dev DelEvent) error {
   uid := s.findEventUserId(dev.Id)
   if dev.UserId == uid {
       s.cache.RWMutex.Lock()
       delete(s.cache.Cache, dev.Id)
       s.cache.RWMutex.Unlock()
       return nil
   } else if uid == 0 {
       return fmt.Errorf("Event with this ID not found")
   } else {
       return fmt.Errorf("Wrong UserID")
   }
}

//=======DAY=======
func (s *Server) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
	respError(w, fmt.Errorf("Wrong HTTP Method"), http.StatusBadRequest)
	return
    }
    uid, err := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 64)
    if err != nil {
	respError(w, fmt.Errorf("Invalid user id"), http.StatusBadRequest)
    }
    date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
    if err != nil {
	respError(w, fmt.Errorf("Invalid date"), http.StatusBadRequest)
    }
    s.cache.RWMutex.RLock()
    defer s.cache.RWMutex.RUnlock()
    var evs []Event
    for _, ev := range s.cache.Cache {
	if ev.Date.date.Year() == date.Year() && 
	   ev.Date.date.Month() == date.Month() &&
	   ev.Date.date.Day() == date.Day() &&
	   ev.UserId == uid {
	       evs = append(evs, ev)
	   }
    }
    respResult(w, fmt.Sprintf("Events for day(%s)", date.Format("2006-01-02")), evs, http.StatusOK)
}

//=======WEEK=======
func (s *Server) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
	respError(w, fmt.Errorf("Wrong HTTP Method"), http.StatusBadRequest)
	return
    } 
    uid, err := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 64)
    if err != nil {
	respError(w, fmt.Errorf("Invalid user id"), http.StatusBadRequest)
    }
    date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
    if err != nil {
	respError(w, fmt.Errorf("Invalid date"), http.StatusBadRequest)
    }
    s.cache.RWMutex.RLock()
    defer s.cache.RWMutex.RUnlock()
    var evs []Event
    year, week := date.ISOWeek()
    for _, ev := range s.cache.Cache {
	evYear, evWeek := ev.Date.date.ISOWeek()
	if evYear == year &&
	   evWeek == week &&
	   ev.UserId == uid {
	       evs = append(evs, ev)
	   }
    }
    respResult(w, fmt.Sprintf("Events for week(%s)", date.Format("2006-01-02")), evs, http.StatusOK)
}

//=======MONTH=======
func (s *Server) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
	respError(w, fmt.Errorf("Wrong HTTP Method"), http.StatusBadRequest)
	return
    }
    uid, err := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 64)
    if err != nil {
	respError(w, fmt.Errorf("Invalid user id"), http.StatusBadRequest)
    }
    date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
    if err != nil {
	respError(w, fmt.Errorf("Invalid date"), http.StatusBadRequest)
    }
    s.cache.RWMutex.RLock()
    defer s.cache.RWMutex.RUnlock()
    var evs []Event
    for _, ev := range s.cache.Cache {
	if ev.Date.date.Year() == date.Year() && 
	   ev.Date.date.Month() == date.Month() &&
	   ev.UserId == uid {
	       evs = append(evs, ev)
	   }
    }
    respResult(w, fmt.Sprintf("Events for month(%s)", date.Format("2006-01")), evs, http.StatusOK)
}

