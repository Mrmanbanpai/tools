package randomstr

import (
	"errors"
	"math/rand"
	"time"
)

type RandomModel int

func (s RandomModel) Check() error {
	if !s.IsValid() {
		return errors.New("unknown random model")
	}

	return nil
}

func (s RandomModel) IsValid() bool {
	switch s {
	case Number, UpperByte, LowerByte, UpperByteWithNumber, LowerByteWithNumber, UpperAndLowerByte, AllByte:
		return true
	default:
		return false
	}
}

const (
	// Number 数字
	Number RandomModel = iota + 1
	// UpperByte 大写字母
	UpperByte
	// LowerByte 小写字母
	LowerByte
	// UpperByteWithNumber 大写字母和数字
	UpperByteWithNumber
	// LowerByteWithNumber 小写字母和数字
	LowerByteWithNumber
	// UpperAndLowerByte 大小写字母
	UpperAndLowerByte
	// AllByte 全部包含
	AllByte
)

const (
	numberByte = "1234567890"
	upperByte  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerByte  = "abcdefghijklmnopqrstuvwxyz"
)

func RandomString(randomModel RandomModel, length int) string {
	// 判断random model 是否合法 不合法就返回空字符串
	if !randomModel.IsValid() || length <= 0 {
		// 不合法且长度小于等于0 返回空字符串
		return ""
	}

	// 长度不等于0 返回固定长度的
	baseString := judgeRandomModel(randomModel)
	if len(baseString) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	var resString string

	for {
		b := []byte(baseString)[rand.Intn(len(baseString))]
		resString += string(b)
		if len(resString) == length {
			break
		}
	}
	// 处理 handle
	// 长度等于0 返回空字符
	return resString
}

func judgeRandomModel(randomModel RandomModel) string {
	switch randomModel {
	case Number:
		return numberByte
	case UpperByte:
		return upperByte
	case LowerByte:
		return lowerByte
	case UpperByteWithNumber:
		return upperByte + numberByte
	case LowerByteWithNumber:
		return lowerByte + numberByte
	case UpperAndLowerByte:
		return upperByte + lowerByte
	case AllByte:
		return upperByte + lowerByte + numberByte
	default:
		return ""
	}
}
