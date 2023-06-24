package initialize

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// GlobalVars 全局环境变量，可能会被其它方法频繁引用
var GlobalVars EnvParam

// EnvParam 全局环境变量的结构体，如果需要初始化什么参数时，可以在这里面增加
type EnvParam struct {
	Env     string
	Migrate bool
}

// 这是一个钩子函数，可以跳过，也可以在里面控制全局变量
func setGlobalVarHook(env string, migrate bool) {

	//判断启动参数是否正确（这个if可以随时按照自己的意愿改写）
	if env == "dev" || env == "fat" || env == "pro" {
		fmt.Println("启动日志：env环境依赖读取完成...")

		//为全局变量赋值
		GlobalVars.Env = env
		GlobalVars.Migrate = migrate
	} else {
		fmt.Println("启动日志：--env参数错误，只能是dev、fat、pro...")
		os.Exit(3)
	}
}

// 标准化参数名称（容错：下划线、点号、减号、等等）
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

// InitCmd 初始化环境
func InitCmd() {

	/*
	 * 命令行参数（启动时可以通过 --xxx=aaa 的方式来调用，优先级最高，可以覆盖config/default.yaml中的变量）
	 *
	 * 考虑到GoMessage自身的元数据对用户来说不怎么关注；
	 * 因此这个默认设置为true，也就是说每次都迁移数据库进行幂等性对齐，这样可以对元数据的存储更具有广泛的支持性
	 * 如果不需要可以通过参数--migrate=false来取消。
	 */
	var env = pflag.StringP("env", "e", "", "启动环境（dev、fat、uat、pro）")
	var ginMode = pflag.StringP("ginMode", "g", "", "Gin框架运行模式（debug、test、release）")
	var migrate = pflag.BoolP("migrate", "m", true, "是否迁移数据库（true、false）")
	var log2es = pflag.BoolP("log2es", "l", false, "是否上传数据到es中（true、false）")

	//设置命令行参数标准化兼容函数（防止用户手滑填写错误，可以兼容：下划线、点号、减号、等等）
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	//解析命令行参数
	pflag.Parse()

	if len(*env) != 0 {
		//判断--env是否为空，如果不为空则覆盖到viper中
		//因为main.go的init中先执行InitConfig，后执行InitEnv，随后才陆续启动其它组件。
		//因此在这里拿到pflag参数覆盖到viper中时；后续调用viper中的变量，就是已经被pflag覆盖后的值。
		viper.Set("global.env", *env)
	}

	if len(*ginMode) != 0 {
		//判断--ginMode是否为空，如果不为空则覆盖到viper中
		//因为main.go的init中先执行InitConfig，后执行InitEnv，随后才陆续启动其它组件。
		//因此在这里拿到pflag参数覆盖到viper中时；后续调用viper中的变量，就是已经被pflag覆盖后的值。
		viper.Set("global.ginMode", *ginMode)
	}

	if *log2es == true {
		//判断--log2es是否为true，如果是则覆盖掉viper中的值
		//因为main.go的init中先执行InitConfig，后执行InitEnv，随后才陆续启动其它组件。
		//因此在这里拿到pflag参数覆盖到viper中时；后续调用viper中的变量，就是已经被pflag覆盖后的值。
		viper.Set("log.log2es", *log2es)
	}

	//判断viper中env的值是否为空
	if len(viper.GetString("global.env")) != 0 {
		//这是一个钩子函数
		//需要注意：migrate参数代表是否迁移一次数据库，该参数没有进入到viper中，只放在了GlobalVars中了。
		setGlobalVarHook(viper.GetString("global.env"), *migrate)

	} else {
		fmt.Println("启动日志：--env参数不能为空，请输入 --help 命令查看启动参数；(或在default.yaml中配置global.env变量)...")
		os.Exit(3)
	}
}
