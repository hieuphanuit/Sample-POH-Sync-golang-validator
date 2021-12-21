package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"example_poh.com/dataType"
)

type Config struct {
	Address        string               `json:"address"`
	Ip             string               `json:"ip"`
	Port           int                  `json:"port"`
	Validators     []dataType.Validator `json:"validators"`
	HashPerSecond  int                  `json:"hash_per_second"`
	TickPerSecond  int                  `json:"tick_per_second"`
	TickPerSlot    int                  `json:"tick_per_slot"`
	BlockStackSize int                  `json:"block_stack_size"`
	TimeOutTicks   int                  `json:"time_out_ticks"` // how many tick validator should wait before create virture block
}

func loadConfig() Config {
	var config Config
	raw, err := ioutil.ReadFile("config/conf.json")
	if err != nil {
		log.Fatalf("Error occured while reading config")
	}
	json.Unmarshal(raw, &config)
	log.Printf("Config loaded: %v\n", config)
	return config
}

var AppConfig = loadConfig()