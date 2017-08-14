package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	reader, _ := cli.Container.Logs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
	})
	defer reader.Close()

	//read the first 8 bytes to ignore the HEADER part from docker container logs
	p := make([]byte, 8)
	reader.Read(p)
	content, _ := ioutil.ReadAll(reader)

	var codeOutput MyJSONStruct
	if err := json.NewDecoder(strings.NewReader(string(content))).Decode(&codeOutput); err != nil {
		// hande error
	}
	//set some other value in stru
	codeOutput.ContainerID = containerID
	fmt.Println("vim-go")
}
