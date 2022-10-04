package sys_config

import (
	"encoding/json"
	"errors"

	"github.com/go-ginseng/sql"
	"gorm.io/gorm"
)

/**
 * 	System config provides a standard way to configure the system.
 *	Each config is a key-value pair. This package ensures that the key is unique.
 *	The value can be string or a struct (will be serialized to json).
 *	The config can be separated into different groups. (groupID)
 *	Each group can be separated into different users. (userID)
 *	The find logic is:
 *		has (groupID + userID) ? -> has (groupID) ? -> has (default) ?
 *	This package will not handle the security of the config. It should be implemented by the project.
 */

/**
 *	| groupID | userID | level
 *	|---------|--------|-------
 *	| 0       | 0      | default
 *	| 1       | 0      | group
 *	| 1       | 1      | user
 * 	| 0       | 1      | X (not allowed)
 */

type SystemConfig struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	GroupID uint   `json:"group_id"`
	UserID  uint   `json:"user_id"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

// Register the config key-value pair
func Register(key string, defaultValue string) {
	check, _ := _findByGroupIDAndUserIDAndKey(db, 0, 0, key)
	if check == nil {
		config := &SystemConfig{
			GroupID: 0,
			UserID:  0,
			Key:     key,
			Value:   defaultValue,
		}
		_, err := sql.Create(db, config)
		if err != nil {
			panic(err)
		}
	}
}

// Register the config struct to the engine
func RegisterStruct(key string, defaultValue interface{}) {
	value, err := _structToJSON(defaultValue)
	if err != nil {
		panic(err)
	}
	Register(key, value)
}

// Set the config value
func Set(groupID uint, userID uint, key string, value string) error {
	if groupID == 0 && userID != 0 {
		return errors.New("user should belongs to a group")
	}
	config, _ := _findByGroupIDAndUserIDAndKey(db, groupID, userID, key)
	if config == nil {
		config = &SystemConfig{
			GroupID: groupID,
			UserID:  userID,
			Key:     key,
			Value:   value,
		}
		_, err := sql.Create(db, config)
		if err != nil {
			return err
		}
	} else {
		config.Value = value
		_, err := sql.Update(db, config)
		if err != nil {
			return err
		}
	}
	SyncMem()
	return nil
}

// Set the config struct
func SetStruct(groupID uint, userID uint, key string, value interface{}) error {
	j, err := _structToJSON(value)
	if err != nil {
		return err
	}
	return Set(groupID, userID, key, j)
}

// Get the config value
func Get(groupID uint, userID uint, key string) (string, error) {
	config, _ := _findByGroupIDAndUserIDAndKey(mem, groupID, userID, key)

	if config == nil && userID != 0 {
		config, _ = _findByGroupIDAndUserIDAndKey(mem, groupID, 0, key)
	}

	if config == nil && groupID != 0 {
		config, _ = _findByGroupIDAndUserIDAndKey(mem, 0, 0, key)
	}

	if config == nil {
		return "", errors.New("config not found")
	}

	return config.Value, nil
}

// Get the config struct
func GetStruct[T any](groupID uint, userID uint, key string) (*T, error) {
	value, err := Get(groupID, userID, key)
	if err != nil {
		return nil, err
	}
	return _jsonToStruct[T](value)
}

func SyncMem() {
	configs, err := sql.FindAll[SystemConfig](db, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	for _, config := range configs {
		err := mem.Save(&config).Error
		if err != nil {
			panic(err)
		}
	}
}

func _findByGroupIDAndUserIDAndKey(tx *gorm.DB, groupID uint, userID uint, key string) (*SystemConfig, error) {
	return sql.FindOne[SystemConfig](tx, sql.And(sql.Eq("group_id", groupID), sql.Eq("user_id", userID), sql.Eq("`key`", key)))
}

func _structToJSON(value interface{}) (string, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func _jsonToStruct[T any](value string) (*T, error) {
	result := new(T)
	err := json.Unmarshal([]byte(value), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
