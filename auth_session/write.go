package auth_session

import (
	"time"

	"github.com/go-ginseng/ginseng-plugin-core/sql_mem"
)

func DeleteExpiredSession() error {
	return sql_mem.MEM.Delete(&SessionTable{}, "expired_at < ?", time.Now().Unix()).Error
}

func CreateSession(sessionID string, duration int64, data map[string]interface{}) (*Session, error) {
	s := &Session{
		BaseSession: BaseSession{
			SessionID: sessionID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			Duration:  duration,
			ExpiredAt: time.Now().Add(time.Duration(duration) * time.Second).Unix(),
		},
		Data: data,
	}
	err := sql_mem.MEM.Create(&SessionTable{
		BaseSession: BaseSession{
			SessionID: s.SessionID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
			ExpiredAt: s.ExpiredAt,
			Duration:  s.Duration,
		},
		Data: s.DataToJSON(),
	}).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}
