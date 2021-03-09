# doh - proxy
> Write in go, lightweight and flexible

This project developed with dns.google doh service as a template.


---
## Quickstart

```docker run -dp 80:80 initdc/doh-proxy```

## Config 
it has some `ENV` variable you can custom
```
 UPSTREAM for upstream, default: dns.google

 UPPATH for upstream dns query path, dns-query

 

 PATH1  server listen path1, query 

 PATH2  server listen path2, resolve

 

 QUERY upstream path for /query listening

 RESOVLE upstream path for /resolve listening
```

## Licence
MIT