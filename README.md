# Analisis Kompleksitas Kode

##Kompleksitas Waktu
Kompleksitas waktu dihitung berdasarkan seberapa lama fungsi berjalan relatif terhadap ukuran inputnya. Panjang string bracket akan dilambangkan dengan 'n'.

Loop Utama:

Fungsi ini memiliki satu loop for yang mengiterasi setiap karakter dalam string bracket.
Karena panjang string bracket adalah n, loop ini akan melakukan n iterasi.
Operasi dalam Loop:

Setiap operasi dalam loop (pengecekan switch, penambahan atau penghapusan dari slice tmp, dan pengecekan panjang slice) memiliki kompleksitas konstan, yaitu O(1).
Dengan demikian, kompleksitas waktu dari fungsi ini adalah O(n)

##Kompleksitas Ruang
Kompleksitas ruang dihitung berdasarkan seberapa banyak memori tambahan yang digunakan fungsi relatif terhadap ukuran inputnya.

Penggunaan Memori:
Fungsi ini menggunakan sebuah slice tmp yang pada kasus terburuk dapat menampung seluruh karakter dalam string bracket.
Dalam kasus terburuk, semua karakter dalam bracket adalah karakter pembuka (misalnya, "((({{{[[["), sehingga slice tmp akan menyimpan n karakter.
Dengan demikian, kompleksitas ruang dari fungsi ini adalah O(n)

Penjelasan ini menunjukkan bahwa baik waktu eksekusi maupun penggunaan memori dari fungsi balancedBracket akan meningkat secara linear seiring dengan bertambahnya panjang input bracket.
