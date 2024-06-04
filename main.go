package main

import (
	"flag"
	"fmt"
	"log"

	twecon "github.com/theobori/teeworlds-econ"
	"github.com/theobori/terraform-teeworlds/command"
)

func main() {
	// Configure then parse the CLI arguments using the `flag` module
	host := flag.String("host", "127.0.0.1", "Econ server host")
	port := flag.Uint("port", 7000, "Econ server port")
	password := flag.String("password", "hello_world", "Econ server password")
	dir := flag.String("dir", ".", "Terraform directory")
	verbose := flag.Bool("v", false, "Verbose")
	broadcast := flag.Bool("broadcast", false, "In-game broadcast the destroyed Terraform resource name")

	flag.Parse()

	// Econ server configuration
	config := twecon.EconConfig{
		Host:     *host,
		Port:     uint16(*port),
		Password: *password,
	}

	// Create the econ controller
	econ := twecon.NewEcon(&config)

	// Connect to the econ server
	if err := econ.Connect(); err != nil {
		log.Fatalln(err)
	}

	// Authenticate to the econ server
	if _, err := econ.Authenticate(); err != nil {
		log.Fatalln(err)
	}

	t := command.NewTerraform(*dir)

	// Create a new event for flag captured (Teeworlds 0.7)
	capture := twecon.EconEvent{
		Name:  "flag_captured",
		Regex: `\[game\]: flag_capture`,
		Func: func(econ *twecon.Econ, eventPayload string) any {
			resourceName, err := t.DestroyRandom()
			if err != nil {
				return nil
			}

			msg := fmt.Sprintf("Destroyed %s", resourceName)
			
			if *verbose {
				log.Println(msg)
			}

			if *broadcast {
				_ = econ.Send("broadcast " + msg + " !")
			}

			return nil
		},
	}

	// Register the event
	if err := econ.EventManager.Register(&capture); err != nil {
		log.Fatalln(err)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)

	log.Printf("Listening event of %s", address)

	econ.HandleEvents()
}
