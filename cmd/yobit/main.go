package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "bool, boooool,b",
			Usage:  "bool",
			EnvVar: "BOOL,HALPPLZ",
		},
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "APP_LANG",
		},
		cli.StringFlag{
			Name:   "YOBIT_KEY, k",
			Value:  "english",
			Usage:  "Yobit access key",
			EnvVar: "YOBIT_KEY",
		},
		cli.StringFlag{
			Name:   "YOBIT_SCT, c",
			Value:  "english",
			Usage:  "Yobit ",
			EnvVar: "YOBIT_SCT",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.Bool("bool") {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		client = yobit.New(c.String("YOBIT_KEY"), c.String("YOBIT_SCT"))
		CoinPairList, err := client.GetTickers()
		//"github.com/bols-blu-org/yobit_bot"
		//CoinPairList, err := yobit.GetCoinPairList()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(CoinPairList)
		for {
		}
		return nil
	}
	app.Run(os.Args)
}
