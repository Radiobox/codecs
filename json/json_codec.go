package json

import (
	jsonEncoding "encoding/json"
	"github.com/stretchrcom/web"
)

// JsonCodec converts objects to and from JSON.
type JsonCodec struct{}

// Converts an object to JSON.
func (c *JsonCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return jsonEncoding.Marshal(object)
}

// Unmarshal converts JSON into an object.
func (c *JsonCodec) Unmarshal(data []byte, obj interface{}) error {
	return jsonEncoding.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (c *JsonCodec) ContentType() string {
	return web.ContentTypeJson
}

// FileExtensions returns the file extension for this codec.
func (c *JsonCodec) FileExtensions() string {
	return web.FileExtensionJson
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *JsonCodec) CanMarshalWithCallback() bool {
	return false
}
