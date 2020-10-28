create database company;
use company;

create table emp (
emp_id int not null,name varchar(20) not null,department varchar(20) not null,primary key (emp_id)
);

insert into emp (emp_id, name, department) values (1, 'a', 'go');
insert into emp (emp_id, name, department) values (2, 'b', 'fe');
insert into emp (emp_id, name, department) values (3, 'c', 'sql');
insert into emp (emp_id, name, department) values (4, 'd', 'sqlH');
insert into emp (emp_id, name, department) values (5, 'e', 'go');

update emp set department='saleforce' where emp_id=3;
update emp set name='server' where emp_id=5;
select * from emp;

delete from emp where emp_id=5;
