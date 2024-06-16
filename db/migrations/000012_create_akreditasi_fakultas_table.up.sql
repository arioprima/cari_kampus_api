create table akreditasi_fakultas(
    id varchar(36) primary key,
    akreditasi varchar(10),
    akreditasi_id varchar(36),
    fakultas_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (fakultas_id) references fakultas(id),
    foreign key (akreditasi_id) references akreditasi(id)
);