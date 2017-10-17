
# mnemo for go

[![Build Status](https://secure.travis-ci.org/murdockq/mnemo.svg)](http://travis-ci.org/murdockq/mnemo)

A port of [mnemo](https://github.com/flon-io/mnemo) from C to Go lang. Mnemo and mnemo for go are inspired by the work of [rufus-mnemo](https://github.com/jmettraux/rufus-mnemo) and the ruby port [munemo](https://github.com/jmettraux/munemo).

Munemo uses the following syllables to map integer numbers to strings.
```go
    var syllables = []string{
        "ba", "bi", "bu", "be", "bo",
        "cha", "chi", "chu", "che", "cho",
        "da", "di", "du", "de", "do",
        "fa", "fi", "fu", "fe", "fo",
        "ga", "gi", "gu", "ge", "go",
        "ha", "hi", "hu", "he", "ho",
        "ja", "ji", "ju", "je", "jo",
        "ka", "ki", "ku", "ke", "ko",
        "la", "li", "lu", "le", "lo",
        "ma", "mi", "mu", "me", "mo",
        "na", "ni", "nu", "ne", "no",
        "pa", "pi", "pu", "pe", "po",
        "ra", "ri", "ru", "re", "ro",
        "sa", "si", "su", "se", "so",
        "sha", "shi", "shu", "she", "sho",
        "ta", "ti", "tu", "te", "to",
        "tsa", "tsi", "tsu", "tse", "tso",
        "wa", "wi", "wu", "we", "wo",
        "ya", "yi", "yu", "ye", "yo",
        "za", "zi", "zu", "ze", "zo",
    }
```
The syllable `xa` prefixes negative numbers.

Interface:
```go
import  "github.com/murdockq/mnemo"

// mnemo.ToString(int)

mnemo.ToString(0) # => 'ba'
mnemo.ToString(1) // => 'bi'
mnemo.ToString(99) // => 'zo'
mnemo.ToString(100) // => 'biba'
mnemo.ToString(101) // => 'bibi'
mnemo.ToString(392406) // => 'kogochi'
mnemo.ToString(25437225) // => 'haleshuha'

mnemo.ToString(-1) // => 'xabi'
mnemo.ToString(-99) // => 'xazo'
mnemo.ToString(-100) // => 'xabiba'


// mnemo.ToInt(string)

mnemo.ToInt('blah blah blah')
  // => ArgumentError: "unknown syllable 'bl'"

mnemo.ToInt('xabixabi')
  // => ArgumentError: "unknown syllable 'xa'"

mnemo.ToInt('yoshida') // => 947110
mnemo.ToInt('bajo') // => 34
mnemo.ToInt('xabaji') // => -31
mnemo.ToInt('tonukatsu') // => 79523582

// mnemo.Convert(var)

mnemo.Convert(99) // => 'zo'
mnemo.ToString('zo') // => 99 

```

## LICENSE

MIT, see [LICENSE.txt](LICENSE.txt)
