package structure

import (
	"fmt"
	"sort"
	"strings"
)

type KV struct {
	Key   string
	Value string
}

func (kv *KV) String() string {
	return fmt.Sprintf("%s=%s", kv.Key, kv.Value)
}

type KVList []KV

func (kvs *KVList) EncodeWithSort() string {
	sort.Slice(*kvs, func(i, j int) bool {
		return (*kvs)[i].Key < (*kvs)[j].Key
	})
	return kvs.Encode()
}

func (kvs *KVList) Set(key, value string) {
	*kvs = append(*kvs, KV{key, value})
}

func (kvs *KVList) Encode() string {
	list := make([]string, len(*kvs))
	for i, kv := range *kvs {
		list[i] = kv.String()
	}
	return strings.Join(list, "&")
}
