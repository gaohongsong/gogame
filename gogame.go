package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gmaclinuxer/gogame/cg"
	"github.com/gmaclinuxer/gogame/ipc"
	"os"
	"strconv"
	"strings"
)

var (
	Version     string
	BuildCommit string
	BuildTime   string
	GoVersion   string
)

// printVersion print version info from makefile
func printVersion() {
	fmt.Printf(`Version:      %s
Go version:   %s
Git commit:   %s
Built:        %s
`, Version, GoVersion, BuildCommit, BuildTime)
}

var centerCli *cg.CenterClient

func startCenterService() error {

	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerCli = &cg.CenterClient{client}

	return nil
}

func Help(args []string) int {
	fmt.Println(`Commands:
	login  <username> <level> <exp>
	logout <username>
	send   <message>
	list
	quit(q)
	help(h)`)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("usage: logout <username>")
	}

	centerCli.RemovePlayer(args[1])

	return 0
}

func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("usage: login <username> <level> <exp>")
		return 0
	}

	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("invalid parameter: <level> should be an integer")

	}

	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("invalid parameter: <exp> should be an integer")

	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp

	err = centerCli.AddPlayer(player)
	if err != nil {
		fmt.Println("failed add player, ", err)
	}

	return 0

}

func ListPlayer(args []string) int {

	if ps, err := centerCli.ListPlayer(""); err == nil {
		for i, v := range ps {
			fmt.Println(i+1, ":", v)
		}
	} else {
		fmt.Println("list failed, ", err)
	}

	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], " ")

	if err := centerCli.BroadCast(message); err != nil {
		fmt.Println("failed, ", err)
	}

	return 0

}

func main() {

	version := flag.Bool("v", false, "print version info")
	flag.Parse()

	// print version info only
	if *version {
		printVersion()
		return
	}

	// go on
	handlers := map[string]func([]string) int{
		"help":   Help,
		"h":      Help,
		"quit":   Quit,
		"q":      Quit,
		"login":  Login,
		"logout": Logout,
		"list":   ListPlayer,
		"send":   Send,
	}

	fmt.Println("Start center server...")
	startCenterService()
	fmt.Println("Start center server successfully")

	Help(nil)

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("command> ")

		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		if handler, ok := handlers[tokens[0]]; ok {
			res := handler(tokens)
			if res != 0 {
				break
			}
		} else {
			fmt.Println("unknown command: ", tokens[0])
		}

	}

}
