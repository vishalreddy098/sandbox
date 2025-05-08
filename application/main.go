package main

import (
    "sandbox/application/routes"
)

func main() {
    r := routes.SetupRouter()
    r.Run(":8080")
}
