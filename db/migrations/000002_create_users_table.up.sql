create table users(
    id varchar(36) primary key,
    nama varchar(255) not null,
    email varchar(255),
    password varchar(255) not null,
    role_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp,
    foreign key (role_id) references user_role(id)
)ENGINE = InnoDB;