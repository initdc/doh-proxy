# doh - proxy
> Write in go, lightweight and flexible

> ! This project is still developing...

---
## Quickstart

```console
docker run -dp 80:80 initdc/doh-proxy 
```

## Config 

This project developed with dns.google doh service as a template.

it has some `ENV` variable you can custom

```console
 UPSTREAM for upstream, default: dns.google

 UPPATH for upstream dns query path, dns-query

 

 PATH1  server listen path1, query 

 PATH2  server listen path2, resolve

 

 QUERY upstream path for /query listening

 RESOVLE upstream path for /resolve listening

```

## Licence
MIT