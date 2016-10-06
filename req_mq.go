package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"log"
	"github.com/streadway/amqp"
)
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}


func main() {

	url := "http://14.140.201.163/travelengine/api/travelengine"

	payload := strings.NewReader("{\n    \"_service\": \"HOTEL_LIST\",\n    \"_version\": \"1.0\",\n    \"service_request\": {\n        \"channel_info\": {\n            \"requester\": \"11\",\n            \"instrument\": \"3\",\n            \"channel_instance_id\": \"01\",\n            \"channel_consumer_id\": \"9988776655\"\n        },\n        \"device_info\": {\n            \"device_os\": \"Android\",\n            \"device_os_version\": \"4.4.0\",\n            \"device_id\": \"87612736878417868\",\n            \"imei\": \"383257047135376\",\n            \"longitude\": \"17.0021\",\n            \"latitude\": \"15.0021\",\n            \"model_name\": \"moto g3\"\n        },\n        \"transaction_info\": {\n            \"time_stamp\": \"2015-12-31T13:54:59.123+05:30\",\n            \"request_id\": \"12345678901234567890\"\n        },\n        \"user\": {\n            \"mdn\": \"919717915661\"\n        },\n        \"hotel\": {\n            \"city\": \"1871465973541478292\",\n            \"checkin_date\": \"2016-09-28\",\n            \"checkout_date\": \"2016-09-29\",\n            \"rooms\": [\n                {\n                    \"adults\": \"2\",\n                    \"children\": {\n                        \"Children_age\": [\n                            1\n                        ]\n                    }\n                }\n            ],\n            \"sort_by\": \"priceup\",\n            \"page_no\": \"0\"\n        },\n        \"service_info\": {\n            \"operator\": \"oxigen\",\n            \"category\": \"hotel\"\n        },\n        \"partner_info\": [\n            {\n                \"code\": \"301\",\n                \"name\": \"goibibo\"\n            }\n        ],\n        \"params\": {\n            \"param1\": \"\",\n            \"param2\": \"\",\n            \"param3\": \"\",\n            \"param4\": \"\",\n            \"param5\": \"\"\n        }\n    }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authorization", "Basic bmV3OmxvZw==")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "7b995861-2927-fd5c-60d5-ef1fe8bf1e4d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	newbody, _ := ioutil.ReadAll(res.Body)

	body := string(newbody)

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()
q, err := ch.QueueDeclare(
  "hello1", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")


err = ch.Publish(
  "",     // exchange
  q.Name, // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing {
    ContentType: "text/plain",
    Body:        []byte(body),
  })
failOnError(err, "Failed to publish a message")

}
