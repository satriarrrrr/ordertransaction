USE store;

TRUNCATE TABLE  `products`;
INSERT INTO 
`products` (name, quantity, price)
VALUES
("Kaos Hitam", 15, 45000),
("Kaos Putih", 1, 30000),
("Celana Jeans Panjang", 5, 100000),
("Jaket Denim", 20, 250000),
("Celana Kargo", 6, 60000);

TRUNCATE TABLE  `discount_type`;
INSERT INTO 
`discount_type` (id, name)
VALUES
(1, "Diskon Nominal"),
(2, "Diskon Persentase");

TRUNCATE TABLE  `coupons`;
INSERT INTO 
`coupons` (code, discount_type, discount_nominal, max_used, valid_date_start, valid_date_end)
VALUES
("DISKON15", 2, 15, 10, "2017-12-01 00:00:00", "2018-01-01 00:00:00"),
("DISKON10", 2, 10, 5, "2017-12-01 00:00:00", "2018-05-01 00:00:00"),
("DISKON10K", 1, 10000, 5, "2017-12-01 00:00:00", "2018-05-01 00:00:00"),
("DISKON100K", 1, 100000, 2, "2017-12-01 00:00:00", "2018-01-01 00:00:00");

TRUNCATE TABLE  `order_status`;
INSERT INTO 
`order_status` (id, name, description)
VALUES
(1, "Active", "Active order, customer can add/update/remove item"),
(2, "Waiting Payment", "Order already submitted, customer can not add/update/remove item. Waiting for payment"),
(3, "Process", "Payment already received, order in process"),
(4, "Shipment", "Order in shipment"),
(5, "Success", "Products already delivered, order is success"),
(6, "Cancel", "Invalid order, order is canceled");

TRUNCATE TABLE  `payment_type`;
INSERT INTO 
`payment_type` (id, name)
VALUES
(1, "Bank Transfer");

TRUNCATE TABLE  `payment_status`;
INSERT INTO 
`payment_status` (id, name, description)
VALUES
(1, "Waiting", "Waiting for payment"),
(2, "Verification", "Transfer proof is already submitted, payment is waiting for verification"),
(3, "Rejected", "Invalid payment, payment rejected by admin. Order will be canceled."),
(4, "Paid", "Payment received and marked as valid by admin. Order is ready to be processed");

TRUNCATE TABLE  `shipping_operators`;
INSERT INTO 
`shipping_operators` (id, name)
VALUES
(1, "JNE"),
(2, "TIKI");

TRUNCATE TABLE  `shipping_status`;
INSERT INTO 
`shipping_status` (id, name, description)
VALUES
(1, "Shipped", "Process order shipment via logistic partner"),
(2, "Delivered", "Product delivered, order is success");