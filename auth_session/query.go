package auth_session

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CurrentSession(ctx *gin.Context) *Session {
	token := strings.TrimPrefix(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		return nil
	}
	session, err := GetSession(token)
	if err != nil {
		return nil
	}
	if session.IsExpired() {
		return nil
	}
	return session
}

func GetSession(sessionID string) (*Session, error) {
	s := &SessionTable{}
	err := db.First(s, "session_id = ?", sessionID).Error
	if err != nil {
		return nil, err
	}
	session := &Session{
		BaseSession: BaseSession{
			SessionID: s.SessionID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
			ExpiredAt: s.ExpiredAt,
			Duration:  s.Duration,
		},
		Data: make(map[string]interface{}),
	}
	session.JSONToData(s.Data)
	return session, nil
}
