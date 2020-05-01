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
    
    
    