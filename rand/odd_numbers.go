package rand

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// 使用自增ID生成基础字母部分
func encodeToAlphabet(num int64) string {
	var sb strings.Builder
	base := int64(len(alphabet))

	for num > 0 {
		rem := num % base
		sb.WriteByte(alphabet[rem])
		num /= base
	}

	return sb.String()
}

// 从哈希值中不同部分组合字符，确保唯一性
func generatePadding(num int64, length int) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d", num)))
	encoded := encodeToAlphabetFromBytes(hash[:])

	// 通过从哈希的不同部分组合字符，确保更加随机
	var sb strings.Builder
	for i := 0; i < length; i++ {
		// 选择哈希中的不同位置（可以根据自增ID的特定属性选择位置）
		index := (i * 3) % len(encoded)
		sb.WriteByte(encoded[index])
	}
	return sb.String()
}

// 从字节数组生成字母序列
func encodeToAlphabetFromBytes(data []byte) string {
	var sb strings.Builder
	for _, b := range data {
		sb.WriteByte(alphabet[b%byte(len(alphabet))])
	}
	return sb.String()
}

// 根据一个数字ID，生成一个纯大写字母的单号，长度由length参数控制
// 不足length部分由哈希补足
func GenerateUniqueID(id int64, length int) string {
	// 基于自增ID生成基础ID部分
	baseID := encodeToAlphabet(id)

	// 如果不足位数，用哈希填充
	if len(baseID) < length {
		padding := generatePadding(id, length-len(baseID))
		baseID += padding
	}

	// 返回最终8位的唯一ID
	return baseID[:length]
}
