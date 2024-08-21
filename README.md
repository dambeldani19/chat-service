step menjalankan chat service

1. buat database baru di local, import database dump-chat-service-202408220217.sql 
2. jalan kan service run man.go
3. import file openapi.yaml di build-openapi/openapi.yaml, di postman
4. untuk akses api, user harus register terlebih dahulu
5. setelah itu login menggunakan email untuk mendapatkan Token
6. step untuk send message
   - di butuh kan Bearer Token yang di dapat dari Login
   - sesuaikan sender_id dengan user login

docker untuk menjalankan docker
1. docker build --tag chat-service .
2. docker compose up



   
 