package auth_session

import (
	"encoding/json"
	"time"
)

type BaseSession struct {
	SessionID string `json:"session_id" gorm:"primary_key"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	ExpiredAt int64  `json:"expired_at"`
	Duration  int64  `json:"duration"`
}

type Session struct {
	BaseSession
	Data map[string]interface{} `json:"data"`
}

type SessionTable struct {
	BaseSession
	Data string
}

func (s *Session) Refresh() {
	s.UpdatedAt = time.Now().Unix()
	s.ExpiredAt = time.Now().Add(time.Duration(s.Duration) * time.Second).Unix()
}

func (s *Session) IsExpired() bool {
	return s.ExpiredAt < time.Now().Unix()
}

func (s *Session) DataToJSON() string {
	b, err := json.Marshal(s.Data)
	if err != nil {
		return ""
	}
	return string(b)
}

func (s *Session) JSONToData(jsonStr string) {
	err := json.Unmarshal([]byte(jsonStr), &s.Data)
	if err != nil {
		s.Data = make(map[string]interface{})
	}
}
