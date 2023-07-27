package utils

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func Sigs() chan os.Signal {
	// cap at 3
	sem := make(chan os.Signal, 3)
	sigs := make(chan os.Signal, 3)
	go func(sem chan os.Signal) {
		for sig := range sem {
			log.Warn().Str("signal", sig.String()).Msg("signal received")
			sigs <- sig
		}
	}(sem)
	signal.Notify(sem,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	return sigs
}

func WaitForTermination(
	try_close func() error,
	sigs chan os.Signal,
) error {
	var err error
	for i := 0; i <= 3; i++ {
		<-sigs
		if err = try_close(); err == nil {
			return nil
		}
	}
	panic(err)
}
