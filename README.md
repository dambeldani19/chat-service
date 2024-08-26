# Chat Service

## Step menjalankan Chat Service

1. **Setup .env File:**
   - Buat file `.env` di root direktori proyek Anda.
   - Tambahkan konfigurasi berikut ke dalam file `.env` (sesuaikan dengan kebutuhan Anda):
     ```
     PORT = 8080
     DB_USERNAME="root"
     DB_PASSWORD="root"
     DB_URL="127.0.0.1:3306"
     DB_DATABASE="chat-service"
     ```

2. **Buat Database:**
   - Buat database baru dengan nama `chat-service` di local.
   - Import file `database dump` dengan nama `dump-chat-service-202408261708.sql` ke dalam database yang telah dibuat.

3. **Jalankan Service:**
   - Jalankan service dengan perintah:
     ```bash
     go run man.go
     ```

4. **Import OpenAPI:**
   - Import file `openapi.yaml` yang terletak di `build-openapi/openapi.yaml` ke dalam Postman.

5. **Akses API:**
   - User harus melakukan register terlebih dahulu.
   - Setelah register, login menggunakan email untuk mendapatkan Bearer Token.

6. **Langkah untuk Mengirim Pesan:**
   - Gunakan Bearer Token yang didapat dari proses login.
   - Sesuaikan `sender_id` dengan ID pengguna yang melakukan login.

## OpenAPI

1. **Build OpenAPI:**
   - Untuk membangun OpenAPI, jalankan perintah:
     ```bash
     make openapi-gen
     ```

## API-Test

1. **Jalankan API-Test:**
   - Untuk menjalankan API-test, gunakan perintah:
     ```bash
     go test -v ./test/api_test.go
     ```
   - Atau jalankan dengan `make`:
     ```bash
     make go-test
     ```

## Docker

1. **Build Docker Image:**
   - Untuk membangun Docker image, jalankan perintah:
     ```bash
     docker build --tag chat-service .
     ```

2. **Jalankan Docker Compose:**
   - Untuk menjalankan Docker Compose, gunakan perintah:
     ```bash
     docker compose up
     ```

users
 id
 name
 image
 created_at

#api users
 - login
 - register
 - get detail users by id

#categoris
 id
 name
 created_at

#artikel_categoris
 artikel_id
 category_id


artikels
 id
 creator_id // refer user_id
 title
 image
 content
 like
 created_at

 #api artikel
 - get list artikel
 - get detail artikel by id 
 - comment artikel by artikel_id dan user_login_id
 

comment_artikels
 id
 artikel_id
 comment
 created_at

emoji_artikels
 id
 artikel_id
 type

#api send emoji by artikel_id


