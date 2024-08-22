# Api Gateway

There is required fields that need to be set in the `.env` file.

| Name              | Required | Default | Description                           |
|-------------------|----------|---------|---------------------------------------|
| SERVER_PORT       | yes      |         | :8001                                 |
| READ_TIMEOUT      | yes      |         | 1s                                    |
| AUTH_API_URL      | yes      |         | http://auth-api:8081                  |
| RESOURCES_API_URL | yes      |         | http://resources-api:8081             |

Available endpoints:
1. POST /api/v1/users/sign_in
2. GET /api/v1/books
3. GET /api/v1/users - with authorization