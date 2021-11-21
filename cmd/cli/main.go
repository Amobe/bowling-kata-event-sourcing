package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/amobe/bowling-kata-event-sourcing/src/handler"
	"github.com/amobe/bowling-kata-event-sourcing/src/storage"
	"github.com/spf13/cobra"
)

func main() {
	r := storage.NewInmemEventStorage()
	h := handler.NewHandler(r)

	app := &cobra.Command{
		Use:   "bowling-cli",
		Short: "play a bowling game!",
	}
	app.AddCommand(rollCommand(h))
	err := app.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func rollCommand(h handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "roll [hit number between 0 ~ 10]",
		Short: "roll the ball with hit number",
		Run:   rollAction(h),
	}
}

func rollAction(h handler.Handler) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.PrintErr(fmt.Errorf("roll action: invalid arguments number"))
			return
		}
		hit, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			cmd.PrintErr(fmt.Errorf("roll action: invalid hit number: %w", err))
			return
		}
		err = h.Roll(uint32(hit))
		if err != nil {
			cmd.PrintErr(fmt.Errorf("roll action: handler error: %w", err))
			return
		}
		return
	}
}
