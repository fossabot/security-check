package http_gateway

import (
	"gotest.tools/assert"
	"testing"
)

func TestHTTPGatewayFake(t *testing.T) {
	expectedResponse := HTTPGetResponse{
		ResponseBody: "",
		ResponseCode: 200,
	}
	url := "https://example.com"
	gateway := HTTPGatewayFake{
		ResponseStub: expectedResponse,
	}

	actualResponse := gateway.Get(url)

	assert.Equal(t, actualResponse, expectedResponse)
}
