create table jurusan(
    id varchar(36) primary key,
    nama varchar(255),
    kode_prodi varchar(100),
    jenjang varchar(255),
    gambar varchar(255),
    profile varchar(255),
    fakultas_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (fakultas_id) references fakultas(id)
);