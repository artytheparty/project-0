create table users (
    usrID VARCHAR PRIMARY KEY,
    username varchar unique,
    passwd VARCHAR,
    fName varchar,
    lName VARCHAR
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
    accNum VARCHAR,
    usrID VARCHAR REFERENCES users(usrID),
    accType VARCHAR,
    accBal FLOAT,
    PRIMARY KEY (accNum,usrID)
);

insert into accounts values ('1', '1', 'c', 1120.23);
insert into accounts values ('2', '2', 'c', 3000.23);
insert into accounts values ('3', '3', 'c', 4500.23);
insert into accounts values ('4', '4', 'c', 6500.23);
insert into accounts values ('5', '5', 'c', 300.23);
insert into accounts values ('6', '6', 'c', 530.23);
insert into accounts values ('7', '7', 'c', 9500.23);
insert into accounts values ('8', '8', 'c', 1320.23);
insert into accounts values ('9', '9', 'c', 15000.23);
insert into accounts values ('10', '10', 'c', 54000.23);
insert into accounts values ('11', '1', 's', 5060.23);
insert into accounts values ('12', '7', 's', 5040.23);
insert into accounts values ('13', '3', 's', 5000.23);