namespace go echo
namespace py echo

struct Message {
  1: optional string text;
}

service Echo {
  string echo(1:Message msg);
}

