/**
 * @Author: Resynz
 * @Date: 2021/12/28 11:32
 */
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	ConfMaps     *ConfMap
	Mode         string
	Port         uint16
	ConfPath     string
	ConfFileName string
	Envs         []string
)

func initConfMaps() error {
	dirs, err := os.ReadDir(ConfPath)
	if err != nil {
		return err
	}
	for _, v := range dirs {
		env := v.Name()
		Envs = append(Envs, env)
		fp := fmt.Sprintf("%s/%s/%s", ConfPath, env, ConfFileName)
		con, err := os.ReadFile(fp)
		if err != nil {
			return err
		}
		var r map[string]interface{}
		if err = json.Unmarshal(con, &r); err != nil {
			return err
		}
		ConfMaps.SetByEnv(env, r)
	}
	return nil
}

func InitEnv() error {
	Mode = "debug"
	if d := os.Getenv("Mode"); d != "" {
		Mode = d
	}
	Port = 8080
	if d, err := strconv.Atoi(os.Getenv("Port")); err == nil && d > 0 && d < 65536 {
		Port = uint16(d)
	}
	ConfPath = "./configs"
	if d := os.Getenv("ConfPath"); d != "" {
		ConfPath = d
	}
	ConfFileName = "config.json"
	if d := os.Getenv("ConfFileName"); d != "" {
		ConfFileName = d
	}

	Envs = make([]string, 0)

	ConfMaps = &ConfMap{
		Mutex: &sync.Mutex{},
		maps:  make(map[string]map[string]interface{}),
	}
	if err := initConfMaps(); err != nil {
		return err
	}
	return nil
}
