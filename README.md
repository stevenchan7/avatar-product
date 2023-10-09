# avatar-product

## Public routes
- /products
  
  Endpoint for getting all products. Return a JSON containing:
  {"success": true, "data": []Product}
  

## Public auth routes
- /auth/register
  
  Endpoint for register new admin account. Required request body:
  1. username: string
  2. password: string
- /auth/login
  
  Endpoint for log into existing admin account. Required request body:
  1. username: string
  2. password: string
- /auth/logout
  Endpoint for log out, removing existing cookie.

## Admin routes (must login)
- /admin/add-product
  
  Endpoint for adding new product into database. Required request body:
  1. title: string
  2. desc: string
  3. image: string (link to image drive)
  4. playstore: string (play store link)
  5. appstore: string (app store link)
- /admin/edit-product
  Endpoint for update existing product in database. Required request body:
  1. title: string
  2. desc: string
  3. image: string (link to image drive)
  4. playstore: string (play store link)
  5. appstore: string (app store link)
  6. prodID: string (corresponding product's ID)
