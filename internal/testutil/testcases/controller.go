package testcases

type ControllerTest[T any] struct {
	Name               string
	Data               T
	ExpectedStatusCode int
	ExpectedResponse   any
}
