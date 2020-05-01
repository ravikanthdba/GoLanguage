create database employee;

create user 'go'@'%' identified by 'employee';

grant all on employee.* to 'go'@'%';

use employee;

drop table if exists employee;

create table employee (
id integer primary key auto_increment,
empname varchar(30),
age integer default 20,
salary integer);

insert into employee (empname, age,salary) values ("Daniel",23, 55000), ("Arin",25, 65000), ("Juan",24, 72000), ("Shen",26, 64000),  ("Myke",27, 58000), ("McLeod",26, 72000), ("James",32, 35000);

drop table if exists phonenumbers;

create table phonenumbers (
phoneid integer primary key auto_increment,
empid integer references employee(id) ON DELETE CASCADE,
phonenumber integer unique key);

insert into phonenumbers (empid, phonenumber) values (1, 1234567), (1, 345672), (2, 345), (4, 456), (5, 123098), (3, 111), (3, 222), (3, 333), (6, 765), (6, 101), (7, 209);


drop table if exists placementlocation;

create table placementlocation (
locationid integer primary key auto_increment,
empid integer references employee(id) ON DELETE CASCADE,
location varchar(30));

insert into placementlocation  (empid, location) values (1, "Hyderabad"), (2, "Bangalore"), (3, "Chennai"), (4, "Delhi"), (5, "Mumbai"), (6, "Kolkata"), (7, "Pune");

drop table if exists department;

create table department (
departmentid integer primary key auto_increment,
empid integer references employee(id) ON DELETE CASCADE,
department varchar(30));

insert into department  (empid, department) values (1, "HR"), (2, "IT"), (3, "Transport"), (4, "Marketing"), (5, "Employee Health"), (6, "IT"), (7, "Finance");


select * from employee;
select * from phonenumbers;
select * from placementlocation;
select * from department;

