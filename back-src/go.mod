module back-src

go 1.15

require (
	github.com/gin-contrib/size v0.0.0-20200815104238-dc717522c4e2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.1.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/gin-contrib/size v0.0.0-20200815104238-dc717522c4e2 => github.com/ashkan-khd/size v0.0.0-20200911190856-a442161182eb
