package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ModelStats struct {
	Count int64
}

func (m *Model) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now
	return nil
}

func (m *Model) PreUpdate(s gorp.SqlExecutor) error {
	m.UpdatedAt = time.Now()
	return nil
}
