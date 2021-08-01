package handler

import (
	"encoding/json"
	"reflect"
	"testing"
	"unsafe"
)

type model struct {
	Id           string `json:"id"` //士兵id
	Rarity       string //士兵稀有度
	CombatPoints string //战力
	Cvc          string //版本
	UnlockArena  string //当前解锁阶段

}

//string转byte[]
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
func TestFindSoldierByRaUnCv(t *testing.T) {
	arity, unlockArena, cvc := "2", "3", "100"
	ret, _ := FindSoldierByRaUnCv(arity, unlockArena, cvc)
	retBytes := String2Bytes(ret)
	m := make(map[string]string)
	_ = json.Unmarshal(retBytes, &m)
	for _, value := range m {
		m2 := model{}
		vs := String2Bytes(value)
		_ = json.Unmarshal(vs, &m2)
		if m2.Rarity == arity && m2.UnlockArena <= unlockArena && m2.Cvc == cvc {
			t.Log("success")
			return
		} else {
			t.Log("failed")
			t.Error("failed")
		}

	}

}
func TestFindSoldierRaById(t *testing.T) {
	id := "13206"
	ra := "3"
	ret := FindSoldierRaById(id)
	if ra != ret {
		t.Error(ret)
	}
}

func TestFindSoldierCoById(t *testing.T) {
	id := "13206"
	co := "3000"
	ret := FindSoldierCoById(id)
	if co != ret {
		t.Error(ret)
	}
}

func TestFindSoldierByCv(t *testing.T) {
	cvc := "1000"
	ret, _ := FindSoldierByCv(cvc)
	retBytes := String2Bytes(ret)
	m := make(map[string]string)
	_ = json.Unmarshal(retBytes, &m)
	for _, value := range m {
		m2 := model{}
		vs := String2Bytes(value)
		_ = json.Unmarshal(vs, &m2)
		if m2.Cvc == cvc {
			//t.Log("success")
			return
		} else {
			//t.Log("failed")
			t.Error("failed")
		}

	}
}

func TestFindSoldierByUn(t *testing.T) {
	un := "1"
	ret, _ := FindSoldierByUn()
	retBytes := String2Bytes(ret)
	m := map[string][]string{}
	_ = json.Unmarshal(retBytes, &m)
	m2 := model{}
	s := "UnlockArena: "
	vs := String2Bytes(m[s+un][0])
	_ = json.Unmarshal(vs, &m2)
	if m2.UnlockArena == un {
		return
	} else {
		t.Error("failed")
	}
}
