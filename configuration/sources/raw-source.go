package configurationSources

import (
	"github.com/goarchitecture/go-interfaces/configuration"
)

const SourceTypeRaw = "raw"

type RawSource []byte

// объявляем пустую переменную типа configuration.Source, чтобы убедится что наш тип RawSource
// ВСЕГДА имплиментирует интерфейсы. У этого есть плюсы:
//
// + если структура будет переписана и мы изменим наш метод так что интерфейс больше не будет имплементироваться
//   IDE подсветит нам, а компилятор не даст скомпилировать (потому что эта пустая переменная не сработает)
// + если интерфейс изменится и мы больше не будем соответстововать - тоже самое
//
// и минусы:
//
// - У нас есть прямая ЗАВИСИМОСТЬ от интерфейса (то есть мы импортируем пакет, в котором находится интерфейс)
// 	   Это не всегда хорошо, иногда нам не нужны лишние зависимости (нужно помнить что значит из пакета configuration
//                                     мы не можем импортировать пакеты Source-ов)
//
var _ configuration.Source = RawSource{}

func NewRawSource(s string) RawSource {
	// просто обворачиваем строку в
	return RawSource(s)
}

func (s RawSource) Contents() ([]byte, error) {
	return s, nil
}

// Кроме имплементирования интерфейса (или нескольких) тип может иметь ещё сколько угодно методов:

// GetFileName возвращает имя файла
func (s RawSource) String() string {
	return string(s)
}
