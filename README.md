# Hooper Looper!!

JS-like loops in golang. There's no data type restriction just use what you want as a `interface{}`. 

## Simple Usage!
### Loop over Slice.
```go
  var dump []string
	dump = []string{
		"v1", "v2", "v3",
	}

	iter := NewIterator(dump)

	data := iter.Slice(func(i, v interface{}) interface{} {
    // i = index
    // v = value of index
    iAsStr := strconv.Atoi(i)
		return v.(string) + " replaced on " + iAsStr
	})

	fmt.Println(data)
```

<hr>

### Loop over Map
```go
  dump := map[string][]string{
		"t1": {"v1", "v2", "v3"},
		"t2": {"v1", "v2", "v3"},
		"t3": {"v1", "v2", "v3"},
		"t4": {"v1", "v2", "v3"},
	}
  
  iter := NewIterator(dump)
	data := iter.Map(func(k, v interface{}) interface{} {
    // if you're sure that map's values are slice then you can
    // call .Slice() too
		//nIter := NewIterator(v)
		//nIter.Slice(func(sV interface{}) interface{} {
		//	fmt.Println(sV)
		//	return sV
		//})
		return v
	})

```


## Features~
* [] Filter
* [] Every
* [] Fill
* [] Find
* [] Join
* [] Keys
* [] Reduce
* [] Reverse
* [] Sort
