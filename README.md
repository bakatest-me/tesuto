### Roadmap:

0.1.0

-   [x] Support run test with folder
-   [ ] Support upload file

0.2.0

-   [ ] Summary testcase report
-   [ ] Export report
    -   [ ] JSON
    -   [ ] CSV

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
