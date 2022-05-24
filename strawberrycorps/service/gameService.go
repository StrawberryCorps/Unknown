package service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"unknown/strawberrycorps/ustruct"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetGamesService(c *gin.Context) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container)
		c.JSON(200, container)
	}
}

func CreateGameService() {

}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetGameLogsService(game ustruct.GameUri, c *gin.Context) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ContainerLogs(ctx, game.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: false,
		Timestamps: false,
		Details:    false,
		Follow:     true})

	if err != nil {
		panic(err)
	}

	ws, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	dec := transform.NewReader(reader, charmap.Windows1252.NewDecoder())

	scanner := bufio.NewScanner(dec)
	for scanner.Scan() {
		err = ws.WriteMessage(1, scanner.Bytes())
		if err != nil {
			fmt.Println(scanner.Text())
			break
		}
	}
}
