# Load Balancer dengan NGINX

Proyek ini mengatur load balancer menggunakan NGINX untuk mendistribusikan lalu lintas ke tiga kontainer Docker yang menjalankan aplikasi Go. Konfigurasi ini menggunakan algoritma default NGINX untuk load balancing dan merupakan pengaturan dasar.

## Konfigurasi Aplikasi Go

Aplikasi Go ini menggunakan flag di command line untuk mengatur konten yang ditampilkan.

## Dockerfile

Dockerfile ini digunakan untuk membangun image aplikasi Go. Konten dan port dapat dikonfigurasi melalui flag.

## Cara Menjalankan Aplikasi

1. **Membangun Image Docker:**

    ```bash
    docker build -t golang-app ./app
    ```

2. **Menjalankan Tiga Kontainer:**

    ```bash
    docker run -d --name app1 -p 8081:8080 golang-app -content="Hello, Golang!" -port=8080
    docker run -d --name app2 -p 8082:8080 golang-app -content="Hello, Docker!" -port=8080
    docker run -d --name app3 -p 8083:8080 golang-app -content="Hello, Linux!" -port=8080
    ```

3. **Konfigurasi NGINX:**

    Salin file `nginx.conf` ke direktori konfigurasi NGINX Anda, lalu restart NGINX.

4. **Mengakses Load Balancer:**

    ```bash
    curl http://localhost/
    ```

NGINX akan mendistribusikan permintaan ke ketiga kontainer sesuai dengan algoritma default.
