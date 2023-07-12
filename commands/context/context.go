package context

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ClientContext struct {
	Name    string
	Host    string
	Port    string
	current bool
}

func (ctx ClientContext) URL() string {
	return ctx.Name + ":" + ctx.Port
}

var (
	homeDir, _         = os.UserHomeDir()
	doptctlContextDir  = homeDir + "/.doptctl/contexts/"
	doptctlContextFile = homeDir + "/.doptctl/contexts/current.json"
	clientContext      *ClientContext
)

func LoadContext() (ClientContext, error) {
	ctx, error := readContext(doptctlContextFile)

	if os.IsNotExist(error) {
		log.Fatal("No Current context set. Try running doptctl context configure or doptctl context set")
		os.Exit(1)
	}

	setContext(ctx)

	return *ctx, nil
}

func initializeContextDir() {
	_, error := os.Stat(doptctlContextDir)
	if os.IsNotExist(error) {
		log.Println("Dir doesn't exist, creating it")
		error := os.MkdirAll(doptctlContextDir, os.ModePerm)
		if error != nil {
			log.Fatal("Couldn't create context dir")
			os.Exit(1)
		}

	}
}

func createNewContext() *ClientContext {
	var newContext ClientContext
	var confirm string

	initializeContextDir()

	fmt.Println("Context Name: ")
	fmt.Scanln(&newContext.Name)

	newContextFilePath := doptctlContextDir + newContext.Name + ".json"

	_, error := os.Stat(newContextFilePath)

	if error == nil {
		fmt.Println("Context already exists, do you wanna override it[Y/n]?")
		fmt.Scanln(&confirm)
		if confirm == "n" {
			return nil
		}
	}

	fmt.Print("D-Optimas Host: ")
	fmt.Scanln(&newContext.Host)

	fmt.Print("D-Optimas Port: ")
	fmt.Scanln(&newContext.Port)

	writeContext(newContext)

	return &newContext
}

func overrideCurrentContext(contextName string) {
	currentContextFile, error := os.Create(doptctlContextFile)

	if error != nil {
		log.Fatal("Couldn't open Current context file")
		os.Exit(1)
	}
	defer currentContextFile.Close()

	ctxFile, error := os.Open(doptctlContextDir + contextName + ".json")

	if error != nil {
		log.Fatal("Couldn't open " + contextName + " context file")
		os.Exit(1)
	}
	defer ctxFile.Close()

	io.Copy(currentContextFile, ctxFile)
}

func getAllContexts() []ClientContext {
	files, error := os.ReadDir(doptctlContextDir)
	var doptctlContexts []ClientContext

	if error != nil {
		log.Fatal("Could not read doptctl context dir")
		os.Exit(1)
	}

	current, _ := readContext("current.json")

	for _, file := range files {
		if file.IsDir() || file.Name() == "current.json" {
			continue
		}
		fmt.Println(file.Name())
		ctx, error := readContext(file.Name())
		if error == nil {
			if current != nil && ctx.Name == current.Name {
				ctx.current = true
			}
			doptctlContexts = append(doptctlContexts, *ctx)
		}
	}

	return doptctlContexts
}

func writeContext(ctx ClientContext) {
	ctxFilePath := doptctlContextDir + ctx.Name + ".json"
	ctxBytes, marshalError := json.Marshal(ctx)
	newContextFile, fileError := os.Create(ctxFilePath)

	if marshalError != nil {
		log.Fatal("Couldn't marshal context to json")
		os.Exit(1)
	}

	if fileError != nil {
		log.Fatal("Couldn't create new context file ", ctxFilePath)
		os.Exit(1)
	}
	defer newContextFile.Close()

	newContextFile.Write(ctxBytes)
}

func readContext(fileName string) (*ClientContext, error) {
	var context ClientContext
	file, error := os.Open(doptctlContextDir + fileName)

	if error != nil {
		return nil, error
	}
	defer file.Close()

	contextBytes, error := io.ReadAll(file)

	if error != nil {
		return nil, error
	}

	error = json.Unmarshal(contextBytes, &context)

	if error != nil {
		log.Println("Couldn't not unmarshal data")
		return nil, error
	}

	return &context, nil
}

func setContext(ctx *ClientContext) {
	clientContext = ctx
	clientContext.current = true
}
