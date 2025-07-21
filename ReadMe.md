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

**Response Body:**
```json
{
"message": "User created"
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

**Response Body:**
```json
{
  "token": "Example of token"
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
**Response Body:**
```json
{
  "adId": "7a3f74ba-ef2c-4549-9cf4-934e56a26e63",
  "adName": "Дора",
  "description": "",
  "imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
  "price": 10500,
  "username": "hu123"
}
```

Добавляет новое объявление.

---

### 4. `GET /getAll`

**Request Parameters:**

- `limit` — количество объявлений на страницу
- `offset` — смещение (для пагинации)
- `date_sort` — сортировка по дате (`ASC` / `DESC`)
- `price_sort` — сортировка по цене (`ASC` / `DESC`)
- `min_price` — минимальная цена
- `max_price` — максимальная цена

**Response Body without bearer token:**
```json
[
{
"adId": "7a3f74ba-ef2c-4549-9cf4-934e56a26e63",
"name": "Дора",
"description": "",
"imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
"price": 10500,
"username": "hu123"
},
{
"adId": "929c9416-1534-4403-84a7-ab7adc4cfbab",
"name": "Дора",
"description": "",
"imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
"price": 10500,
"username": "hu123"
},
{
"adId": "217cb890-d290-4870-8512-ae34e1c89469",
"name": "Дора",
"description": "",
"imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
"price": 10500,
"username": "hu123"
}
]
```
**Response Body with bearer token:**
```json
[
    {
        "adId": "7a3f74ba-ef2c-4549-9cf4-934e56a26e63",
        "name": "Дора",
        "description": "",
        "imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
        "price": 10500,
        "username": "hu123",
        "mine": true
    },
    {
        "adId": "929c9416-1534-4403-84a7-ab7adc4cfbab",
        "name": "Дора",
        "description": "",
        "imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
        "price": 10500,
        "username": "hu123",
        "mine": true
    },
    {
        "adId": "217cb890-d290-4870-8512-ae34e1c89469",
        "name": "Дора",
        "description": "",
        "imageUrl": "https://images.pexels.com/photos/9218709/pexels-photo-9218709.jpeg",
        "price": 10500,
        "username": "hu123",
        "mine": true
    }
]
 
```
Возвращает список объявлений по заданным фильтрам.

Для запроса на добавление объявления нужен bearer token

Для запроса на получение ленты объявлений bearer token не обязателен