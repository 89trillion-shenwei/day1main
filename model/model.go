package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type model struct {
	Id           string `json:"id"` //士兵id
	Rarity       string //士兵稀有度
	CombatPoints string //战力
	Cvc          string //版本
	UnlockArena  string //当前解锁阶段

}

func GetMap() (mm map[string]model) {
	bytes, err1 := ioutil.ReadFile("../conf/new.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	m := make(map[string]model)
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	return m
}
