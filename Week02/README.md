学习笔记

作业

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

ErrNoRows:
    GET http://127.0.0.1:8080/api1/item/0

other errors:
    GET http://127.0.0.1:8080/api1/item/1

success:
    GET http://127.0.0.1:8080/api1/item/2