create table kampus(
    id varchar(36) primary key,
    nama varchar(255) not null,
    alamat varchar(255) not null,
    nomor_telepon varchar(15) not null,
    logo varchar(255) not null,
    gambar varchar(255) not null,
    website varchar(255) not null,
    status varchar(10) not null,
    akreditasi_id varchar(36) not null,
    tentang_kampus_id varchar(36) not null,
    provinsi varchar(255) not null,
    kabupaten_kota varchar(255) not null,
    kode_pos varchar(10) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (akreditasi_id) references akreditasi(id),
    foreign key (tentang_kampus_id) references tentang_kampus(id)
);