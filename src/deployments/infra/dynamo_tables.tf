resource "aws_dynamodb_table" "directories" {
  name           = "${local.service_name}_directories"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "S"
  }

}

resource "aws_secretsmanager_secret" "directories_table" {
  name = "${local.service_name}_directories_table"
}

resource "aws_secretsmanager_secret_version" "directories_table" {
  secret_id     = aws_secretsmanager_secret.directories_table.id
  secret_string = aws_dynamodb_table.directories.name
}