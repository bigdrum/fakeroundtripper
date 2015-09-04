# common
--
    import "github.com/bigdrum/fakeroundtripper"


## Usage

#### type FakeRoundTripper

```go
type FakeRoundTripper struct {
}
```

FakeRoundTripper for unitests.

#### func  NewFakeRoundTripper

```go
func NewFakeRoundTripper() *FakeRoundTripper
```
NewFakeRoundTripper creates a new FakeRoundTripper.

#### func (*FakeRoundTripper) BindContent

```go
func (f *FakeRoundTripper) BindContent(url string, content string)
```
BindContent binds the content of a string to a URL.

#### func (*FakeRoundTripper) BindFile

```go
func (f *FakeRoundTripper) BindFile(url string, testDataFileName string)
```
BindFile binds the content of a file to a URL.

#### func (*FakeRoundTripper) BindHandler

```go
func (f *FakeRoundTripper) BindHandler(prefix string, handler http.Handler)
```
BindHandler binds a handler to a URL.

#### func (*FakeRoundTripper) RoundTrip

```go
func (f *FakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error)
```
RoundTrip implementation.
