// package main

// import (
//     "github.com/gin-gonic/gin"
//     "github.com/99designs/gqlgen/graphql/handler"
//     "github.com/99designs/gqlgen/graphql/handler/transport"
//     "github.com/99designs/gqlgen/graphql/playground"
//     "GraphQLModule/graph"
// )

// func setupServer() *gin.Engine {
//     // Initialize Gin server
//     r := gin.Default()

//     // GraphQL handler setup
//     srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
//     srv.AddTransport(transport.POST{})
//     srv.AddTransport(transport.GET{})

//     // Register GraphQL endpoint
//     r.POST("/query", gin.WrapH(srv))

//     // Register GraphQL Playground endpoint
//     r.GET("/", playgroundHandler())

//     return r
// }

// func main() {
//     // Start the server
//     r := setupServer()
//     r.Run(":8080")
// }

// // playgroundHandler serves the GraphQL Playground
// func playgroundHandler() gin.HandlerFunc {
//     h := playground.Handler("GraphQL Playground", "/query")
//     return func(c *gin.Context) {
//         h.ServeHTTP(c.Writer, c.Request)
//     }
// }

package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/handler/transport"
    "GraphQLModule/graph"
)

func setupServer() *gin.Engine {
    // Initialize Gin server
    r := gin.Default()

    // GraphQL handler setup
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
    srv.AddTransport(transport.POST{})

    // Register GraphQL endpoint for POST requests
    r.POST("/query", func(c *gin.Context) {
        srv.ServeHTTP(c.Writer, c.Request)
    })

    // Register a welcome message for GET requests
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to GraphQL server!")
    })

    return r
}

func main() {
    // Start the server
    r := setupServer()
    r.Run(":8080")
}
