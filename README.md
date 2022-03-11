### Создание образа
`docker build . -f ./DockerFile -t main-go-image`
> Образ собирается в 2 этапа. Сначала билдится, потом на облегченный докер образ закидывается готовый бинарник.
> 
> При каждом изменении исходников проекта, нужно пересобирать образ, и перезапускать контейнер
### Запуск в контейнере
`docker-compose up -d --force-recreate`

### Как проверить работу
выполнить http запрос `http://localhost:8090/hi` Веб сервер вернет строку из docker-compose.yaml environment WHAT_SAY

### Полезные ссылки
* Понятная статейка по докеру
https://selectel.ru/blog/what-is-docker/
* Мультисборка контейнеров 
https://docs.docker.com/develop/develop-images/multistage-build/
* Инфа по docer-compose
https://habr.com/ru/company/ruvds/blog/450312/
