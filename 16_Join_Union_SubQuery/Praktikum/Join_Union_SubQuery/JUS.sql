-- ========= SOAL JOIN, SUBQUERY, UNION ===========

-- Soal 1
select *
from  transaction
where  user_id = 1
union 
select  *
from  transaction
where  user_id = 2;

-- Soal 2
select sum(total_price) as total_harga from transaction where user_id = 1;

-- Soal 3
select  sum(total_price) as total_price_ProductType_3
from transaction_details td
join transaction t on td.transaction_id = t.id
join  product p on td.product_id = p.id
where  p.product_type_id = 3;

-- Soal 4
select p.*, pt.name  from product p  join product_type pt on p.product_type_id = pt.id order by p.product_type_id  ;

-- Soal 5
select  transaction.*, product.name as Product_Name, users.name as User_Name
from transaction 
join users ON transaction.user_id = users.id 
join  transaction_details ON transaction.id = transaction_details.transaction_id 
join  product ON transaction_details.product_id = product.id;

-- Soal 6
DELIMITER $$
create trigger delete_transaction_details  
    after delete  
    on `transaction`
    for  each row  
BEGIN
   delete from transaction_details td where td.transaction_id  = old.id; 
END$$
DELIMITER ;

-- Soal 7
DELIMITER
$$
CREATE TRIGGER update_total_qty2
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE `transaction`
    SET total_qty = total_qty - OLD.qty
    WHERE id = OLD.transaction_id;
end
$$
DELIMITER ;

-- Soal 8

SELECT *
FROM product
WHERE id NOT IN (
    SELECT product_id
    FROM transaction_details
);

-- ========= AKHIR SOAL JOIN, SUBQUERY, UNION ==========