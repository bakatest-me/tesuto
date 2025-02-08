const baseURL = "http://localhost:8080";

const setup = {
    url: `${baseURL}/api/v1/users/{id}`,
    method: "GET",
    headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer 1234567890",
    },
};

var testcase = {
    "should return 200": {
        params: {
            id: "1234567890",
        },
        expected: (res) => res.status === 200,
    },
};
