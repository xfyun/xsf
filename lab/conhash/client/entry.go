package main

import (
	"flag"
	"fmt"
	"git.iflytek.com/AIaaS/xsf/client"
	"git.iflytek.com/AIaaS/xsf/utils"
	"log"
	"time"
)

const (
	cname = "xsf-client" //配置文件的主段名

	clientCfg  = "client.toml"
	cfgUrl     = "http://10.1.87.69:6868"
	cfgPrj     = "guiderAllService"
	cfgGroup   = "gas"
	cfgService = "xsf-client" //服务发现的服务名
	cfgVersion = "0.0.0"      //配置文件的版本号
	apiVersion = "1.0.0"      //api版本号，一般不用修改

	cacheService = true
	cacheConfig  = true
	cachePath    = "./findercache" //配置缓存路径
)

var (
	tm   = flag.Int64("tm", 1000, "timeout")
	mode = flag.Int64("mode", 0, "0:native;1:center")
)

func main() {
	flag.Parse()
	cli, cliErr := xsf.InitClient(
		cname,
		func() xsf.CfgMode {
			switch *mode {
			case 0:
				{
					fmt.Println("about to init native client")
					return utils.Native
				}
			case 2:
				{
					{
						fmt.Println("about to init custom client")
						return utils.Custom
					}
				}
			default:
				{
					fmt.Println("about to init centre client")
					return utils.Centre
				}
			}
		}(),
		utils.WithCfgCacheService(cacheService),
		utils.WithCfgCacheConfig(cacheConfig),
		utils.WithCfgCachePath(cachePath),
		utils.WithCfgName(clientCfg),
		utils.WithCfgURL(cfgUrl),
		utils.WithCfgPrj(cfgPrj),
		utils.WithCfgGroup(cfgGroup),
		utils.WithCfgService(cfgService),
		utils.WithCfgVersion(cfgVersion),
	)
	if cliErr != nil {
		log.Fatal("main | InitCient error:", cliErr)
	}

	callConHashTest(xsf.NewCaller(cli), time.Millisecond*time.Duration(*tm))
}
