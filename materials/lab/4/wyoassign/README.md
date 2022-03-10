Aram Maljanian

Option 2: modify/harden the wyoassign.go endpoints

1. added functionality to prevent duplicate id in create
2. added output for cases where no matching id is found in GetAssignment
3. Prevented deletions and updates without proper key ("key" = 453), also returned json message explaining why
    Example of a put with the key: http://127.0.0.1:8080/assignments/Mike1A?key=453&id=Mike1A