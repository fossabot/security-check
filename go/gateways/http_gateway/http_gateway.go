package http_gateway

type HTTPGetResponse struct {
	ResponseBody string
	ResponseCode int
}

type HTTPGateway interface {
	Get(url string) HTTPGetResponse
}

// Fake implementation of HTTPGateway for use during testing. Responses stubs
// are defined when the fake is used.
type HTTPGatewayFake struct {
	ResponseStub HTTPGetResponse
}

func (httpGatewayStub *HTTPGatewayFake) Get(url string) HTTPGetResponse {
	return httpGatewayStub.ResponseStub
}
