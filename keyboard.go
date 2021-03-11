// Пакет keyboard предназначен для получения вводных данных с клавиатуры от пользователя
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat читает данные в формате float64
// В случае ошибки вернет 0
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
