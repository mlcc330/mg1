基本操作
=
登录
-

    > mongodb
    
 查看和使用数据库
 -
 
    > show dbs
    > use mydb
    
 常用查询
 -
 
    - 查看所有数据
    > db.tablename.find()
    - 更新数据
    > db.tablename.update({name:"lgonline"},{$set:{age:33}})
    - 删除数据
    > db.tablename.remove({age:36})
    
 