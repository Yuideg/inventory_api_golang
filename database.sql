--///  USE LATEST VERSION POSTGRESQL  NOW IAM USING POSGRESQL 13 CURRENT VERSION/////
--//BEFORE EXECUTING BELOW POSTGRESQL STATEMENTS FISRT YOU NEED TO ADD TWO EXTENSIONS TO YOU DATABASE ///
-- THE TWO EXTENSION USED FOR UUID AND ISSN SPECIAL DATATYPE FOR POSTGRESQL DATABASE 
--THESE TWO TYPES ARE NOT SUPPORTED BY DEFUAT ,SO YOU HAVE TO ADD MANUALY BY YOURSELF///
 --///////////////////COMMANDS//////
--   CREATE EXTENSION ISN;
--   CREATE EXTENSION IF NOT EXISTS "uuid-ossp";



create table role (
id  uuid  DEFAULT uuid_generate_v4 () ,
name varchar(20) not null unique,
PRIMARY KEY(id)
);
create table user_tb (
    id  uuid  DEFAULT uuid_generate_v4 () ,
    username  varchar(50) not null unique ,
    password varchar(50) not null ,
    first_name VARCHAR(255) not null ,
    last_name VARCHAR(255) not null,
    gender varchar(255) not null ,
    email varchar(255) not null unique ,
    country varchar(255) not null ,
    city  varchar(123),
    state varchar(255) ,
    zipcode int ,
    street  varchar(255),
    latitude numeric (10,6),
    longitude numeric (10,4),
    phone varchar(128) not null ,
    role_id uuid not null,
    CONSTRAINT fk_role_id FOREIGN KEY(role_id) REFERENCES role(id) on update cascade  on delete cascade ,
    created_on    timestamp not null DEFAULT CURRENT_TIMESTAMP,
    updated_on timestamp not null DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
create table supplier(
    id  uuid  DEFAULT uuid_generate_v4 () ,
    supplier_name varchar(70) not null,
    country varchar(255) not null ,
    city  varchar(123),
    state varchar(255) ,
    zipcode int ,
    street  varchar(255),
    latitude numeric (10,6),
    longitude numeric (10,4),
    fax varchar(45),
    po_box  varchar(6),
    email  varchar(255) unique ,
    website  varchar(255),
    phone  varchar (50) unique ,
    created_on timestamp not null DEFAULT CURRENT_TIMESTAMP,
    updated_on timestamp not null DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

create table stock (
    id  uuid  DEFAULT uuid_generate_v4 () ,
    products_common_code varchar(70) not null unique not null,
    products_common_name varchar(70) not null unique  not null,
    number_of_products  bigint not null,
    supplier_id uuid not null,
    country varchar(255) not null ,
    city  varchar (123),
    state varchar(255) ,
    zipcode int ,
    street  varchar(255),
    latitude numeric (10,6),
    longitude numeric (10,4),
    CONSTRAINT fk_supplier_id FOREIGN KEY(supplier_id)
    REFERENCES supplier(id) on update cascade  on delete cascade ,
    created_on timestamp not null DEFAULT CURRENT_TIMESTAMP,
    updated_on TIMESTAMP not null DEFAULT  CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

create table product (
    id  uuid  DEFAULT uuid_generate_v4 (),
    barcode ISSN not null unique ,
    product_name varchar(100) not null,
    stock_id uuid not null,
    brand varchar(255) not null,
    image    varchar(255) not null,
    cost  decimal(10,3) not null,
    price  decimal(10,3),
    unit    varchar(10),
    tax_rate numeric (4,2),
    discount   numeric(3,2) default 0,
    mfg_date  timestamp not null  ,
    expired_on  timestamp not null  ,
    CONSTRAINT fk_stock_id FOREIGN KEY(stock_id)
    REFERENCES stock(id) on update cascade  on delete cascade ,
    created_on timestamp not null DEFAULT CURRENT_TIMESTAMP,
    updated_on TIMESTAMP not null DEFAULT  CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
create table orders(
    id  uuid  DEFAULT uuid_generate_v4 () ,
    quantity numeric(10,3) not null,
    unit    varchar(10),
    confirm_by uuid  not null, CONSTRAINT fk_confirm_by FOREIGN KEY(confirm_by)
    REFERENCES user_tb(id) on update cascade  on delete cascade ,
    user_id uuid  not null ,CONSTRAINT fk_customer_id FOREIGN KEY(user_id)
    REFERENCES user_tb(id) on update cascade  on delete cascade ,
    product_id uuid  not null, CONSTRAINT fk_product_id FOREIGN KEY(product_id)
    REFERENCES product(id) on update cascade  on delete cascade,
    status varchar(16) not null,
    expired_on timestamp not null,
    created_on timestamp not null DEFAULT CURRENT_TIMESTAMP,
    updated_on TIMESTAMP not null DEFAULT  CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- delete from orders where id=1;
-- ALTER TABLE product ALTER COLUMN tax_rate TYPE numeric (4,2);
-- ALTER TABLE product ADD COLUMN expired_at TIMESTAMP;
-- ALTER TABLE user ADD CONSTRAINT user_username UNIQUE (username);
-- ALTER TABLE user ADD CONSTRAINT user_email UNIQUE (hashed_email);
-- CREATE INDEX user_username_idx ON user (username, hashed_email,"password");
--
-- ALTER TABLE product ADD CONSTRAINT profile_id UNIQUE (id);
-- ALTER TABLE "product" ADD CONSTRAINT profile_username UNIQUE (username);
-- CREATE INDEX profile_username_idx ON "profiles" (username);
-- DROP TABLE IF exists users;
-- ALTER TABLE supplier ALTER COLUMN fax TYPE VARCHAR(50);
-- postgresql 13
-- ALTER TABLE supplier ALTER COLUMN phone TYPE VARCHAR(50);