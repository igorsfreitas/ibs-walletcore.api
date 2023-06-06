package createtransaction

import (
	"testing"

	"github.com/igorsfreitas/ibs-walletcore.api/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) GetByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Client1", "client1@teste.com.br")
	account1 := entity.NewAccount(client1)
	account1.Deposit(1000)

	client2, _ := entity.NewClient("Client2", "client2@teste.com.br")
	account2 := entity.NewAccount(client2)
	account2.Deposit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("GetByID", account1.ID).Return(account1, nil)
	mockAccount.On("GetByID", account2.ID).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := &CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount)
	outputDto, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, outputDto)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "GetByID", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
