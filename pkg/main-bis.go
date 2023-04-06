// package main

// import (
// 	// "encoding/json"
// 	"context"
// 	"time"
// 	"fmt"
// 	"log"
// 	"errors"
// 	"os"

// 	"github.com/sashabaranov/go-openai"
// 	"google.golang.org/api/compute/v1"
// 	"google.golang.org/api/option"
// 	"google.golang.org/api/monitoring/v3"
// )

// type OpenAIClient struct {
// 	client *openai.Client
// }

// func (c *OpenAIClient) Configure(token string) error {
// 	client := openai.NewClient(token)
// 	if client == nil {
// 		return errors.New("error creating openai client.")
// 	}
// 	c.client = client
// 	return nil
// }

// func (c *OpenAIClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
// 	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
// 		Model: openai.GPT3Dot5Turbo,
// 		Messages: []openai.ChatCompletionMessage{
// 			{
// 				Role:    "user",
// 				Content: prompt,
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp.Choices[0].Message.Content, nil
// }

// func requestGPTSuggestions(instanceInfo string) (string, error) {
// 	apiKey := os.Getenv("OPENAI_API_KEY")
// 	if apiKey == "" {
// 		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
// 	}

// 	client := &OpenAIClient{}
// 	err := client.Configure(apiKey)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to configure OpenAI client: %v", err)
// 	}

// 	prompt := instanceInfo + " Please suggest some improvements for cost optimization."
// 	ctx := context.Background()

// 	completion, err := client.GetCompletion(ctx, prompt)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to request GPT API: %v", err)
// 	}

// 	return completion, nil
// }




// func listComputeInstances(projectID, zone string) ([]*compute.Instance, error) {
// 	ctx := context.Background()

// 	computeService, err := compute.NewService(ctx, option.WithScopes("https://www.googleapis.com/auth/cloud-platform.read-only"))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create Compute Engine client: %v", err)
// 	}

// 	instanceList, err := computeService.Instances.List(projectID, zone).Do()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to list Compute Engine instances: %v", err)
// 	}

// 	return instanceList.Items, nil
// }

// func getResourceUsageMetrics(projectID, instanceName string) (map[string]float64, error) {
// 	ctx := context.Background()

// 	monitoringService, err := monitoring.NewService(ctx, option.WithScopes("https://www.googleapis.com/auth/monitoring.read"))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create Cloud Monitoring client: %v", err)
// 	}

// 	endTime := time.Now().UTC()
// 	startTime := endTime.Add(-1 * time.Hour)

// 	filter := fmt.Sprintf(
// 		"metric.type = starts_with(\"compute.googleapis.com/instance/\") AND resource.type=\"gce_instance\" AND resource.labels.instance_name=\"%s\"",
// 		instanceName)

// 	listCall := monitoringService.Projects.TimeSeries.List("projects/" + projectID).
// 		IntervalStartTime(startTime.Format(time.RFC3339)).
// 		IntervalEndTime(endTime.Format(time.RFC3339)).
// 		Filter(filter)

// 	resp, err := listCall.Do()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get resource usage metrics: %v", err)
// 	}

// 	metrics := make(map[string]float64)
// 	for _, series := range resp.TimeSeries {
// 		metricName := series.Metric.Type
// 		value := series.Points[0].Value.DoubleValue
// 		metrics[metricName] = *value
// 	}

// 	return metrics, nil
// }

// func main() {
// 	projectID := "agicap-tech-dev"
// 	zone := "europe-west1-b"

// 	instances, err := listComputeInstances(projectID, zone)
// 	if err != nil {
// 		log.Fatalf("Error listing Compute Engine instances: %v", err)
// 	}
	
// 	fmt.Printf("Compute Engine instances in project %s, zone %s:\n", projectID, zone)
// 	for _, instance := range instances {
// 		// fmt.Printf("- Name: %s\n", instance.Name)
// 		// fmt.Printf("  Machine Type: %s\n", instance.MachineType)
// 		// fmt.Printf("  Zone: %s\n", instance.Zone)
// 		// fmt.Printf("  Status: %s\n", instance.Status)
// 		// fmt.Printf("  Disks: %v\n", instance.Disks)
// 		// fmt.Printf("  Network Interfaces: %v\n", instance.NetworkInterfaces)
// 		// fmt.Printf("  Labels: %v\n", instance.Labels)
// 		// fmt.Printf("  Tags: %v\n", instance.Tags)
// 		// fmt.Printf("  Creation Timestamp: %s\n", instance.CreationTimestamp)

// 		// metrics, err := getResourceUsageMetrics(projectID, instance.Name)
// 		// if err != nil {
// 		// 	log.Printf("Error getting resource usage metrics for instance %s: %v", instance.Name, err)
// 		// 	continue
// 		// }

// 		// fmt.Printf("  Resource usage metrics:\n")
// 		// for metric, value := range metrics {
// 		// 	fmt.Printf("    %s: %.2f\n", metric, value)
// 		// }

// 		instanceInfo := fmt.Sprintf("Instance name: %s, Machine Type: %s, Status: %s, Zone: %s, Creation Timestamp: %s.",
// 		instance.Name, instance.MachineType, instance.Status, instance.Zone, instance.CreationTimestamp)

// 		suggestion, err := requestGPTSuggestions(instanceInfo)
// 		if err != nil {
// 			log.Printf("Error getting GPT suggestions for instance %s: %v", instance.Name, err)
// 			continue
// 		}

// 		fmt.Printf("GPT suggested improvements for instance %s: %s\n", instance.Name, suggestion)
// 	}
// }
