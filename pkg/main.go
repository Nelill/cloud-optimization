package main

import (
	"fmt"
	"log"

	"github.com/Nelill/cloud-optimization/ai"
	"github.com/Nelill/cloud-optimization/gcp"
)

func main() {
	// Configurez l'accès à l'API Compute Engine
	projectID := "agicap-tech-dev" // Remplacez par votre ID de projet GCP
	zone := "europe-west1-b"

	err := gcp.Configure(projectID, zone)
	if err != nil {
		log.Fatalf("Failed to configure GCP: %v", err)
	}

	// Récupérez les instances Compute Engine
	instances, err := gcp.GetInstances("us-central1-a")
	if err != nil {
		log.Fatalf("Failed to get instances: %v", err)
	}

	// Analysez chaque instance et demandez des suggestions d'optimisation des coûts
	for _, instance := range instances {
		instanceInfo := fmt.Sprintf("Instance Name: %s, Instance Type: %s", instance.Name, instance.MachineType)
		fmt.Println("Analyzing instance:", instanceInfo)

		suggestion, err := ai.RequestGPTSuggestions(instanceInfo)
		if err != nil {
			log.Printf("Error getting GPT suggestions for instance %s: %v", instance.Name, err)
			continue
		}

		fmt.Printf("GPT suggested improvements for instance %s: %s\n", instance.Name, suggestion)
	}
}
