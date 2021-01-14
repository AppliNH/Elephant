package dockermng

import (
	. "github.com/applinh/elephant/models"
	"bufio"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func ReadLogs(cli *client.Client, containers map[string]RunningContainer) {

	c := make(chan *bufio.Reader, len(containers))
	readers := map[RunningContainer]*bufio.Reader{}

	for _, container := range containers {
		hi, _ := cli.ContainerAttach(context.Background(), container.ID, types.ContainerAttachOptions{Stdout: true, Logs: true, Stderr: true, Stream: true})
		readers[container] = hi.Reader
	}

	for container, reader := range readers {
		go follow(reader, container, c)
	}
	// ui.Init()
	// w := createLogBoxes(containers)
	// ui.Render(w...)

	for {
		select {
		case <-c:
			fmt.Println()
		}
	}

}

func createLogBoxes(containers map[string]RunningContainer) []termui.Drawable {

	widgetsList := []termui.Drawable{}
	var x int
	for _, container := range containers {

		p := widgets.NewParagraph()
		p.Title = container.Name
		p.SetRect(x, 0, 25, 50)
		widgetsList = append(widgetsList, p)
		x += 30
	}
	return widgetsList
}

func follow(r *bufio.Reader, container RunningContainer, c chan *bufio.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(container.Name + ": " + scanner.Text())
	}

	c <- r
}
