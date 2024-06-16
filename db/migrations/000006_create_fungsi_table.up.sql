create table fungsi(
    id varchar(36) primary key,
    tentang_kampus_id varchar(36),
    fungsi text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (tentang_kampus_id) references tentang_kampus(id)
);