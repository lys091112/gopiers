package redisson

import "github.com/gomodule/redigo/redis"

// Set key value
func (r *Redisson) Set(key,value string) error {
	_, err := r.Do("SET",key,value)
	return err
}

// Get key
func (r *Redisson) Get(key string) (string,error) {
	b , err := redis.Bytes(r.Do("GET",key))
	if err != nil {
		return "" , err
	}

	return string(b),nil
}

// Append append value to key's origin value
// Integer reply: the length of the string after the append operation.
func (r *Redisson) Append(key,appedValue string)(int, error) {
	return redis.Int(r.Do("APPEND", key,appedValue))
}

// BitCount Count the number of set bits (population counting) in a string O(N)
// Integer reply The number of bits set to 1.
// 统计字符串中二进制为1的个数，BitMap位图使用
func (r *Redisson) BitCount(key string)(int, error) {
	return redis.Int(r.Do("BITCOUNT",key))
}

// BitCountWithRange same as BitCount
func (r *Redisson) BitCountWithRange(key string,start,end int)(int, error) {
	return redis.Int(r.Do("BITCOUNT",key,start,end))
}

// SetBit  在指定偏移offsert位设置值
// Integer reply: the original bit value stored at offset.
func (r *Redisson) SetBit(key string,offset,value int) (int,error) {
	return redis.Int(r.Do("SETBIT", key,offset,value))
}

// GetBit  获取指定偏移offsert位值
// Integer reply: the original bit value stored at offset.
func (r *Redisson) GetBit(key string,offset int) (int,error) {
	return redis.Int(r.Do("GETBIT", key,offset))
}

