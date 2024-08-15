package main
import (
    "net/http"
    "encoding/json"
    "time"
    "log"
    "os"
    "sync"
)
/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Event struct {
    Id      uint64  `json:"id"`
    UserId  uint64  `json:"user_id"`
    Name    string  `json:"name"`
    Date    nDate  `json:"date"`
}

type nDate struct {
    date time.Time
}

func (d *nDate) MarshalJSON() ([]byte, error) {
    s := d.date.Format("2006-01-02")
    return json.Marshal(s)
}

func (d *nDate) UnmarshalJSON(input []byte) error {
    var err error
    d.date, err = time.Parse(`"2006-01-02"`, string(input))
    return err
}

type EventCache struct {
    RWMutex *sync.RWMutex
    Cache map[uint64]Event
}

func MiddlewareLog(hf http.HandlerFunc) http.HandlerFunc{
    return http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
	    hf.ServeHTTP(w, r)
	    log.Printf("Method: %s, Path: %s\n", r.Method, r.URL.Path)
	},
    )
}

type Server struct {
    cache EventCache
    mux *http.ServeMux
}

func NewServer() *Server {
    return &Server{
	cache: EventCache{
	    RWMutex: new(sync.RWMutex),
	    Cache: make(map[uint64]Event),
	},
	mux: new(http.ServeMux),
    }
}

func (s *Server) Start(port string) {
    s.mux.HandleFunc("/create_event", MiddlewareLog(s.CreateEventHandler))
    s.mux.HandleFunc("/update_event", MiddlewareLog(s.UpdateEventHandler))
    s.mux.HandleFunc("/delete_event", MiddlewareLog(s.DeleteEventHandler))

    s.mux.HandleFunc("/events_for_day", MiddlewareLog(s.EventsForDayHandler))
    s.mux.HandleFunc("/events_for_week", MiddlewareLog(s.EventsForWeekHandler))
    s.mux.HandleFunc("/events_for_month", MiddlewareLog(s.EventsForMonthHandler))
    log.Fatal(http.ListenAndServe("localhost:"+port, s.mux))
}

type Config struct {
    Port string `json:"port"`
}

func main() {
    srv := NewServer()
    file, err := os.Open("config.json")
    if err != nil {
	log.Fatal(err)
    }
    decoder := json.NewDecoder(file)
    cfg := Config{}
    err = decoder.Decode(&cfg)
    if err != nil {
	log.Fatal(err)
    }
    file.Close()
    srv.Start(cfg.Port)
}
