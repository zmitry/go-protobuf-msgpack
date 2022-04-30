using Go = import "/go.capnp";
@0xb7eeedc03c4a0186;
$Go.package("foo");
$Go.import("foo/product");

struct Product {
    id @0 :UInt64;
    name @1 :Text;
    description @2 :Text;
    price @3 :Int32;
    colors @4 :List(Text);
}