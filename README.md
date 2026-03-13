# ClasseInfoService
ClasseInfo service provides information about CLASSE users.
```
# lookup by uid
curl "http://localhost:8302/translate?uid=abc"

# lookup by full user name
curl "http://localhost:8302/translate?name=First+Last"

# lookup by uid number
curl "http://localhost:8302/translate?uidNumber=1622"

# lookup by user's email
curl "http://localhost:8302/translate?email=user%40cornell.edu"
```

