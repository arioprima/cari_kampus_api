create table akreditasi(
    id varchar(36) primary key,
    akreditasi_nasioanal varchar(10),
    akreditasi_internasional varchar(10),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp
);