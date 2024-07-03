package core

import (
	"flag"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	_ "github.com/flipped-aurora/gin-vue-admin/server/packfile"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println("config changed")
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}

	// 获取当前文件跟路径
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("")
	return v
}
