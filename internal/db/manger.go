package db

import (
	"fmt"
	"gin-gonic-gorm-boilerplate/configs"
	"gin-gonic-gorm-boilerplate/internal/model"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
	"gorm.io/gorm"
	"math/rand"
)

type Manager struct {
	Master   *gorm.DB
	Replicas []*gorm.DB
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Init(master configs.DBConfig, replicas []configs.DBConfig) {
	var err error
	m.Master, err = Connection(master)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to Master database: %v", err))
	}
	m.Master.AutoMigrate(&model.User{})

	var reader *gorm.DB
	for i, r := range replicas {
		reader, err = Connection(r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to connect to Replica #%d: %s database: %v", i, r.Host, err))
		}
		m.Replicas = append(m.Replicas, reader)
	}
}

func (m *Manager) Writer() *gorm.DB {
	return m.Master
}

func (m *Manager) Reader() *gorm.DB {
	return m.Replicas[rand.Intn(len(m.Replicas))]
}

func (m *Manager) ReaderChoice(i int) *gorm.DB {
	if i >= len(m.Replicas) {
		return m.Replicas[0]
	}
	return m.Replicas[i]
}

func (m *Manager) Close() error {
	var err error

	err = Close(m.Master)

	for _, r := range m.Replicas {
		err = Close(r)
	}

	return err
}
