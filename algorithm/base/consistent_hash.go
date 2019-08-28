package base

import (
	"errors"
	"hash/crc32"
	"sort"
	"sync"
)

// 一致性hash

// TODO 考虑和其他算法的合并，例如和最少链接策略共同作用，防止过载
// TODO 考虑添加虚拟节点，

type ConsistentHash struct {
	nums      int
	sortNodes []uint32 // TODO 排序好的节点，可以使用其他数据结构代替
	nodes     map[uint32]string
	mu        sync.Mutex
}

func New() *ConsistentHash {
	return &ConsistentHash{
		nums:      0,
		sortNodes: make([]uint32, 0),
		nodes:     make(map[uint32]string)}
}

func hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *ConsistentHash) Add(host string) (uint32, error) {

	hash := hashKey(host)

	if _, ok := c.nodes[hash]; ok {
		return 1, errors.New("this key has exist")
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.nums += 1
	c.nodes[hash] = host

	c.sortNodes = append(c.sortNodes, hash)

	sort.Slice(c.sortNodes, func(i, j int) bool {
		return c.sortNodes[i] < c.sortNodes[j]
	})
	return hash, nil
}

// 查询节点
func (c *ConsistentHash) Get(key string) (string, error) {
	if c.nums <= 0 {
		return "", errors.New("nodes is empty")
	}

	// 找到最接近的一个key
	nodeIdx := searchNearNode(c.sortNodes, hashKey(key))
	return c.nodes[nodeIdx], nil
}

func searchNearNode(sortNodes []uint32, h uint32) uint32 {
	for _, v := range sortNodes {
		if h <= v {
			return v
		}
	}
	return sortNodes[0]
}

// 删除节点
func (c *ConsistentHash) Remove(host string) error {

	if len(host) <= 0 {
		return errors.New("host can't be blank")
	}

	hash := hashKey(host)
	if _, ok := c.nodes[hash]; ok {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.nums -= 1
		delete(c.nodes, hash)
		c.sortNodes = reSlice(c.sortNodes, hash)
		return nil
	} else {
		return errors.New("host not exist ")
	}
}

func reSlice(sortNodes []uint32, hash uint32) []uint32 {

	for i, v := range sortNodes {
		if v == hash {
			return append(sortNodes[:i], sortNodes[i+1:]...)
		}
	}

	return sortNodes
}
