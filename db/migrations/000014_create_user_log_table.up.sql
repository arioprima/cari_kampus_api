create table user_log(
    id varchar(36) primary key,
    user_id varchar(36),
    status_login varchar(10),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (user_id) references users(id)
);