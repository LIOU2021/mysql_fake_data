use test_db;

select count(*) from users ;

select * from users limit 10;

explain select id, name from users limit 10; # 没中主key
explain select id from users limit 10; # 中主key

explain select * from users where status = 1 limit 10; # 没中
explain select name from users where status = 1 limit 10; # 中
explain select name from users where status = 1 and address = "ffff" limit 10; # 没中
explain select name from users where name = "abc" and address = "ffff" limit 10; # 中
explain select name from users where address = "ffff" and name = "abc" limit 10; # 中
explain select * from users where name = "Esta Kiehn" limit 10; # 中
explain select * from users where name = "Esta Kiehn" and status = 1 limit 10; # 中
explain select * from users where status = 1 and name = "Esta Kiehn" limit 10; # 中
explain select * from users where name = "Esta Kiehn" and status = 1 and address = "fff" limit 10; # 中
explain select * from users where name = "Esta Kiehn" and address = "fff" limit 10; # 中
explain select * from users where status = 1 and address = "fff" limit 10; # 没中
