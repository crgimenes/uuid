package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/gosidekick/goconfig"
)

func main() {
	type configFlags struct {
		CarriageReturn bool   `json:"cr" cfg:"cr" cfgDefault:"false"`
		LineFeed       bool   `json:"lf" cfg:"lf" cfgDefault:"false"`
		Uppercase      bool   `json:"u" cfg:"u" cfgDefault:"false"`
		Armored        bool   `json:"a" cfg:"a" cfgDefault:"false"`
		ArmorChar      string `json:"ac" cfg:"ac" cfgDefault:"\""`
		NtoGenerate    int    `json:"n" cfg:"n" cfgDefault:"1"`
		IDSeparator    string `json:"ids" cfg:"ids" cfgDefault:""`
	}

	cfg := configFlags{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	cr := ""
	lf := ""
	ac := ""
	if cfg.CarriageReturn {
		cr = "\r"
	}
	if cfg.LineFeed {
		lf = "\n"
	}
	if cfg.Armored {
		ac = cfg.ArmorChar
	}
	for n := 0; n < cfg.NtoGenerate; n++ {
		id := uuid.New().String()
		if cfg.Uppercase {
			id = strings.ToUpper(id)
		}
		fmt.Printf("%v%v%v%v%v%v", ac, id, ac, cfg.IDSeparator, cr, lf)
	}
}
