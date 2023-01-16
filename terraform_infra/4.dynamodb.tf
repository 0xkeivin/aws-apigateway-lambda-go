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

// user-notes

resource "aws_dynamodb_table" "user_notes" {
  name           = "user-notes"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key = "user"
  range_key = "create_date"  
  
  attribute {
    name = "user"
    type = "S"
  }
  attribute {
    name = "create_date"
    type = "S"
  }
}
resource "aws_dynamodb_table_item" "user1" {
  table_name = aws_dynamodb_table.user_notes.name
  hash_key   = aws_dynamodb_table.user_notes.hash_key
  range_key = aws_dynamodb_table.user_notes.range_key

  item = <<ITEM
{
  "user": {"S": "useremail1@gmail.com"},
  "create_date": {"S": "2023-01-01T00:00:00Z"},
  "text": {"S": "sample note1"}
}
ITEM
}

resource "aws_dynamodb_table_item" "user2" {
  table_name = aws_dynamodb_table.user_notes.name
  hash_key   = aws_dynamodb_table.user_notes.hash_key
  range_key = aws_dynamodb_table.user_notes.range_key

  item = <<ITEM
{
  "user": {"S": "useremail2@gmail.com"},
  "create_date": {"S": "2023-01-02T00:00:00Z"},
  "text": {"S": "sample note2"}
}
ITEM
}