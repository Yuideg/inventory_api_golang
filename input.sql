-- INSERT ROLES
INSERT INTO role(name)VALUES('ADMIN'),('USER');
-- 	   INSERT CUSTOMER
INSERT INTO user_tb(username,password,first_name,last_name,gender,email,country,city,state,zipcode,street,latitude,longitude,phone,role_id)
VALUES(
'john32','90901234typo','Yohannes','Ketema','Male','yohannesketem@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'algeria street',12.564,3.677,'+251 91 093 7962','064a3d9f-5032-48d2-9645-b488b282ab96'),
('alefew2121','33097322','Alewfew','Yimer','Male','alefewyimer2121@gmail.com','Ethiopia','Adama','Oromo',4590,'abdisa street',34.9022,26.009,'+251 90 683 7962','064a3d9f-5032-48d2-9645-b488b282ab96'),
('get04','33609087','Getahun','Honelet','Male','getme04@gmail.com','Ethiopia','Bahir Dar','Amahara',5678,'theodrosII street',11.2202,24.0034,'+251 94 992 2604','ee280bee-42d9-456b-88f3-3bfca1022fe3'),
('Hana12','h456788','Hana','Yosef','Female','hanayosef@gmail.com','Ethiopia','Hawassa','Sidama',2345,'main street',1.2202,14.9034,'+251 94 334 2104','064a3d9f-5032-48d2-9645-b488b282ab96'),
('kalkidan34','235677','Kalkidan','Abebe','Female','kalabe32@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'algeria street',12.564,3.677,'+251 97 836 0030','064a3d9f-5032-48d2-9645-b488b282ab96'),
('kefale','9045334','Kefale','Aytegeb','Male','kefaleaytegeb99@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'adisu gebeya street',12.564,3.677,'+251 93 881 0448','ee280bee-42d9-456b-88f3-3bfca1022fe3'),
('yaregal','235677','Yaregal','Seyfe','Male','yaregalseyfe@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'piassa street',12.564,3.677,'+251 94 836 0030','064a3d9f-5032-48d2-9645-b488b282ab96'),
 ('yuidem','2378yuidegm','Yuideg','Misganew','Male','misganewendeg879@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'russia street',12.564,3.677,'+251 93 457 8879','064a3d9f-5032-48d2-9645-b488b282ab96'),
('yabsira','Yoni2345','Yonithan','Abyu','Male','yonithanabyu@gmail.com','Ethiopia','Addis Ababa','Addis Ababa',1000,'6killo street',12.564,3.677,'+251 98 890 0030','064a3d9f-5032-48d2-9645-b488b282ab96');

-- INSERT SUPPLIER
INSERT INTO supplier(supplier_name,country,city,state,zipcode,street,latitude,longitude,fax,po_box,email,website,phone)
VALUES('KERCHANSHE TRADING P.L.C','Ethiopia','Addis Ababa','Addis Ababa',1000,'South Africa Street',12.33,34.66,'234','453','kerchanshe@gmail.com','https://www.kerchanshe.com/','+251 11 371 6370'),
     ('Oromia Coffee Farmers Cooperative','Ethiopia','Adama','oromia',1222,'algeria street',11,33.66,'234','390','oromiacoffeeunion@gmail.com','https://www.oromiacoffeeunion.org/','+251 11 391 9070'),
     ('Sidama Coffee Farmers Cooperative','Ethiopia','Hawassa','sidama',2098,'Africa Street ',7.33,3.66,'407166','122062','sidama23coffee@gmail.com','','+251 11 440 7166'),
     ('Kata Muduga Multipurpose','Ethiopia','Addis Ababa','Addis Ababa',1000,'Bole street',9.33,4.6,'23300122','909877','mullerkata@gmail.com','https://www.mullerkata.com/','+251 11 786 0987'),
     ('Daye Bensa Coffee Exp. Plc','Ethiopia','Arbaminch','Souther ethiopia',4357,'kefa street',6.33,7.06,'0078655','13454','dayebensacoffee@gmail.com','http://www.dayebensacoffee.com/','+251 11 098 1267');



-- INSERT STOCK
INSERT INTO stock(products_common_code,products_common_name, number_of_products,supplier_id,country,city,state,zipcode,street,latitude,longitude)
VALUES('oro-112','coffees',3400,'4685585e-3ff8-4913-9132-0a9ac070990a','Ethiopia','Addis Ababa','Addis Ababa',1000,'South Africa Street',12.33,34.66),
('sidama-092','sida_coffees',300,'8ace0fad-8fb1-47ca-a464-cb4a26179c95','Ethiopia','Adama','oromia',1222,'algeria street',11,33.66),
('ker-001','plastics',1200,'43ec8ba7-8b74-432e-9d82-9f51c3816053','Ethiopia','Hawassa','sidama',2098,'Africa Street ',7.33,3.66),
('kem-mul-11','Fruits',4000,'a54159a5-2cf0-47cf-bd11-4acbbb2e2c27','Ethiopia','Addis Ababa','Addis Ababa',1000,'Bole street',9.33,4.6),
('itecce-111','kefa_coffee',2000,'861b3e01-a3f2-44fb-a093-c21c159ad20e','Ethiopia','Arbaminch','Southern ethiopia',4357,'kefa street',6.33,7.06);





-- INSERT PRODUCT
INSERT INTO product(barcode, product_name,stock_id,brand,image,cost,price,unit,tax_rate,discount,mfg_date,expired_on)
VALUES('0317-8471','orange','6ac9b0d9-4f2d-4b81-84e0-852515aa7b4f','normal','',40.87,60.23,'kg',10.0,0,'2021-04-30','2021-06-03'),
      ('1050-124X','coffee_oro','c3babacf-e8bd-4e83-b9d1-026399fef8ca','oro-cc12','',100.10,120.27,'kg',20.5,0,'2021-01-11','2021-04-30'),
      ('3691-2689','plate','1b6a2b7e-5c19-4823-83da-c8fef4502535','new','',20.00,30.55,'gram',3.0,0,'2021-04-30','2022-04-30'),
      ('1991-8941','coffee_sidama','1b6a2b7e-5c19-4823-83da-c8fef4502535','kefa','',120.00,160.34,'kg',23.8,0,'2021-04-30','2022-04-30');


-- INSERT ORDERS
INSERT INTO orders(quantity,unit,confirm_by,user_id,product_id,status,expired_on)
VALUES(2,'kg','bde1da30-8da9-4e3d-ac55-11690dde3125','96bb196d-6466-4575-bd62-50b1ad03198c','dea3ac6d-e1eb-4ff2-9e02-ab880eaeba3b','unpaid','2021-07-30'),
(9,'kg','8eb5d584-13c4-4b40-a1ec-416fac2dd523','47574f62-b0ee-4ca8-8ea0-0de6610b2620','cda8c6aa-452c-47e1-b8cf-012a493b3efb','unpaid','2021-08-30'),
(34,'kg','8eb5d584-13c4-4b40-a1ec-416fac2dd523','47574f62-b0ee-4ca8-8ea0-0de6610b2620','4ee355e9-0895-48ca-a4ad-ee3570508047','paid','2021-06-14'),
(11,'kg','d9393a48-b2dc-4e48-8a0c-5e2c2ca76c1a','96bb196d-6466-4575-bd62-50b1ad03198c','0c57f52a-8678-4346-8590-e8cb5a13a24e','unpaid','2021-07-10'),
(12,'gram','d9393a48-b2dc-4e48-8a0c-5e2c2ca76c1a','47574f62-b0ee-4ca8-8ea0-0de6610b2620','0c57f52a-8678-4346-8590-e8cb5a13a24e','unpaid','2021-06-23'),
(21,'kg','13d3d7ee-63ab-4e50-aebb-bb349b82871a','96bb196d-6466-4575-bd62-50b1ad03198c','cda8c6aa-452c-47e1-b8cf-012a493b3efb','paid','2021-08-03');















