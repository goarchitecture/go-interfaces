package main

import (
	"flag"
	"fmt"
	"github.com/goarchitecture/go-interfaces/configuration"
	configurationSources "github.com/goarchitecture/go-interfaces/configuration/sources"
	httpConfigSource "github.com/goarchitecture/go-interfaces/third-party/plugins/http-config-source"
	"os"
)

var (
	configurationFrom = flag.String("configuration-from", configurationSources.SourceTypeRaw, "configuration source type")
)

func main() {
	flag.Parse()

	var config *configuration.Configuration
	var err error

	switch *configurationFrom {
	case configurationSources.SourceTypeRaw:
		fmt.Println("Введите содержимое конфигурации (raw json): ")
		var rawContents configurationSources.RawSource
		if _, err = fmt.Scan(&rawContents); err != nil {
			fmt.Printf("ошибка ввода содержимого конфигурации: %s\n", err)
			os.Exit(1)
		}

		config, err = configuration.Load(rawContents)
	case configurationSources.SourceTypeFile:
		fmt.Println("Введите путь к файлу содержащиму конфигурации (json): ")
		var fileName string
		if _, err = fmt.Scan(&fileName); err != nil {
			fmt.Printf("ошибка ввода пути к файлу: %s\n", err)
			os.Exit(1)
		}

		config, err = configuration.Load(configurationSources.NewFileSource(fileName))

	case httpConfigSource.SourceTypeHttp:
		fmt.Println("Введите url для загрузки конфигурации (json): ")
		var urlConfiguration string
		if _, err = fmt.Scan(&urlConfiguration); err != nil {
			fmt.Printf("ошибка ввода пути к файлу: %s\n", err)
			os.Exit(1)
		}

		config, err = configuration.Load(httpConfigSource.NewHttpConfigSource(urlConfiguration))
	}

	if err != nil {
		fmt.Printf("ошибка чтения конфигурации: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Конфиг прочтён: ")
	fmt.Println(config)
}
