# jorm

jorm is an interface wrapper around [gorm](http://gorm.io/)

jorm can replace gorm with very little code change, and allows the benefit of easier testing via mocks and the ability to trace a query with using the [OpenTracing](https://opentracing.io/) instrumentation by [opentracing-gorm](https://github.com/smacker/opentracing-gorm)

To set the context simply call `db.WithContext(ctx)` and use the db returned by that function
