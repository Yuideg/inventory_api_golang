-- INSERT ROLES
INSERT INTO role(name)VALUES('ADMIN'),('USER');
-- 	   INSERT CUSTOMER
INSERT INTO customer(username,password,first_name,last_name,gender,email,country,city,state,zipcode,street,latitude,longitude,phone,role_id)
VALUES('john32','90901234typo','Yohannes','Ketema','Male','yohannesketem@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'algeria street',12.564,3.677,'+251 91 093 7962','8935fc25-68ab-4284-9424-374036f71d08'),
      ('alefew2121','33097322','Alewfew','Yimer','Male','alefewyimer2121@gmail.com','Ethiopia','Adama','Oromo',4590,'abdisa street',34.9022,26.009,'+251 90 683 7962','8935fc25-68ab-4284-9424-374036f71d08'),
       ('get04','33609087','Getahun','Honelet','Male','getme04@gmail.com','Ethiopia','Bahir Dar','Amahara',5678,'theodrosII street',11.2202,24.0034,'+251 94 992 2604','8935fc25-68ab-4284-9424-374036f71d08'),
      ('Hana12','h456788','Hana','Yosef','Female','hanayosef@gmail.com','Ethiopia','Hawassa','Sidama',2345,'main street',1.2202,14.9034,'+251 94 334 2104','8935fc25-68ab-4284-9424-374036f71d08');

-- INSERT SUPPLIER
INSERT INTO supplier(supplier_name,country,city,state,zipcode,street,latitude,longitude,fax,po_box,email,website,phone)
VALUES('KERCHANSHE TRADING P.L.C','Ethiopia','Addis Ababa','Addis Ababa',1000,'South Africa Street',12.33,34.66,'234','453','kerchanshe@gmail.com','https://www.kerchanshe.com/','+251 11 371 6370'),
     ('Oromia Coffee Farmers Cooperative','Ethiopia','Adama','oromia',1222,'algeria street',11,33.66,'234','390','oromiacoffeeunion@gmail.com','https://www.oromiacoffeeunion.org/','+251 11 391 9070'),
     ('Sidama Coffee Farmers Cooperative','Ethiopia','Hawassa','sidama',2098,'Africa Street ',7.33,3.66,'407166','122062','sidama23coffee@gmail.com','','+251 11 440 7166'),
     ('Kata Muduga Multipurpose','Ethiopia','Addis Ababa','Addis Ababa',1000,'Bole street',9.33,4.6,'23300122','909877','mullerkata@gmail.com','https://www.mullerkata.com/','+251 11 786 0987'),
     ('Daye Bensa Coffee Exp. Plc','Ethiopia','Arbaminch','Souther ethiopia',4357,'kefa street',6.33,7.06,'0078655','13454','dayebensacoffee@gmail.com','http://www.dayebensacoffee.com/','+251 11 098 1267');


-- INSERT STOCK
INSERT INTO stock(products_code,products_name, number_of_products,supplier_id,country,city,state,zipcode,street,latitude,longitude)
VALUES('oro-112','coffees',3400,'8e7d4412-c146-4e1d-ab4b-3514c13ff7f0','Ethiopia','Addis Ababa','Addis Ababa',1000,'South Africa Street',12.33,34.66),
('sidama-092','coffees',300,'5f6496b1-a2cf-4438-a42c-0fc308ae3763','Ethiopia','Adama','oromia',1222,'algeria street',11,33.66),
('ker-001','plastics',1200,'5f6496b1-a2cf-4438-a42c-0fc308ae3763','Ethiopia','Hawassa','sidama',2098,'Africa Street ',7.33,3.66),
('kem-mul-11','Fruits',4000,'a6803e39-c01e-467b-b1ff-682cd980efa3','Ethiopia','Addis Ababa','Addis Ababa',1000,'Bole street',9.33,4.6),
('dbcce-111','coffees',2000,'b1d9f58f-1dc7-4783-b0f7-030f4dae3e6c','Ethiopia','Arbaminch','Souther ethiopia',4357,'kefa street',6.33,7.06);

-- INSERT PRODUCT
INSERT INTO product(barcode, product_name,stock_id,brand,image,cost,price,unit,tax_rate,discount,mfg_date,expired_on)
VALUES('0317-8471','orange','8264313e-4f7b-48c2-93e7-1895c11f1f98','normal','',40.87,60.23,'kg',10.0,0,'2021-04-30','2021-06-03'),
      ('1050-124X','coffee','69cd2c2d-2fd1-4e7e-9ec5-ca424c51a3b8','sidama','',100.10,120.27,'kg',20.5,0,'2021-01-11','2021-04-30'),
      ('3691-2689','plate','8264313e-4f7b-48c2-93e7-1895c11f1f98','new','',20.00,30.55,'gram',3.0,0,'2021-04-30','2022-04-30'),
      ('1991-8941','coffee','2a0b04a4-60c7-4037-9c94-1d0751aaee1d','kefa','',120.00,160.34,'kg',23.8,0,'2021-04-30','2022-04-30');


-- INSERT STAFFS
INSERT INTO staff(username,password,first_name,last_name,gender,email,country,city,state,zipcode,street,latitude,longitude,phone,role_id,salary)
VALUES('kalkidan34','235677','Kalkidan','Abebe','Female','kalabe32@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'algeria street',12.564,3.677,'+251 97 836 0030','3d9948a7-e151-448b-a7b0-1c3a432a876e',9000),
('kefale','9045334','Kefale','Aytegeb','Male','kefaleaytegeb99@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'adisu gebeya street',12.564,3.677,'+251 93 881 0448','3d9948a7-e151-448b-a7b0-1c3a432a876e',8922),
('yaregal','235677','Yaregal','Seyfe','Male','yaregalseyfe@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'piassa street',12.564,3.677,'+251 94 836 0030','3d9948a7-e151-448b-a7b0-1c3a432a876e',4500),
 ('yuidem','2378yuidegm','Yuideg','Misganew','Male','misganewendeg879@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'russia street',12.564,3.677,'+251 93 457 8879','3d9948a7-e151-448b-a7b0-1c3a432a876e',10223),
('yabsira','Yoni2345','Yonithan','Abyu','Male','yonithanabyu@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'6killo street',12.564,3.677,'+251 98 890 0030','3d9948a7-e151-448b-a7b0-1c3a432a876e',23000);

-- INSERT ORDERS
INSERT INTO orders(quantity,unit,confirm_by,customer_id,product_id,status,expired_on)
VALUES(2,'kg','0364b9f5-e151-4812-a3e6-2e4f19c73e67','4816fd9a-cce7-4068-b893-88a317c431eb','03f63881-9601-4788-9fda-752d815853a7','unpaid','2021-07-30'),
(9,'kg','6b8009cf-94ce-4c62-8fc7-486e559be497','483d9830-ebc6-4e27-863f-ad6d8fa8e68f','d5f29826-9d3b-4115-bb18-e6d4afbd4351','unpaid','2021-08-30'),
(34,'kg','d45d2a6c-197a-4109-a00e-23ce9fafa14d','f0e3dd10-4b56-4ff5-af24-ae06eabcb1d4','c70b8b01-f110-4792-a9e7-0b74ec9610c9','paid','2021-06-14'),
(11,'kg','fed84964-6678-461b-990a-acd5fc2257a0','e23184eb-941d-463d-843d-461510a781ef','c70b8b01-f110-4792-a9e7-0b74ec9610c9','unpaid','2021-07-10'),
(12,'gram','b8e9a041-4249-40ec-8491-d5f039ac2d28','483d9830-ebc6-4e27-863f-ad6d8fa8e68f','c70b8b01-f110-4792-a9e7-0b74ec9610c9','unpaid','2021-06-23'),
(21,'kg','d45d2a6c-197a-4109-a00e-23ce9fafa14d','4816fd9a-cce7-4068-b893-88a317c431eb','ba616562-0207-45e1-bf61-02c4ef8afc58','paid','2021-08-03');















