# Komikku API
Restful API Manga bahasa Indonesia built with ❤️ and Go

# Usage
1. Clone this repository
    ```bash
    git clone https://github.com/Romi666/komikku-api.git
    ```
2. Init module
    ```
   go mod init komikku-api
   ```

3. Installing dependencies
    ```
   go mod download
   ```
4. Build binary file
   ```
   make build
   ```
5. Run program
   ```
   make start
   ```
# Using Docker Compose
#### Build image and run container
   ```
   make up
   ```

#### Remove image and stop container
   ```
   make down
   ```

# API Collection

POSTMAN COLLECTION = https://www.getpostman.com/collections/4c984c36d27bb591c445


# API Output Example

## Get All Comic

#### Get Comic List

   ```
   http://localhost:3011/comic/list
   http://localhost:3011/comic/list?filter=manga
   http://localhost:3011/comic/list?filter=manhwa
   http://localhost:3011/comic/list?filter=manhua
   ```

#### Get Popular Comic
   ```
   http://localhost:3011/comic/popular?page=1
   ```

#### Get Recommended Comic
   ```
   http://localhost:3011/comic/recommended?page=1
   ```

#### Get Newest Comic
   ```
   http://localhost:3011/comic/newest?page=2
   ```

### API Example
   ```
   {
    "success": true,
    "data": [
        {
            "title": ".hack//G.U.+",
            "image": "https://thumbnail.komiku.id/wp-content/uploads/Manga-hackGU.jpg",
            "endpoint": "/manga/hack-g-u/"
        }
    ],
    "message": "Get All Comic",
    "code": 200
}
   ```

## Get Comic Info
   ```
   http://localhost:3011/comic{endpoint}
   ```

### API Example
   ```
   http://localhost:3011/comic/manga/hack-g-u/
   ```
   ```
   {
    "success": true,
    "data": {
        "thumbnail": "https://thumbnail.komiku.id/wp-content/uploads/Komik-hackGU.jpg",
        "title": ".hack//G.U.+",
        "type": "Manga",
        "author": "Tatsuya Hamazaki & Yuzuka Morita",
        "status": "Ongoing",
        "rating": "15 Tahun (minimal)",
        "genre": [
            "Action",
            "Adventure",
            "Fantasy",
            "Magic",
            "Mecha"
        ],
        "chapter_list": [
            {
                "name": "Chapter 6",
                "endpoint": "/ch/hack-g-u-chapter-6/"
            },
            {
                "name": "Chapter 5",
                "endpoint": "/ch/hack-g-u-chapter-5/"
            },
            {
                "name": "Chapter 4",
                "endpoint": "/ch/hack-g-u-chapter-4/"
            },
            {
                "name": "Chapter 3",
                "endpoint": "/ch/hack-g-u-chapter-3/"
            },
            {
                "name": "Chapter 2",
                "endpoint": "/ch/hack-g-u-chapter-2/"
            },
            {
                "name": "Chapter 1",
                "endpoint": "/ch/hack-g-u-chapter-1/"
            }
        ]
    },
    "message": "Get Comic Info",
    "code": 200
}
   ```

## Search Comic
   ```
   http://localhost:3011/comic/search?query={comicName}
   ```

### API Example
   ```
   http://localhost:3011/comic/search?query=one%20punch
   ```
   ```
   {
    "success": true,
    "data": [
        {
            "title": "One Punch Man",
            "image": "https://thumbnail.komiku.id/wp-content/uploads/Manga-One-Punch-Man.png",
            "endpoint": "/manga/one-punch-man/"
        },
        {
            "title": "Onepunchman Saitama vs God",
            "image": "https://thumbnail.komiku.id/wp-content/uploads/Manga-Onepunchman-Saitama-vs-God.jpg",
            "endpoint": "/manga/onepunchman-saitama-vs-god/"
        },
        {
            "title": "One-Punch Man (ONE)",
            "image": "https://thumbnail.komiku.id/wp-content/uploads/Manga-One-Punch-Man-ONE.jpg",
            "endpoint": "/manga/one-punch-man-one/"
        },
        {
            "title": "Here’s a Punch",
            "image": "https://thumbnail.komiku.id/wp-content/uploads/2021/09/Komik-Heres-a-Punch.png",
            "endpoint": "/manga/heres-a-punch/"
        },
        {
            "title": "Fire Punch",
            "image": "https://thumbnail.komiku.id/wp-content/uploads/2020/09/Komik-Fire-Punc.jpg",
            "endpoint": "/manga/fire-punch-2/"
        }
    ],
    "message": "Search manga",
    "code": 200
}
   ```

## Get Chapter Detail
   ```
   http://localhost:3011/comic/chapter{endpoint chapter}
   ```

### API Example
   ```
   http://localhost:3011/comic/manga/hack-g-u/
   ```
   ```
   {
    "success": true,
    "data": {
        "title": ".hack//G.U.+ Chapter 6",
        "image": [
            "https://cdn.komiku.co.id/wp-content/uploads/85998-1.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-2.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-3.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-4.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-5.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-6.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-7.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-8.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-9.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-10.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-11.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-12.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-13.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-14.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-15.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-16.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-17.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-18.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-19.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-20.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-21.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-22.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-23.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-24.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-25.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-26.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-27.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-28.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-29.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-30.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-31.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-32.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-33.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-34.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-35.jpg",
            "https://cdn.komiku.co.id/wp-content/uploads/85998-36.jpg"
        ]
    },
    "message": "Get Chapter Detail",
    "code": 200
}
   ```