package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ozlog "github.com/usthooz/oozlog/go"
)

const (
	// 主命令
	exec = "json2go"
	// version 当前版本
	version = "v1.0"
)

var (
	// command 命令
	command string
	// workPath current work path
	workPath string
	// jsonFile json文件名称
	jsonFile string
	// outputFile 输出文件名称
	outFile string
	// outType 输出类型
	outType string
)

var (
	// commandsMap 命令集
	commandMap map[string]*Command
)

// Command
type Command struct {
	Name   string
	Detail string
	Func   func(name, detail string)
}

func init() {
	flag.StringVar(&jsonFile, "json_file", "json2go.json", "json file.")
	flag.StringVar(&outType, "out_type", "file", "struct out type.")
	flag.StringVar(&outFile, "out_file", "json2go_types.go", "output file.")
}

// getHelp get this project's help
func getHelp(name, detail string) {
	commands := make([]string, 0, len(commandMap))
	for _, v := range commandMap {
		commands = append(commands, fmt.Sprintf("%s\t%s", v.Name, v.Detail))
	}
	outputHelp(fmt.Sprintf("Usage: %s <command>", exec), commands, []string{
		"-json_file\t json文件, 默认json文件为json2go.json",
		"-out_type\t 输出类型, 默认输出方式为输出到文件file/可选print、file",
		"-out_file\t 输出文件, 默认输出文件为json2go_types.go",
	}, []string{
		"json2go gen_type",
		"json2go gen_type -out_type=print",
		"json2go gen_type -out_type=file",
		"json2go gen_type -out_type=file -out_file=out_types.go",
	})
}

// initCommands
func initCommands() {
	for i, v := range os.Args {
		switch i {
		case 1:
			command = v
		}
	}

	// 初始化命令列表
	commandMap = map[string]*Command{
		"v": &Command{
			Name:   "v",
			Detail: "查看当前版本号",
			Func:   getVersion,
		},
		"help": &Command{
			Name:   "help",
			Detail: "查看帮助信息",
			Func:   getHelp,
		},
		"gen_type": &Command{
			Name:   "gen_type",
			Detail: "根据json文件自动生成struct",
			Func:   getHelp,
		},
	}
}

// checkArgs check common is nil?
func checkArgs() bool {
	if len(command) == 0 {
		getHelp("help", commandMap["help"].Detail)
		return false
	}
	return true
}

func outputHelp(usage string, commands, options, examples []string) {
	fmt.Println("\n", usage)
	if len(commands) > 0 {
		fmt.Println("\n Commands:")
		for _, s := range commands {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}
	if len(options) > 0 {
		fmt.Println("\n Options:")
		for _, s := range options {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}
	if len(examples) > 0 {
		fmt.Println("\n Examples:")
		for _, s := range examples {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}
	fmt.Println()
}

// getVersion 查看当前版本
func getVersion(name, detail string) {
	fmt.Println(version)
}

// getWorkDir get current work dir
func getWorkDir() {
	// get current dir
	currentDir, err := os.Getwd()
	if err != nil {
		ozlog.Fatalf("%s", err)
	}
	// gei this window all gopath
	pathList := strings.Split(os.Getenv("GOPATH"), ":")
	for _, path := range pathList {
		if strings.HasPrefix(currentDir, path) {
			var (
				prefix string
			)
			// check this path ends
			if strings.HasSuffix(path, "/") {
				// /Users/admin/work/goProject/
				prefix = path + "src/"
			} else {
				// /Users/admin/work/goProject
				prefix = path + "/src/"
			}
			// output: github.com/usthooz/oozgorm
			currentDir = currentDir[len(prefix):]
		}
		workPath = currentDir
	}
}

func main() {
	// 获取当前目录
	getWorkDir()
	// 初始化命令
	initCommands()
	if len(os.Args) < 2 {
		getHelp("help", commandMap["help"].Detail)
		return
	}
	flag.CommandLine.Parse(os.Args[3:])
	if !checkArgs() {
		return
	}
	c := commandMap[command]
	if c == nil {
		getHelp("help", commandMap["help"].Detail)
		return
	} else {
		c.Func(c.Name, c.Detail)
	}
}
