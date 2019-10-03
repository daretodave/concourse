// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/worker"
)

type FakeFetcher struct {
	FetchStub        func(context.Context, lager.Logger, db.ContainerMetadata, worker.Worker, worker.ContainerSpec, runtime.ProcessSpec, resource.Resource, atc.VersionedResourceTypes, db.ContainerOwner, string, worker.ImageFetchingDelegate, db.UsedResourceCache, string) (worker.GetResult, worker.Volume, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1  context.Context
		arg2  lager.Logger
		arg3  db.ContainerMetadata
		arg4  worker.Worker
		arg5  worker.ContainerSpec
		arg6  runtime.ProcessSpec
		arg7  resource.Resource
		arg8  atc.VersionedResourceTypes
		arg9  db.ContainerOwner
		arg10 string
		arg11 worker.ImageFetchingDelegate
		arg12 db.UsedResourceCache
		arg13 string
	}
	fetchReturns struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetcher) Fetch(arg1 context.Context, arg2 lager.Logger, arg3 db.ContainerMetadata, arg4 worker.Worker, arg5 worker.ContainerSpec, arg6 runtime.ProcessSpec, arg7 resource.Resource, arg8 atc.VersionedResourceTypes, arg9 db.ContainerOwner, arg10 string, arg11 worker.ImageFetchingDelegate, arg12 db.UsedResourceCache, arg13 string) (worker.GetResult, worker.Volume, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1  context.Context
		arg2  lager.Logger
		arg3  db.ContainerMetadata
		arg4  worker.Worker
		arg5  worker.ContainerSpec
		arg6  runtime.ProcessSpec
		arg7  resource.Resource
		arg8  atc.VersionedResourceTypes
		arg9  db.ContainerOwner
		arg10 string
		arg11 worker.ImageFetchingDelegate
		arg12 db.UsedResourceCache
		arg13 string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	fake.recordInvocation("Fetch", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.fetchReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeFetcher) FetchCalls(stub func(context.Context, lager.Logger, db.ContainerMetadata, worker.Worker, worker.ContainerSpec, runtime.ProcessSpec, resource.Resource, atc.VersionedResourceTypes, db.ContainerOwner, string, worker.ImageFetchingDelegate, db.UsedResourceCache, string) (worker.GetResult, worker.Volume, error)) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = stub
}

func (fake *FakeFetcher) FetchArgsForCall(i int) (context.Context, lager.Logger, db.ContainerMetadata, worker.Worker, worker.ContainerSpec, runtime.ProcessSpec, resource.Resource, atc.VersionedResourceTypes, db.ContainerOwner, string, worker.ImageFetchingDelegate, db.UsedResourceCache, string) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	argsForCall := fake.fetchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10, argsForCall.arg11, argsForCall.arg12, argsForCall.arg13
}

func (fake *FakeFetcher) FetchReturns(result1 worker.GetResult, result2 worker.Volume, result3 error) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetcher) FetchReturnsOnCall(i int, result1 worker.GetResult, result2 worker.Volume, result3 error) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 worker.GetResult
			result2 worker.Volume
			result3 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetcher) recordInvocation(key string, args []interface{}) {
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

var _ worker.Fetcher = new(FakeFetcher)