CREATE TABLE person
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name varchar(30),
    last_name  varchar(40)
);