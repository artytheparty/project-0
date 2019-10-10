create table users (
    usrID VARCHAR PRIMARY KEY,
    username varchar(80) unique not null,
    passwd VARCHAR not null,
    fName varchar not null,
    lName VARCHAR not NULL
);

insert into users values ('1','akhv','password','Artem','Khvan');
insert into users values ('2','jjohna','password','John','Johna');
insert into users values ('3','mmk','password','Michael','Moore');
insert into users values ('4','tcook','password','Tim','Cook');
insert into users values ('5','sjobs','password','Stev-o','Jobs');
insert into users values ('6','ckent','password','Clark','Kent');
insert into users values ('7','omg','password','Olly','Gunderson');
insert into users values ('8','imbored','password','Impending','Boredom');
insert into users values ('9','pparker','password','Peter','Parker');
insert into users values ('10','jonm','password','Jon','Martinez');

create table accounts (
    accNum VARCHAR PRIMARY KEY,
    accType VARCHAR,
    accBal FLOAT,
    usrID VARCHAR FOREIGN KEY REFERENCES users
);

insert into accounts values ('1', 'c', 1120.23, '1');
insert into accounts values ('2', 'c', 3000.23, '2');
insert into accounts values ('3', 'c', 4500.23, '3');
insert into accounts values ('4', 'c', 6500.23, '4');
insert into accounts values ('5', 'c', 300.23, '5');
insert into accounts values ('6', 'c', 530.23, '6');
insert into accounts values ('7', 'c', 9500.23, '7');
insert into accounts values ('8', 'c', 1325000.23, '8');
insert into accounts values ('9', 'c', 15000.23, '9');
insert into accounts values ('10', 'c', 54000.23, '10');
insert into accounts values ('11', 's', 5060.23, '1');
insert into accounts values ('12', 's', 20005040.23, '7');
insert into accounts values ('13', 's', 5000.23, '8');





