
--1. Creat database alta_online_shop
create database alta_online_shop

--2. Creat table user
CREATE table user (
    ID int not null primary key auto_increment,
    name varchar(100),
    address varchar(255),
    birth_date date,
    status_user varchar(50),
    gender enum('laki-laki','wanita'),
    created_at datetime,
    updated_at datetime
);
-- Creat table item
create table item (
    id int primary key auto_increment,
    product varchar(50) not null,
    product_type varchar(50) not null,
    operaotors varchar(50) not null,
    payment_method varchar(50) not null,
    product_description varchar(50) not null
);
-- Creat table transaction
create table transaction (
    id int primary key auto_increment,
    user_id int not null,
    detailTransaction_id int not null,
    transaction_date date,
    total_price int not null,
    foreign key (user_id) references user(id),
    foreign key (detailTransaction_id) references transaction_detail(id)
);
-- creat table transaction_detail
create table transaction_detail (
    id int primary key auto_increment,
    item_id int not null,
    quantity int not null,
    foreign key (item_id) references item(id)
);

--3. Creat table kurir
create table kurir (
    id int primary key auto_increment,
    name varchar(50) not null,
    created_at datetime,
    updated_at datetime
);

--4. alter table kurir creat column ongkos_dasar
alter table kurir add ongkos_dasar int;


--5. alter table rename kurir to shipping
alter table kurir rename to shipping;

--6 drop table shipping
drop table shipping;

--7. 

-- one to one
create table payment_method(
    id int primary key auto_increment,
    name varchar(50) not null,
    description_id int not null,
    created_at datetime,
    updated_at datetime,
    foreign key (description_id) references description(id)
);

create table description (
    id int primary key auto_increment,
    description varchar(255) not null,
    created_at datetime,
    updated_at datetime
);



-- one to many table user and alamat

create table alamat (
    id int primary key auto_increment,
    user_id int not null,
    alamat varchar(255) not null,
    created_at datetime,
    updated_at datetime,
    foreign key (user_id) references user(id)
);

-- many to many table user and payment_method_detail
create table payment_method_detail (
    id int primary key auto_increment,
    user_id int not null,
    payment_method_id int not null,
    created_at datetime,
    updated_at datetime,
    foreign key (user_id) references user(id),
    foreign key (payment_method_id) references payment_method(id)
);
