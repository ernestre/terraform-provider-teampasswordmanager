terraform {
  required_providers {
    tpmsync = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/tpmsync"
    }
  }
}

provider "tpmsync" {
  host        = "http://localhost:8081"
  public_key  = "1356a192b7913b04c54574d18c28d46e6395428ab44f2ef0cabc9347835b9ea5"
  private_key = "5c005bc16db8b0e9f407c6747d4656fc48bbf0d6773e681f47fd86e1e7d6009b"
}

# Creating projects
resource "tpmsync_project" "fp" {
  name = "first-project"
  tags = local.tags
}

locals {
  password_value = "s3cr3t"
  tags           = []
}

resource "tpmsync_project" "fpp" {
  name      = "child-project"
  notes     = "the note"
  tags      = local.tags
  parent_id = tpmsync_project.fp.id
}

resource "tpmsync_project" "fppp" {
  name      = "child-project"
  notes     = "the note"
  tags      = local.tags
  parent_id = tpmsync_project.fpp.id
}

resource "tpmsync_password" "password_fp" {
  name       = "child-project"
  password   = local.password_value
  project_id = tpmsync_project.fp.id
}

resource "tpmsync_password" "password_fpp" {
  name       = "child-project"
  password   = local.password_value
  project_id = tpmsync_project.fpp.id
}

resource "tpmsync_password" "password_fppp" {
  name       = "child-project"
  password   = local.password_value
  project_id = tpmsync_project.fppp.id
}
