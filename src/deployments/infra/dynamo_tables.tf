resource "aws_dynamodb_table" "directories" {
  name           = "${local.service_name}_directories"
  billing_mode   = "PROVISIONED"
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "S"
  }

}