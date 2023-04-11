# Sample Go API Gateway
The solution represents API Gateway implemented as microservice written in Go. The Gateway provides `/v1/users` endpoint that enables a consumer to:
- Read user data by `id` (ex. GET `/v1/users/1`)
- Create new users (ex. POST `/v1/users` with JSON body `{"firstName": "Rob", "lastName": "Smith", "Role": "Manager"}`)

## Assumptions ##
The solution assumes that it is backed by the UserService gRPC server that implements remote procedures described in `proto/user.proto` file. The UserService gRPC servise address should be provided using the `USER_SERVICE_ADDR` environment variable in the format of `host:port`.
In order to interact with API the gateway it awaits a HS256 JWT token to be attached in form of Authorization header to every HTTP request. The secret is required be defined as environment variable `JWT_SECRET`. It is expected that `id` and `role` claims are defined within JWT.

## Environment variables ##
There's a list of configuration options that can be defined in form of environment variables:
- `PORT` - An integer value holding the port number used by the Gateway to spin up the HTTP server
- `USER_SERVICE_ADDR` - A string in the `host:port` format that represents the address of a backing UserService gRPC server
- `JWT_SECRET` - A string holding the secret for decryption of JWT token using HS256 algorithm

## Solution description ##
The solution separates concerns.
There's http module that represents HTTP server based on Gin web framework.
The HTTP server uses separate handlers for each action.
JWT-based authentication is handled by custom `auth_handler` middleware.
Action handlers use concrete UserService implementation via the UserService interface.
Solution uses envitonment variables based configuration according to 12-factor app methodology.

## Testing ##
The code is covered with tests:
```
internal/http  coverage: 90.0% of statements
internal/services  coverage: 85.7% of statements
```

Mocking is used to test http package separately from the UserService client implementation and fake UserService gRPC server is used to test the UserService client implementation.

## Further development ##
The list of the features planned for the further implementation:
- [ ] Add ability to add users
- [ ] Add ability to update user
- [ ] Return proper 404 response if user not found by backing UserService
- [ ] Frobid role "Customer" from user creation (authorization middleware)
- [ ] Allow role "Customer" to only get user info about itself (id in JWT and in GET /user should match) (authorization middleware)
- [ ] Allow only "Admin" to create admins and managers (authorization middleware)
- [ ] Add email to user
- [ ] Add birthday to user
- [ ] Map particular UserService's errors to 400 in case of absence of required or invalid params
- [ ] Map particular UserService's errors to 422 responses in case of for example try to create user with already existing email