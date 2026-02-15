package main

import (
	"fmt"
	"log"

	"ovh-orphans/models"

	"github.com/ovh/go-ovh/ovh"
)

func main() {
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		log.Fatalf("Error creating OVH client : %s", err)
	}
	var proj []string
	err = client.Get("/cloud/project", &proj)
	if err != nil {
		log.Fatalf("Error fetching projects: %s", err)
	}

	fmt.Println("Searching for orphan volumes...")

	for _, projectID := range proj {
		fmt.Println("Looking in project id :", projectID, "...")

		var volumes []models.Volume
		endpoint := fmt.Sprintf("/cloud/project/%s/volume", projectID)

		if err := client.Get(endpoint, &volumes); err != nil {
			log.Printf("Could not fetch volumes for project %s: %s", projectID, err)
			continue
		}

		for _, vol := range volumes {
			// An orphan volume has no attachments
			if len(vol.AttachedTo) == 0 {
				fmt.Println("---------------------------------")
				fmt.Printf("| ORPHAN FOUND - volume unattached !\n")
				fmt.Printf("| Project: %s\n", projectID)
				fmt.Printf("| ID:      %s\n", vol.Id)
				fmt.Printf("| Name:    %s\n", vol.Name)
				fmt.Printf("| Size:    %d GB\n", vol.Size)
				fmt.Printf("| Region:  %s\n", vol.Region)
				fmt.Println("---------------------------------")
			}
		}
	}
}
