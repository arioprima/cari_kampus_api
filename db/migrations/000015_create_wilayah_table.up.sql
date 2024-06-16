create table wilayah(
    kode varchar(13) NOT NULL PRIMARY KEY,
    nama varchar(255) default null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp
);