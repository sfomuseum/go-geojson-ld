package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-geojson-ld"
	"github.com/tidwall/pretty"
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

		body, err := geojson.AsGeoJSONLDWithReader(ctx, fh)

		if err != nil {
			log.Fatal(err)
		}

		body = pretty.Pretty(body)

		fmt.Println(string(body))
	}

}
