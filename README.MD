# A. System Design
## 1. Flowchart Sistem
![Flowchart Sistem](./assets/dokumentation/image/Flowchart_Bioskop.drawio.png)

##  2. 
Status awal tiket adalah avaible dan saat tiket dibeli, ubah status menjadi pending. Jika pembayaran berhasil status menjadi succes. jika gagal status failed dan terlalu lama sampai waktu kadaluarsa tiket tidak dibayar, status akan expired. Tiket yang expired atau gagal akan di update statusnya lewat cron job menjadi avaible sehingga bisa dibeli lagi. Jika user melakukan refund, status yang succes tadi di perbarui menjadi avaible.

# B. Database Design
![ERD Bioskop](./assets/dokumentation/image/ERD_Bioskop.drawio.png)

# C. Project
Database migration on folder db

## Migration
On root type this command, one by one

```bash
migrate create -ext sql -dir db create_table_users
migrate create -ext sql -dir db create_table_movies
migrate create -ext sql -dir db create_table_studio_schedules
migrate create -ext sql -dir db create_table_studio_seats
migrate create -ext sql -dir db create_table_transactions
```

Then migrate table to your database with this command

```bash
migrate -path db -database "postgresql://username:db-password@host:port/db-name?sslmode=disable" up
```

## Run Project
To run this project, run a command

```bash
go mod tidy
```

and run the project with
```bash
go run cmd/web/main.go
```


## Postman File Export
[Download Postman Collection](./assets/dokumentation/postman/be_latihan_mkp.postman_collection.json)
