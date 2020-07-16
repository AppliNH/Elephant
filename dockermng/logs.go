package dockermng

import (
	. "applinh/elephant/models"
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func ReadLogs(cli *client.Client, containers map[string]RunningContainer) {
	fmt.Println(len(containers))
	c := make(chan map[string]*bufio.Reader)
	readers := map[string]*bufio.Reader{}

	for _, container := range containers {
		hi, _ := cli.ContainerAttach(context.Background(), container.ID, types.ContainerAttachOptions{Stdout: true, Logs: true, Stderr: true, Stream: true})
		readers[container.ID] = hi.Reader
	}

	for id, reader := range readers {
		go follow(map[string]*bufio.Reader{id: reader}, c)
	}
	// ui.Init()
	// w := createLogBoxes(containers)
	// ui.Render(w...)
	for msg := range c {
		if msg != nil {
			fmt.Println("pop it")
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

func follow(r map[string]*bufio.Reader, c chan map[string]*bufio.Reader) {
	fmt.Println("follow")
	for _,v := range r {
		io.Copy(os.Stdout,v)
	}
	c <- r
}
