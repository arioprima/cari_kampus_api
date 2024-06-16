create table tentang_kampus(
    id varchar(36) primary key,
    tahun_berdiri varchar(4),
    visi text,
    tugas_pokok text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp
);
