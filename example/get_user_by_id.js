const baseURL = "http://localhost:8080";

const setup = {
    url: `${baseURL}/api/v1/users/{id}`,
    method: "GET",
    headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer 1234567890",
    },
};

const testID = "1234567890";
var testcase = {
    "should return 200": {
        params: {
            id: testID,
        },
        expected: (res) => res.status === 200,
    },
    "should return valid name": {
        params: {
            id: testID,
        },
        expected: (res) => res.body.data.name === "caption america",
    },
    "should be hulk": {
        params: {
            id: testID,
        },
        expected: (res) => res.body.data.name === "hulk",
    },
};
