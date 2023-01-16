resource "aws_dynamodb_table" "token-email-lookup" {
  name           = "token-email-lookup"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "token"

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
  hash_key       = "user"
  range_key      = "create_date"

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
  range_key  = aws_dynamodb_table.user_notes.range_key

  for_each = {
    "e845d3aa-9556-11ed-a1eb-0242ac120002" = {
      create_date = "2023-01-01T00:00:00Z"
      text        = "sample note1"
    }
    "e2da05b2-9556-1x6f-a1eb-0242ac120002" = {
      create_date = "2023-01-02T00:00:00Z"
      text        = "sample note2"
    }
    "e2da05b2-9556-2x3d-2s31-0242ac120002" = {
      create_date = "2023-01-11T00:00:00Z"
      text        = "sample note3"
    }

    "e2da05b2-9556-11ed-a1eb-0242ac120002" = {
      create_date = "2023-01-10T00:00:00Z"
      text        = "sample note4"
    }
    "e2da05b2-9556-1s61-a1eb-0242ac120002" = {
      create_date = "2023-01-09T00:00:00Z"
      text        = "sample note5"
    }
    "e2da05b2-9556-12f6-a1eb-0242ac120002" = {
      create_date = "2023-01-08T00:00:00Z"
      text        = "sample note6"
    }
    "e2da05b2-1s4f-11ed-a1eb-0242ac120002" = {
      create_date = "2023-01-07T00:00:00Z"
      text        = "sample note7"
    }
    "e2da05b2-9556-1s6v-a1eb-0242ac120002" = {
      create_date = "2023-01-06T00:00:00Z"
      text        = "sample note8"
    }
    "e2da05b2-9556-11ed-2d2s-0242ac120002" = {
      create_date = "2023-01-05T00:00:00Z"
      text        = "sample note9"
    }
    "e2da05b2-9556-11ed-2ds6-0242ac160002" = {
      create_date = "2023-01-04T00:00:00Z"
      text        = "sample note10"
    }
    "e2da05b2-9556-1s4s-a1eb-0242ac121002" = {
      create_date = "2023-02-03T00:00:00Z"
      text        = "sample note11"
    }
    "e2da05b2-9556-1s5s-a1eb-0242ac1210002" = {
      create_date = "2022-01-02T00:00:00Z"
      text        = "sample note12"
    }
  }
  item = <<ITEM
{
  "user": {"S": "useremail1@gmail.com"},
  "id": {"S": "${each.key}"},
  "create_date": {"S": "${each.value.create_date}"},
  "text": {"S": "${each.value.text}"}
}
ITEM
}

resource "aws_dynamodb_table_item" "user2" {
  table_name = aws_dynamodb_table.user_notes.name
  hash_key   = aws_dynamodb_table.user_notes.hash_key
  range_key  = aws_dynamodb_table.user_notes.range_key


  for_each = {
    "e3da05b2-9556-11ed-a1eb-0242ac120002" = {
      create_date = "2023-01-01T00:00:00Z"
      text        = "sample note1"
    }
    "e3da05b2-9556-1x6f-a1eb-0242ac120002" = {
      create_date = "2023-01-02T00:00:00Z"
      text        = "sample note2"
    }
    "e3da05b4-9556-2x3d-2s31-0242ac120002" = {
      create_date = "2023-01-11T00:00:00Z"
      text        = "sample note3"
    }

    "e3da05b2-9556-11ed-a1eb-0242ac120002" = {
      create_date = "2023-01-10T00:00:00Z"
      text        = "sample note4"
    }
    "e3da05b2-9556-1s61-a1eb-0242ac120002" = {
      create_date = "2023-01-09T00:00:00Z"
      text        = "sample note5"
    }
    "e3da05b2-9556-12f6-a1eb-0242ac120002" = {
      create_date = "2023-01-08T00:00:00Z"
      text        = "sample note6"
    }
    "e3da05b2-1s4f-11ed-a1eb-0242ac120002" = {
      create_date = "2023-01-07T00:00:00Z"
      text        = "sample note7"
    }
    "e3da05b2-9556-1s6v-a1eb-0242ac120002" = {
      create_date = "2023-01-06T00:00:00Z"
      text        = "sample note8"
    }
    "e3da05b2-9556-11ed-2d2s-0242ac120002" = {
      create_date = "2023-01-05T00:00:00Z"
      text        = "sample note9"
    }
    "e3da05b2-9556-11ed-2ds6-0242ac160002" = {
      create_date = "2023-01-04T00:00:00Z"
      text        = "sample note10"
    }
    "e3da05b2-9556-1s4s-a1eb-0242ac121002" = {
      create_date = "2023-02-03T00:00:00Z"
      text        = "sample note11"
    }
    "e3da05b2-9556-1s5s-a1eb-0242ac1210002" = {
      create_date = "2022-01-02T00:00:00Z"
      text        = "sample note12"
    }
  }
  item = <<ITEM
{
  "user": {"S": "useremail2@gmail.com"},
  "id": {"S": "${each.key}"},
  "create_date": {"S": "${each.value.create_date}"},
  "text": {"S": "${each.value.text}"}
}
ITEM
}
