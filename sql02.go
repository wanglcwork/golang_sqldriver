// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

// -- 开始事务
START TRANSACTION;

// -- 假设账户 A 的 ID 为 1，账户 B 的 ID 为 2
SET @from_account_id = 1;
SET @to_account_id = 2;
SET @amount = 100;

// -- 检查账户 A 的余额是否足够
SELECT balance INTO @from_balance FROM accounts WHERE id = @from_account_id;

IF @from_balance >= @amount THEN
	// -- 扣除账户 A 的余额

	UPDATE accounts SET balance = balance - @amount WHERE id = @from_account_id;
	
	// -- 增加账户 B 的余额
	UPDATE accounts SET balance = balance + @amount WHERE id = @to_account_id;
	
	// -- 记录转账信息
	INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (@from_account_id, @to_account_id, @amount);
	
	// -- 提交事务
	COMMIT;
ELSE
	// -- 余额不足，回滚事务
	ROLLBACK;
END IF;

// -- 结束事务
// -- END;
// -- 注意：具体的事务控制语句可能因数据库系统而异，上述示例适用于 MySQL。

