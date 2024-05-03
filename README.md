# Measure Pefromance of Different Language

### Grafana
```bash
docker run  -p 3001:3000 grafana/grafana
```

### Graphite
```bash
docker run  \
 --name graphite \
 --restart=always \
 -p 80:80 \
 -p 2003-2004:2003-2004 \
 -p 2023-2024:2023-2024 \
 -p 8125:8125/udp \
 -p 8126:8126 \
 graphiteapp/graphite-statsd
```

### Jmeter

### Application 
- Go 
- Rust
- Express
- TODO : Add more


