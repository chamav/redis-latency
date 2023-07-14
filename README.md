# Log Loader

Измеритель Latency по ping и pong до redis в микросекундах в формате prometheus

### Как запустить
go run  ./cmd/redis-latency/redis-latency.go

Входим по локальному адресу и в параметрах url указываем адрес до redis
http://127.0.0.1:8085/?url=localhost:6379
возвращает для примера

redis_latency 2784 1689339123

### Как скопилировать
go build -tags osusergo,netgo ./cmd/redis-latency/redis-latency.go

### Как обновить
После обновления кода не забыть изменить тэг, чтобы изменения могли быть выкачены:

```git
git commit ...
git tag v*.*.*
git push
git push --tags
```