package geojsonld

import (
	"bytes"
	"context"
	_ "fmt"
	"io"
	"os"
	"testing"
)

func TestAsGeoJSONLDWithReader(t *testing.T) {

	path_geojson := "fixtures/102527513.geojson"
	path_geojsonld := "fixtures/102527513.geojsonld"

	r, err := os.Open(path_geojson)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path_geojson, err)
	}

	defer r.Close()

	ld_r, err := os.Open(path_geojsonld)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path_geojsonld, err)
	}

	defer ld_r.Close()

	ctx := context.Background()

	body, err := AsGeoJSONLDWithReader(ctx, r)

	if err != nil {
		t.Fatalf("Failed to convert %s, %v", path_geojson, err)
	}

	expected, err := io.ReadAll(ld_r)

	if err != nil {
		t.Fatalf("Failed to read expected GeoJSON-LD, %v", err)
	}

	if !bytes.Equal(body, expected) {
		t.Fatalf("Unexpected output")
	}

}
