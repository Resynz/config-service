/**
 * @Author: Resynz
 * @Date: 2021/12/28 11:29
 */
package config

import "sync"

type ConfMap struct {
	*sync.Mutex
	maps map[string]map[string]interface{}
}

func (c *ConfMap) GetByEnv(env string) map[string]interface{} {
	val, ok := c.maps[env]
	if !ok {
		return nil
	}
	return val
}

func (c *ConfMap) SetByEnv(env string, val map[string]interface{}) {
	c.Lock()
	defer c.Unlock()
	c.maps[env] = val
}

func (c *ConfMap) Get(env, key string) interface{} {
	data, ok := c.maps[env]
	if !ok {
		return nil
	}
	val, ok := data[key]
	if !ok {
		return nil
	}
	return val
}

func (c *ConfMap) Set(env, key string, val interface{}) {
	c.Lock()
	defer c.Unlock()
	data, ok := c.maps[env]
	if !ok {
		data = make(map[string]interface{})
	}
	data[key] = val
	c.maps[env] = data
}
