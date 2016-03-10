# mysql2mongo

``` python
cd mysql2mongo

go build
 
./mysql2mongo --conf /tmp/mysql2mongo.cnf
``` 

# conf example
``` python
mysql :

 url : "root:123456@tcp(localhost:3306)/data"
  
 sql : "select f1,f2,f3 from order limit 100;"
    
 dbname : "data"
	  
 tablename : "order"
	    
 ncloumns : ["Paytype","Dorm_id"]
		  

mongo :

 url : "localhost:27017"
		    
 dbname : "test"
			  
 collection : "t1"
```
