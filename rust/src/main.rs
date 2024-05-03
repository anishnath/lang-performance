use actix_web::{web, App, HttpServer, Responder, HttpResponse};
use serde::{Serialize};

#[derive(Serialize)]
struct Message {
    text: String,
}

#[derive(Serialize)]
struct FactorialResponse {
    input: u64,
    result: u64,
}

async fn ping() -> impl Responder {
    let message = Message { text: String::from("pong") };
    HttpResponse::Ok().json(message)
}

fn factorial(n: u64) -> u64 {
    match n {
        0 | 1 => 1,
        _ => n * factorial(n - 1)
    }
}

async fn factorial_handler(path: web::Path<u64>) -> impl Responder {
    let input = path.into_inner();
    let result = factorial(input);
    let response = FactorialResponse { input, result };
    HttpResponse::Ok().json(response)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/ping", web::get().to(ping))
            .route("/factorial/{number}", web::get().to(factorial_handler))
    })
        .bind("127.0.0.1:18081")?
        .run()
        .await
}
