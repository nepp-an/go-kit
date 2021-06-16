package main

import (
"fmt"
"sync"
)

type PushToken struct {
    id string
    time int64
}

type TokenCache struct {
    token        map[string]*PushToken
    lock         sync.Mutex
    lastModified int64
}

var AllCLToken []*TokenCache
var CLIndex = make(map[string]int)

func NewAllCLTokenCache(SupportCL []string) {
    for i := range SupportCL {
        CLIndex[SupportCL[i]] = i+1
        AllCLToken = append(AllCLToken, &TokenCache{
            make(map[string]*PushToken),
            sync.Mutex{},
            0,
        })
    }
}

func AddToken(CL, id string, token *PushToken) {
    index := CLIndex[CL]
    if index == 0 {
        return
    }
    tokenCache := AllCLToken[index-1]
    tokenCache.lock.Lock()
    defer tokenCache.lock.Unlock()
    tokenCache.token[id] = token
    tokenCache.lastModified = token.time
}

func getLastModified(CL string) int64 {
    index := CLIndex[CL]
    if index == 0 {
        return 0
    }
    tokenCache := AllCLToken[index-1]
    tokenCache.lock.Lock()
    defer tokenCache.lock.Unlock()
    return tokenCache.lastModified
}

func rangeTokenCache(CL string) {
    index := CLIndex[CL]
    if index == 0 {
        return
    }
    tokenCache := AllCLToken[index-1]
    tokenCache.lock.Lock()
    defer tokenCache.lock.Unlock()
    for i := range tokenCache.token {
        fmt.Printf("%v", tokenCache.token[i])
    }
}


