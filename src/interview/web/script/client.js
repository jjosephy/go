function Client() {

}

Client.SaveInterview = function(i, success, error) {
    var s = JSON.stringify(i);
    $.ajax({
        url : 'http://localhost:8080/interview',
        type: 'POST',
        data : s,
        headers: {
            "Api-Version":  1.0
        },
        success: success,
        error: error,
    });
}

Client.GetInterview = function(id, cname, success, error) {
    $.ajax({
        url : 'http://localhost:8080/interview?id=' + id,
        type: 'GET',
        headers: {
            "Api-Version":  1.0
        },
        success: success,
        error: error,
    });
}