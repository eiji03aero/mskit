DROP TABLE IF EXISTS restaurants;
CREATE TABLE restaurants (
  id VARCHAR NOT NULL,
  name VARCHAR NOT NULL
);

DROP TABLE IF EXISTS restaurant_menu_items;
CREATE TABLE restaurant_menu_items (
  id VARCHAR NOT NULL,
  restaurant_id VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  price int NOT NULL
);
