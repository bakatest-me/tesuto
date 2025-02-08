const baseURL = "http://localhost:8080";
const setup = {
    url: `${baseURL}/api/v1/login`,
    method: "POST",
    headers: {
        "Content-Type": "application/json",
    },
};

var testcase = {
    "should return 200": {
        body: {
            username: "test_username",
            password: "test_password",
        },
        expected: (res) => res.status === 200,
    },
    "should return 401": {
        body: {
            username: "test_username",
            password: "failed_password",
        },
        expected: (res) => res.status === 401,
    },
};
