package main

import (
	"context"
	"fmt"
	"internal/awsTools"
	"internal/events"
	"sync"
)


func TestAWS() {
	getParamLoc := "/qa/sqs/rs-post-service-ImageScrape/url"
	testPutParamLoc := "/test/mkiss/myValue"

	awsServices, err := awsTools.NewAWSService()
	if err != nil {
		panic(err)
	}

	value, err := awsServices.GetSSMParameter(getParamLoc, false)
	if err != nil {
		panic(err)
	}
	// Testing get param
	fmt.Println(value)

	// Testing put param
	err = awsServices.PutSSMParameter(testPutParamLoc, "Hello from GO ^_^!!!", "String", false)
	if err != nil {
		panic(err)
	}
}

var topics = []string{
	"merchant.get-token",
	"merchant.token-received",
}

func main() {
	var enableAWSTest = false

	var wg sync.WaitGroup
	defer wg.Wait()
	// Event topics should be in a centralized location in my opinion

	// Which is why its required to initialize an event system
	eventSys := events.NewEventSystem(topics)
	//eventSys.PrintRegisteredTopics()
	ctx := context.Background()

	// Register Events
	helloEvent := events.NewHelloEvent(eventSys.Bus, "helloEvent")

	helloEvent.Start(&wg)
	defer helloEvent.Stop()

	for i := 0; i < 5; i++ {
		fmt.Printf("Emitting event #%d\n", i)
		data := fmt.Sprintf("Michael-%d", i)
		err := eventSys.Bus.Emit(
			ctx,
			"merchant.get-token",
			data,
		)
		if err != nil {
			fmt.Println("ERROR >>>>", err)
		}
	}

	if enableAWSTest {
		TestAWS()
	}

	fmt.Println("\nmain exited")
}