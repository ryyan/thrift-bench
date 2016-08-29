service Echo {
  string echo(1:Message msg);
}

struct Message {
  1: optional string text;
}
