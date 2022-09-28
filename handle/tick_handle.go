// auth: kunlun
// date: 2019-01-24
// description:
package handle

import (
	"common"
	"encoding/json"
	"entity"
	"fmt"
	"strconv"
	"time"
)

// tick事件处理器
func TickHandle(data []byte) {
	tick := entity.Tick{}
	json.Unmarshal(data, &tick)
	key := strconv.Itoa(int(tick.Ts))
	val, err := common.RedisGetString(key)
	now := time.Now().UnixNano()
	if err != nil {
		fmt.Printf("Get tick data from redis faild")
	}
	// 如果库中不存在则新增
	if val == "" {
		storeData := entity.StoreData{}
		storeData.Id = tick.Ts
		storeData.Symbol = tick.Symbol
		storeData.Open = tick.Data[0].Price
		storeData.CreateTime = now
		storeData.UpdateTime = now
		var vol float32
		for _, item := range tick.Data {
			vol += item.Amount
			storeData.High = item.Price
			storeData.Low = item.Price
			storeData.Close = item.Price
		}
		storeData.Amount = vol
		storeData.Dir = tick.Data[0].Dir
		//设置别名
		storeData.SymbolAlias = tick.Symbol
		bytes, err := json.Marshal(storeData)
		if err != nil {
			fmt.Printf("JSON marshal err: %v", err)
			return
		}
		common.RedisSet(key, string(bytes))
	} else {
		redisData := &entity.StoreData{}
		err := json.Unmarshal([]byte(val), redisData)
		if err != nil {
			fmt.Printf("JSON unmarshal err: %v", err)
			return
		}
		data := tick.Data
		redisData.UpdateTime = now
		for _, item := range data {
			redisData.Amount += item.Amount
			redisData.Close = item.Price
			if item.Price <= redisData.High {
				redisData.High = item.Price
			}
			if item.Price <= redisData.Low {
				redisData.Low = item.Price
			}
			redisData.Close = item.Price
		}

		bytes, err := json.Marshal(redisData)
		if err != nil {
			fmt.Printf("JSON marshal err: %v", err)
			return
		}
		common.RedisSet(key, string(bytes))
	}
}
