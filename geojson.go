package geojsonld

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io"
	"strings"
)

// NS_GEOJSON is the default namespace for GeoJSON-LD
const NS_GEOJSON string = "https://purl.org/geojson/vocab#"

// DefaultGeoJSONLDContext return a dictionary mapping GeoJSON property keys to their GeoJSON-LD @context equivalents.
func DefaultGeoJSONLDContext() map[string]interface{} {

	bbox := map[string]string{
		"@container": "@list",
		"@id":        "geojson:bbox",
	}

	coords := map[string]string{
		"@container": "@list",
		"@id":        "geojson:coordinates",
	}

	features := map[string]string{
		"@container": "@set",
		"@id":        "geojson:features",
	}

	ctx := map[string]interface{}{
		"geojson":            NS_GEOJSON,
		"Feature":            "geojson:Feature",
		"FeatureCollection":  "geojson:FeatureCollection",
		"GeometryCollection": "geojson:GeometryCollection",
		"LineString":         "geojson:LineString",
		"MultiLineString":    "geojson:MultiLineString",
		"MultiPoint":         "geojson:MultiPoint",
		"MultiPolygon":       "geojson:MultiPolygon",
		"Point":              "geojson:Point",
		"Polygon":            "geojson:Polygon",
		"bbox":               bbox,
		"coordinates":        coords,
		"features":           features,
		"geometry":           "geojson:geometry",
		"id":                 "@id",
		"properties":         "geojson:properties",
		"type":               "@type",
	}

	return ctx
}

// AsGeoJSONLDWithReader convert GeoJSON Feature data contained in r in to GeoJSON-LD.
func AsGeoJSONLDWithReader(ctx context.Context, r io.Reader) ([]byte, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read Feature data, %w", err)
	}

	return AsGeoJSONLD(ctx, body)
}

// AsGeoJSONLDWithReader convert GeoJSON Feature data contained in body in to GeoJSON-LD.
func AsGeoJSONLD(ctx context.Context, body []byte) ([]byte, error) {

	geojson_ctx := DefaultGeoJSONLDContext()

	props_rsp := gjson.GetBytes(body, "properties")

	if !props_rsp.Exists() {
		return nil, fmt.Errorf("Missing properties element")
	}

	for k, _ := range props_rsp.Map() {

		parts := strings.Split(k, ":")

		var k_fq string

		if len(parts) == 2 {

			ns := parts[0]
			pred := parts[1]

			// sudo make this dynamic / a callback / equivalent

			k_fq = fmt.Sprintf("https://github.com/whosonfirst/whosonfirst-properties/tree/master/properties/%s#%s", ns, pred)
		} else {

			k_fq = fmt.Sprintf("x-urn:geojson:properties#%s", k)
		}

		geojson_ctx[k_fq] = k
	}

	body, err := sjson.SetBytes(body, "@context", geojson_ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to assign @context, %w", err)
	}

	id_rsp := gjson.GetBytes(body, "id")

	if id_rsp.Exists() {

		body, err = sjson.SetBytes(body, "id", id_rsp.String())

		if err != nil {
			return nil, fmt.Errorf("Failed to assign id, %w", err)
		}
	}

	return body, nil
}
