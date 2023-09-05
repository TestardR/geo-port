package cli

import (
	"context"
	"encoding/json"
	"github.com/TestardR/geo-port/internal/application/command"
	"os"
)

type portHandler interface {
	HandleAddOrUpdatePort(ctx context.Context, command command.AddOrUpdatePort) error
}

type errorLogger interface {
	Error(format string, v ...interface{})
}

type stdLogger interface {
	Printf(format string, v ...interface{})
}

type portCLI struct {
	portHandler     portHandler
	consoleOutputer stdLogger
	errorLogger     errorLogger
}

func NewPortCLI(
	portHandler portHandler,
	consoleOutputer stdLogger,
	errorLogger errorLogger,
) *portCLI {
	return &portCLI{
		portHandler:     portHandler,
		consoleOutputer: consoleOutputer,
		errorLogger:     errorLogger,
	}
}

func (c *portCLI) Handle(ctx context.Context, filePath string) {
	portsFile, err := os.Open(filePath)
	if err != nil {
		c.errorLogger.Error("failed to open file with file path %q: %v", filePath, err)

		return
	}
	defer func() {
		err := portsFile.Close()
		if err != nil {
			c.errorLogger.Error("failed to close file with path %q: %v", err)

			return
		}
	}()

	decoder := json.NewDecoder(portsFile)
	_, err = decoder.Token()
	if err != nil {
		c.errorLogger.Error("failed to decode file with path %q: %v", err)

		return
	}

	numPortsProcessed := 0
	for decoder.More() {
		select {
		case <-ctx.Done():
			return
		default:
			token, err := decoder.Token()
			if err != nil {
				c.errorLogger.Error("failed to parse JSON in file with path %q: %v", filePath, err)

				return
			}

			inputPortID := token.(string)
			var inputPort port
			if err := decoder.Decode(&inputPort); err != nil {
				c.errorLogger.Error("failed to parse JSON in file with path %q: %v", filePath, err)

				return
			}

			addOrUpdatePort := command.NewAddOrUpdatePort(
				inputPortID,
				inputPort.Name,
				inputPort.City,
				inputPort.Country,
				inputPort.Aliases,
				inputPort.Regions,
				inputPort.Coordinates[1], // latitude
				inputPort.Coordinates[0], // longitude
				inputPort.Province,
				inputPort.Timezone,
				inputPort.Unlocs,
				inputPort.Code,
			)

			err = c.portHandler.HandleAddOrUpdatePort(ctx, addOrUpdatePort)
			if err != nil {
				c.errorLogger.Error(err.Error())

				return
			}

			numPortsProcessed++
		}
	}

	c.consoleOutputer.Printf("Number of ports processed: %d", numPortsProcessed)
}
