terraform {
  required_providers {
    demo2 = {
      source  = "local/ondrejsika/demo2"
      version = "0.0.1"
    }
  }
}

provider "demo2" {
  api_origin = "http://127.0.0.1:8000"
  token      = "xxx"
}

resource "demo2_box" "prague" {
  name = "Prague"
}

resource "demo2_text" "cncf" {
  box_id = demo2_box.prague.id
  text   = "Cloud Native Computing Foundation Prague"
}

resource "demo2_text" "pycon" {
  box_id = demo2_box.prague.id
  text   = "Pycon"
}
