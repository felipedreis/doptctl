package context

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type ClientContext struct {
	Name    string
	Host    string
	Port    string
	current bool
}

func (ctx ClientContext) URL() string {
	return ctx.Host + ":" + ctx.Port
}

var (
	homeDir, _         = os.UserHomeDir()
	doptctlContextDir  = filepath.Join(homeDir, ".doptctl", "contexts")
	doptctlContextFile = filepath.Join(homeDir, ".doptctl", "contexts", "current.json")
	clientContext      *ClientContext
)

func LoadContext() (ClientContext, error) {
	ctx, err := readContext("current.json")

	if os.IsNotExist(err) {
		return ClientContext{}, fmt.Errorf("no current context set. Try running doptctl context configure or doptctl context set")
	}

	if err != nil {
		return ClientContext{}, fmt.Errorf("couldn't load current context: %w", err)
	}

	setContext(ctx)

	return *ctx, nil
}

func initializeContextDir() error {
	_, err := os.Stat(doptctlContextDir)
	if os.IsNotExist(err) {
		log.Println("Dir doesn't exist, creating it")
		err := os.MkdirAll(doptctlContextDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("couldn't create context dir: %w", err)
		}
	}
	return nil
}

func createNewContext() (*ClientContext, error) {
	var newContext ClientContext
	var confirm string

	err := initializeContextDir()
	if err != nil {
		return nil, err
	}

	fmt.Println("Context Name: ")
	fmt.Scanln(&newContext.Name)

	newContextFilePath := filepath.Join(doptctlContextDir, newContext.Name+".json")

	_, err = os.Stat(newContextFilePath)

	if err == nil {
		fmt.Println("Context already exists, do you wanna override it[Y/n]?")
		fmt.Scanln(&confirm)
		if confirm == "n" {
			return nil, nil
		}
	}

	fmt.Print("D-Optimas Host: ")
	fmt.Scanln(&newContext.Host)

	fmt.Print("D-Optimas Port: ")
	fmt.Scanln(&newContext.Port)

	err = writeContext(newContext)
	if err != nil {
		return nil, err
	}

	return &newContext, nil
}

func overrideCurrentContext(contextName string) error {
	currentContextFile, err := os.Create(doptctlContextFile)

	if err != nil {
		return fmt.Errorf("couldn't open current context file: %w", err)
	}
	defer currentContextFile.Close()

	ctxFile, err := os.Open(filepath.Join(doptctlContextDir, contextName+".json"))

	if err != nil {
		return fmt.Errorf("couldn't open %s context file: %w", contextName, err)
	}
	defer ctxFile.Close()

	_, err = io.Copy(currentContextFile, ctxFile)
	return err
}

func getAllContexts() ([]ClientContext, error) {
	files, err := os.ReadDir(doptctlContextDir)
	var doptctlContexts []ClientContext

	if err != nil {
		return nil, fmt.Errorf("could not read doptctl context dir: %w", err)
	}

	current, _ := readContext("current.json")

	for _, file := range files {
		if file.IsDir() || file.Name() == "current.json" {
			continue
		}
		ctx, err := readContext(file.Name())
		if err == nil {
			if current != nil && ctx.Name == current.Name {
				ctx.current = true
			}
			doptctlContexts = append(doptctlContexts, *ctx)
		}
	}

	return doptctlContexts, nil
}

func writeContext(ctx ClientContext) error {
	ctxFilePath := filepath.Join(doptctlContextDir, ctx.Name+".json")
	ctxBytes, marshalError := json.Marshal(ctx)
	if marshalError != nil {
		return fmt.Errorf("couldn't marshal context to json: %w", marshalError)
	}

	newContextFile, fileError := os.Create(ctxFilePath)
	if fileError != nil {
		return fmt.Errorf("couldn't create new context file %s: %w", ctxFilePath, fileError)
	}
	defer newContextFile.Close()

	_, err := newContextFile.Write(ctxBytes)
	return err
}

func readContext(fileName string) (*ClientContext, error) {
	var context ClientContext
	file, error := os.Open(filepath.Join(doptctlContextDir, fileName))

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
