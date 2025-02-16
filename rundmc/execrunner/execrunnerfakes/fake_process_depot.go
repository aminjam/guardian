// Code generated by counterfeiter. DO NOT EDIT.
package execrunnerfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/execrunner"
	"code.cloudfoundry.org/lager"
)

type FakeProcessDepot struct {
	CreateProcessDirStub        func(lager.Logger, string, string) (string, error)
	createProcessDirMutex       sync.RWMutex
	createProcessDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	createProcessDirReturns struct {
		result1 string
		result2 error
	}
	createProcessDirReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessDepot) CreateProcessDir(arg1 lager.Logger, arg2 string, arg3 string) (string, error) {
	fake.createProcessDirMutex.Lock()
	ret, specificReturn := fake.createProcessDirReturnsOnCall[len(fake.createProcessDirArgsForCall)]
	fake.createProcessDirArgsForCall = append(fake.createProcessDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateProcessDir", []interface{}{arg1, arg2, arg3})
	fake.createProcessDirMutex.Unlock()
	if fake.CreateProcessDirStub != nil {
		return fake.CreateProcessDirStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createProcessDirReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcessDepot) CreateProcessDirCallCount() int {
	fake.createProcessDirMutex.RLock()
	defer fake.createProcessDirMutex.RUnlock()
	return len(fake.createProcessDirArgsForCall)
}

func (fake *FakeProcessDepot) CreateProcessDirCalls(stub func(lager.Logger, string, string) (string, error)) {
	fake.createProcessDirMutex.Lock()
	defer fake.createProcessDirMutex.Unlock()
	fake.CreateProcessDirStub = stub
}

func (fake *FakeProcessDepot) CreateProcessDirArgsForCall(i int) (lager.Logger, string, string) {
	fake.createProcessDirMutex.RLock()
	defer fake.createProcessDirMutex.RUnlock()
	argsForCall := fake.createProcessDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeProcessDepot) CreateProcessDirReturns(result1 string, result2 error) {
	fake.createProcessDirMutex.Lock()
	defer fake.createProcessDirMutex.Unlock()
	fake.CreateProcessDirStub = nil
	fake.createProcessDirReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) CreateProcessDirReturnsOnCall(i int, result1 string, result2 error) {
	fake.createProcessDirMutex.Lock()
	defer fake.createProcessDirMutex.Unlock()
	fake.CreateProcessDirStub = nil
	if fake.createProcessDirReturnsOnCall == nil {
		fake.createProcessDirReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createProcessDirReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createProcessDirMutex.RLock()
	defer fake.createProcessDirMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessDepot) recordInvocation(key string, args []interface{}) {
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

var _ execrunner.ProcessDepot = new(FakeProcessDepot)
