create table akreditasi_jurusan(
    id varchar(36) primary key,
    akreditasi varchar(10),
    akreditasi_id varchar(36),
    jurusan_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (jurusan_id) references jurusan(id),
    foreign key (akreditasi_id) references akreditasi(id)
);