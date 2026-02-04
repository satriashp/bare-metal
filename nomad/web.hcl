job "web-demo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    network {
      port "http" {
        static = 8080
      }
    }

    task "python-http" {
      driver = "exec"

      config {
        command = "/bin/sh"
        args = ["-c", "python3 -m http.server 8080 2>&1"]
      }

      service {
        name = "web-demo"
        port= "http"

        check {
          type = "http"
          path = "/"
          interval = "5s"
          timeout = "2s"
        }
      }
    }
  }
}