// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	api "github.com/ava-labs/avalanche-network-runner/api"
	admin "github.com/ava-labs/avalanchego/api/admin"

	avm "github.com/ava-labs/avalanchego/vms/avm"

	evm "github.com/ava-labs/coreth/plugin/evm"

	health "github.com/ava-labs/avalanchego/api/health"

	indexer "github.com/ava-labs/avalanchego/indexer"

	info "github.com/ava-labs/avalanchego/api/info"

	keystore "github.com/ava-labs/avalanchego/api/keystore"

	mock "github.com/stretchr/testify/mock"

	platformvm "github.com/ava-labs/avalanchego/vms/platformvm"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AdminAPI provides a mock function with given fields:
func (_m *Client) AdminAPI() admin.Client {
	ret := _m.Called()

	var r0 admin.Client
	if rf, ok := ret.Get(0).(func() admin.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(admin.Client)
		}
	}

	return r0
}

// CChainAPI provides a mock function with given fields:
func (_m *Client) CChainAPI() evm.Client {
	ret := _m.Called()

	var r0 evm.Client
	if rf, ok := ret.Get(0).(func() evm.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evm.Client)
		}
	}

	return r0
}

// CChainEthAPI provides a mock function with given fields:
func (_m *Client) CChainEthAPI() api.EthClient {
	ret := _m.Called()

	var r0 api.EthClient
	if rf, ok := ret.Get(0).(func() api.EthClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.EthClient)
		}
	}

	return r0
}

// CChainIndexAPI provides a mock function with given fields:
func (_m *Client) CChainIndexAPI() indexer.Client {
	ret := _m.Called()

	var r0 indexer.Client
	if rf, ok := ret.Get(0).(func() indexer.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(indexer.Client)
		}
	}

	return r0
}

// HealthAPI provides a mock function with given fields:
func (_m *Client) HealthAPI() health.Client {
	ret := _m.Called()

	var r0 health.Client
	if rf, ok := ret.Get(0).(func() health.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(health.Client)
		}
	}

	return r0
}

// InfoAPI provides a mock function with given fields:
func (_m *Client) InfoAPI() info.Client {
	ret := _m.Called()

	var r0 info.Client
	if rf, ok := ret.Get(0).(func() info.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(info.Client)
		}
	}

	return r0
}

// KeystoreAPI provides a mock function with given fields:
func (_m *Client) KeystoreAPI() keystore.Client {
	ret := _m.Called()

	var r0 keystore.Client
	if rf, ok := ret.Get(0).(func() keystore.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Client)
		}
	}

	return r0
}

// PChainAPI provides a mock function with given fields:
func (_m *Client) PChainAPI() platformvm.Client {
	ret := _m.Called()

	var r0 platformvm.Client
	if rf, ok := ret.Get(0).(func() platformvm.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(platformvm.Client)
		}
	}

	return r0
}

// PChainIndexAPI provides a mock function with given fields:
func (_m *Client) PChainIndexAPI() indexer.Client {
	ret := _m.Called()

	var r0 indexer.Client
	if rf, ok := ret.Get(0).(func() indexer.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(indexer.Client)
		}
	}

	return r0
}

// XChainAPI provides a mock function with given fields:
func (_m *Client) XChainAPI() avm.Client {
	ret := _m.Called()

	var r0 avm.Client
	if rf, ok := ret.Get(0).(func() avm.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(avm.Client)
		}
	}

	return r0
}

// XChainWalletAPI provides a mock function with given fields:
func (_m *Client) XChainWalletAPI() avm.WalletClient {
	ret := _m.Called()

	var r0 avm.WalletClient
	if rf, ok := ret.Get(0).(func() avm.WalletClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(avm.WalletClient)
		}
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
