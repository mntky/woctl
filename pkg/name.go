package pkg

import (
	"math/rand"
	"strconv"
)

func Naming() string {
	containername := strconv.Itoa(int(rand.Int63()))

	//TODO 現在のコンテナ名とかぶってないかの確認。
	//かぶってたらもっかいNaming()

	return containername
}
