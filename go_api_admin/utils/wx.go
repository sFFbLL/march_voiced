package utils

import (
	"github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"net/http"
	"project/common/global"
	"time"
)

func GetAccess() (string, error) {
	token, err := global.Wx.GetAccessToken()
	if err != nil {
		return token, err
	}
	err = global.Rdb.Set("wxAccessToken", token, 300*time.Second).Err()
	if err != nil {
		return token, err
	}
	return token, nil
}

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// SaveWxHead 处理保存微信头像
func SaveWxHead(imgUrl string) (string, error) {
	//imgUrl := "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIomz1vPJX1xeY2pMaOvXQ6ticGJfQWaJw6wjoiaicYoIjwAOg2vFvhMdOianQ7A4OxicJ8Ml76N2an8Nw/132"
	//获取远端图片
	res, err := http.Get(imgUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)
	guid := uuid.New().String()
	fileName := guid + ".png"
	singleFile := "static/uploadfile/" + fileName
	err = ioutil.WriteFile(singleFile, data, 0666) //buffer输出文件中（不做处理，直接写到文件）
	if err != nil {
		return "", err
	}
	return fileName, nil
}
