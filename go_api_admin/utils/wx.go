package utils

import (
	"math/rand"
	"project/common/global"
	"time"
)

func GetAccess()(string,error){
	token, err := global.Wx.GetAccessToken()
	if err != nil {
		return token,err
	}
	err = global.Rdb.Set("wxAccessToken",token,300*time.Second).Err()
	if err != nil {
		return token,err
	}
	return token,nil
}
const (
	KC_RAND_KIND_NUM   = 0  // 纯数字
	KC_RAND_KIND_LOWER = 1  // 小写字母
	KC_RAND_KIND_UPPER = 2  // 大写字母
	KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}

