INSERT INTO users
(id, "name", email, password_hash, is_admin, created_at, updated_at)
VALUES('891543f5-7b38-4d70-92f0-5ed4940e7c3b'::uuid, NULL, 'daniel@hi2.in', decode('24326124313424367255316B65754564756E454B5A4C45476941642F652E4D495053685472646653697066336F692E4E6D75516F4B6A424A4A6F646D','hex'), false, '2024-12-24 15:16:16.848', '2024-12-24 15:16:16.848');
INSERT INTO users
(id, "name", email, password_hash, is_admin, created_at, updated_at)
VALUES('e3376a06-bfa0-4a25-b346-a8b32d868f98'::uuid, NULL, 'daniel@mail.com', decode('243261243134247667722F597254584C316C39713145476B6A7448666573306A5A74655375703930724F5878513050334D6C536C4357327931742F61','hex'), true, '2024-12-24 15:34:14.620', '2024-12-24 15:34:14.620');


INSERT INTO products
(id, "name", description, price, stock, created_at, updated_at)
VALUES('e548526d-0f92-4784-a223-20866e9e14f4'::uuid, 'cera Ve', 'body lotion', 4.00, 100, '2024-12-24 20:03:00.725', '2024-12-24 20:03:00.725');
INSERT INTO products
(id, "name", description, price, stock, created_at, updated_at)
VALUES('f20b703f-3435-4f52-9e46-3744ee5cbe46'::uuid, 'cera Ve', 'retinol', 4.00, 100, '2024-12-24 20:03:09.460', '2024-12-24 20:03:09.460');


INSERT INTO orders
(id, user_id, status, total_amount, created_at, updated_at)
VALUES('7fed63c7-8826-425b-a8b2-84a5b584cdb9'::uuid, '891543f5-7b38-4d70-92f0-5ed4940e7c3b'::uuid, 'cancelled', 20.00, '2024-12-24 23:02:17.239', '2024-12-24 23:02:17.239');


INSERT INTO order_products
(order_id, product_id, quantity)
VALUES('7fed63c7-8826-425b-a8b2-84a5b584cdb9'::uuid, 'e548526d-0f92-4784-a223-20866e9e14f4'::uuid, 5);