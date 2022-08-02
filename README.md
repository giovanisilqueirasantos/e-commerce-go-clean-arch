# e-commerce-go-clean-arch
a backend of an e-commerce to study clean architecture in go lang.

## mysql commit to up the database:
docker run --detach --name=gocleanarch-db --env="MYSQL_ROOT_PASSWORD=rootpass" --env="MYSQL_PASSWORD=password" --env="MYSQL_USER=user" --env="MYSQL_DATABASE=gocleanarch" --publish 3306:3306 --volume=$(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql mysql:5.7

## routes of the aplication

/signup

```json
{
	"login": "user@test.com",
	"password": "Password123$",
	"email": "user@test.com",
	"firstName": "test",
	"lastName": "test",
	"phoneNumber": "(11) 98888-8888",
	"address": {
		"city": "city",
		"state": "state",
		"neighborhood": "neighborhood",
		"street": "street",
		"number": "432",
		"zipcode": "zipcode"
	}
}
```

/login

```json
{
	"login": "user@test.com",
	"password": "Password123$"
}
```

/forgotpass/code

```json
{
	"login": "user@test.com"
}
```

/forgotpass/reset

```json
{
	"login": "user@test.com",
  "code": "n90VAn",
	"newPassword": "Password1234$"
}
```

/products/:uuid  Header (Authorization = Token)
