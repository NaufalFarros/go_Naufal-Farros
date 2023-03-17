## Rangkuman Materi DDL

* DDL atau Data Definition Language adalah bahasa pemrograman yang digunakan untuk membuat, mengubah, dan menghapus struktur database seperti tabel, kolom, dan indeks. Berikut adalah beberapa perintah DDL yang umum digunakan dalam SQL:
```
CREATE TABLE nama_tabel (
  kolom1 datatype,
  kolom2 datatype,
  kolom3 datatype,
  ...
);

ALTER TABLE nama_tabel
ADD kolom_baru datatype;

ALTER TABLE nama_tabel
DROP kolom_yang_tidak_diperlukan;

DROP TABLE nama_tabel;

CREATE INDEX nama_indeks
ON nama_tabel (kolom1, kolom2);

ALTER INDEX nama_indeks
RENAME TO nama_indeks_baru;

DROP INDEX nama_indeks;



```
