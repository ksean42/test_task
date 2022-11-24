# test_task

## Тестовое задание

Запуск:
````
docker-compose up
````
- get и backup методы работают на порту 8070
- set метод работает на порту 8071

Реализовано основное задание + метод бэкапа


- **Метод получения данных по user_id**

Запрос:

````
curl -X GET "localhost:8070/get?user_id=1"
````
Ответ:

````
Status 200
{
	"Result": {
		"user_id": "1",
		"postpaid_limit": 2545,
		"spp": 2321,
		"shipping_fee": 20,
		"return_fee": 29
	}
}
````

- **Метод получения бэкапа данных**

Запрос:

````
localhost:8070/backup
````

В ответ получаем файл csv.gz со временем генерации бэкапа в первой строке

- **Метод set**

Запрос:

````
curl -X POST localhost:8071/set --user admin:admin -H "Content-Type: application/json" -d '{ "user_id": "1", "postpaid_limit": 254, "spp": 2321, "shipping_fee": 20, "return_fee": 29 }'
````

Ответ:
````
Status 200
{
    "Result":"OK"
}
````
## Задание
1. Программа минимум

- 2 http роутера на разных портах
- На 1 порту метод set, на другом порту метод get.

set - post запрос, в body запроса json модели UserGrade. Сохраняем структуру в storage

get - get запрос, параметр в urlencoded ?user_id=, на выходе json модели UserGrade из storage

- Реализовать пакет storage. Задача пакета - хранить в рам структуру UserGrade по стринговому ключу. Ключ UserId

Имеет публичные методы set, get

type UserGrade struct {

UserId        string json:"user_id" validate:"required"

PostpaidLimit int    json:"postpaid_limit"

Spp           int    json:"spp"

ShippingFee   int    json:"shipping_fee"

ReturnFee     int    json:"return_fee"

}

- Реализовать middleware с basic auth, закрыть им метод set
- В set могут присылать данные частями. Одним запросом Spp, след ShippingFee и т.д.

2.  Программа максимум (hard skill)

- Реализовать репликацию в методе set. Для репликации используем брокер сообщений (nats streaming/kafka на ваш вкус - поднимайте локлько в докере)

При получении данных в метод set, сервис публикует сообщение в канал. В горутине сервис подписывается на этот же канал.

Отфильтровывает свои сообщения и обрабатывает сообщения других реплик.

- Реализовать метод /backup. Метод при запросе генерит дамп файл локальных данных в формате csv.gz, передает в response

В бекап зашить время, когда он был сгенерирован.

- При старте приложения мы дергаем метод /backup реплики. Заполняем данными storage.

Подписываемся к каналу с того времени, которое указано в бекапе.

- Подумайте над порядком запуска функций, бекап может быть большим и восстановление может занять время