-- 1. Redis 
KEYS user:*

-- 2. Neo4j
MATCH (u:User) RETURN u;

-- 3. Casandra
SELECT * FROM users;