package configurationSources

import (
	"github.com/goarchitecture/go-interfaces/configuration"
	"io/ioutil"
)

const SourceTypeFile = "file"

type FileSource struct {
	fileName string
}

// объявляем пустую переменную типа configuration.Source, чтобы убедится что наша структура FileSource
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
var _ configuration.Source = &FileSource{}

func NewFileSource(fileName string) *FileSource {
	return &FileSource{fileName: fileName}
}

func (s *FileSource) Contents() ([]byte, error) {
	return ioutil.ReadFile(s.fileName)
}

// Кроме имплементирования интерфейса (или нескольких) структура может иметь ещё сколько угодно методов:

// GetFileName возвращает имя файла
func (s *FileSource) GetFileName() string {
	return s.fileName
}
