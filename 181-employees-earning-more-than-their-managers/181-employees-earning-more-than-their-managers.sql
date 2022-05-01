# Write your MySQL query statement below
select a.name as Employee from Employee a, Employee b where a.salary > (select b.salary where b.id = a.managerId);