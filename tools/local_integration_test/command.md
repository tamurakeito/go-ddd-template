ローカルでインテグレーションテストする際にコピーして使ってください

**`hello-world`**

```
curl http://localhost:8080/hello-world
```

**`sign-in`**

```
curl -X POST \
-H "Content-Type: application/json" \
-d '{
  "user_id": "exampleUser",
  "password": "examplePassword",
}' \
http://localhost:8080/sign-in
```

**`sign-up`**

```
curl -X POST \
-H "Content-Type: application/json" \
-d '{
  "user_id": "exampleUser",
  "password": "examplePassword",
  "name": "Example Name"
}' \
http://localhost:8080/sign-up
```
