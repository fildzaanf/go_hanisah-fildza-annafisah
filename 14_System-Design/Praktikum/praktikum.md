# Praktikum System Design

## Entity Relationship Diagram
Tap for details [ERD](https://drive.google.com/file/d/1R24EimbW6N4U9EPIf1aVtms0MA7unhMx/view?usp=sharing) 

---- 

## Use Case Diagram
Tap for details [UCD](https://drive.google.com/file/d/1R24EimbW6N4U9EPIf1aVtms0MA7unhMx/view?usp=sharing) 

----
## Query

### SQL
```sql
SELECT * FROM users;
```

### Redis
```r
SCAN 0 MATCH users:* 
```

### Neo4j
```
MATCH (u:Users)
RETURN u
```

### Cassandra
```cs
SELECT * FROM users;
```

