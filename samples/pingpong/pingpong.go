package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"math"
	"os"
	"strings"
	"strconv"
)

func main() {

	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <pub_topic> <sub_topic>\n",
			os.Args[0])
		os.Exit(1)
	}

	broker := os.Args[1]
	group := os.Args[2]
	pub_topic := os.Args[3]
	sub_topic := os.Args[4]

	fmt.Printf("cmdline of %s: broker=%s, pub_topic=%s, sub_topic=%s\n",group,broker,pub_topic,sub_topic)
	
	//setup consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    broker,
		"group.id":             group,
		"session.timeout.ms":   6000,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	defer c.Close()

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics([]string{sub_topic}, nil)

	//setup producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer p.Close()
	
	fmt.Printf("Created Producer %v\n", p)

	//as Ponger, initiate the msg flow
	cnt := 1
	msg := "Hello Go!"
	value := fmt.Sprintf("%s %d",msg,cnt)
	if group == "ponger" {
		err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &pub_topic, Partition: kafka.PartitionAny}, Value: []byte(value)},nil)
		if err != nil {
			fmt.Println("error for produce: ",err)
		}
		fmt.Println("ponger send first msg")
	}

	run := true
	for run==true {
		ev := c.Poll(100)
		if ev == nil {
			continue
		}
		fmt.Println("---------recv 1")
		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				fmt.Printf("%s failed to send msg: %v\n", group, e.TopicPartition.Error)
			} else {
				fmt.Printf("- %s recv on %s: %s\n", group, e.TopicPartition, string(e.Value))
				data := strings.Split(string(e.Value)," ")
				cnt, err = strconv.Atoi(data[len(data)-1])
				if err != nil {
					fmt.Println("xxx String/Int conversion failure:",e)
					break
				}
				cnt = (cnt+1)%math.MaxInt32
				value = fmt.Sprintf("%s %d",msg,cnt)
				fmt.Printf("- %s send: %s\n", group, value)
				err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &pub_topic, Partition: kafka.PartitionAny}, Value: []byte(value)},nil)
				if err != nil {
					fmt.Println("error for produce: ",err)
				}
			}
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
	fmt.Println("--------- ",group," exits ----------")
}

