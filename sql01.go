// 题目1：基本CRUD操作
//  假设有一个名为students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

// 插入一条新记录
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');	

// 查询所有年龄大于 18 岁的学生信息
SELECT * FROM students WHERE age > 18;

// 更新姓名为 "张三" 的学生年级
UPDATE students SET grade = '四年级' WHERE name = '张三';

// 删除年龄小于 15 岁的学生记录
DELETE FROM students WHERE age < 15;






