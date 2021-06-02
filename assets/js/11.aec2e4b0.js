(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{365:function(e,a,n){"use strict";n.r(a);var t=n(25),s=Object(t.a)({},(function(){var e=this,a=e.$createElement,n=e._self._c||a;return n("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[n("h1",{attrs:{id:"benchmarks"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#benchmarks"}},[e._v("#")]),e._v(" Benchmarks")]),e._v(" "),n("p",[e._v("Following are the benchmarks comparing Kiwi with BuntDB on a MacBook Pro (8th gen\nIntel i5 2.4GHz processor, 8GB RAM).")]),e._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[e._v("❯ go test -bench=. -test.benchmem ./benchmark\ngoos: darwin\ngoarch: amd64\npkg: github.com/sdslabs/kiwi/benchmark\nBenchmarkBuntDB_Update-8        11777931                96.6 ns/op            48 B/op          1 allocs/op\nBenchmarkBuntDB_View-8          23310963                47.1 ns/op            48 B/op          1 allocs/op\nBenchmarkKiwi_Update-8          10356004               115 ns/op              48 B/op          3 allocs/op\nBenchmarkKiwi_Get-8             21910110                53.2 ns/op             0 B/op          0 allocs/op\nPASS\nok      github.com/sdslabs/kiwi/benchmark       6.216s\n")])])]),n("p",[e._v("Following are the key differences due to which Kiwi is a little slow:")]),e._v(" "),n("ol",[n("li",[e._v("BuntDB supports transactions, i.e., it locks the database once to apply all\nthe operations (and this is what is tested).")]),e._v(" "),n("li",[e._v("Kiwi supports dynamic data-types, which means, allocation on heap at runtime\n("),n("code",[e._v("interface{}")]),e._v(") whereas BuntDB is statically typed.")])]),e._v(" "),n("p",[e._v('The above two differences are what makes Kiwi unique and suitable to use on\nmany occasions. Due to the aforementioned reasons, Kiwi can support typed values\nand not everything is just another "string".')]),e._v(" "),n("p",[e._v("There are places where we could improve more. Some performance issues also lie\nin the implementation of values. For example, when updating a string, not returning\nthe updated string avoids an extra allocation.")])])}),[],!1,null,null,null);a.default=s.exports}}]);