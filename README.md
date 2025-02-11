### Now Feature v0.0.1:

-   [x] Support run test.js file or with folder
-   [x] Generate curl command from testcase
-   [x] Support debug request and response
-   [x] Support test api with http method

### Upcoming Feature:

-   [ ] Pre-request (setup, testcase)
-   [ ] Post-response (setup, testcase)
-   [ ] Build in function
    -   [ ] faker https://github.com/go-faker/faker
-   [ ] Support upload file
-   [ ] Load env to test.js file
-   [ ] Summary testcase report
-   [ ] Export report
    -   [ ] JSON
    -   [ ] CSV

### Installation

```bash
# Download the latest release for your platform
wget https://github.com/bakatest-me/tesuto/releases/download/v0.0.1/tesuto-darwin.tar.gz
# or other platform
# https://github.com/bakatest-me/tesuto/releases

# Extract the archive
tar -xzvf tesuto-darwin.tar.gz

# Move the tesuto binary to /usr/local/bin
sudo mv tesuto /usr/local/bin/tesuto

# Verify the installation
tesuto version
```

### Usage:

Run testcase file:

```bash
tesuto run example/get_user_by_id.js
```

Run testcases in folder:

```bash
tesuto run example/
```

### Example:

-   example/get_user_by_id.js
-   example/query_user.js
-   example/login.js

_Mock api with mockoon file `example/mockoon.json` for testing example files_

### Quick Start:

1. Create js file for test script

```bash
touch testcase/get_user_by_id.js
```

2. Setup Object:

```js
const baseURL = "http://localhost:8080";
const setup = {
    url: `${baseURL}/api/v1/users/{id}`,
    method: "GET",
    headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer 1234567890",
    },
};
```

The setup object contains the configuration for the HTTP request:

-   url: The endpoint for fetching user data, where {id} is a placeholder for the user ID.
-   method: Specifies that the request method is GET.
-   headers: Sets the request headers, including the content type and an authorization token.

3. Testcase Object:

```js
var testcase = {
    "should return 200": {
        params: {
            id: "1234567890",
        },
        expected: (res) => res.status === 200,
    },
};
```

The testcase object contains the test cases:

-   "should return 200": The name of the test case.
-   params: The parameters for the test case.
-   expected: The expected result of the test case.

Finally file:

```js
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
```

4. Run testcase:

```bash
tesuto run testcase/get_user_by_id.js
```

### Full config testcase:

-   params: params with replace `{key}` with value
-   query: will add to url as query string `?key=value`

```js
var testcase = {
    "should return 200": {
        params: {
            id: "1234567890",
        },
        query: {
            name: "book",
        },
        expected: (res) => res.status === 200,
    },
};
```

### Flag options:

```bash
tesuto run <testcase_file.js or directory>
```

#### debug

show debug request and response
`--debug`
`-d`

#### curl

generate curl command from testcase
`--curl`
`-c`

Reference:

-   HTTP Request [Resty](https://github.com/go-resty/resty) for now use v3
