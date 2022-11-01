resource "aws_dynamodb_table" "directories" {
  name           = "${local.service_name}_directories"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "S"
  }

}