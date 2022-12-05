from rust:1.61 AS builder

WORKDIR /usr/src/myapp

COPY . .

RUN cargo install --path app 

FROM debian:buster-slim

WORKDIR /rust/bin

COPY --from=builder /usr/src/myapp/app/target/release/app myapp

CMD ["./myapp"]
