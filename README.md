# Mini Commerce API

Tech stack

1. Go
2. Gin Gonic
3. Go JWT
4. GORM
5. MySQL
6. Redis
7. RabbitMQ

## ABOUT

Here is an online store API that includes two roles for authorization: user and admin. Admin has the ability to perform CRUD operations on a product, while users are only allowed to retrieve product information. The authentication is implemented using JWT, with two tokens issued upon login: an access token and a refresh token (stored in Redis). RabbitMQ is utilized for publishing and consuming data when payment status updates occur (for instance, transitioning from pending to paid).

## GET STARTED

1. Configure the `.env` file according to the `.env.example` file.
2. Ensure you have installed Go by running the command `go version`.
3. Ensure that you are in the develop branch by executing the command `git checkout develop`.
4. Run the command `go run main.go`.

## URL

SERVER

```
http://localhost:8080
```

### DOCUMENTATION

#### List RESTful API

1. [Auth](#Auth)
2. [Product](#Product)
3. [Cart](#Cart)
4. [Transaction](#Transaction)
5. [Address](#Address)

#### Global Error Response

Forbidden Access:

```
{
    "code": 403,
    "message": "Forbidden",
    "error": {
        "message": "user role not authorized"
    }
}
```

Token Expired:

```
{
    "code": 401,
    "message": "Unauthorized",
    "error": {
        "message": "Token is expired"
    }
}
```

Unauthorized:

```
{
    "code": 401,
    "message": "Unauthorized",
    "error": {
        "message": "unauthorized user"
    }
}
```

Invalid Token:

```
{
    "code": 401,
    "message": "Unauthorized",
    "error": {
        "message": "signature is invalid"
    }
}
```

Not found:

```
{
    "code": 404,
    "message": "Not found",
    "error": {
        "message": "record not found"
    }
}
```

#### Auth

1. ##### Register User `POST METHOD`
   endpoint : `/auth/register`  
   json request body :
   ```
   {
   "name": "taufiq 4",
   "email": "taufiq4@gmail.com",
   "phone_number": "0812738128",
   "password": "123456"
   }
   ```
   json response :
   ```
    {
    "code": 201,
    "message": "Success create new user",
    }
    "data": <data>
   ```
   json response validation required :
   ```
   {
    "code": 400,
    "message": "All field are required",
    "error": {
        "message": [
            "Key: 'UserInput.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'required' tag"
        ]
    }
    }
   ```
   json response validation email :
   ```
   {
        "code": 400,
        "message": "Bad Request",
        "error": {
            "message": "Email already exists"
        }
    }
   ```
2. ##### Login User `POST METHOD`
   endpoint : `/auth/login`  
   json request body :
   ```
   {
    "email": "taufiq4@gmail.com",
    "password": "123456"
    }
   ```
   json response :
   ```
    {
    "code": 200,
    "message": "Success login",
    "data": {
        "access_token": <access_token>,
        "refresh_token": <refresh_token>
        "user": <uuser>
    }
    }
   ```
   json response validation required :
   ```
   {
    "code": 400,
    "message": "All fields are required",
    "error": {
        "message": [
            "Key: 'UserLogin.Password' Error:Field validation for 'Password' failed on the 'required' tag"
        ]
    }
    }
   ```
   json response invalid email or password :
   ```
   {
    "code": 400,
    "message": "Invalid email or password",
    "error": {
        "message": "invalid password"
    }
    }
   ```
3. ##### Refresh Token User `POST METHOD`

   endpoint : `/auth/refresh-token`  
   json request body :

   ```
    "refresh_token": "refresh_token"
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "Token refreshed",
    "data": {
        "access_token": <access_token>
    }
    }
   ```

   json response invalid refresh token:

   ```
   {
    "code": 401,
    "message": "Unauthorized",
    "error": {
        "message": "Refresh token not found or invalid"
    }
    }
   ```

4. ##### Logout User `POST METHOD`
   endpoint : `/auth/logout`  
    request header :
   ```
   {
   "Authorization": <access_token>
   }
   ```
   json response :
   ```
    {
    "code": 200,
    "message": "Successfully logged out",
    "data": {}
    }
   ```
   json response invalid token :
   ```
    {
    "code": 401,
    "message": "Unauthorized",
    "error": {
        "message": "signature is invalid"
    }
    }
   ```

#### Product

1. ##### Get All Product `GET METHOD`
   endpoint : `/products`  
   json request body :
   ```
    not needed
   ```
   json response :
   ```
    {
    "code": 200,
    "message": "success get all product",
    "data": [
        {
            "id": 1,
            "name": "iPhone 14 Pro Max 128 GB Updated",
            "description": "Garansi resmi IBOX updated",
            "photo_product": "iphone14_image.jpg",
            "price": 1400000,
            "stock": 9,
            "user_id": 4
        },
        {
            "id": 2,
            "name": "iPhone 15 Pro Max 128 GB",
            "description": "Garansi resmi IBOX",
            "photo_product": "iphone15_image.jpg",
            "price": 1700000,
            "stock": 8,
            "user_id": 4
        }
    ]
    }
   ```
2. ##### Get Product By Id `GET METHOD`
   endpoint : `/product/:productId`  
   json request body :
   ```
   not needed
   ```
   json response :
   ```
    {
    "code": 200,
    "message": "success get product by id",
    "data": {
        "id": 1,
        "name": "iPhone 14 Pro Max 128 GB Updated",
        "description": "Garansi resmi IBOX updated",
        "photo_product": "test",
        "price": 1400000,
        "stock": 9,
        "user_id": 4
    }
    }
   ```
3. ##### Create Product `POST METHOD`

   endpoint : `/product`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "name": "iPhone 15 Pro Max 128 GB",
    "description": "Garansi resmi IBOX",
    "photo_product": "iphone_image.jpg",
    "price": 1700000,
    "stock": 2
    }
   ```

   json response :

   ```
    {
    "code": 201,
    "message": "success create product",
    "data": <data>
    }
   ```

   json response field required:

   ```
   {
    "code": 400,
    "message": "All field are required",
    "error": {
        "error": [
            "Key: 'ProductInput.PhotoProduct' Error:Field validation for 'PhotoProduct' failed on the 'required' tag"
        ]
    }
    }
   ```

4. ##### Update Product `PUT METHOD`

   endpoint : `/product/:productId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "name": "iPhone 15 Pro Max 128 GB Updated"
    }
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "Successfully update product",
    "data": <data>
    }
   ```

   json response product not found :

   ```
    {
    "code": 404,
    "message": "not found",
    "error": "record not found"
    }
   ```

5. ##### Delete Product `DEL METHOD`

   endpoint : `/product/:productId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
   not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "product successfully deleted"
    }
   ```

   json response product not found :

   ```
    {
    "code": 404,
    "message": "not found",
    "error": "record not found"
    }
   ```

#### Cart

1. ##### Get My Cart `GET METHOD`
   endpoint : `/products`  
   json request body :
   ```
    not needed
   ```
   json response :
   ```
    {
    "code": 200,
    "message": "success get my cart",
    "data": [
        {
            "id": 7,
            "user_id": 6,
            "product_id": 1,
            "product_name": "iPhone 14 Pro Max 128 GB Updated",
            "quantity": 1,
            "total_price": 1400000
        }
    ]
    }
   ```
2. ##### Add Product to Cart `POST METHOD`

   endpoint : `/cart/:productId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "quantity": 1
    }
   ```

   json response :

   ```
    {
    "code": 201,
    "message": "success create cart",
    "data": {
        "id": 7,
        "user_id": 6,
        "product_id": 1,
        "product_name": "iPhone 14 Pro Max 128 GB Updated",
        "quantity": 1,
        "total_price": 1400000
    }
    }
   ```

   json response over quantity:

   ```
    {
    "code": 400,
    "message": "Bad Request",
    "error": "quantity is greater than stock"
    }
   ```

3. ##### Update Product in Cart `PUT METHOD`

   endpoint : `cart/:cartId/product/:productId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "quantity": 2
    }
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "Successfully update cart",
    "data": <data>
    }
   ```

   json response not found :

   ```
    {
    "code": 404,
    "message": "product or cart not found",
    "error": "record not found"
    }
   ```

4. ##### Delete Product From Cart `DEL METHOD`

   endpoint : `/cart/:cartId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
   not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "success delete cart",
    "data": "cart successfully deleted"
    }
   ```

   json response cart not found :

   ```
    {
    "code": 404,
    "message": "Cart not found",
    "error": "record not found"
    }
   ```

#### Transaction

1. ##### Get All Transaction `GET METHOD`

   endpoint : `/transactions`

   request header:

   ```
   "Authorization": <access_token>
   ```

   json request body :

   ```
    not needed
   ```

   json response :

   ```
   {
    "code": 200,
    "message": "success get all transaction",
    "data": [
        {
            "id": 7,
            "date": "2024-01-05T14:54:43.308+07:00",
            "user_id": 6,
            "address_id": 3,
            "cart_id": 5,
            "product_id": 1,
            "product_name": "iPhone 14 Pro Max 128 GB Updated",
            "quantity": 1,
            "total_price": 1400000,
            "payment_method": "transfer",
            "status": "paid",
            "address": {
                "id": 3,
                "receiver": "Putri Dzakiyah Updated",
                "phone_receiver": "081293819190",
                "address_detail": "Jl. Happy selalu no. 35",
                "province": "DKI Jakarta",
                "city": "Jakarta Selatan",
                "user_id": 6
            }
        },
        {
            "id": 9,
            "date": "2024-01-05T15:27:47.295+07:00",
            "user_id": 6,
            "address_id": 1,
            "cart_id": 6,
            "product_id": 1,
            "product_name": "iPhone 14 Pro Max 128 GB Updated",
            "quantity": 1,
            "total_price": 1400000,
            "payment_method": "transfer",
            "status": "paid",
            "address": {
                "id": 1,
                "receiver": "Taufiqurrahman Saleh",
                "phone_receiver": "081293812321",
                "address_detail": "Jl. Bahagia selalu no. 24",
                "province": "DKI Jakarta",
                "city": "Jakarta Utara",
                "user_id": 6
            }
        }
    ]
    }
   ```

2. ##### Get Detail Transaction`GET METHOD`

   endpoint : `/transaction/:transactionId`

   request header:

   ```
   "Authorization": <access_token>
   ```

   json request body :

   ```
   not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "success get transaction detail",
    "data": {
        "id": 7,
        "date": "2024-01-05T14:54:43.308+07:00",
        "user_id": 6,
        "address_id": 3,
        "cart_id": 5,
        "product_id": 1,
        "product_name": "iPhone 14 Pro Max 128 GB Updated",
        "quantity": 1,
        "total_price": 1400000,
        "payment_method": "transfer",
        "status": "paid",
        "address": {
            "id": 3,
            "receiver": "Putri Dzakiyah Updated",
            "phone_receiver": "081293819190",
            "address_detail": "Jl. Happy selalu no. 35",
            "province": "DKI Jakarta",
            "city": "Jakarta Selatan",
            "user_id": 6
        }
    }
    }
   ```

3. ##### Create Product `POST METHOD`

   endpoint : `/product`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "address_id": 1,
    "cart_id": 6,
    "payment_method": "transfer"
    }
   ```

   json response :

   ```
    {
    "code": 201,
    "message": "success create transaction",
    "data": <data>
    }
   ```

   json response field required:

   ```
   {
    "code": 400,
    "message": "All field are required",
    "error": {
        "message": [
            "Key: 'TransactionInput.PaymentMethod' Error:Field validation for 'PaymentMethod' failed on the 'required' tag"
        ]
    }
    }
   ```

   json response not found:

   ```
   {
   "code": 404,
   "message": "not found",
   "error": "cart not found"
   }
   ```

4. ##### Payment Transaction `PUT METHOD`

   endpoint : `/transaction/payment/:transactionId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "nominal": 1400000
    }
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "success payment transaction",
    "data": {}
    }
   ```

   json response invalid nominal :

   ```
   {
    "code": 400,
    "message": "Bad request",
    "error": "Invalid nominal total price"
    }
   ```

   json response product not found :

   ```
    {
    "code": 404,
    "message": "not found",
    "error": "record not found"
    }
   ```

5. ##### Get All User Transaction **Only Admin** `GET METHOD`

   endpoint : `/transactions-user`

   request header:

   ```
   "Authorization": <access_token>
   ```

   json request body :

   ```
    not needed
   ```

   json response :

   ```
   {
    "code": 200,
    "message": "success get all user transaction",
    "data": [<all_data_transactions>]
    }
   ```

6. ##### Delete Transaction **(Only Admin)** `DEL METHOD`

   endpoint : `/transaction/:transactionId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
   not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "transaction successfully deleted"
    }
   ```

   json response transaction not found :

   ```
    {
    "code": 404,
    "message": "not found",
    "error": "record not found"
    }
   ```

#### Address

1. ##### Get All Address `GET METHOD`

   endpoint : `/addresses`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "success get address by user id",
    "data": [
        {
            "id": 1,
            "receiver": "Taufiqurrahman Saleh",
            "phone_receiver": "081293812321",
            "address_detail": "Jl. Bahagia selalu no. 24",
            "province": "DKI Jakarta",
            "city": "Jakarta Utara",
            "user_id": 6
        },
        {
            "id": 3,
            "receiver": "Putri Dzakiyah Updated",
            "phone_receiver": "081293819190",
            "address_detail": "Jl. Happy selalu no. 35",
            "province": "DKI Jakarta",
            "city": "Jakarta Selatan",
            "user_id": 6
        }
    ]
    }
   ```

2. ##### Get Address By Id `GET METHOD`

   endpoint : `/address/:addressId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
   not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "success get address by id",
    "data": {
        "id": 1,
        "receiver": "Taufiqurrahman Saleh",
        "phone_receiver": "081293812321",
        "address_detail": "Jl. Bahagia selalu no. 24",
        "province": "DKI Jakarta",
        "city": "Jakarta Utara",
        "user_id": 6
    }
    }
   ```

   json response not found:

   ```
   {
    "code": 404,
    "message": "Not found",
    "error": {
        "message": "record not found"
    }
    }
   ```

3. ##### Create Address `POST METHOD`

   endpoint : `/address`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "receiver": "Putri Dzakiyah Baru",
    "phone_receiver": "081293819190",
    "address_detail": "Jl. Happy selalu no. 35",
    "province": "DKI Jakarta",
    "city": "Jakarta Selatan"
    }
   ```

   json response :

   ```
    {
    "code": 201,
    "message": "success create new address",
    "data": <data>
    }
   ```

   json response field required:

   ```
   {
    "code": 400,
    "message": "All field are required",
    "error": {
        "message": [
            "Key: 'AddressInput.Province' Error:Field validation for 'Province' failed on the 'required' tag"
        ]
    }
    }
   ```

4. ##### Update Address `PUT METHOD`

   endpoint : `/address/:addressId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
    {
    "receiver": "Putri Dzakiyah Updated"
    }
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "Successfully update product",
    "data": <data>
    }
   ```

   json response product not found :

   ```
    {
    "code": 404,
    "message": "not found",
    "error": "record not found"
    }
   ```

5. ##### Delete Address `DEL METHOD`

   endpoint : `/address/:addressId`

   request header :

   ```
    "Authorization": <access_token>
   ```

   json request body :

   ```
   not needed
   ```

   json response :

   ```
    {
    "code": 200,
    "message": "address successfully deleted"
    }
   ```

   json response address not found :

   ```
    {
    "code": 404,
    "message": "not found",
    "error": "record not found"
    }
   ```
