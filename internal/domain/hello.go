package domain

import "time"

type Hello struct {
	ID        string
	Message   string
	SaidAt    time.Time
}

type HelloSaid struct {
	ID      string
	Message string
	SaidAt  time.Time
}

func (h *Hello) Apply(event HelloSaid) {
	h.ID = event.ID
	h.Message = event.Message
	h.SaidAt = event.SaidAt
}

func ParseTime(val interface{}) time.Time {
    switch v := val.(type) {
    case string:
        t, _ := time.Parse(time.RFC3339, v)
        return t
    case time.Time:
        return v
    default:
        return time.Time{}
    }
}
