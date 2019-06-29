package tcell_helpers

import (
  "testing"
)

func TestTextBoxFormatString(t *testing.T) {
  var result = textBoxFormatString("hello, this is a text to format", 10);
  println("in tests");

  for _, thing := range result {
  println(thing);
}
  t.Error("whaaat");
}
