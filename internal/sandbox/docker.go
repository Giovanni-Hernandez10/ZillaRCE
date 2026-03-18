package sandbox

// docker orchestration logic -> spinning up containers, adding security, resource limits

import (
	// "context"

	"github.com/moby/moby/client"
)

// Struct to hold the created docker client
// has img and language associated with it
type ClientSandbox struct {
	cli      *client.Client
	img      string
	language string
}

// Method to create a new docker client
// Pass in the img and language to create the docker client for
func CreateDockerClient(img string, language string) (*ClientSandbox, error) {
	cli, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &ClientSandbox{cli: cli, img: img, language: language}, nil
}

// // Method to run the container with the desired image and command
// func (c *ClientSandbox) RunContainer(ctx context.Context, image string, cmd []string) {

// }
