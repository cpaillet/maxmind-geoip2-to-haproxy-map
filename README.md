## Generate a geoip file for use with HAProxy

Inspired by https://github.com/andyjack/geoip2-haproxy-ranges/

Download the free versions from https://dev.maxmind.com/geoip/geoip2/geolite2/.

## Running

```sh
git clone https://github.com/cpaillet/maxmind-geoip2-to-haproxy-map
cd maxmind-geoip2-to-haproxy-map
go get github.com/oschwald/maxminddb-golang
go run . --db path/to/GeoLite2-Country.mmdb --destDir output
```

The file look like
```
::200:9600/119 FR|PDL|48.858200|2.338700
::200:9800/120 FR|IDF|48.776400|2.290300
(...)
5.66.249.0/24 GB|ENG|53.966700|-1.083300
5.66.250.0/25 GB|ENG|54.438500|-0.871200
(...)
```

In your haproxy frontend configuration, you can now add

```
http-request set-header x-country %[src,map_ip(/etc/haproxy/maps/geoip.map),field(1,|)]
http-request set-header x-subdivisions %[src,map_ip(/etc/haproxy/maps/geoip.map),field(2,|)]
http-request set-header x-latitude %[src,map_ip(/etc/haproxy/maps/geoip.map),field(3,|)]
http-request set-header x-longitude %[src,map_ip(/etc/haproxy/maps/geoip.map),field(4,|)]
```

