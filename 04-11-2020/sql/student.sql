use Student;


create table stumarks (
rollno int not null,
marks int not null,
class varchar(2),
primary key(rollno));

insert into stumarks (rollno, marks, result) values (1, name1, 12th);
insert into stumarks (rollno, marks, result) values (2, name2, 10th);
insert into stumarks (rollno, marks,  result) values (3, name3,12th );
insert into stumarks (rollno, marks, result) values (4, name4, 10th);
insert into stumarks (rollno, marks, result) values (5, name5, 5th);

SELECT studet.name, stumarks.result
FROM stumark
INNER JOIN stumarks ON studet.rollno = stumarks.rollno;
SELECT stumarks.rollno, studet.name, studet.dateofjoining
FROM stumarks
RIGHT JOIN studet ON studet.rollno = stumarks.rollno
ORDER BY studet.rollno;
SELECT studet.name, stumarks.rollno
FROM studet
LEFT JOIN stumarks ON studet.rollno = stumarks.rollno
ORDER BY studet.name;
SELECT studet.name, stumarks.rollno
FROM studet
FULL OUTER JOIN stumarks ON studet.rollno=stumarks.rollno
ORDER BY studet.name;
SELECT studet.name AS Name, B.CustomerName AS CustomerName2, A.City
FROM Customers A, Customers B
WHERE studet.rollno = stumarks.rollno
AND studet.name = stumarks.name
ORDER BY studet.name;