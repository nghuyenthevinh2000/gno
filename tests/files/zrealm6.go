// PKGPATH: gno.land/r/test
package test

import (
	"github.com/gnolang/gno/_test/timtadh/data-structures/tree/avl"
	"github.com/gnolang/gno/_test/timtadh/data-structures/types"
)

var tree *avl.AvlNode

func init() {
	tree, _ = tree.Put(types.String("key0"), "value0")
	tree, _ = tree.Put(types.String("key1"), "value1")
}

func main() {
	var updated bool
	tree, updated = tree.Put(types.String("key3"), "value3")
	println(updated, tree.Size())
}

// Output:
// false 3

// Realm:
// c[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:8]=(gno.StructValueImage) {
//  ObjectInfo: (gno.ObjectInfoImage) {
//   _RealmID: (gno.RealmID) RIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30,
//   NewTime: (uint64) 8,
//   _OwnerNewTime: (uint64) 4,
//   _ModTime: (uint64) 0,
//   _RefCount: (int) 1
//  },
//  Fields: ([]gno.TypedValueImage) (len=5 cap=5) {
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 63CDE69354F70377B65D4C6BDDBD1D23A8AF7217,
//    ValueImage: (gno.PrimitiveValueImage) (len=4 cap=8) {
//     00000000  6b 65 79 33                                       |key3|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 473287F8298DBA7163A897908958F7C0EAE733E2,
//    ValueImage: (gno.PrimitiveValueImage) (len=6 cap=8) {
//     00000000  76 61 6c 75 65 33                                 |value3|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 6DA88C34BA124C41F977DB66A4FC5C1A951708D2,
//    ValueImage: (gno.PrimitiveValueImage) (len=8 cap=8) {
//     00000000  01 00 00 00 00 00 00 00                           |........|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 0000000000000000000000000000000000000000,
//      ValueImage: (gno.ValueImage) <nil>
//     }
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 0000000000000000000000000000000000000000,
//      ValueImage: (gno.ValueImage) <nil>
//     }
//    }
//   }
//  }
// }
//
// u[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:4]=(gno.StructValueImage) {
//  ObjectInfo: (gno.ObjectInfoImage) {
//   _RealmID: (gno.RealmID) RIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30,
//   NewTime: (uint64) 4,
//   _OwnerNewTime: (uint64) 3,
//   _ModTime: (uint64) 5,
//   _RefCount: (int) 1
//  },
//  Fields: ([]gno.TypedValueImage) (len=5 cap=5) {
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 63CDE69354F70377B65D4C6BDDBD1D23A8AF7217,
//    ValueImage: (gno.PrimitiveValueImage) (len=4 cap=8) {
//     00000000  6b 65 79 31                                       |key1|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 473287F8298DBA7163A897908958F7C0EAE733E2,
//    ValueImage: (gno.PrimitiveValueImage) (len=6 cap=8) {
//     00000000  76 61 6c 75 65 31                                 |value1|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 6DA88C34BA124C41F977DB66A4FC5C1A951708D2,
//    ValueImage: (gno.PrimitiveValueImage) (len=8 cap=8) {
//     00000000  02 00 00 00 00 00 00 00                           |........|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 0000000000000000000000000000000000000000,
//      ValueImage: (gno.ValueImage) <nil>
//     }
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 4AF0F175D54357F0FEEAE4CF180A42BE848369E8,
//      ValueImage: (gno.RefImage) {
//       RealmID: (gno.RealmID) RID0000000000000000000000000000000000000000,
//       NewTime: (uint64) 8,
//       Hash: (gno.ValueHash) {
//        Hashlet: (gno.Hashlet) (len=20 cap=20) {
//         00000000  52 b4 56 4d 4f d3 70 c9  a7 89 4b 7e 06 36 d7 42  |R.VMO.p...K~.6.B|
//         00000010  d8 98 dc 0a                                       |....|
//        }
//       }
//      }
//     }
//    }
//   }
//  }
// }
//
// u[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:3]=(gno.StructValueImage) {
//  ObjectInfo: (gno.ObjectInfoImage) {
//   _RealmID: (gno.RealmID) RIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30,
//   NewTime: (uint64) 3,
//   _OwnerNewTime: (uint64) 0,
//   _ModTime: (uint64) 6,
//   _RefCount: (int) 1
//  },
//  Fields: ([]gno.TypedValueImage) (len=5 cap=5) {
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 63CDE69354F70377B65D4C6BDDBD1D23A8AF7217,
//    ValueImage: (gno.PrimitiveValueImage) (len=4 cap=8) {
//     00000000  6b 65 79 30                                       |key0|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 473287F8298DBA7163A897908958F7C0EAE733E2,
//    ValueImage: (gno.PrimitiveValueImage) (len=6 cap=8) {
//     00000000  76 61 6c 75 65 30                                 |value0|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 6DA88C34BA124C41F977DB66A4FC5C1A951708D2,
//    ValueImage: (gno.PrimitiveValueImage) (len=8 cap=8) {
//     00000000  03 00 00 00 00 00 00 00                           |........|
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 0000000000000000000000000000000000000000,
//      ValueImage: (gno.ValueImage) <nil>
//     }
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 4AF0F175D54357F0FEEAE4CF180A42BE848369E8,
//      ValueImage: (gno.RefImage) {
//       RealmID: (gno.RealmID) RID0000000000000000000000000000000000000000,
//       NewTime: (uint64) 4,
//       Hash: (gno.ValueHash) {
//        Hashlet: (gno.Hashlet) (len=20 cap=20) {
//         00000000  20 75 46 62 fc a5 cf 7a  20 cb bd 03 67 4f b9 53  | uFb...z ...gO.S|
//         00000010  67 f9 47 98                                       |g.G.|
//        }
//       }
//      }
//     }
//    }
//   }
//  }
// }
//
// u[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:0]=(gno.BlockValueImage) {
//  ObjectInfo: (gno.ObjectInfoImage) {
//   _RealmID: (gno.RealmID) RIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30,
//   NewTime: (uint64) 0,
//   _OwnerNewTime: (uint64) 0,
//   _ModTime: (uint64) 7,
//   _RefCount: (int) 0
//  },
//  ParentID: (gno.ObjectID) OIDNONE:0,
//  Values: ([]gno.TypedValueImage) (len=3 cap=3) {
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 0BA050DA455A6AAD7074EB2148D53ECD5BECC26D,
//    ValueImage: (gno.FuncValueImage) {
//     TypeID: (gno.TypeID) (len=20 cap=20) 0BA050DA455A6AAD7074EB2148D53ECD5BECC26D,
//     IsMethod: (bool) false,
//     Name: (gno.Name) (len=6) "init.0",
//     ClosureRef: (gno.RefImage) {
//      RealmID: (gno.RealmID) RID0000000000000000000000000000000000000000,
//      NewTime: (uint64) 0,
//      Hash: (gno.ValueHash) {
//       Hashlet: (gno.Hashlet) (len=20 cap=20) {
//        00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
//        00000010  00 00 00 00                                       |....|
//       }
//      }
//     },
//     FileName: (gno.Name) (len=16) "files/zrealm6.go",
//     PkgPath: (string) (len=15) "gno.land/r/test"
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) 0BA050DA455A6AAD7074EB2148D53ECD5BECC26D,
//    ValueImage: (gno.FuncValueImage) {
//     TypeID: (gno.TypeID) (len=20 cap=20) 0BA050DA455A6AAD7074EB2148D53ECD5BECC26D,
//     IsMethod: (bool) false,
//     Name: (gno.Name) (len=4) "main",
//     ClosureRef: (gno.RefImage) {
//      RealmID: (gno.RealmID) RID0000000000000000000000000000000000000000,
//      NewTime: (uint64) 0,
//      Hash: (gno.ValueHash) {
//       Hashlet: (gno.Hashlet) (len=20 cap=20) {
//        00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
//        00000010  00 00 00 00                                       |....|
//       }
//      }
//     },
//     FileName: (gno.Name) (len=16) "files/zrealm6.go",
//     PkgPath: (string) (len=15) "gno.land/r/test"
//    }
//   },
//   (gno.TypedValueImage) {
//    TypeID: (gno.TypeID) (len=20 cap=20) E6E0E2CE563ADB23D6A4822DD5FC346A5DE899A0,
//    ValueImage: (gno.PointerValueImage) {
//     TypedValue: (gno.TypedValueImage) {
//      TypeID: (gno.TypeID) (len=20 cap=20) 4AF0F175D54357F0FEEAE4CF180A42BE848369E8,
//      ValueImage: (gno.RefImage) {
//       RealmID: (gno.RealmID) RID0000000000000000000000000000000000000000,
//       NewTime: (uint64) 3,
//       Hash: (gno.ValueHash) {
//        Hashlet: (gno.Hashlet) (len=20 cap=20) {
//         00000000  51 21 c7 7c 18 73 2a 69  1c 6b 2d ad d7 ee ba b5  |Q!.|.s*i.k-.....|
//         00000010  f1 a2 b0 cd                                       |....|
//        }
//       }
//      }
//     }
//    }
//   }
//  }
// }
//
