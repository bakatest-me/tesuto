const baseURL = "http://localhost:8080";
const setup = {
    url: `${baseURL}/api/v1/users`,
    method: "GET",
    headers: {
        "Content-Type": "application/json",
    },
};

var testcase = {
    "should return 200": {
        query: {
            name: "book",
        },
        expected: (res) => res.status === 200,
    },
    "should return valid name": {
        query: {
            name: "book",
        },
        expected: (res) => res.body.data[0].name === "book",
    },
};
