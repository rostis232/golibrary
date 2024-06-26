# GoLibrary

GoLibrary is a web portal containing information about educational materials for Golang. The project uses the Echo framework and the Templ package for HTML page generation.

## Uses

- Golang 1.22

## Requirements

- Docker
- Docker Compose

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/rostis232/golibrary.git
   cd golibrary
   ```

2. Rename the file .env.example to .env

   ```sh
   mv .env.example .env
   ```

3. Fill in the environment variables:
   ```
   PG_PORT=
   PG_PASS=
   PG_USER=
   PG_DB_NAME=
   PORT=
   ```
4. Start Docker Compose:
   ```sh
   docker-compose up --build
   ```

## Usage

The web portal will be available once Docker Compose is up and running.

## Contact

- Telegram: [rostis232](https://t.me/rostis232)
- Email: [rostislav.pylypiv@gmail.com](mailto:rostislav.pylypiv@gmail.com)
- LinkedIn: [Rostyslav Pylypiv](https://www.linkedin.com/in/rostyslav-pylypiv/)

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
