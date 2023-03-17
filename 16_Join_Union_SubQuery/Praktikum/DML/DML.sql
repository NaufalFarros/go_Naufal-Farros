
-- SOAL DML
-- 1 a insert into table operators
insert into operators (name, created_at, updated_at) values ("faros", "2023-03-15 22:03:47", "2023-03-15 22:03:47");
insert into operators (name, created_at, updated_at) values ("iqbal", "2023-03-15 22:03:47", "2023-03-15 22:03:47");
insert into operators (name, created_at, updated_at) values ("dayat", "2023-03-15 22:03:47", "2023-03-15 22:03:47");
insert into operators (name, created_at, updated_at) values ("firman", "2023-03-15 22:03:47", "2023-03-15 22:03:47");
insert into operators (name, created_at, updated_at) values ("tasya", "2023-03-15 22:03:47", "2023-03-15 22:03:47");

-- 1 b insert into table product type
insert into product_type (name, created_at, updated_at) values ("makanan","2023-03-15 22:09:00","2023-03-15 22:09:00");
insert into product_type (name, created_at, updated_at) values ("minuman","2023-03-15 22:09:00","2023-03-15 22:09:00");
insert into product_type (name, created_at, updated_at) values ("kosmetik","2023-03-15 22:09:00","2023-03-15 22:09:00");

-- 1 c insert into table product
insert into product (product_type_id, operators_id,name) values ("1","3","Kentang Goreng");
insert into product (product_type_id, operators_id,name) values ("1","3","SilverQuuen");

-- 1 d insert into product 
insert into product (product_type_id, operators_id,name) values ("2","1","Sprite");
insert into product (product_type_id, operators_id,name) values ("2","1","Kopi ABC");
insert into product (product_type_id, operators_id,name) values ("2","1","C 1000");

-- 1 e insert into product
insert into product (product_type_id, operators_id,name) values ("3","4","Lipstik");
insert into product (product_type_id, operators_id,name) values ("3","4","Bedak");
insert into product (product_type_id, operators_id,name) values ("3","4","Cream");

-- 1 f product description
insert into product_descriptions (product_id, description) values ("1","kentang goreng nikmat dimakan sama nasi");
insert into product_descriptions (product_id, description) values ("2","SilverQueen nikmat sekali coklatnya lumer");
insert into product_descriptions (product_id, description) values ("3","Sprite segar sekali");
insert into product_descriptions (product_id, description) values ("8","Kopi ABC kopi asli");
insert into product_descriptions (product_id, description) values ("9","Vitamin C di C1000");
insert into product_descriptions (product_id, description) values ("10","Lipstik digunakan untuk mempercantik");
insert into product_descriptions (product_id, description) values ("13","Bedak Asli tanpa Mercuri");
insert into product_descriptions (product_id, description) values ("14","Cream kecantikan nomor satu");

-- 1 g payment method
insert  into payment_methods (name) values ("Bank BCA");
insert  into payment_methods (name) values ("Bank BRI");
insert  into payment_methods (name) values ("Bank Tunai");

--1 h users

insert into users (name,gender, created_at) values ("Faros","M", "2023-03-15 00:00:00");
insert into users (name,gender , created_at) values ("Dayat","M","2023-03-14 00:00:00");
insert into users (name,gender, created_at) values ("Iqbal","M", "2023-03-12 00:00:00");
insert into users (name,gender, created_at) values ("Tasya","F", "2023-03-10 00:00:00");
insert into users (name,gender , created_at) values ("Shinta","F","2023-03-9 00:00:00");

-- 1 i transaction
insert into transaction (user_id, payment_method_id, total_qty,total_price) values ("1","1","3","75000");
insert into transaction (user_id, payment_method_id, total_qty,total_price) values ("2","2","2","100000");
insert into transaction (user_id, payment_method_id, total_qty,total_price) values ("3","3","5","50000");

-- 1 j transaction detail
insert into transaction_details (transaction_id, product_id , qty, price) values ("1","1","3","25000");
insert into transaction_details (transaction_id, product_id , qty, price) values ("2","14","2","50000");
insert into transaction_details (transaction_id, product_id , qty, price) values ("3","3","5","10000");

-- Soal 2 a
select name from users u where u.gender = "M";

-- Soal 2 b
select * from product p where id = 3;

-- Soal 2 c
select  *
from  users
where  created_at between  DATE_SUB(NOW(), interval  7 day) and  NOW()
and  name like  '%a%';

-- Soal 2 d 
select count(gender) as jumlah_female  from users u where gender = 'F' ;

-- Soal 2 e
select name from users order by name asc ;

-- Soal 2 f
select * from  product p limit 5;

-- Soal 3 a
update product set name = 'product dummy' where id = 1;

-- Soal 3 b
update transaction_details set qty = '3' where product_id = 1;

-- Soal 4 a
delete from product where id = 1;

-- Soal 4 b
delete from product  where product_type_id = 1;

-- ========= AKHIR SOAL DML ===========
