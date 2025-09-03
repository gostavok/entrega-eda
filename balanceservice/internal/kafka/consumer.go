package kafka

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

)

func StartBalanceConsumer(db *sql.DB) {
	topic := "balances"
	group := "balanceservice_app"
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:29092"
	}
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	}
	consumer, err := ckafka.NewConsumer(&configMap)
	if err != nil {
		log.Fatalf("Erro ao criar consumer Kafka: %v", err)
	}
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Erro ao inscrever no tópico: %v", err)
	}
	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("Kafka read error: %v", err)
				continue
			}
			log.Printf("Mensagem recebida bruta: %s", string(msg.Value))
			// Novo formato esperado
			type Payload struct {
				AccountIDFrom        string  `json:"account_id_from"`
				AccountIDTo          string  `json:"account_id_to"`
				BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
				BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
			}
			type BalanceUpdatedMsg struct {
				Name    string  `json:"Name"`
				Payload Payload `json:"Payload"`
			}
			var event BalanceUpdatedMsg
			err = json.Unmarshal(msg.Value, &event)
			if err != nil {
				log.Printf("Parse error: %v", err)
				continue
			}
			// Atualiza ambos os balances
			_, err = db.Exec("INSERT INTO balances (account_id, balance) VALUES (?, ?) ON DUPLICATE KEY UPDATE balance = ?", event.Payload.AccountIDFrom, event.Payload.BalanceAccountIDFrom, event.Payload.BalanceAccountIDFrom)
			if err != nil {
				log.Printf("DB update error (from): %v", err)
			}
			_, err = db.Exec("INSERT INTO balances (account_id, balance) VALUES (?, ?) ON DUPLICATE KEY UPDATE balance = ?", event.Payload.AccountIDTo, event.Payload.BalanceAccountIDTo, event.Payload.BalanceAccountIDTo)
			if err != nil {
				log.Printf("DB update error (to): %v", err)
			}
		}
	}()
}
