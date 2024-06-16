create table transport(
    id varchar(36) primary key,
    nama varchar(255),
    gambar varchar(255),
    kampus_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (kampus_id) references kampus(id)
);