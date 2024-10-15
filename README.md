# Crisis Core Materia Fusion API

An API based on the game [Crisis Core: Final Fantasy VII](https://en.wikipedia.org/wiki/Crisis_Core:_Final_Fantasy_VII) for simulating fusion of materia.

To try it out, skip to [Getting Started](#getting-started).

<img width="1062" alt="Screenshot" src="https://github.com/user-attachments/assets/83f2b32c-1856-41d9-85eb-6dfc1b76987f">

## Motivation for building this project

### The game doesn't make it easy to guess the fusion output

- The patterns and rules used to determine the fusion output of two materia are surprisingly complex. For example, the output can change if Materia 1 and Materia 2 are interchanged, or if one or both materia are 'Mastered'. These complex rules, when applied to 144 distinct materia, lead to over **280,000** possible permutations!
- Even trial and error is challenging, as you must first obtain the input materia in the game before you can see what the fusion output might be.

This API seeks to address both these pain points.

### An excuse to learn stuff

- I wanted to learn Go and PostgreSQL. The prospect of building something practical, useful AND for one of my favourite game series of all time was way more fun than going through yet another YouTube course.
- This will also explain why some of the [features](#technical-features) of the API feel like overkill.
- I've written about the journey I went through and the things I learnt while building this project (link coming soon).

## Getting started

- [OpenAPI documentation and testing](https://crisis-core-materia-fusion-api-546461677134.us-central1.run.app/docs). Try it out, no authentication required.
- Use the [Discussions tab](https://github.com/RayMathew/crisis-core-materia-fusion-api/discussions) for questions and comments.

### Running locally

Assumptions: You already have a PostgreSQL DB up and running. If you want to create a local instance in MacOS I recommend [Homebrew](https://formulae.brew.sh/formula/postgresql@16#default). If you want to create a remote instance, [try this](https://www.cockroachlabs.com/docs/stable/deploy-app-gcr).

1. After downloading the repo, create a `.env` file in the root directory with the following keys:

    - HTTP_PORT (recommended value 4444)

    - DB_DSN (format: `"<username>:<password>@<db_host_url>:<db_port>/<db_name>?sslmode=disabled"`). It should **not** be prefixed with `postgres://`

    - For DB_DSN, you can also use `sslmode=verify-full&sslrootcert=<path_to_cert>.crt` if connecting to a remote DB.

    - API_TIMEOUT_SECONDS (suggested value 5)

    - API_CALLS_ALLOWED_PER_SECOND (suggested value 1)

2. Run:

    ```sh
    go mod tidy
    go mod vendor
    make run
    ```

3. Optional: install golangci-lint

    ```sh
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    ```

### Running locally - with Docker

The Dockerfile is currently being used only for deploying on Google Cloud Run. If you wish to use it locally, do the following:

1. Create the `.env` file, just like in the above section.
2. Add the line `COPY .env /app/.env` to the Dockerfile to copy your secret environment variables into the Docker image.
3. Run:

    ```sh
    make dbuild
    make drun
    ```

### Testing locally

- Use `<rootfolder>/api/docs/swagger.json` (or the [hosted version](https://crisis-core-materia-fusion-api-546461677134.us-central1.run.app/docs/doc.json)) to import all endpoint definitions into tools like Postman.

## Tech stack

|                     |             |
|---------------------|-------------|
| Language            | Go          |
| DB                  | PostgreSQL  |
| Web server platform | [Google Cloud Run](https://cloud.google.com/run?hl=en)  |
| DB platform         | [CockroachDB](https://www.cockroachlabs.com/) |
| Documentation       | [go-swagger](https://goswagger.io/go-swagger/)     |

## Technical Features

1. The brilliant [Autostrada.dev](https://autostrada.dev/) was used to generate the project's boilerplate code.
2. The API uses an **in-server cache**, since both `/materia` and `/fusion` endpoints use the same data, and the data is static. So, requests to the DB happen only once after the server is redeployed.
3. The API has middleware for:

    - panic recovery
    - 'Content-Type'checking
    - rate limiting
    - api timeout protection

4. The inputs for the `/fusion` endpoint have validations for mandatory keys, data types, and materia names.
5. A `.golangci.yaml` file is available for lint checks in local. Install [golanci-lint](https://golangci-lint.run/) first.
6. A Dockerfile is available for building the API and running as a container. It is currently being used only for deploying on Cloud Run. To use locally, see the comments in `<rootfolder>/Makefile`.
7. The code constitutes over **1300 rules** to determine fusion ouputs.

## Future Roadmap

1. Add more unit tests.
2. Create a separate Dockerfile for local builds, and move PostgreSQL setup into it.
3. Reduce the setup required for `.env` file.
4. Add `.golangci.yaml` into the pipeline and fine-tune the individual linter checks. This might incur some costs, so it is the lowest priority for now.
5. The big one - create a new endpoint for materia fusion with items.

## Thanks

A huge thanks to the mysterious [ZeoKnight](https://gamefaqs.gamespot.com/community/ZeoKnight), who did most of the heavylifting in [figuring out the rules](https://gamefaqs.gamespot.com/psp/925138-crisis-core-final-fantasy-vii/faqs/75088/materia-combination-list) back in 2017. I did end up modifying a few of his rules that were either incorrect or incomplete, but it doesn't take anything away from the brilliance of his work.

## License

This project is licensed under the terms of the [GNU General Public License](https://www.gnu.org/licenses/gpl-3.0.en.html).
