package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-geojsonld"
	"log"
	"os"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	for _, path := range flag.Args() {

		// go-cloud blob/bucket me

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		body, err := geojsonld.AsGeoJSONLDWithReader(ctx, fh)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}

}
