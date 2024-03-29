// Code generated by counterfeiter. DO NOT EDIT.
package datafakes

import (
	"chat/pkg/data"
	"chat/pkg/rest/models"
	"context"
	"sync"
)

type FakeDBClientInterface struct {
	AddNewUserStub        func(context.Context, *models.User) error
	addNewUserMutex       sync.RWMutex
	addNewUserArgsForCall []struct {
		arg1 context.Context
		arg2 *models.User
	}
	addNewUserReturns struct {
		result1 error
	}
	addNewUserReturnsOnCall map[int]struct {
		result1 error
	}
	GetChatroomByIdStub        func(context.Context, int32) (*models.Chatroom, error)
	getChatroomByIdMutex       sync.RWMutex
	getChatroomByIdArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	getChatroomByIdReturns struct {
		result1 *models.Chatroom
		result2 error
	}
	getChatroomByIdReturnsOnCall map[int]struct {
		result1 *models.Chatroom
		result2 error
	}
	GetUserByUsernameStub        func(context.Context, string) (*models.User, error)
	getUserByUsernameMutex       sync.RWMutex
	getUserByUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getUserByUsernameReturns struct {
		result1 *models.User
		result2 error
	}
	getUserByUsernameReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	GetUserChatroomsByUserIdStub        func(context.Context, int32) ([]*models.UserChatroom, error)
	getUserChatroomsByUserIdMutex       sync.RWMutex
	getUserChatroomsByUserIdArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	getUserChatroomsByUserIdReturns struct {
		result1 []*models.UserChatroom
		result2 error
	}
	getUserChatroomsByUserIdReturnsOnCall map[int]struct {
		result1 []*models.UserChatroom
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBClientInterface) AddNewUser(arg1 context.Context, arg2 *models.User) error {
	fake.addNewUserMutex.Lock()
	ret, specificReturn := fake.addNewUserReturnsOnCall[len(fake.addNewUserArgsForCall)]
	fake.addNewUserArgsForCall = append(fake.addNewUserArgsForCall, struct {
		arg1 context.Context
		arg2 *models.User
	}{arg1, arg2})
	stub := fake.AddNewUserStub
	fakeReturns := fake.addNewUserReturns
	fake.recordInvocation("AddNewUser", []interface{}{arg1, arg2})
	fake.addNewUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDBClientInterface) AddNewUserCallCount() int {
	fake.addNewUserMutex.RLock()
	defer fake.addNewUserMutex.RUnlock()
	return len(fake.addNewUserArgsForCall)
}

func (fake *FakeDBClientInterface) AddNewUserCalls(stub func(context.Context, *models.User) error) {
	fake.addNewUserMutex.Lock()
	defer fake.addNewUserMutex.Unlock()
	fake.AddNewUserStub = stub
}

func (fake *FakeDBClientInterface) AddNewUserArgsForCall(i int) (context.Context, *models.User) {
	fake.addNewUserMutex.RLock()
	defer fake.addNewUserMutex.RUnlock()
	argsForCall := fake.addNewUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDBClientInterface) AddNewUserReturns(result1 error) {
	fake.addNewUserMutex.Lock()
	defer fake.addNewUserMutex.Unlock()
	fake.AddNewUserStub = nil
	fake.addNewUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBClientInterface) AddNewUserReturnsOnCall(i int, result1 error) {
	fake.addNewUserMutex.Lock()
	defer fake.addNewUserMutex.Unlock()
	fake.AddNewUserStub = nil
	if fake.addNewUserReturnsOnCall == nil {
		fake.addNewUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addNewUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBClientInterface) GetChatroomById(arg1 context.Context, arg2 int32) (*models.Chatroom, error) {
	fake.getChatroomByIdMutex.Lock()
	ret, specificReturn := fake.getChatroomByIdReturnsOnCall[len(fake.getChatroomByIdArgsForCall)]
	fake.getChatroomByIdArgsForCall = append(fake.getChatroomByIdArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	stub := fake.GetChatroomByIdStub
	fakeReturns := fake.getChatroomByIdReturns
	fake.recordInvocation("GetChatroomById", []interface{}{arg1, arg2})
	fake.getChatroomByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDBClientInterface) GetChatroomByIdCallCount() int {
	fake.getChatroomByIdMutex.RLock()
	defer fake.getChatroomByIdMutex.RUnlock()
	return len(fake.getChatroomByIdArgsForCall)
}

func (fake *FakeDBClientInterface) GetChatroomByIdCalls(stub func(context.Context, int32) (*models.Chatroom, error)) {
	fake.getChatroomByIdMutex.Lock()
	defer fake.getChatroomByIdMutex.Unlock()
	fake.GetChatroomByIdStub = stub
}

func (fake *FakeDBClientInterface) GetChatroomByIdArgsForCall(i int) (context.Context, int32) {
	fake.getChatroomByIdMutex.RLock()
	defer fake.getChatroomByIdMutex.RUnlock()
	argsForCall := fake.getChatroomByIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDBClientInterface) GetChatroomByIdReturns(result1 *models.Chatroom, result2 error) {
	fake.getChatroomByIdMutex.Lock()
	defer fake.getChatroomByIdMutex.Unlock()
	fake.GetChatroomByIdStub = nil
	fake.getChatroomByIdReturns = struct {
		result1 *models.Chatroom
		result2 error
	}{result1, result2}
}

func (fake *FakeDBClientInterface) GetChatroomByIdReturnsOnCall(i int, result1 *models.Chatroom, result2 error) {
	fake.getChatroomByIdMutex.Lock()
	defer fake.getChatroomByIdMutex.Unlock()
	fake.GetChatroomByIdStub = nil
	if fake.getChatroomByIdReturnsOnCall == nil {
		fake.getChatroomByIdReturnsOnCall = make(map[int]struct {
			result1 *models.Chatroom
			result2 error
		})
	}
	fake.getChatroomByIdReturnsOnCall[i] = struct {
		result1 *models.Chatroom
		result2 error
	}{result1, result2}
}

func (fake *FakeDBClientInterface) GetUserByUsername(arg1 context.Context, arg2 string) (*models.User, error) {
	fake.getUserByUsernameMutex.Lock()
	ret, specificReturn := fake.getUserByUsernameReturnsOnCall[len(fake.getUserByUsernameArgsForCall)]
	fake.getUserByUsernameArgsForCall = append(fake.getUserByUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetUserByUsernameStub
	fakeReturns := fake.getUserByUsernameReturns
	fake.recordInvocation("GetUserByUsername", []interface{}{arg1, arg2})
	fake.getUserByUsernameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDBClientInterface) GetUserByUsernameCallCount() int {
	fake.getUserByUsernameMutex.RLock()
	defer fake.getUserByUsernameMutex.RUnlock()
	return len(fake.getUserByUsernameArgsForCall)
}

func (fake *FakeDBClientInterface) GetUserByUsernameCalls(stub func(context.Context, string) (*models.User, error)) {
	fake.getUserByUsernameMutex.Lock()
	defer fake.getUserByUsernameMutex.Unlock()
	fake.GetUserByUsernameStub = stub
}

func (fake *FakeDBClientInterface) GetUserByUsernameArgsForCall(i int) (context.Context, string) {
	fake.getUserByUsernameMutex.RLock()
	defer fake.getUserByUsernameMutex.RUnlock()
	argsForCall := fake.getUserByUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDBClientInterface) GetUserByUsernameReturns(result1 *models.User, result2 error) {
	fake.getUserByUsernameMutex.Lock()
	defer fake.getUserByUsernameMutex.Unlock()
	fake.GetUserByUsernameStub = nil
	fake.getUserByUsernameReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeDBClientInterface) GetUserByUsernameReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.getUserByUsernameMutex.Lock()
	defer fake.getUserByUsernameMutex.Unlock()
	fake.GetUserByUsernameStub = nil
	if fake.getUserByUsernameReturnsOnCall == nil {
		fake.getUserByUsernameReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.getUserByUsernameReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeDBClientInterface) GetUserChatroomsByUserId(arg1 context.Context, arg2 int32) ([]*models.UserChatroom, error) {
	fake.getUserChatroomsByUserIdMutex.Lock()
	ret, specificReturn := fake.getUserChatroomsByUserIdReturnsOnCall[len(fake.getUserChatroomsByUserIdArgsForCall)]
	fake.getUserChatroomsByUserIdArgsForCall = append(fake.getUserChatroomsByUserIdArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	stub := fake.GetUserChatroomsByUserIdStub
	fakeReturns := fake.getUserChatroomsByUserIdReturns
	fake.recordInvocation("GetUserChatroomsByUserId", []interface{}{arg1, arg2})
	fake.getUserChatroomsByUserIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDBClientInterface) GetUserChatroomsByUserIdCallCount() int {
	fake.getUserChatroomsByUserIdMutex.RLock()
	defer fake.getUserChatroomsByUserIdMutex.RUnlock()
	return len(fake.getUserChatroomsByUserIdArgsForCall)
}

func (fake *FakeDBClientInterface) GetUserChatroomsByUserIdCalls(stub func(context.Context, int32) ([]*models.UserChatroom, error)) {
	fake.getUserChatroomsByUserIdMutex.Lock()
	defer fake.getUserChatroomsByUserIdMutex.Unlock()
	fake.GetUserChatroomsByUserIdStub = stub
}

func (fake *FakeDBClientInterface) GetUserChatroomsByUserIdArgsForCall(i int) (context.Context, int32) {
	fake.getUserChatroomsByUserIdMutex.RLock()
	defer fake.getUserChatroomsByUserIdMutex.RUnlock()
	argsForCall := fake.getUserChatroomsByUserIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDBClientInterface) GetUserChatroomsByUserIdReturns(result1 []*models.UserChatroom, result2 error) {
	fake.getUserChatroomsByUserIdMutex.Lock()
	defer fake.getUserChatroomsByUserIdMutex.Unlock()
	fake.GetUserChatroomsByUserIdStub = nil
	fake.getUserChatroomsByUserIdReturns = struct {
		result1 []*models.UserChatroom
		result2 error
	}{result1, result2}
}

func (fake *FakeDBClientInterface) GetUserChatroomsByUserIdReturnsOnCall(i int, result1 []*models.UserChatroom, result2 error) {
	fake.getUserChatroomsByUserIdMutex.Lock()
	defer fake.getUserChatroomsByUserIdMutex.Unlock()
	fake.GetUserChatroomsByUserIdStub = nil
	if fake.getUserChatroomsByUserIdReturnsOnCall == nil {
		fake.getUserChatroomsByUserIdReturnsOnCall = make(map[int]struct {
			result1 []*models.UserChatroom
			result2 error
		})
	}
	fake.getUserChatroomsByUserIdReturnsOnCall[i] = struct {
		result1 []*models.UserChatroom
		result2 error
	}{result1, result2}
}

func (fake *FakeDBClientInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addNewUserMutex.RLock()
	defer fake.addNewUserMutex.RUnlock()
	fake.getChatroomByIdMutex.RLock()
	defer fake.getChatroomByIdMutex.RUnlock()
	fake.getUserByUsernameMutex.RLock()
	defer fake.getUserByUsernameMutex.RUnlock()
	fake.getUserChatroomsByUserIdMutex.RLock()
	defer fake.getUserChatroomsByUserIdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDBClientInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ data.DBClientInterface = new(FakeDBClientInterface)
