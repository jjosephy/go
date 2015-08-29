package handler

import (
    "bytes"
    "encoding/json"
    "fmt"
    "interview/contract/v1"
    "interview/httperror"
    "io"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

type MockInterviewRepository struct {
}

func(r *MockInterviewRepository) GetInterview(id string, name string) string {
    return "Mock"
}

var h http.HandlerFunc
var ts *httptest.Server

func ValidateC1(t *testing.T, c contract.InterviewContractV1) {
    if c.Candidate != "Bob" {
        t.Fatal("Candiate name not correct")
    }

    if p := len(c.Comments); p != 3 {
        t.Fatal("Invalid Number of comments")
    }
}

func TestMain(m *testing.M) {
    h = InterviewHandler(new(MockInterviewRepository))
    ts = httptest.NewServer(http.HandlerFunc(h))
    defer ts.Close()
    os.Exit(m.Run())
}

func stringFromStream(b io.ReadCloser) string {
    buf := new(bytes.Buffer)
    buf.ReadFrom(b)
    return buf.String()
}

func readErrorResponse(b io.ReadCloser) (contract.ErrorContractV1, error) {
    var errDetail contract.ErrorContractV1
    body, err := ioutil.ReadAll(b)
    if err != nil {
        return errDetail, err
    }
    json.Unmarshal(body, &errDetail)

    return errDetail, nil
}

func readSuccessResponse(b io.ReadCloser) (contract.InterviewContractV1, error) {
    var c contract.InterviewContractV1
    body, err := ioutil.ReadAll(b)
    if err != nil {
        return c, err
    }
    json.Unmarshal(body, &c)
    return c, nil
}

func assertErrorEqual(
    t *testing.T,
    e contract.ErrorContractV1,
    code int,
    msg string,
    statusExpected int,
    statusReceived int) {
    if e.Code != code {
        t.Fatalf("Error Code is not correct expected %d, got %d", code, e.Code)
    }

    if e.Message != msg {
        t.Fatalf("Error Messages do not match: %s", e.Message)
    }

    if statusExpected != statusReceived {
        t.Fatalf(
            "Unexpected status code returned expected %d : received %d",
            statusExpected,
            statusReceived)
    }
}

func validateRequest(
    t *testing.T,
    uri string,
    headers map[string]string,
    errCode int,
    errMsg string,
    expectedHttpStatus int) {

    client := &http.Client {}
    req, err := http.NewRequest("GET", uri, nil)

    if err != nil {
        t.Fatalf("Unexpected error trying to create a request %v", err)
    }

    for k, v:= range headers {
        req.Header.Add(k, v)
    }

    resp, err := client.Do(req)
    if err != nil {
        t.Fatalf("Unexpected error %v", err)
    }

    switch (expectedHttpStatus) {
        case http.StatusOK:
            c, err := readSuccessResponse(resp.Body)
            if err != nil {
                t.Fatalf("Unexpected error reading body %v", err)
            } else {
                defer resp.Body.Close()
            }
            ValidateC1(t, c)
        case http.StatusBadRequest:
            errDetail, err := readErrorResponse(resp.Body)
            if err != nil {
                t.Fatalf("Unexpected error reading body %v", err)
            } else {
                defer resp.Body.Close()
            }

            assertErrorEqual(
                t,
                errDetail,
                errCode,
                errMsg,
                expectedHttpStatus,expectedHttpStatus)
    }
}

func Test_BadRequest_UnSupportedVersion_V1(t *testing.T) {
    headers := map[string]string{
        "Api-Version": "3.1",
    }

    validateRequest(
        t,
        fmt.Sprint(ts.URL, "?id=2"),
        headers,
        httperror.BADREQUEST_UNSUPPORTEDVERSION,
        httperror.UNSUPPORTED_VERSION,
        http.StatusBadRequest)
}

func TestBadRequest_NoQueryParameters_V1(t *testing.T) {
    headers := map[string]string{
        "Api-Version": "1.0",
    }

    validateRequest(
        t,
        ts.URL,
        headers,
        httperror.BADREQUEST_NOINPUTPARAMETERS,
        httperror.NO_PARAMETERS_PROVIDED,
        http.StatusBadRequest)
}

func TestBadRequest_InvalidVersion_V1(t *testing.T) {
    headers := map[string]string{
        "Api-Version": "invalid",
    }

    validateRequest(
        t,
        ts.URL,
        headers,
        httperror.BADREQUEST_INVALIDVERSION,
        httperror.INVALID_VERSION,
        http.StatusBadRequest)
}

func TestBadRequest_NoVersion_V1(t *testing.T) {
    headers := map[string]string{
        "No-Version": "",
    }

    validateRequest(
        t,
        ts.URL,
        headers,
        httperror.BADREQUEST_NOVERSION,
        httperror.NO_VERSION_PROVIDED,
        http.StatusBadRequest)
}

func Test_Success_ValidRequest_V1(t *testing.T) {
    headers := map[string]string{
        "Api-Version": "1.0",
    }

    validateRequest(
        t,
        fmt.Sprint(ts.URL, "?id=2"),
        headers,
        0,
        "",
        http.StatusOK)
}
