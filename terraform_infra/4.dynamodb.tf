resource "aws_dynamodb_table" "token-email-lookup" {
  name           = "token-email-lookup"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key = "token"
  
  attribute {
    name = "token"
    type = "S"
  }
}
// add items to the table
resource "aws_dynamodb_table_item" "token1" {
  table_name = aws_dynamodb_table.token-email-lookup.name
  hash_key   = aws_dynamodb_table.token-email-lookup.hash_key

  item = <<ITEM
{
  "token": {"S": "token1"},
  "email": {"S": "useremail1@gmail.com"}
}
ITEM
}

resource "aws_dynamodb_table_item" "token2" {
  table_name = aws_dynamodb_table.token-email-lookup.name
  hash_key   = aws_dynamodb_table.token-email-lookup.hash_key

  item = <<ITEM
{
  "token": {"S": "token2"},
  "email": {"S": "useremail2@gmail.com"}
}
ITEM
}