# Техническое задание

____

## Сводное HTTP API

API системы доступен по пути верхнего уровня `/api`.

### Проверка работоспособности сервера

- **GET** `/api/ping` – проверка работоспособности сервиса.

### Система работы с пользователем должна предоставлять следующие HTTP-хендлеры:

### Система работы с пользователем доступна по пути среднего уровня `/consumers`.

- **POST** `/api/consumers/registration` — регистрация пользователя.
- **POST** `/api/consumers/login` — аутентификация пользователя.
- **DELETE** `/api/consumers/delete` - удаление пользователя.
- **GET** `/api/consumers/get_me` — получение профиля аутентифицированного пользователя.
- **PUT** `/api/consumers/update_password` - изменение пароля пользователя.
- **POST** `/api/consumers/refresh-token` — перевыпуск пары авторизационных токенов.

#### Предполагается:

- **GET** `/api/consumers/all` — получение списка пользователей;

### Система работы с тестами должна предоставлять следующие HTTP-хендлеры:

Система работы с тестами доступна по пути среднего уровня `/tests`.

- **GET** `/api/test/all` - получение всех тестов, которые имеются в системе.
- **GET** `/api/test/:test_id` - получение теста по его идентификатору.

### Система работы с решенными тестами должна предоставлять следующие HTTP-хендлеры:

Система работы с тестами доступна по пути среднего уровня `/resolved`.


- **POST** `/api/resolved/create` - создание решенного теста.
- **GET** `/api/resolved/:resolved_id` - получение решенного теста по его идентификатору.


### Система работы с результатами тестирования должна предоставлять следующие HTTP-хендлеры:

Система работы с результатами тестирования доступна по пути среднего уровня `/results`.

- **GET** `/api/results/all` - получение всех результатов пользователя.
- **POST** `/api/results/create` - сохранение результата тестирования.
- **GET** `/api/results/:result_id` - получение результата тестирования по его идентификатору.
- **GET** `/api/results/my` - получение результатов тестирования по идентификатору пользователя. 
- **POST** `/api/results/send_on_email` - отправка результата тестирования на e-mail адрес.

____

### Метод для проверки работоспособности сервиса.

#### Хендлер: **GET** `/api/ping`.

- ##### Проверка работоспособности сервиса.

Формат запроса:

```
GET /api/ping HTTP/1.1
Content-Type application/json
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

### Описание методов системы для работы с решенными тестами.

#### Хендлер: **POST** `/api/resolved/create`.

- ##### Создание решенного теста.

Формат запроса:

```
GET /api/resolved/create HTTP/1.1
Content-Type application/json
...

{
  "test_type": "first_order_test",
  "questions": [
    {
      "question_order": "<question_order>",
      "question": "<question>",
      "question_answer": "<answer>"
    },
    ...
  ]
}
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "id": "<id>",
  "questions": [
    {
      "question_order": "<question_order>",
      "question": "<question>",
      "question_answer": "<question_answer>",
      "mark": "<mark>"
    },
    ...
  ]
}
```

#### Хендлер: **GET** `/api/resolved/:resolved_id`.

- ##### Получение решенного теста по его идентификатору.

Формат запроса:

```
GET /api/resolved/:resolved_id HTTP/1.1
Content-Type application/json
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "id": "<id>",
  "user_id": "<user_id>",
  "resolved_type": "<resolved_type>",
  "is_active": "<is_active>",
  "passed_at": "<passed_at>",
  "questions": [
    {
      "resolved_id": "<resolved_id>",
      "question_order": "<question_order>",
      "issue": "<issue>",
      "question_answer": "<question_answer>",
      "image_location": "<image_location>",
      "mark": "<mark>"
    },
    ...
  ]
}
```

### Описание методов системы для работы с пользователем.

- #### Регистрация пользователя.

##### Хендлер: **POST** `/api/consumers/registration`.

Формат запроса:

```
POST /api/consumers/registration HTTP/1.1
Content-Type application/json
...

{
  "login": "<login>",
  "password": "<password>"
} 
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

- #### Аутентификация пользователя

##### Хендлер: **POST** `/api/consumers/login`.

Формат запроса:

```
POST /api/consumers/login HTTP/1.1
Content-Type application/json
...

{
  "login": "<login>",
  "password": "<password>"
}
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

- #### Удаление пользователя

##### Хендлер: **DELETE** `/api/consumers/delete`.

Формат запроса:

```
DELETE /api/consumers/delete HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...
```

- #### Получение профиля аутентифицированного пользователя.

##### Хендлер: **GET** `/api/consumers/get_me`.

Формат запроса:

```
GET /api/consumers/get_me HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "id": "<consumer_id>",
  "login": "<login>",
  "created_at": "<timestamp>",
}
```

- #### Изменение пароля пользователя.

##### Хендлер: **PUT** `/api/consumers/update_password`.

Формат запроса:

```
PUT /api/consumers/update_password HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...

{
  "old_password": "<old_password>",
  "new_password": "<new_password>"
}
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...
```

- #### Перевыпуск пары авторизационных токенов.

##### Хендлер: **POST** `/api/consumers/refresh-token`.

Формат запроса:

```
POST /api/consumers/refresh-token HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

### Описание методов системы для работы с тестами.

- #### Получение всех тестов, которые имеются в системе.

##### Хендлер: **GET** `/api/test/all`.

Формат запроса:

```
GET /api/test/all HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "tests": [
    {
      "id": "<id>",
      "name": "<name>",
      "description": "<description",
      "questions": [
        {
          "order": "<order>",
          "question": "<question>"
        },
        ...
      ]
    },
    ...
  ]
}
```

- #### Получение теста по его идентификатору.

##### Хендлер: **GET** `/api/test/:test_id`.

Формат запроса:

```
GET /api/test/:test_id HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "id": "<id>",
  "name": "<name>",
  "description": "<description",
  "questions": [
    {
      "order": "<order>",
      "question": "<question>"
    },
    ...
  ]
}
```

### Описание методов системы для работы с результатами тестирования.

- #### Сохранение результата тестирования.

##### Хендлер: **POST** `/api/results/create`.

Формат запроса:

```
POST /api/results/create HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...

{
  "resolved_id": "<resolved_id>",
  "test_type": "<test_type>",
  "image_location": "<image_location>",
  "questions": [
    {
      "question_order": "<question_order>",
      "mark": "<mark>"
    },
    ...
  ]
}
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "id": "<id>",
  "resolved_id": "<resolved_id>",
  "image_location": "<image_location>",
  "professions": [
    "<profession>",
    "<profession>",
    "<profession>"
  ]
}
```

- #### Получение результата тестирования по его идентификатору.

##### Хендлер: **GET** `/api/results/:result_id`.

Формат запроса:

```
GET /api/results/:result_id HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...

{
  "resolved_id": "<resolved_id>",
  "test_type": "<test_type>",
  "image_location": "<image_location>",
  "questions": [
    {
      "question_order": "<question_order>",
      "mark": "<mark>"
    },
    ...
  ]
}
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  "id": "<id>",
  "resolved_id": "<resolved_id>",
  "image_location": "<image_location>",
  "professions": [
    "<profession>",
    "<profession>",
    "<profession>"
  ]
}
```

- #### Получение результатов тестирования по идентификатору пользователя.

##### Хендлер: **GET** `/api/results/my`.

Формат запроса:

```
GET /api/results/my HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...

{
  Results: [
    {
      "id": "<id>",
      "user_id": "<user_id>",
      "resolved_id": "<resolved_id>",
      "image_location": "<image_location>",
      "professions": [
        "<profession>",
        "<profession>",
        "<profession>"
      ],
      "created_at": "<created_at>"
    },
    ...
  ]
}
```

- #### Отправка результата тестирования на e-mail адрес.

##### Хендлер: **POST** `/api/results/send_on_email`.

Формат запроса:

```
POST /api/results/send_on_email HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...

{
  "test_name": "<test_name>",
  "email": "<email>",
  "professions": [
    "<profession>",
    "<profession>",
    "<profession>"
  ],
}
```

Формат ответа:

```
HTTP/1.1
Content-Type application/json
...
```
