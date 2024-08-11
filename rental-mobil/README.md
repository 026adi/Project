# Project Rental Mobil

## Deskripsi
Project Rental Mobil adalah aplikasi untuk mengelola rental mobil termasuk peminjaman dan fitur terkait. Proyek ini mencakup fitur-fitur seperti registrasi dan login pengguna, autentikasi dasar menggunakan JWT, serta operasi CRUD untuk data master mobil dan transaksi rental.

## Fitur
1. **Autentikasi Pengguna:**
   - **Register:** Pengguna baru dapat mendaftar dengan username dan password.
   - **Login:** Pengguna yang terdaftar dapat masuk dan menerima token JWT untuk autentikasi.

2. **Manajemen Mobil:**
   - **CRUD Mobil:** Admin dapat membuat, membaca, memperbarui, dan menghapus data mobil yang tersedia untuk disewakan.

3. **Rental Mobil:**
   - **CRUD Rental:** Pengguna dapat melakukan pemesanan rental mobil, melihat detail rental, memperbarui informasi rental, dan membatalkan rental.

4. **Pembayaran:**
   - **CRUD Pembayaran:** Mengelola pembayaran untuk rental, termasuk pembuatan catatan pembayaran dan pengecekan status pembayaran.

5. **Autorisasi:**
   - Setiap endpoint dilindungi oleh JWT, memastikan hanya pengguna yang terautentikasi yang dapat mengakses fitur-fitur tertentu.