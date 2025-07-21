# VK Internship Project

Проект представляет из себя задание для попадания на стажировку в ВКонтакте.


## Запуск проекта

Для запуска проекта используйте команду:
docker compose up --build

Перед запуском необходимо создать файл .env в корне проекта. Пример .env файла

```code
DB_URL=postgres://pavel:durov@db:5432/vk_server?sslmode=disable
DB_SOURCE=file:///app/migrations
CONN_STRING=host=db user=pavel password=durov dbname=vk_server sslmode=disable
JWT_SECRET=WcSzP3nnRNzjw8HkQtal0QcfDC0VbSeGW49i8Zz2Kl0bBgqEbf5eBbz1fgWZWGuJ
DB_NAME=vk_server
DB_USER=pavel
DB_PASSWORD=durov
```


## Серверные эндпоинты

### 1. `POST /register`

**Request Body:**

```json
{
  "username": "hu123",
  "password": "1234567890"
}
```

Регистрирует нового пользователя.

---

### 2. `POST /login`

**Request Body:**

```json
{
  "username": "hu123",
  "password": "1234567890"
}
```

Авторизация пользователя.

---

### 3. `POST /addAd`

**Request Body:**

```json
{
  "adName": "Дора",
  "price": 10500,
  "description": "",
  "imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg"
}
```

Добавляет новое объявление.

---

### 4. `GET /getAll`

**Request Parameters:**

- `limit` — количество объявлений на страницу
- `offset` — смещение (для пагинации)
- `date_sort` — сортировка по дате (`asc` / `desc`)
- `price_sort` — сортировка по цене (`asc` / `desc`)
- `min_price` — минимальная цена
- `max_price` — максимальная цена

Возвращает список объявлений по заданным фильтрам.