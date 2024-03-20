package servers

import (
	"fmt"
	"net/http"
	"strconv"
)

type ToursServer struct {
	BaseURL string
}

func NewToursServer(baseURL string) *ToursServer {
	return &ToursServer{BaseURL: baseURL}
}

func (c *ToursServer) UpdateCheckpointEncounter(checkpointID, encounterID string, isSecretPrerequisite bool) error {
	url := c.BaseURL + "/v1/tours/checkpoint/encounter/" + checkpointID + "/" + encounterID + "/" + strconv.FormatBool(isSecretPrerequisite)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status code, handle accordingly
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error, status code :", resp.StatusCode)
	}

	return nil
}
