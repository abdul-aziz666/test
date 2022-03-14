-- for design DB please check this link 
-- https://dbdiagram.io/d/622ea3f361d06e6eadf46c83

CREATE TYPE "trans_status" AS ENUM (
    'PENDING',
    'SUCCESS',
    'FAILED'
);

CREATE TABLE "transaction" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "order_date" timestamp NOT NULL,
    "status" trans_status NOT NULL DEFAULT 'PENDING',
    "payment_date" timestamp,
    "user_id" int,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp
);

CREATE TABLE "transaction_detail" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "transaction_id" int,
    "price" int,
    "pcs" int,
    "sub_total" int,
    "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "transaction_detail" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction" ("id");
CREATE INDEX "idx_transaction_user_id" ON "transaction" ("user_id");
CREATE INDEX "idx_transaction_status" ON "transaction" ("status");
CREATE INDEX "idx_transaction_detail_transaction_id" ON "transaction_detail" ("transaction_id");


-- for dummy data 
INSERT INTO public."transaction" (order_date,status,payment_date,user_id,created_at,updated_at) VALUES
		('2022-03-14 09:26:43.288677','PENDING',NULL,100,'2022-03-14 09:26:43.288677',NULL),
		('2022-03-14 09:26:43.288677','SUCCESS','2022-03-14 09:27:03.1176',100,'2022-03-14 09:26:43.288677','2022-03-14 09:26:43.288');

INSERT INTO public.transaction_detail (transaction_id,price,pcs,sub_total,created_at) VALUES
		(1,10000,2,20000,'2022-03-14 09:27:35.482082'),
		(1,5000,1,5000,'2022-03-14 09:27:35.482082'),
		(2,10000,5,50000,'2022-03-14 09:27:35.482082'),
		(2,50000,2,100000,'2022-03-14 09:27:35.482082');


-- for query please check this link 
SELECT t.order_date, t.status, t.payment_date, sum(td.sub_total), sum(td.pcs)
FROM "transaction" t 
JOIN transaction_detail td ON t.id = td.transaction_id 
WHERE t.user_id = 100
GROUP  BY t.id