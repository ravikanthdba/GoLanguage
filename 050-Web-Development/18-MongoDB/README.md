Databases are of 2 types:

1) SQL
2) NOSQL

SQL:
    - has a structure
    - It has a schema
    - It's not scalable horizontally
    - Example: Oracle, MySQL etc
    - Consistancy and Available
    
    
NoSQL:
    - Doesn't have a structure
    - Partitioned and Available
    - Consistency takes some heat, which means when we insert data into the DB, it takes a bit of time to update, hence consistency is impacted.
    - Horizontally scalable
    - Example: Redis, CouchBase, Mongo DB etc
    
    4 types of databases:
    
    - KV database: DB stores only Key/Value pairs. Example: Redis
    - Document Store: stores documents. documents are Complex JSON structures, which has nested structures in it.
    - Columnar store Databases: SQL databases stores data in Row, but in this case, it is inverted. It's stored in columns. Example: Google Big Table, Cassandra, HBASE
    - Graph Databases: interconnected data with relationships
     

RDBMS <- collection of databases <- collection of tables <- collection of records <- collection of strings
Mongo DB <- collection of databases <- collection of documents <- collection of fields <- collection of embeds <- _id is the default key


Commands of Mongo DB:

    - show dbs: It gives the list of all the databases
    - use <db>: If the database doesn't exist, this commands creates the database
    - use <db>: If the database exists, it will enter that databases.
    - show collections: Collections are equivalent to tables in RDBMS. This command lists the collections in the database.
    - db.createCollection('<collectionName>', <collection properties list>) - creates a collection with required properties.
    - db.<collection>.insert([<document1>, <document2>, <document3>]) - If collection doesn't exists, it creates a collection and inserts documents into the collection. documents are equivalent to rows in RDBMS
    - db.<collection>.find() - select data from the collection
    - db.<collection>.find().query() - selects data from the collection and returns in pretty format
    - db.<collection>.drop() - drops the collection
    - db.dropDatabase() - drops the database
    - db.<collection>.findOne() - selects the first record
    - db.<collection>.find({key: "value"}) - selects the record where key matches the value
    - db.<collection>.findOne({key: "value"}) - if there are multiple values that match the key, then findOne gives the first record of that collection
    - db.<collection>.find({$and: [{key1: "value1"}, {key2: "value2"}]}) - selects data from collection where key1 = value1 and key2 = value2
    - Example2: db.<collection>.find({$and: [{key1, "value1"}, {key2: {$lt: value2}}]}) - selects data from collection where key1 = value1 and key2 < value2
    - Example3: db.<collection>.find({$and: [{key1, "value1"}, {key2: {$gt: value2}}]}) - selects data from collection where key1 = value1 and key2 > value2
    - db.<collection>.find({$or: [{key1: "value1"}, {key2: "value2"}]}) -  selects data from collection where key1 = value1 or key2 = value2
    - Example2: db.<collection>.find({$or: [{key1: "value1"}, {key2: {$lt: value2}}]}) - selects data from collection where key1 = value1 or key2 < value2
    - Example3: db.<collection>.find({$or: [{key1: "value1"}, {key2: {$gt: value2}}]}) - selects data from collection where key1 = value1 or key2 > value2
    - db.<collection>.find({$or: [
        {$and: [{key1: "value1"}, {key2: "value2"}, key3: "value3"]},
        {$and: [{key9: "value9"}, {key8: "value8"}, key7: "value7"]},
    ]}) - this query selects data with clause: (key1 = value1 and key2 = value2 and key3 = value3) or (key9 = value9 and key8 = value8 = key7 = value7)
    - db.<collection>.find({key: {$regex: ^M}}) - selects data whose name starts with "M" - regex cheat sheet: https://lse-my472.github.io/week08/regular-expressions-cheat-sheet-v2.pdf
    - db.<collection>.update(<criteria>, {$set: {key: "value"}}) - updates a record
    - Example1: db.employee.update({name: "Ravikanth"}, {$set: {salary: 500000}} - Updates the salary for Ravikanth to 500000
    - Example2: db.employee.update({name: "Swarna"}, {$set: {role: "citizen", salary: 500000}}) - Updates the role and salary of Swarna to citizen and 500000 respectively
    - Example3: db.<collection>.update({age: {$gt: 35}}, {$set: {role: "citizen", "salary": 500000}}) - This just updates the first record of all the set of records having age > 35 and updates the role to citizen and salary to 500000
    - Example4: db.<collection>.update({age: {$gt: 35}}, {$set: {role: "citizen", "salary": 500000}}, {multi: true}) - This updates all the set of records having age > 35 and updates the role to citizen and salary to 500000
    - Example5: db.<collection>.update({}, {$set: {role: "citizen"}, {multi: true}) - This updates all the set of records role to citizen
    - db.<collection>.save - this option inserts the record, if it doesn't exist in the collection (or) if the record already exists, then just updates it
    - Example1: db.employee.save({name: "raj", "location": dubai}) - since the record doesn't exist, it has created a new record
    - Example2: db.employee.save("_id" : ObjectId("5eac53cea9cf6f265a4e516d"),name: "hello", age: 32, location:"dubai"}) - updated the existing record
    - db.<collection>.remove({key: "value"}) - remove the data where key = value
    - db.<collection>.remove({key: "value"}, 1) - remove the first record matching the criteria
    - db.<collection>.remove({}) - removes everybody
    - db.<collection>.find({<selection criteria>, <list of fields with toggle 0 or 1>}) - This command selects only limited columns from a bunch of columns in a collection
    - Example1: db.employee.find({name: "Ravikanth"}, {"_id": 0, "name": 1}) - selects name from the employee where name="Ravikanth"
    - Example2: db.employee.find({}, {"_id": 0, "name": 1}).pretty() - selects only name from all the other columns in the collection employee
    - Example3: db.employee.find({age: {$gt: 32}}, {"_id": 0, "name": 1}).pretty() - selects only name from the collection employee where the age > 32
    - Example4: db.employee.find({age: {$gt: 32}}, {"_id": 0, "name": 1, "Location": 1}).pretty() - prints name and location from collections employee where age > 32
    - db.<collection>.find({key: "value"}).limit(n) - selects all the data where key = value and limits the values on display
    - Example1: db.employee.find({age: {$gte: 32}}, {"_id": 0, "name": 1, "Location": 1}).pretty().limit(2) - From the collection employees, selects all the records where age >= 32 and while displaying limits the number of records to 2
    - db.<collection>.find({key: "value"}).sort({<field to sort on>:<1 for ascending, -1 for descending>}) - This command selects data from a collection with selection criteria and sorts by the key with 1 as ascending and -1 as descending
    - Example1: db.employee.find({}).sort({"age": -1}) - selects data from collection employee and sorts by age in descending order
    - Example2: db.employee.find({}, {name: 1, _id: 0, age: 1}).sort({"age": 1}) - selects data from collection employee, displays only name and age, and data is sorted by age in ascending order
    - db.<collection>.createIndex({field: <toggle to turn on - 1 for createIndex>}) - this will make the queries performant. There is an overhead of regularly updating the index.
    - Example1: db.oscars.createIndex("title": 1) - created an index on the column title on the collections: oscars
    - db.<collection>.find().count() - prints the count of records in the collection
    - Example1: db.oscars.find().count() - prints the count of records in the collection "oscars"
    - Example2: db.oscars.find({releaseYear: 1999}).count() - prints the count of records in the collection "oscars", where the releaseYear is 1999
    - Example3: db.oscars.find({releaseYear: {$gt: 1988}}).count() - prints the count of records in the collection "oscars" where releaseYear > 1988
    - db.<collection>.distinct('field') - shows the distinct set of records for that field
    - Example1: db.inventory.distinct('dept') - in the inventory, deptA is given 3 times and B is given 2 times, so the query just gives [A, B]
    - Example2: db.inventory.distinct('item.sku') - displays the distinct item.skus
    - Complex Aggregation: db.<collection name>.aggregate([{<match, sort, geoNear>},{<group>}])
    - Example1: db.orders.aggregate([{$match:{status:"A"}},{$group:{_id: "$cust_id",total: {$sum:"$amount"}}}]) - pulls all records where status: A and then groups by cust_id and sums up the amount
    - use admin; db.createUser({"user" : <username>, pwd: <password>, roles: [{role: role, db: <databases>}]}) - creates a user with username and password and assigns roles to the database 
    
    
     
    
    
Operators:

    - equality: {key: "value"}
    - less than: {key: {$lt: "value"}}
    - greater than: {key: {$gt: "value"}}
    - less than equals: {key: {$lte: "value"}}
    - greater than equals: {key: {$gte: "value"}}
    - not equals: {key: {$ne: "value"}}
    
    
    
    
    
    