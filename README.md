## Go OTP request and Authentication service

Run Server
```
go run cmd/main.go
```

Send OTP
``` 
curl -H "Content-Type: application/json" -X POST
 -d `{"phoneNumber": "+12384747473"}` http://localhost:8000/otp
```

Verify OTP 
```
curl -H "Content-Type application/json" -X POST
 -d `{"phoneNumber": "+12384747473", "code": "123456"}`
  http://localhost:8000/verify-otp
```