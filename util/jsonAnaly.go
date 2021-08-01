package util

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
)

type model struct {
	Id           string `json:"id"` //士兵id
	Rarity       string //士兵稀有度
	CombatPoints string //战力
	Cvc          string //版本
	UnlockArena  string //当前解锁阶段

}

// JsonAnaly 通过pflag获取路径
func JsonAnaly() {
	path1 := pflag.String("path1", "../conf/util.army.model.json", "old json path")
	path2 := pflag.String("path2", "../conf/news.json", "new json path")
	pflag.Parse()
	var s1, s2 string = *path1, *path2
	//将路径传参
	JsonAnaly1(s1, s2)
}

// JsonAnaly1 根据路径优化json文件
func JsonAnaly1(path1, path2 string) {
	bytes, err1 := ioutil.ReadFile(path1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	m := make(map[string]model)
	err2 := json.Unmarshal(bytes, &m)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	mjson, err3 := json.Marshal(&m)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	err4 := ioutil.WriteFile(path2, mjson, 0666)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
}
