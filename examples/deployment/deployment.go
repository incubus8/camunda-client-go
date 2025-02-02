package main

import (
	"fmt"
	"os"
	"time"

	camunda_client_go "github.com/incubus8/camunda-client-go"
)

func main() {
	client := camunda_client_go.NewClient(camunda_client_go.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	file, err := os.Open("HelloWorld.bpmn")
	if err != nil {
		fmt.Printf("Error read file: %s\n", err)
		return
	}
	result, err := client.Deployment.Create(camunda_client_go.ReqDeploymentCreate{
		DeploymentName: "HelloWorldProcessDemo",
		Resources: map[string]interface{}{
			"HelloWorld.bpmn": file,
		},
	})
	if err != nil {
		fmt.Printf("Error deploy process: %s\n", err)
		return
	}

	fmt.Printf("Result: %#+v\n", result)
}
