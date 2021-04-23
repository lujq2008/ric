

package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	Http struct {
		Port int
	}
	Logging struct{
		LogLevel int
	}
	GlobalRicId struct{
		Mcc	string
		Mnc	string
		RicId	string
	}
	nricmgmthost string
}

func ParseConfiguration(dir string) *Configuration{
	viper.SetConfigType("yaml")
	viper.SetConfigName("configuration")
	viper.AddConfigPath("KubernetesSimulator/resources/")
	viper.AddConfigPath("./resources/")  //For production
	viper.AddConfigPath("../resources/") //For test under Docker
	viper.AddConfigPath("/app/pkg/"+dir+"/resources/") //For test under Docker
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("#configuration.ParseConfiguration - failed to read configuration file: %s\n", err))
	}

	config := Configuration{}
	config.fillHttpConfig(viper.Sub("http"))
	config.fillLogLevelConfig(viper.Sub("logging"))
	config.fillGlobalRicId(viper.Sub("globalRicId"))
	return &config
}

func (c *Configuration)fillHttpConfig(httpConfig *viper.Viper) {
	if httpConfig == nil {
		panic(fmt.Sprintf("#configuration.fillHttpConfig - failed to fill HTTP configuration: The entry 'http' not found\n"))
	}
	c.Http.Port = httpConfig.GetInt("port")
}

func (c *Configuration)fillLogLevelConfig(LogLevelConfig *viper.Viper) {
	if LogLevelConfig == nil {
		panic(fmt.Sprintf("#configuration.LogLevelConfig - failed to fill LogLevel configuration: The entry 'logging' not found\n"))
	}
	c.Logging.LogLevel = LogLevelConfig.GetInt("logLevel")
}


func (c *Configuration)fillGlobalRicId(GlobalRicId *viper.Viper) {
	if GlobalRicId == nil {
		panic(fmt.Sprintf("#configuration.LogLevelConfig - failed to fill LogLevel configuration: The entry 'logging' not found\n"))
	}
	c.GlobalRicId.Mcc = GlobalRicId.GetString("mcc")
	c.GlobalRicId.Mnc = GlobalRicId.GetString("mnc")
	c.GlobalRicId.RicId = GlobalRicId.GetString("ricId")
}
