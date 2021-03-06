package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"example_poh.com/dataType"
)

type GenesisBlockInfo struct {
	InitAddress string  `json:"init_address"`
	InitPubkey  string  `json:"init_pubkey"`
	InitBalance float64 `json:"init_balance"`
}

type Config struct {
	Address                    string               `json:"address"`
	Ip                         string               `json:"ip"`
	Port                       int                  `json:"port"`
	NodeType                   string               `json:"node_type"`
	Validators                 []dataType.Validator `json:"validators"`
	HashPerSecond              int                  `json:"hash_per_second"`
	TickPerSecond              int                  `json:"tick_per_second"`
	TickPerSlot                int                  `json:"tick_per_slot"`
	BlockStackSize             int                  `json:"block_stack_size"`
	TimeOutTicks               int                  `json:"time_out_ticks"` // how many tick validator should wait before create virture block
	TransactionPerHash         int                  `json:"transaction_per_hash"`
	NumberOfValidatePohRoutine int                  `json:"number_of_validate_poh_routine"`
	AccountDBPath              string               `json:"account_db_path"`
	GenesisBlockInfo           GenesisBlockInfo     `json:"genesis_block_info"`
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
