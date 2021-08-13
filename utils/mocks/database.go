package mocks

import (
	"github.com/henrique502/go-repo-seed/domain"
	"github.com/jackc/pgconn"
)

type DatabaseMock struct {
	ConnectFunc           func()
	CloseFunc             func()
	AlertUpSertFunc       func(alert domain.Alert) pgconn.CommandTag
	IntegrationUpSertFunc func(integration domain.Integration) pgconn.CommandTag
	TeamUpSertFunc        func(team domain.Team) pgconn.CommandTag
}

func (m *DatabaseMock) Connect() {
	m.CloseFunc()
}

func (m *DatabaseMock) Close() {
	m.CloseFunc()
}

func (m *DatabaseMock) AlertUpSert(alert domain.Alert) pgconn.CommandTag {
	return m.AlertUpSertFunc(alert)
}

func (m *DatabaseMock) IntegrationUpSert(alert domain.Alert) pgconn.CommandTag {
	return m.IntegrationUpSert(alert)
}

func (m *DatabaseMock) TeamUpSert(alert domain.Alert) pgconn.CommandTag {
	return m.TeamUpSert(alert)
}
