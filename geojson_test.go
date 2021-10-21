package geojsonld

import (
	"context"
	"fmt"
	"testing"
	"os"
)

func TestAsGeoJSONLDWithReader(t *testing.T) {
	
	path := "fixtures/102527513.geojson"
	
	r, err := os.Open(path)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path, err)
	}

	defer r.Close()

	ctx := context.Background()
	
	body, err := AsGeoJSONLDWithReader(ctx, r)

	if err != nil {
		t.Fatalf("Failed to convert %s, %v", path, err)
	}

	fmt.Println(string(body))
}
