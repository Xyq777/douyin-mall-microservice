CREATE DATABASE orders;
CREATE DATABASE auth;
CREATE DATABASE cart;
CREATE DATABASE checkout;
CREATE DATABASE payment;
CREATE DATABASE product;
CREATE DATABASE user;

INSERT INTO `category` (
    `category_name`,
    `description`,
    `created_at`,
    `updated_at`
) VALUES ('sticker',
          'miku',
          NOW(),
          NOW()
         );
-- 插入一条产品数据（假设自增 id 为 1）
INSERT INTO `product` (
  `name`,
  `description`,
  `picture`,
  `price`,
  `created_at`,
  `updated_at`
) VALUES (
  'miku',
  'a miku',
  'https://ssn7gazd0.hn-bkt.clouddn.com/miku1.jpeg',
  5999.99,
  NOW(),
  NOW()
);
-- 关联分类 id=1 和产品 id=1
INSERT INTO `product_category` (
    `product_id`,
    `category_id`
) VALUES (
             1,
             1
         );
